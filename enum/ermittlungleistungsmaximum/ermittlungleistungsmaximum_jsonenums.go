// Code generated by jsonenums --type ErmittlungLeistungsmaximum; DO NOT EDIT.

package ermittlungleistungsmaximum

import (
	"encoding/json"
	"fmt"
)

var (
	_ErmittlungLeistungsmaximumNameToValue = map[string]ErmittlungLeistungsmaximum{
		"VERWENDUNG_HOCHLASTFENSTER":       VERWENDUNG_HOCHLASTFENSTER,
		"KEINE_VERWENDUNG_HOCHLASTFENSTER": KEINE_VERWENDUNG_HOCHLASTFENSTER,
	}

	_ErmittlungLeistungsmaximumValueToName = map[ErmittlungLeistungsmaximum]string{
		VERWENDUNG_HOCHLASTFENSTER:       "VERWENDUNG_HOCHLASTFENSTER",
		KEINE_VERWENDUNG_HOCHLASTFENSTER: "KEINE_VERWENDUNG_HOCHLASTFENSTER",
	}
)

func init() {
	var v ErmittlungLeistungsmaximum
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ErmittlungLeistungsmaximumNameToValue = map[string]ErmittlungLeistungsmaximum{
			interface{}(VERWENDUNG_HOCHLASTFENSTER).(fmt.Stringer).String():       VERWENDUNG_HOCHLASTFENSTER,
			interface{}(KEINE_VERWENDUNG_HOCHLASTFENSTER).(fmt.Stringer).String(): KEINE_VERWENDUNG_HOCHLASTFENSTER,
		}
	}
}

// MarshalJSON is generated so ErmittlungLeistungsmaximum satisfies json.Marshaler.
func (r ErmittlungLeistungsmaximum) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ErmittlungLeistungsmaximumValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ErmittlungLeistungsmaximum: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ErmittlungLeistungsmaximum satisfies json.Unmarshaler.
func (r *ErmittlungLeistungsmaximum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ErmittlungLeistungsmaximum should be a string, got %s", data)
	}
	v, ok := _ErmittlungLeistungsmaximumNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ErmittlungLeistungsmaximum %q", s)
	}
	*r = v
	return nil
}
