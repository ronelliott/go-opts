package opts

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestOptionSetStruct struct {
	Args []string `positional:"true"`

	SnoringChihuahua string

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

type TestInvalidOptionSetStruct struct {
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

type TestEmptyOptionSetStruct struct{}

type TestParseDefaultArgsStruct struct {
	CoverageOut     string `long:"test.outputdir"`
	CoverageProfile string `long:"test.coverprofile"`
	Verbose         bool   `long:"test.v"`
}

type TestInvalidDefaultOptionSetStruct struct {
	Verbose bool `
        default:"Nein"
        description:"Use verbose logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`
}

func TestNewOptionSet(t *testing.T) {
	set, err := NewOptionSet(&TestOptionSetStruct{})
	require.Nil(t, err)
	require.Equal(t, 3, len(set.Options))
	_, ok := set.Options["Name"]
	require.True(t, ok)
	_, ok = set.Options["Verbose"]
	require.True(t, ok)
	_, ok = set.Options["Ducks"]
	require.False(t, ok)
	_, ok = set.Options["SnoringChihuahua"]
	require.False(t, ok)
}

func TestNewOptionSet_NonPointer(t *testing.T) {
	_, err := NewOptionSet(TestOptionSetStruct{})
	require.NotNil(t, err)
}

func TestNewOptionSet_InvalidDefault(t *testing.T) {
	_, err := NewOptionSet(&TestInvalidDefaultOptionSetStruct{})
	require.NotNil(t, err)
}

func TestOptionSet_HasOptions(t *testing.T) {
	set, err := NewOptionSet(&TestOptionSetStruct{})
	require.Nil(t, err)
	require.True(t, set.HasOptions())

	set, err = NewOptionSet(&TestEmptyOptionSetStruct{})
	require.Nil(t, err)
	require.False(t, set.HasOptions())
}

func TestOptionSetParse_DefaultArgs(t *testing.T) {
	opts := TestParseDefaultArgsStruct{}
	set, err := NewOptionSet(&opts)
	require.Nil(t, err)
	err = set.Parse(nil)
	require.Nil(t, err)
}

func TestOptionSetParse_Defaults(t *testing.T) {
	opts := TestOptionSetStruct{}
	set, err := NewOptionSet(&opts)
	require.Nil(t, err)
	err = set.Parse([]string{})
	require.Nil(t, err)
	require.Equal(t, "foo", opts.Name)
	require.False(t, opts.Verbose)
}

func TestOptionSetParse_Short(t *testing.T) {
	opts := TestOptionSetStruct{}
	set, err := NewOptionSet(&opts)
	require.Nil(t, err)
	err = set.Parse([]string{"-v", "-n", "bar"})
	require.Nil(t, err)
	require.Equal(t, "bar", opts.Name)
	require.True(t, opts.Verbose)
}

func TestOptionSetParse_Long(t *testing.T) {
	opts := TestOptionSetStruct{}
	set, err := NewOptionSet(&opts)
	require.Nil(t, err)
	err = set.Parse([]string{"--verbose", "--name", "bar"})
	require.Nil(t, err)
	require.Equal(t, "bar", opts.Name)
	require.True(t, opts.Verbose)
}

func TestOptionSetParse_NonstandardLong(t *testing.T) {
	opts := TestOptionSetStruct{}
	set, err := NewOptionSet(&opts)
	require.Nil(t, err)
	err = set.Parse([]string{"-verbose", "-name", "bar"})
	require.Nil(t, err)
	require.Equal(t, "bar", opts.Name)
	require.True(t, opts.Verbose)
}

func TestOptionSetParse_Undeclared(t *testing.T) {
	opts := TestOptionSetStruct{}
	set, err := NewOptionSet(&opts)
	require.Nil(t, err)
	err = set.Parse([]string{"-ducks"})
	require.NotNil(t, err)
}

func TestOptionSetParse_LeftoverArgs(t *testing.T) {
	opts := TestOptionSetStruct{}
	set, err := NewOptionSet(&opts)
	require.Nil(t, err)
	err = set.Parse([]string{"-verbose", "-name", "bar", "far", "zar"})
	require.Nil(t, err)
	require.Equal(t, []string{"far", "zar"}, opts.Args)
}

func TestOptionSetParse_LeftoverArgs_InvalidType(t *testing.T) {
	opts := TestInvalidOptionSetStruct{}
	_, err := NewOptionSet(&opts)
	require.NotNil(t, err)
}

func TestOptionSetWriteHelp(t *testing.T) {
	set, err := NewOptionSet(&TestOptionSetStruct{})
	require.Nil(t, err)
	buf := bytes.Buffer{}
	set.WriteHelp(&buf)

	expected := "  -n string\n    \tThe name to use (default \"foo\")\n  " +
		"-name string\n    \tThe name to use (default \"foo\")\n  " +
		"-v\tUse verbose logging.\n  -verbose\n    \tUse verbose " +
		"logging.\n"

	require.Equal(t, expected, buf.String())
}
