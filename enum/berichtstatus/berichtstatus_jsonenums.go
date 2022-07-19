// Code generated by jsonenums --type BerichtStatus; DO NOT EDIT.

package berichtstatus

import (
	"encoding/json"
	"fmt"
)

var (
	_BerichtStatusNameToValue = map[string]BerichtStatus{
		"ERFOLGREICH": ERFOLGREICH,
		"FEHLER":      FEHLER,
	}

	_BerichtStatusValueToName = map[BerichtStatus]string{
		ERFOLGREICH: "ERFOLGREICH",
		FEHLER:      "FEHLER",
	}
)

func init() {
	var v BerichtStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BerichtStatusNameToValue = map[string]BerichtStatus{
			interface{}(ERFOLGREICH).(fmt.Stringer).String(): ERFOLGREICH,
			interface{}(FEHLER).(fmt.Stringer).String():      FEHLER,
		}
	}
}

// MarshalJSON is generated so BerichtStatus satisfies json.Marshaler.
func (r BerichtStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BerichtStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BerichtStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BerichtStatus satisfies json.Unmarshaler.
func (r *BerichtStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BerichtStatus should be a string, got %s", data)
	}
	v, ok := _BerichtStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BerichtStatus %q", s)
	}
	*r = v
	return nil
}
