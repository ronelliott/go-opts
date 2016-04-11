package opts

import (
    "errors"
    "flag"
    "fmt"
    "reflect"
    "strconv"
)

type Option struct {
    // the default value for the option
    Default string

    // the short description of the option
    Description string

    // the help for the option
    Help string

    // the short tag for the field (i.e. "--verbose")
    Long string

    // the name of the field
    Name string

    // the short tag for the field (i.e. "-v")
    Short string

    // the tags for the field
    Tags TagSet

    // the type of the option
    Type string

    // the pointer to the field
    pointer interface{}
}

// Create a option. Parses the field tags, type and name. Stores a pointer to
// the field value.
func NewOption(fieldType reflect.StructField, fieldValue reflect.Value) (*Option, error) {
    if !fieldValue.CanAddr() {
        return nil, errors.New("Cannot address field value: " + fieldType.Name)
    }

    ptrIface := fieldValue.Addr()

    if !ptrIface.CanInterface() {
        return nil, errors.New(
            "Cannot interface field address: " + fieldType.Name)
    }

    tags := NewTagSet(string(fieldType.Tag))
    opt := Option{
        Default: tags["default"],
        Description: tags["description"],
        Help: tags["help"],
        Long: tags["long"],
        Name: fieldType.Name,
        Short: tags["short"],
        Tags: tags,
        Type: fieldType.Type.String(),
        pointer: ptrIface.Interface(),
    }

    if opt.IsPositional() && opt.Type != "[]string" {
        return nil, errors.New(
            "Invalid type for positional args: " + opt.Type)
    }

    return &opt, nil
}

// Adds this option to the flag set, using the defined short/long flags and
// default value
func (this *Option) AddToFlagSet(set *flag.FlagSet) error {
    var err error

    switch this.Type {
    case "bool":
        var def bool

        if this.Default != "" {
            def, err = strconv.ParseBool(this.Default)

            if err != nil {
                return err
            }
        }

        if this.Short != "" {
            set.BoolVar(
                this.pointer.(*bool),
                this.Short,
                def,
                this.Description)
        }

        if this.Long != "" {
            set.BoolVar(
                this.pointer.(*bool),
                this.Long,
                def,
                this.Description)
        }

    case "float64":
        var def float64

        if this.Default != "" {
            def, err = strconv.ParseFloat(this.Default, 64)

            if err != nil {
                return err
            }
        }

        if this.Short != "" {
            set.Float64Var(
                this.pointer.(*float64),
                this.Short,
                def,
                this.Description)
        }

        if this.Long != "" {
            set.Float64Var(
                this.pointer.(*float64),
                this.Long,
                def,
                this.Description)
        }

    case "int":
        var def int

        if this.Default != "" {
            val, err := strconv.ParseInt(this.Default, 10, 32)

            if err != nil {
                return err
            }

            def = int(val)
        }

        if this.Short != "" {
            set.IntVar(
                this.pointer.(*int),
                this.Short,
                int(def),
                this.Description)
        }

        if this.Long != "" {
            set.IntVar(
                this.pointer.(*int),
                this.Long,
                int(def),
                this.Description)
        }

    case "int64":
        var def int64

        if this.Default != "" {
            def, err = strconv.ParseInt(this.Default, 10, 64)

            if err != nil {
                return err
            }
        }

        if this.Short != "" {
            set.Int64Var(
                this.pointer.(*int64),
                this.Short,
                def,
                this.Description)
        }

        if this.Long != "" {
            set.Int64Var(
                this.pointer.(*int64),
                this.Long,
                def,
                this.Description)
        }

    case "string":
        if this.Short != "" {
            set.StringVar(
                this.pointer.(*string),
                this.Short,
                this.Default,
                this.Description)
        }

        if this.Long != "" {
            set.StringVar(
                this.pointer.(*string),
                this.Long,
                this.Default,
                this.Description)
        }

    case "uint":
        var def uint

        if this.Default != "" {
            val, err := strconv.ParseUint(this.Default, 10, 32)

            if err != nil {
                return err
            }

            def = uint(val)
        }

        if this.Short != "" {
            set.UintVar(
                this.pointer.(*uint),
                this.Short,
                uint(def),
                this.Description)
        }

        if this.Long != "" {
            set.UintVar(
                this.pointer.(*uint),
                this.Long,
                uint(def),
                this.Description)
        }

    case "uint64":
        var def uint64

        if this.Default != "" {
            def, err = strconv.ParseUint(this.Default, 10, 64)

            if err != nil {
                return err
            }
        }

        if this.Short != "" {
            set.Uint64Var(
                this.pointer.(*uint64),
                this.Short,
                def,
                this.Description)
        }

        if this.Long != "" {
            set.Uint64Var(
                this.pointer.(*uint64),
                this.Long,
                def,
                this.Description)
        }

    default:
        return errors.New(
            fmt.Sprintf("Type '%s' cannot be handled.", this.Type))
    }

    return nil
}

// Returns true if this Option is for storing positional args
func (this *Option) IsPositional() bool {
    return this.Tags["positional"] == "true"
}
