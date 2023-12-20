// Code generated by jsonenums --type Steuerkanalsleistungsbeschreibung; DO NOT EDIT.

package steuerkanalsleistungsbeschreibung

import (
	"encoding/json"
	"fmt"
)

var (
	_SteuerkanalsleistungsbeschreibungNameToValue = map[string]Steuerkanalsleistungsbeschreibung{
		"AN_AUS":  AN_AUS,
		"GESTUFT": GESTUFT,
	}

	_SteuerkanalsleistungsbeschreibungValueToName = map[Steuerkanalsleistungsbeschreibung]string{
		AN_AUS:  "AN_AUS",
		GESTUFT: "GESTUFT",
	}
)

func init() {
	var v Steuerkanalsleistungsbeschreibung
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SteuerkanalsleistungsbeschreibungNameToValue = map[string]Steuerkanalsleistungsbeschreibung{
			interface{}(AN_AUS).(fmt.Stringer).String():  AN_AUS,
			interface{}(GESTUFT).(fmt.Stringer).String(): GESTUFT,
		}
	}
}

// MarshalJSON is generated so Steuerkanalsleistungsbeschreibung satisfies json.Marshaler.
func (r Steuerkanalsleistungsbeschreibung) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SteuerkanalsleistungsbeschreibungValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Steuerkanalsleistungsbeschreibung: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Steuerkanalsleistungsbeschreibung satisfies json.Unmarshaler.
func (r *Steuerkanalsleistungsbeschreibung) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Steuerkanalsleistungsbeschreibung should be a string, got %s", data)
	}
	v, ok := _SteuerkanalsleistungsbeschreibungNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Steuerkanalsleistungsbeschreibung %q", s)
	}
	*r = v
	return nil
}
