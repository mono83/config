package config

import (
	"bytes"
	"encoding/json"
	"errors"
)

// ExpandError expands errors with more details
func ExpandError(err error) error {
	return errors.New(ExpandErrorMessage(err))
}

// ExpandErrorMessage expands error message with detailed information
func ExpandErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	if l, ok := err.(LocateAndReadError); ok {
		// Formatting message
		buf := bytes.NewBuffer(nil)
		buf.WriteString(l.Cause)
		buf.WriteString("\nLookup locations for configuration files:\n")
		for _, p := range l.Paths {
			buf.WriteString("  ")
			buf.WriteString(p)
			buf.WriteRune('\n')
		}
		return buf.String()
	}
	if m, ok := err.(MarshalError); ok {
		// Formatting message
		s, _ := json.MarshalIndent(m.Struct, "", "  ")

		buf := bytes.NewBuffer(nil)
		buf.WriteString(m.Cause)
		buf.WriteString("\nConfiguration file: ")
		buf.WriteString(m.FileName)
		buf.WriteString("\nConfiguration structure in JSON format (approximately):\n")
		buf.Write(s)
		buf.WriteByte('\n')

		return buf.String()
	}

	return err.Error()
}

// LocateAndReadError is an error, emitted when configuration file
// was not found or there were read error
type LocateAndReadError struct {
	Paths []string
	Cause string
}

func (l LocateAndReadError) Error() string {
	return l.Cause
}

// MarshalError is an error, emitted on marshalling error
type MarshalError struct {
	FileName string
	Struct   interface{}
	Cause    string
}

func (m MarshalError) Error() string {
	return m.Cause
}
