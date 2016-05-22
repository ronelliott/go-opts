package opts

import "strings"

type TagSet map[string]string

// Creates a new TagSet from the given raw Tag
func NewTagSet(raw string) TagSet {
	set := TagSet{}

	key := ""
	in_key := true

	value := ""
	in_value := false

	for _, rune := range raw {
		if (rune == '\n' || rune == '\r') && !in_value {
			continue
		}

		if rune == ':' && !in_value {
			in_key = false
			in_value = false
			continue
		}

		if rune == '"' && !in_value {
			in_value = true
			continue
		}

		if rune == '"' && in_value {
			set[strings.Trim(key, " ")] = value
			key = ""
			value = ""
			in_value = false
			in_key = true
			continue
		}

		if in_key {
			key += string(rune)
		} else {
			value += string(rune)
		}
	}

	return set
}

// Gets the value defined by the given key, if it exists
func (this *TagSet) Get(key string) string {
	return (*this)[key]
}

// Returns true if this TagSet has a value for the given key
func (this *TagSet) Has(key string) bool {
	_, ok := (*this)[key]
	return ok
}
