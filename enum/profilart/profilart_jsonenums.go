// Code generated by jsonenums --type Profilart; DO NOT EDIT.

package profilart

import (
	"encoding/json"
	"fmt"
)

var (
	_ProfilartNameToValue = map[string]Profilart{
		"ART_STANDARDLASTPROFIL":                   ART_STANDARDLASTPROFIL,
		"ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL": ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL,
		"ART_LASTPROFIL":                           ART_LASTPROFIL,
	}

	_ProfilartValueToName = map[Profilart]string{
		ART_STANDARDLASTPROFIL:                   "ART_STANDARDLASTPROFIL",
		ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL: "ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL",
		ART_LASTPROFIL:                           "ART_LASTPROFIL",
	}
)

func init() {
	var v Profilart
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ProfilartNameToValue = map[string]Profilart{
			interface{}(ART_STANDARDLASTPROFIL).(fmt.Stringer).String():                   ART_STANDARDLASTPROFIL,
			interface{}(ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL).(fmt.Stringer).String(): ART_TAGESPARAMETERABHAENGIGES_LASTPROFIL,
			interface{}(ART_LASTPROFIL).(fmt.Stringer).String():                           ART_LASTPROFIL,
		}
	}
}

// MarshalJSON is generated so Profilart satisfies json.Marshaler.
func (r Profilart) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ProfilartValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Profilart: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Profilart satisfies json.Unmarshaler.
func (r *Profilart) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Profilart should be a string, got %s", data)
	}
	v, ok := _ProfilartNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Profilart %q", s)
	}
	*r = v
	return nil
}
