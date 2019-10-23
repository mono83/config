package config

import (
	"errors"
	"net/url"
	"strings"
)

// URL stands for URL from configuration
type URL string

func (u URL) String() string {
	return string(u)
}

// IsEmpty returns true if URL contains empty string
func (u URL) IsEmpty() bool {
	return len(u) == 0
}

// ToGoURL converts URL struct into Golang URL
func (u URL) ToGoURL() (*url.URL, error) {
	return url.Parse(u.String())
}

// Slashed appends slash to URL if was not has it
func (u URL) Slashed() URL {
	if strings.HasSuffix(u.String(), "/") {
		return u
	}

	return URL(u + "/")
}

// Validate performs value validation
func (u URL) Validate() error {
	if u.IsEmpty() {
		return errors.New("empty URL")
	}

	g, err := u.ToGoURL()
	if err == nil {
		if g.Scheme != "http" && g.Scheme != "https" {
			return errors.New(" expected http(s) scheme, but got " + g.Scheme)
		}
	}
	return err
}
