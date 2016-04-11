package opts

import (
    "flag"
    "github.com/stretchr/testify/assert"
    "reflect"
    "testing"
)

type TestNewOptionStruct struct {
    Verbose bool `
        default:"true"
        description:"Use verbose logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`

    Args []string `positional:"true"`

    InvalidArgs string `positional:"true"`

    nonExported string `short:"f"`

    Count int `
        description:"Use verbose logging."
        help:"Be very talkative when logging"
        long:"verbose"
        short:"v"`
}

func optionTestGetFieldType(num int) reflect.StructField {
    return reflect.TypeOf(&TestNewOptionStruct{}).Elem().Field(num)
}

func optionTestGetFieldValue(num int) reflect.Value {
    return reflect.ValueOf(&TestNewOptionStruct{}).Elem().Field(num)
}

func optionTestNewFlagSet() *flag.FlagSet {
    return flag.NewFlagSet("Tests", flag.ContinueOnError)
}

func TestNewOption_Default(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.Equal(t, "true", opt.Default)
}

func TestNewOption_Default_CurrentValue(t *testing.T) {
    opts := TestNewOptionStruct{
        Count: 192,
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(4)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(4)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "192", opt.Default)
}

func TestNewOption_NonAddressable(t *testing.T) {
    structType := reflect.TypeOf(TestNewOptionStruct{}).Field(3)
    structValue := reflect.ValueOf(TestNewOptionStruct{}).Field(3)
    _, err := NewOption(structType, structValue)
    assert.NotNil(t, err)
    assert.Equal(t, "Cannot address field value: nonExported", err.Error())
}

func TestNewOption_NonInterfaceable(t *testing.T) {
    _, err := NewOption(optionTestGetFieldType(3), optionTestGetFieldValue(3))
    assert.NotNil(t, err)
    assert.Equal(t, "Cannot interface field address: nonExported", err.Error())
}

func TestNewOption_Description(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.Equal(t, "Use verbose logging.", opt.Description)
}

func TestNewOption_Help(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.Equal(t, "Be very talkative when logging", opt.Help)
}

func TestNewOption_Long(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.Equal(t, "verbose", opt.Long)
}

func TestNewOption_Name(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.Equal(t, "Verbose", opt.Name)
}

func TestNewOption_Short(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.Equal(t, "v", opt.Short)
}

func TestNewOption_Type(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.Equal(t, "bool", opt.Type)
}

func TestNewOption_Positional_Default(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.Equal(t, "", opt.Default)
}

func TestNewOption_Positional_Description(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.Equal(t, "", opt.Description)
}

func TestNewOption_Positional_Help(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.Equal(t, "", opt.Help)
}

func TestNewOption_Positional_Long(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.Equal(t, "", opt.Long)
}

func TestNewOption_Positional_Name(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.Equal(t, "Args", opt.Name)
}

func TestNewOption_Positional_Short(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.Equal(t, "", opt.Short)
}

func TestNewOption_Positional_Type(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.Equal(t, "[]string", opt.Type)
}

func TestNewOption_Positional_Type_Invalid(t *testing.T) {
    _, err := NewOption(optionTestGetFieldType(2), optionTestGetFieldValue(2))
    assert.NotNil(t, err)
}

func TestAddToFlagSet_Invalid(t *testing.T) {
    var value bool
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "[]bool",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.NotNil(t, err)
    assert.Nil(t, set.Lookup("v"))
    assert.Nil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Bool(t *testing.T) {
    var value bool
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "bool",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Bool_Invalid(t *testing.T) {
    var value bool
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "Nein",
        Long: "verbose",
        Short: "v",
        Type: "bool",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.NotNil(t, err)
    assert.Nil(t, set.Lookup("v"))
    assert.Nil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Bool_NoDefault(t *testing.T) {
    var value bool
    set := optionTestNewFlagSet()
    opt := Option{
        Long: "verbose",
        Short: "v",
        Type: "bool",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Float64(t *testing.T) {
    var value float64
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "3.14",
        Long: "verbose",
        Short: "v",
        Type: "float64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Float64_Invalid(t *testing.T) {
    var value float64
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "float64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.NotNil(t, err)
    assert.Nil(t, set.Lookup("v"))
    assert.Nil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Float64_NoDefault(t *testing.T) {
    var value float64
    set := optionTestNewFlagSet()
    opt := Option{
        Long: "verbose",
        Short: "v",
        Type: "float64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Int(t *testing.T) {
    var value int
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "-10",
        Long: "verbose",
        Short: "v",
        Type: "int",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Int_Invalid(t *testing.T) {
    var value int
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "int",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.NotNil(t, err)
    assert.Nil(t, set.Lookup("v"))
    assert.Nil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Int_NoDefault(t *testing.T) {
    var value int
    set := optionTestNewFlagSet()
    opt := Option{
        Long: "verbose",
        Short: "v",
        Type: "int",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Int64(t *testing.T) {
    var value int64
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "-10",
        Long: "verbose",
        Short: "v",
        Type: "int64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Int64_Invalid(t *testing.T) {
    var value int64
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "int64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.NotNil(t, err)
    assert.Nil(t, set.Lookup("v"))
    assert.Nil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Int64_NoDefault(t *testing.T) {
    var value int64
    set := optionTestNewFlagSet()
    opt := Option{
        Long: "verbose",
        Short: "v",
        Type: "int64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_String(t *testing.T) {
    var value string
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "string",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_String_NoDefault(t *testing.T) {
    var value string
    set := optionTestNewFlagSet()
    opt := Option{
        Long: "verbose",
        Short: "v",
        Type: "string",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Uint(t *testing.T) {
    var value uint
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "10",
        Long: "verbose",
        Short: "v",
        Type: "uint",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Uint_Invalid(t *testing.T) {
    var value uint
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "uint",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.NotNil(t, err)
    assert.Nil(t, set.Lookup("v"))
    assert.Nil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Uint_NoDefault(t *testing.T) {
    var value uint
    set := optionTestNewFlagSet()
    opt := Option{
        Long: "verbose",
        Short: "v",
        Type: "uint",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Uint64(t *testing.T) {
    var value uint64
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "10",
        Long: "verbose",
        Short: "v",
        Type: "uint64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Uint64_Invalid(t *testing.T) {
    var value uint64
    set := optionTestNewFlagSet()
    opt := Option{
        Default: "true",
        Long: "verbose",
        Short: "v",
        Type: "uint64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.NotNil(t, err)
    assert.Nil(t, set.Lookup("v"))
    assert.Nil(t, set.Lookup("verbose"))
}

func TestAddToFlagSet_Uint64_NoDefault(t *testing.T) {
    var value uint64
    set := optionTestNewFlagSet()
    opt := Option{
        Long: "verbose",
        Short: "v",
        Type: "uint64",
        pointer: &value,
    }

    err := opt.AddToFlagSet(set)
    assert.Nil(t, err)
    assert.NotNil(t, set.Lookup("v"))
    assert.NotNil(t, set.Lookup("verbose"))
}

func TestIsPositional_Is(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(1), optionTestGetFieldValue(1))
    assert.Nil(t, err)
    assert.True(t, opt.IsPositional())
}

func TestIsPositional_IsNot(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(0), optionTestGetFieldValue(0))
    assert.Nil(t, err)
    assert.False(t, opt.IsPositional())
}
