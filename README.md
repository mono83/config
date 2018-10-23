Simple application configuration tool
=====================================

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

## Simple example

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