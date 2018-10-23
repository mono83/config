package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dataProvider = []struct {
	Source  Source
	Folders []string
	Files   []string
}{
	{
		Source:  Source{FileName: "config.yaml"},
		Folders: []string{"."},
		Files:   []string{"config.yaml"},
	},
	{
		Source:  Source{FileName: "config.yaml", FileNames: []string{"alt.json"}, LookupEtc: true},
		Folders: []string{".", "/etc/"},
		Files:   []string{"config.yaml", "alt.json"},
	},
	{
		Source:  Source{FileName: "config", FileNames: []string{"alt.json", "beta"}, Subfolder: "app", LookupEtc: true},
		Folders: []string{".", "./app/", "/etc/app/"},
		Files:   []string{"config.json", "config.yaml", "config.yml", "config.toml", "alt.json", "beta.json", "beta.yaml", "beta.yml", "beta.toml"},
	},
}

func TestSource_GetAllFolders(t *testing.T) {
	for _, data := range dataProvider {
		t.Run(fmt.Sprintf("%v", data), func(t *testing.T) {
			assert.Equal(t, data.Folders, data.Source.GetAllFolders())
		})
	}
}

func TestSource_GetAllFileNames(t *testing.T) {
	for _, data := range dataProvider {
		t.Run(fmt.Sprintf("%v", data), func(t *testing.T) {
			assert.Equal(t, data.Files, data.Source.GetAllFileNames())
		})
	}
}
