package opts

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type TestParseStruct struct {
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

type TestParseInvalidStruct struct {
	Args string `positional:"true"`

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

func TestParse_Valid(t *testing.T) {
	opts := TestParseStruct{}
	err := Parse(&opts, []string{"-v", "--name", "bar", "duck", "sauce"})
	require.Nil(t, err)
	require.True(t, opts.Verbose)
	require.Equal(t, "bar", opts.Name)
	require.Equal(t, []string{"duck", "sauce"}, opts.Args)
}

func TestParse_Invalid(t *testing.T) {
	opts := TestParseInvalidStruct{}
	err := Parse(&opts, []string{"-v", "--name", "foo", "ducks"})
	require.NotNil(t, err)
}
