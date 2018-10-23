package config

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

func iniUnmarshal(data []byte, v interface{}) error {
	c, err := ini.Load(data)
	if err != nil {
		return err
	}
	return c.MapTo(v)
}

func readBytes(filename string, bts []byte, target interface{}) error {
	ft := detectFormat(filename)

	// Initializing marshaller
	marsh := json.Unmarshal

	if ft == fYAML {
		marsh = yaml.Unmarshal
	} else if ft == fTOML {
		marsh = toml.Unmarshal
	} else if ft == fINI {
		marsh = iniUnmarshal
	}

	err := marsh(bts, target)
	if err != nil {
		err = MarshalError{Cause: err.Error(), FileName: filename, Struct: target}
	}
	return err
}
