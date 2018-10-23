package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBytesJSON(t *testing.T) {
	source := `{"id":35, "name": "foo", "enabled": true}`
	var tc testConfig
	if assert.NoError(t, readBytes("foo.json", []byte(source), &tc)) {
		assert.Equal(t, testConfig{A: 35, B: "foo", C: true}, tc)
	}
}

func TestReadBytesYAML(t *testing.T) {
	source := "id: -43\nname: bar\nenabled: true"
	var tc testConfig
	if assert.NoError(t, readBytes("foo.yaml", []byte(source), &tc)) {
		assert.Equal(t, testConfig{A: -43, B: "bar", C: true}, tc)
	}
}

func TestReadBytesTOML(t *testing.T) {
	source := "id = 234\nname = \"baz\"\nenabled = true"
	var tc testConfig
	if assert.NoError(t, readBytes("foo.toml", []byte(source), &tc)) {
		assert.Equal(t, testConfig{A: 234, B: "baz", C: true}, tc)
	}
}

func TestReadBytesINI(t *testing.T) {
	source := "id = 11\nname = \"queue\"\nenabled = true"
	var tc testConfig
	if assert.NoError(t, readBytes("foo.ini", []byte(source), &tc)) {
		assert.Equal(t, testConfig{A: 11, B: "queue", C: true}, tc)
	}
}

type testConfig struct {
	A int    `json:"id" yaml:"id" toml:"id" ini:"id"`
	B string `json:"name" yaml:"name" toml:"name" ini:"name"`
	C bool   `json:"enabled" yaml:"enabled" toml:"enabled" ini:"enabled"`
}
