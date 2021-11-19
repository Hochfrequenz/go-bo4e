// Code generated by jsonenums --type Profiltyp; DO NOT EDIT.

package profiltyp

import (
	"encoding/json"
	"fmt"
)

var (
	_ProfiltypNameToValue = map[string]Profiltyp{
		"SLP_SEP": SLP_SEP,
		"TLP_TEP": TLP_TEP,
	}

	_ProfiltypValueToName = map[Profiltyp]string{
		SLP_SEP: "SLP_SEP",
		TLP_TEP: "TLP_TEP",
	}
)

func init() {
	var v Profiltyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ProfiltypNameToValue = map[string]Profiltyp{
			interface{}(SLP_SEP).(fmt.Stringer).String(): SLP_SEP,
			interface{}(TLP_TEP).(fmt.Stringer).String(): TLP_TEP,
		}
	}
}

// MarshalJSON is generated so Profiltyp satisfies json.Marshaler.
func (r Profiltyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ProfiltypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Profiltyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Profiltyp satisfies json.Unmarshaler.
func (r *Profiltyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Profiltyp should be a string, got %s", data)
	}
	v, ok := _ProfiltypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Profiltyp %q", s)
	}
	*r = v
	return nil
}