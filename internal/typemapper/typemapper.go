package typemapper

import (
	"fmt"
)

// TypeFromValue returns a type value from v (being a string or string pointer)
// by using map m to determine it's enum value.
func TypeFromValue[E any](v interface{}, m map[string]E) (E, error) {
	var empty E

	// if v is nil, an empty value is returned.
	if v == nil {
		return empty, nil
	}

	key := toString(v)

	// if v is a pointer to nil key is an empty string in which case an empty value is returned.
	if key == "" {
		return empty, nil
	}

	if v, ok := m[key]; ok {
		return v, nil
	}

	// If v has a key set but still does not find anything, return an error.
	return empty, fmt.Errorf("unknown key %s", key)
}

func toString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}

	if s, ok := v.(*string); ok && s != nil {
		return *s
	}

	return ""
}
