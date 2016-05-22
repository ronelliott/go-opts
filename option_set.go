package opts

import (
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

type OptionSet struct {
	// the options in this set
	Options map[string]*Option

	// the flags for this set
	flags *flag.FlagSet
}

// Creates a new OptionSet for the given struct
func NewOptionSet(data interface{}) (*OptionSet, error) {
	dataType := reflect.TypeOf(data)

	if dataType.Kind() != reflect.Ptr {
		return nil, errors.New("Data type is not a pointer.")
	}

	dataType = dataType.Elem()
	dataValue := reflect.ValueOf(data).Elem()

	set := OptionSet{
		Options: map[string]*Option{},
		flags:   flag.NewFlagSet(dataType.Name(), flag.ContinueOnError),
	}

	// flags package outputs to os.Stderr in certain cases, stifle this by
	// setting it to write to /dev/null
	set.flags.SetOutput(ioutil.Discard)

	for n := 0; n < dataType.NumField(); n++ {
		fieldType := dataType.Field(n)

		// ignore field without tags
		if fieldType.Tag == "" {
			continue
		}

		fieldValue := dataValue.Field(n)
		opt, err := NewOption(fieldType, fieldValue)

		if err != nil {
			return nil, err
		}

		// skip adding positional args to FlagSet
		if !opt.IsPositional() {
			err = opt.AddToFlagSet(set.flags)

			if err != nil {
				return nil, err
			}
		}

		set.Options[opt.Name] = opt
	}

	return &set, nil
}

// Checks if the OptionSet has options
func (this *OptionSet) HasOptions() bool {
	return len(this.Options) != 0
}

// Parses the given args using this OptionSet
func (this *OptionSet) Parse(args []string) error {
	if args == nil {
		args = os.Args[1:]
	}

	err := this.flags.Parse(args)

	if err != nil {
		return err
	}

	leftovers := this.flags.Args()

	for _, opt := range this.Options {
		if opt.IsPositional() {
			var ptr *[]string = opt.pointer.(*[]string)
			*ptr = leftovers
		}
	}

	return nil
}

// Writes the default options and descriptions to the given io.Writer
func (this *OptionSet) WriteHelp(out io.Writer) {
	this.flags.SetOutput(out)
	this.flags.PrintDefaults()
	// flags package outputs to os.Stderr in certain cases, stifle this by
	// setting it to write to /dev/null
	this.flags.SetOutput(ioutil.Discard)
}
