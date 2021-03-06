// Code generated by jsonenums --type Waehrungseinheit; DO NOT EDIT.

package waehrungseinheit

import (
	"encoding/json"
	"fmt"
)

var (
	_WaehrungseinheitNameToValue = map[string]Waehrungseinheit{
		"EUR": EUR,
		"CT":  CT,
	}

	_WaehrungseinheitValueToName = map[Waehrungseinheit]string{
		EUR: "EUR",
		CT:  "CT",
	}
)

func init() {
	var v Waehrungseinheit
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_WaehrungseinheitNameToValue = map[string]Waehrungseinheit{
			interface{}(EUR).(fmt.Stringer).String(): EUR,
			interface{}(CT).(fmt.Stringer).String():  CT,
		}
	}
}

// MarshalJSON is generated so Waehrungseinheit satisfies json.Marshaler.
func (r Waehrungseinheit) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _WaehrungseinheitValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Waehrungseinheit: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Waehrungseinheit satisfies json.Unmarshaler.
func (r *Waehrungseinheit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Waehrungseinheit should be a string, got %s", data)
	}
	v, ok := _WaehrungseinheitNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Waehrungseinheit %q", s)
	}
	*r = v
	return nil
}
