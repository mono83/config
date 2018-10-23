package config

import "strings"

// format describes file format
type format byte

const (
	fUnknown format = iota
	fJSON
	fYAML
	fTOML
)

func detectFormat(filename string) format {
	if l := strings.LastIndex(filename, "."); l > -1 {
		switch strings.ToLower(filename[l+1:]) {
		case "json":
			return fJSON
		case "yml", "yaml":
			return fYAML
		case "toml":
			return fTOML
		}
	}

	return fUnknown
}

func expandName(filename string) []string {
	if fUnknown != detectFormat(filename) {
		return []string{filename}
	}

	return []string{filename + ".json", filename + ".yaml", filename + ".yml", filename + ".toml"}
}
