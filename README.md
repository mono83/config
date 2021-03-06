Simple application configuration tool
=====================================

# Key features

Here is the list of features you may be interested in:

- :heavy_check_mark: **Simple** 
- :heavy_check_mark: Supports JSON (Using `encoding/json`)
- :heavy_check_mark: Supports YAML (Using https://gopkg.in/yaml.v2)
- :heavy_check_mark: Supports TOML (Using https://github.com/BurntSushi/toml)
- :heavy_check_mark: Supports INI (Using https://gopkg.in/ini.v1)
- :heavy_check_mark: Aliases, subfolders and on-demand lookup in home folder (github.com/mitchellh/go-homedir) and `/etc/`
- :heavy_check_mark: Data validation
- :x: Placeholders
- :x: Configuration file chaining
- :x: Autoreload

## Before you start

First of all, you need some structure, that will hold configuration
for your application. Like this:

```go
type Config struct {
    Token string `json:"token" yaml:"token" toml:"token"`
}
```

## Simplest example

Just import and call `config.ReadDefault`.

```go
package main

import "github.com/mono83/config"

func main() {
  var c Config
  if err := config.ReadDefault(&c); err != nil {
    panic(err)
  }

  fmt.Println(c)
}
```

Module will make attempt to find one of these files: `config.json`,
`config.yaml`, `config.toml` in current folder and unmarshal them using
corresponding unmarshaller.

## More complex usecase

You can specify filename and other options using `config.Source`
structure (event `config.ReadDefault` uses it intenally):

```go
if err := (config.Source{FileName: "app.json"}).Read(&c); err != nil {
  panic(err)
}
```

There can be multiple aliases for configuration files (module reads only
one - that was specified/found first, no joining applied):

```go
if err := (config.Source{
    FileNames: []string{"app.json", "config.json"},
  }).Read(&c); err != nil {
  panic(err)
}
```

You can specify subfolder for configuration files:

```go
if err := (config.Source{
    Subfolder: "configuration",
    FileNames: []string{"app.json", "config.json"},
  }).Read(&c); err != nil {
  panic(err)
}
```

Or even allow lookup in homedir and `/etc/`

```go
if err := (config.Source{
    Subfolder: "configuration",
    FileNames: []string{"app.json", "config.json"},
    LookupHome: true,
    LookupEtc: true,
  }).Read(&c); err != nil {
  panic(err)
}
```

## Validation

If structure, that is used in configuration has method `Validate() error`, 
configuration reader will automatically invoke it:

```go
type Config struct {
    Token string `json:"token" yaml:"token" toml:"token"`
}

// Validate will be automatically invoked by mono83/config
func (c Config) Validate() error {
  if len(c.Token) == 0 {
    return errors.New("empty access token")
  }

  return nil
}
```

This behaviour can be disabled during `config.Source` initialization:

```go 

(config.Source{SkipValidation: true, ...}).Read(...)

```