package opts

import (
    "bytes"
    "github.com/stretchr/testify/assert"
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

type TestParseDefaultArgsStruct struct {
    CoverageOut string `long:"test.outputdir"`
    CoverageProfile string `long:"test.coverprofile"`
    Verbose bool `long:"test.v"`
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
    assert.Nil(t, err)
    assert.Equal(t, 3, len(set.Options))
    _, ok := set.Options["Name"]
    assert.True(t, ok)
    _, ok = set.Options["Verbose"]
    assert.True(t, ok)
    _, ok = set.Options["Ducks"]
    assert.False(t, ok)
    _, ok = set.Options["SnoringChihuahua"]
    assert.False(t, ok)
}

func TestNewOptionSet_NonPointer(t *testing.T) {
    _, err := NewOptionSet(TestOptionSetStruct{})
    assert.NotNil(t, err)
}

func TestNewOptionSet_InvalidDefault(t *testing.T) {
    _, err := NewOptionSet(&TestInvalidDefaultOptionSetStruct{})
    assert.NotNil(t, err)
}

func TestOptionSetParse_DefaultArgs(t *testing.T) {
    opts := TestParseDefaultArgsStruct{}
    set, err := NewOptionSet(&opts)
    assert.Nil(t, err)
    err = set.Parse(nil)
    assert.Nil(t, err)
}

func TestOptionSetParse_Defaults(t *testing.T) {
    opts := TestOptionSetStruct{}
    set, err := NewOptionSet(&opts)
    assert.Nil(t, err)
    err = set.Parse([]string{})
    assert.Nil(t, err)
    assert.Equal(t, "foo", opts.Name)
    assert.False(t, opts.Verbose)
}

func TestOptionSetParse_Short(t *testing.T) {
    opts := TestOptionSetStruct{}
    set, err := NewOptionSet(&opts)
    assert.Nil(t, err)
    err = set.Parse([]string{"-v", "-n", "bar"})
    assert.Nil(t, err)
    assert.Equal(t, "bar", opts.Name)
    assert.True(t, opts.Verbose)
}

func TestOptionSetParse_Long(t *testing.T) {
    opts := TestOptionSetStruct{}
    set, err := NewOptionSet(&opts)
    assert.Nil(t, err)
    err = set.Parse([]string{"--verbose", "--name", "bar"})
    assert.Nil(t, err)
    assert.Equal(t, "bar", opts.Name)
    assert.True(t, opts.Verbose)
}

func TestOptionSetParse_NonstandardLong(t *testing.T) {
    opts := TestOptionSetStruct{}
    set, err := NewOptionSet(&opts)
    assert.Nil(t, err)
    err = set.Parse([]string{"-verbose", "-name", "bar"})
    assert.Nil(t, err)
    assert.Equal(t, "bar", opts.Name)
    assert.True(t, opts.Verbose)
}

func TestOptionSetParse_Undeclared(t *testing.T) {
    opts := TestOptionSetStruct{}
    set, err := NewOptionSet(&opts)
    assert.Nil(t, err)
    err = set.Parse([]string{"-ducks"})
    assert.NotNil(t, err)
}

func TestOptionSetParse_LeftoverArgs(t *testing.T) {
    opts := TestOptionSetStruct{}
    set, err := NewOptionSet(&opts)
    assert.Nil(t, err)
    err = set.Parse([]string{"-verbose", "-name", "bar", "far", "zar"})
    assert.Nil(t, err)
    assert.Equal(t, []string{"far", "zar"}, opts.Args)
}

func TestOptionSetParse_LeftoverArgs_InvalidType(t *testing.T) {
    opts := TestInvalidOptionSetStruct{}
    _, err := NewOptionSet(&opts)
    assert.NotNil(t, err)
}

func TestOptionSetWriteHelp(t *testing.T) {
    set, err := NewOptionSet(&TestOptionSetStruct{})
    assert.Nil(t, err)
    buf := bytes.Buffer{}
    set.WriteHelp(&buf)

    expected := "  -n string\n    \tThe name to use (default \"foo\")\n  " +
        "-name string\n    \tThe name to use (default \"foo\")\n  " +
        "-v\tUse verbose logging.\n  -verbose\n    \tUse verbose " +
        "logging.\n"

    assert.Equal(t, expected, buf.String())
}
