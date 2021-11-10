// Code generated by jsonenums --type Gebiettyp; DO NOT EDIT.

package gebiettyp

import (
	"encoding/json"
	"fmt"
)

var (
	_GebiettypNameToValue = map[string]Gebiettyp{
		"REGELZONE":              REGELZONE,
		"MARKTGEBIET":            MARKTGEBIET,
		"BILANZIERUNGSGEBIET":    BILANZIERUNGSGEBIET,
		"VERTEILNETZ":            VERTEILNETZ,
		"TRANSPORTNETZ":          TRANSPORTNETZ,
		"REGIONALNETZ":           REGIONALNETZ,
		"AREALNETZ":              AREALNETZ,
		"GRUNDVERSORGUNGSGEBIET": GRUNDVERSORGUNGSGEBIET,
		"VERSORGUNGSGEBIET":      VERSORGUNGSGEBIET,
	}

	_GebiettypValueToName = map[Gebiettyp]string{
		REGELZONE:              "REGELZONE",
		MARKTGEBIET:            "MARKTGEBIET",
		BILANZIERUNGSGEBIET:    "BILANZIERUNGSGEBIET",
		VERTEILNETZ:            "VERTEILNETZ",
		TRANSPORTNETZ:          "TRANSPORTNETZ",
		REGIONALNETZ:           "REGIONALNETZ",
		AREALNETZ:              "AREALNETZ",
		GRUNDVERSORGUNGSGEBIET: "GRUNDVERSORGUNGSGEBIET",
		VERSORGUNGSGEBIET:      "VERSORGUNGSGEBIET",
	}
)

func init() {
	var v Gebiettyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_GebiettypNameToValue = map[string]Gebiettyp{
			interface{}(REGELZONE).(fmt.Stringer).String():              REGELZONE,
			interface{}(MARKTGEBIET).(fmt.Stringer).String():            MARKTGEBIET,
			interface{}(BILANZIERUNGSGEBIET).(fmt.Stringer).String():    BILANZIERUNGSGEBIET,
			interface{}(VERTEILNETZ).(fmt.Stringer).String():            VERTEILNETZ,
			interface{}(TRANSPORTNETZ).(fmt.Stringer).String():          TRANSPORTNETZ,
			interface{}(REGIONALNETZ).(fmt.Stringer).String():           REGIONALNETZ,
			interface{}(AREALNETZ).(fmt.Stringer).String():              AREALNETZ,
			interface{}(GRUNDVERSORGUNGSGEBIET).(fmt.Stringer).String(): GRUNDVERSORGUNGSGEBIET,
			interface{}(VERSORGUNGSGEBIET).(fmt.Stringer).String():      VERSORGUNGSGEBIET,
		}
	}
}

// MarshalJSON is generated so Gebiettyp satisfies json.Marshaler.
func (r Gebiettyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _GebiettypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Gebiettyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Gebiettyp satisfies json.Unmarshaler.
func (r *Gebiettyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Gebiettyp should be a string, got %s", data)
	}
	v, ok := _GebiettypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Gebiettyp %q", s)
	}
	*r = v
	return nil
}
