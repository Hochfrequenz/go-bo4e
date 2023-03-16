// Code generated by jsonenums --type MesstechnischeEinordnung; DO NOT EDIT.

package messtechnischeeinordnung

import (
	"encoding/json"
	"fmt"
)

var (
	_MesstechnischeEinordnungNameToValue = map[string]MesstechnischeEinordnung{
		"IMS":           IMS,
		"KME_MME":       KME_MME,
		"KEINE_MESSUNG": KEINE_MESSUNG,
	}

	_MesstechnischeEinordnungValueToName = map[MesstechnischeEinordnung]string{
		IMS:           "IMS",
		KME_MME:       "KME_MME",
		KEINE_MESSUNG: "KEINE_MESSUNG",
	}
)

func init() {
	var v MesstechnischeEinordnung
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MesstechnischeEinordnungNameToValue = map[string]MesstechnischeEinordnung{
			interface{}(IMS).(fmt.Stringer).String():           IMS,
			interface{}(KME_MME).(fmt.Stringer).String():       KME_MME,
			interface{}(KEINE_MESSUNG).(fmt.Stringer).String(): KEINE_MESSUNG,
		}
	}
}

// MarshalJSON is generated so MesstechnischeEinordnung satisfies json.Marshaler.
func (r MesstechnischeEinordnung) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _MesstechnischeEinordnungValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid MesstechnischeEinordnung: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so MesstechnischeEinordnung satisfies json.Unmarshaler.
func (r *MesstechnischeEinordnung) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("MesstechnischeEinordnung should be a string, got %s", data)
	}
	v, ok := _MesstechnischeEinordnungNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid MesstechnischeEinordnung %q", s)
	}
	*r = v
	return nil
}