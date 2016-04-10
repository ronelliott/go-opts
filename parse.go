package opts

// A shortcut function for creating an OptionSet from the given struct, then
// parses the arguments from the given args string slice.
func Parse(data interface{}, args []string) error {
    set, err := NewOptionSet(data)

    if err != nil {
        return err
    }

    return set.Parse(args)
}
