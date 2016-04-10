# go-opts

[![GoDoc](https://godoc.org/github.com/ronelliott/go-opts?status.png)](https://godoc.org/github.com/ronelliott/go-opts)
[![Build Status](https://travis-ci.org/ronelliott/go-opts.svg?branch=master)](https://travis-ci.org/ronelliott/go-opts)
[![Coverage Status](https://img.shields.io/coveralls/ronelliott/go-opts.svg)](https://coveralls.io/r/ronelliott/go-opts?branch=master)

a go library for parsing command line flags

## Installation

    $ go get github.com/ronelliott/go-opts

## Example

```go
package main

import (
    "fmt"
    "github.com/ronelliott/go-opts"
    "os"
)

type Options struct {
    Args []string `positional:"true"`

    Name string `
        default:"foo"
        description:"The name to use"
        help:"What do you want to name this thing?"
        long:"name"
        short:"n"`

    Verbose bool `
        default:"false"
        description:"Use verbose logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`
}

func main() {
    options := Options{}
    err := opts.Parse(&options, nil)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("Options.Args:", options.Args)
    fmt.Println("Options.Name:", options.Name)
    fmt.Println("Options.Verbose:", options.Verbose)
}
```