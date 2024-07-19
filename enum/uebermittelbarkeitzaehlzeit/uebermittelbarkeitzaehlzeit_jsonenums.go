// Code generated by jsonenums --type UebermittelbarkeitZaehlzeit; DO NOT EDIT.

package uebermittelbarkeitzaehlzeit

import (
	"encoding/json"
	"fmt"
)

var (
	_UebermittelbarkeitZaehlzeitNameToValue = map[string]UebermittelbarkeitZaehlzeit{
		"ELEKTRONISCH":       ELEKTRONISCH,
		"NICHT_ELEKTRONISCH": NICHT_ELEKTRONISCH,
	}

	_UebermittelbarkeitZaehlzeitValueToName = map[UebermittelbarkeitZaehlzeit]string{
		ELEKTRONISCH:       "ELEKTRONISCH",
		NICHT_ELEKTRONISCH: "NICHT_ELEKTRONISCH",
	}
)

func init() {
	var v UebermittelbarkeitZaehlzeit
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_UebermittelbarkeitZaehlzeitNameToValue = map[string]UebermittelbarkeitZaehlzeit{
			interface{}(ELEKTRONISCH).(fmt.Stringer).String():       ELEKTRONISCH,
			interface{}(NICHT_ELEKTRONISCH).(fmt.Stringer).String(): NICHT_ELEKTRONISCH,
		}
	}
}

// MarshalJSON is generated so UebermittelbarkeitZaehlzeit satisfies json.Marshaler.
func (r UebermittelbarkeitZaehlzeit) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _UebermittelbarkeitZaehlzeitValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid UebermittelbarkeitZaehlzeit: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so UebermittelbarkeitZaehlzeit satisfies json.Unmarshaler.
func (r *UebermittelbarkeitZaehlzeit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("UebermittelbarkeitZaehlzeit should be a string, got %s", data)
	}
	v, ok := _UebermittelbarkeitZaehlzeitNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid UebermittelbarkeitZaehlzeit %q", s)
	}
	*r = v
	return nil
}
