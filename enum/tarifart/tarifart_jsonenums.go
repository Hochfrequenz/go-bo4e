// Code generated by jsonenums --type Tarifart; DO NOT EDIT.

package tarifart

import (
	"encoding/json"
	"fmt"
)

var (
	_TarifartNameToValue = map[string]Tarifart{
		"Eintarif":          Eintarif,
		"Zweitarif":         Zweitarif,
		"Mehrtarif":         Mehrtarif,
		"SmartMeter":        SmartMeter,
		"Leistungsgemessen": Leistungsgemessen,
	}

	_TarifartValueToName = map[Tarifart]string{
		Eintarif:          "Eintarif",
		Zweitarif:         "Zweitarif",
		Mehrtarif:         "Mehrtarif",
		SmartMeter:        "SmartMeter",
		Leistungsgemessen: "Leistungsgemessen",
	}
)

func init() {
	var v Tarifart
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TarifartNameToValue = map[string]Tarifart{
			interface{}(Eintarif).(fmt.Stringer).String():          Eintarif,
			interface{}(Zweitarif).(fmt.Stringer).String():         Zweitarif,
			interface{}(Mehrtarif).(fmt.Stringer).String():         Mehrtarif,
			interface{}(SmartMeter).(fmt.Stringer).String():        SmartMeter,
			interface{}(Leistungsgemessen).(fmt.Stringer).String(): Leistungsgemessen,
		}
	}
}

// MarshalJSON is generated so Tarifart satisfies json.Marshaler.
func (r Tarifart) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _TarifartValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Tarifart: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Tarifart satisfies json.Unmarshaler.
func (r *Tarifart) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Tarifart should be a string, got %s", data)
	}
	v, ok := _TarifartNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Tarifart %q", s)
	}
	*r = v
	return nil
}