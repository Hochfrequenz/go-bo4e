package typemapper

import (
	"errors"
	"fmt"
	"reflect"
)

// TypeFromValue returns a type value from v (being a string or string pointer)
// by using map m to determine it's enum value.
func TypeFromValue[E any](v interface{}, m map[string]E) (E, error) {
	var empty E

	if v == nil {
		return empty, errors.New("key is nil")
	}

	key := toString(v)

	if v, ok := m[key]; ok {
		return v, nil
	}

	return empty, fmt.Errorf("unknown key %s", key)
}

func toString(v interface{}) string {
	var out string

	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		if v == nil || v.(*string) == nil {
			return ""
		}
		out = *v.(*string)
	} else {
		out = v.(string)
	}

	return out
}
