// Code generated by jsonenums --type Erzeugungsart; DO NOT EDIT.

package erzeugungsart

import (
	"encoding/json"
	"fmt"
)

var (
	_ErzeugungsartNameToValue = map[string]Erzeugungsart{
		"KWK":          KWK,
		"WIND":         WIND,
		"SOLAR":        SOLAR,
		"KERNKRAFT":    KERNKRAFT,
		"WASSER":       WASSER,
		"GEOTHERMIE":   GEOTHERMIE,
		"BIOMASSE":     BIOMASSE,
		"KOHLE":        KOHLE,
		"GAS":          GAS,
		"SONSTIGE":     SONSTIGE,
		"SONSTIGE_EEG": SONSTIGE_EEG,
	}

	_ErzeugungsartValueToName = map[Erzeugungsart]string{
		KWK:          "KWK",
		WIND:         "WIND",
		SOLAR:        "SOLAR",
		KERNKRAFT:    "KERNKRAFT",
		WASSER:       "WASSER",
		GEOTHERMIE:   "GEOTHERMIE",
		BIOMASSE:     "BIOMASSE",
		KOHLE:        "KOHLE",
		GAS:          "GAS",
		SONSTIGE:     "SONSTIGE",
		SONSTIGE_EEG: "SONSTIGE_EEG",
	}
)

func init() {
	var v Erzeugungsart
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ErzeugungsartNameToValue = map[string]Erzeugungsart{
			interface{}(KWK).(fmt.Stringer).String():          KWK,
			interface{}(WIND).(fmt.Stringer).String():         WIND,
			interface{}(SOLAR).(fmt.Stringer).String():        SOLAR,
			interface{}(KERNKRAFT).(fmt.Stringer).String():    KERNKRAFT,
			interface{}(WASSER).(fmt.Stringer).String():       WASSER,
			interface{}(GEOTHERMIE).(fmt.Stringer).String():   GEOTHERMIE,
			interface{}(BIOMASSE).(fmt.Stringer).String():     BIOMASSE,
			interface{}(KOHLE).(fmt.Stringer).String():        KOHLE,
			interface{}(GAS).(fmt.Stringer).String():          GAS,
			interface{}(SONSTIGE).(fmt.Stringer).String():     SONSTIGE,
			interface{}(SONSTIGE_EEG).(fmt.Stringer).String(): SONSTIGE_EEG,
		}
	}
}

// MarshalJSON is generated so Erzeugungsart satisfies json.Marshaler.
func (r Erzeugungsart) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ErzeugungsartValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Erzeugungsart: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Erzeugungsart satisfies json.Unmarshaler.
func (r *Erzeugungsart) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Erzeugungsart should be a string, got %s", data)
	}
	v, ok := _ErzeugungsartNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Erzeugungsart %q", s)
	}
	*r = v
	return nil
}
