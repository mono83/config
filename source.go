package config

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

// Source contains information about configuration source file
type Source struct {
	FileName  string   // File name of configuration file. Can be without extension
	FileNames []string // Aliases for configuation files

	Subfolder string // Subfolder name

	SkipValidation bool // Disable automatic validation for structs with Validate method
	LookupHome     bool // If true, module will make attempt to find config inside user home folder
	LookupEtc      bool // If true, module will make attempt to find config in /etc/
}

// GetAllFileNames returns all file names for configuration files
func (s Source) GetAllFileNames() []string {
	var total []string
	if len(s.FileName) > 0 {
		total = append(total, expandName(s.FileName)...)
	}

	for _, f := range s.FileNames {
		if len(f) > 0 {
			total = append(total, expandName(f)...)
		}
	}

	return total
}

// GetAllFolders returns all folders to lookup for configuration files
func (s Source) GetAllFolders() []string {
	total := []string{"."}

	hasSubfolder := len(s.Subfolder) > 0

	if hasSubfolder {
		total = append(total, "."+string(os.PathSeparator)+s.Subfolder+string(os.PathSeparator))
	}

	if s.LookupHome {
		if home, err := homedir.Dir(); err == nil {
			total = append(total, home)
		}
	}

	if s.LookupEtc && os.PathSeparator == '/' {
		if hasSubfolder {
			total = append(total, "/etc/"+s.Subfolder+"/")
		} else {
			total = append(total, "/etc/")
		}
	}

	return total
}

// GetAllPaths returns all file paths
func (s Source) GetAllPaths() []string {
	var total []string
	for _, p := range s.GetAllFolders() {
		for _, f := range s.GetAllFileNames() {
			total = append(total, pathConcat(p, f))
		}
	}

	return total
}

// Find searches for existing file
func (s Source) Find() (string, bool) {
	for _, p := range s.GetAllPaths() {
		if _, err := os.Stat(p); err == nil {
			return p, true
		}
	}

	return "", false
}

// Read method reads configuration from file
func (s Source) Read(target interface{}) error {
	f, ok := s.Find()
	if !ok {
		return LocateAndReadError{
			Paths: s.GetAllPaths(),
			Cause: "no configuration file found",
		}
	}

	// Reading all bytes
	bts, err := ioutil.ReadFile(f)
	if err != nil {
		return fmt.Errorf(
			`unable to read configuration file "%s" - %s`,
			f, err.Error(),
		)
	}

	err = readBytes(f, bts, target)
	if err == nil && !s.SkipValidation {
		if v, ok := target.(validable); ok {
			err = v.Validate()
			if err != nil {
				err = MarshalError{Cause: err.Error(), FileName: f, Struct: target}
			}
		}
	}

	return err
}

type validable interface {
	Validate() error
}
