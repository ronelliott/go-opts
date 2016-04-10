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
