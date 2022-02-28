// Code generated by jsonenums -type=AvisTyp; DO NOT EDIT.

package avistyp

import (
	"encoding/json"
	"fmt"
)

var (
	_AvisTypNameToValue = map[string]AvisTyp{
		"ABGELEHNTE_FORDERUNG": ABGELEHNTE_FORDERUNG,
		"ZAHLUNGSAVIS":         ZAHLUNGSAVIS,
	}

	_AvisTypValueToName = map[AvisTyp]string{
		ABGELEHNTE_FORDERUNG: "ABGELEHNTE_FORDERUNG",
		ZAHLUNGSAVIS:         "ZAHLUNGSAVIS",
	}
)

func init() {
	var v AvisTyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_AvisTypNameToValue = map[string]AvisTyp{
			interface{}(ABGELEHNTE_FORDERUNG).(fmt.Stringer).String(): ABGELEHNTE_FORDERUNG,
			interface{}(ZAHLUNGSAVIS).(fmt.Stringer).String():         ZAHLUNGSAVIS,
		}
	}
}

// MarshalJSON is generated so AvisTyp satisfies json.Marshaler.
func (r AvisTyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _AvisTypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid AvisTyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AvisTyp satisfies json.Unmarshaler.
func (r *AvisTyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AvisTyp should be a string, got %s", data)
	}
	v, ok := _AvisTypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid AvisTyp %q", s)
	}
	*r = v
	return nil
}
