// Code generated by jsonenums --type Messwerterfassung; DO NOT EDIT.

package messwerterfassung

import (
	"encoding/json"
	"fmt"
)

var (
	_MesswerterfassungNameToValue = map[string]Messwerterfassung{
		"FERN_AUSLESBAR":      FERN_AUSLESBAR,
		"MANUELL_AUSGELESENE": MANUELL_AUSGELESENE,
	}

	_MesswerterfassungValueToName = map[Messwerterfassung]string{
		FERN_AUSLESBAR:      "FERN_AUSLESBAR",
		MANUELL_AUSGELESENE: "MANUELL_AUSGELESENE",
	}
)

func init() {
	var v Messwerterfassung
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MesswerterfassungNameToValue = map[string]Messwerterfassung{
			interface{}(FERN_AUSLESBAR).(fmt.Stringer).String():      FERN_AUSLESBAR,
			interface{}(MANUELL_AUSGELESENE).(fmt.Stringer).String(): MANUELL_AUSGELESENE,
		}
	}
}

// MarshalJSON is generated so Messwerterfassung satisfies json.Marshaler.
func (r Messwerterfassung) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _MesswerterfassungValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Messwerterfassung: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Messwerterfassung satisfies json.Unmarshaler.
func (r *Messwerterfassung) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Messwerterfassung should be a string, got %s", data)
	}
	v, ok := _MesswerterfassungNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Messwerterfassung %q", s)
	}
	*r = v
	return nil
}
