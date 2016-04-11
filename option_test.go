package opts

import (
    "flag"
    "github.com/stretchr/testify/assert"
    "os"
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

    Name string `short:"n" env:"FLABBERGASTED"`

    Duck string `short:"d" env:"FLABBERGASTED" default:"quack"`

    Bool bool `long:"bool"`

    Float64 float64 `long:"float64"`

    Int int `long:"int"`

    Int64 int64 `long:"int64"`

    String string `long:"string"`

    Uint uint `long:"uint"`

    Uint64 uint64 `long:"uint64"`

    Db string `
        long:"db"
        default:"mongodb://localhost:27017/db"
        description:"The db resource to connect to."`
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

func TestNewOption_Default_ResourceUrls(t *testing.T) {
    opt, err := NewOption(
        optionTestGetFieldType(13),
        optionTestGetFieldValue(13))
    assert.Nil(t, err)
    assert.Equal(t, "mongodb://localhost:27017/db", opt.Default)
}

func TestNewOption_Default_CurrentValue_Bool(t *testing.T) {
    opts := TestNewOptionStruct{
        Bool: true,
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(6)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(6)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "true", opt.Default)
}

func TestNewOption_Default_CurrentValue_Float64(t *testing.T) {
    opts := TestNewOptionStruct{
        Float64: 3.14,
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(7)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(7)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "3.14", opt.Default)
}

func TestNewOption_Default_CurrentValue_Int(t *testing.T) {
    opts := TestNewOptionStruct{
        Int: 192,
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(8)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(8)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "192", opt.Default)
}

func TestNewOption_Default_CurrentValue_Int64(t *testing.T) {
    opts := TestNewOptionStruct{
        Int64: 192,
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(9)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(9)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "192", opt.Default)
}

func TestNewOption_Default_CurrentValue_String(t *testing.T) {
    opts := TestNewOptionStruct{
        String: "foo",
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(10)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(10)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "foo", opt.Default)
}

func TestNewOption_Default_CurrentValue_Uint(t *testing.T) {
    opts := TestNewOptionStruct{
        Uint: 192,
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(11)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(11)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "192", opt.Default)
}

func TestNewOption_Default_CurrentValue_Uint64(t *testing.T) {
    opts := TestNewOptionStruct{
        Uint64: 192,
    }

    fieldType := reflect.TypeOf(&opts).Elem().Field(12)
    fieldValue := reflect.ValueOf(&opts).Elem().Field(12)

    opt, err := NewOption(fieldType, fieldValue)
    assert.Nil(t, err)
    assert.Equal(t, "192", opt.Default)
}

func TestNewOption_Default_Env(t *testing.T) {
    os.Setenv("FLABBERGASTED", "happy happy joy joy")
    opt, err := NewOption(optionTestGetFieldType(4), optionTestGetFieldValue(4))
    assert.Nil(t, err)
    assert.Equal(t, "happy happy joy joy", opt.Default)
    os.Unsetenv("FLABBERGASTED")
}

func TestNewOption_Default_EnvOverride(t *testing.T) {
    os.Setenv("FLABBERGASTED", "happy happy joy joy ....")
    opt, err := NewOption(optionTestGetFieldType(5), optionTestGetFieldValue(5))
    assert.Nil(t, err)
    assert.Equal(t, "happy happy joy joy ....", opt.Default)
    os.Unsetenv("FLABBERGASTED")
}

func TestNewOption_Default_EnvOverride_NoEnvSet(t *testing.T) {
    opt, err := NewOption(optionTestGetFieldType(5), optionTestGetFieldValue(5))
    assert.Nil(t, err)
    assert.Equal(t, "quack", opt.Default)
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

func TestNewOption_Description_ResourceUrls(t *testing.T) {
    opt, err := NewOption(
        optionTestGetFieldType(13),
        optionTestGetFieldValue(13))
    assert.Nil(t, err)
    assert.Equal(t, "The db resource to connect to.", opt.Description)
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
