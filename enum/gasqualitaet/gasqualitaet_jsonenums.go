// Code generated by jsonenums --type Gasqualitaet; DO NOT EDIT.

package gasqualitaet

import (
	"encoding/json"
	"fmt"
)

var (
	_GasqualitaetNameToValue = map[string]Gasqualitaet{
		"H_GAS": H_GAS,
		"L_GAS": L_GAS,
	}

	_GasqualitaetValueToName = map[Gasqualitaet]string{
		H_GAS: "H_GAS",
		L_GAS: "L_GAS",
	}
)

func init() {
	var v Gasqualitaet
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_GasqualitaetNameToValue = map[string]Gasqualitaet{
			interface{}(H_GAS).(fmt.Stringer).String(): H_GAS,
			interface{}(L_GAS).(fmt.Stringer).String(): L_GAS,
		}
	}
}

// MarshalJSON is generated so Gasqualitaet satisfies json.Marshaler.
func (r Gasqualitaet) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _GasqualitaetValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Gasqualitaet: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Gasqualitaet satisfies json.Unmarshaler.
func (r *Gasqualitaet) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Gasqualitaet should be a string, got %s", data)
	}
	v, ok := _GasqualitaetNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Gasqualitaet %q", s)
	}
	*r = v
	return nil
}
