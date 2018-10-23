package config

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

func readBytes(filename string, bts []byte, target interface{}) error {
	ft := detectFormat(filename)

	// Initializing marshaller
	marsh := json.Unmarshal

	if ft == fYAML {
		marsh = yaml.Unmarshal
	} else if ft == fTOML {
		marsh = toml.Unmarshal
	}

	err := marsh(bts, target)
	if err != nil {
		err = MarshalError{Cause: err.Error(), FileName: filename, Struct: target}
	}
	return err
}
