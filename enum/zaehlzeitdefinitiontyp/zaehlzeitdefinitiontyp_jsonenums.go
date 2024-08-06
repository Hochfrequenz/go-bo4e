// Code generated by jsonenums --type ZaehlzeitdefinitionTyp; DO NOT EDIT.

package zaehlzeitdefinitiontyp

import (
	"encoding/json"
	"fmt"
)

var (
	_ZaehlzeitdefinitionTypNameToValue = map[string]ZaehlzeitdefinitionTyp{
		"WAERMEPUMPE":            WAERMEPUMPE,
		"NACHTSPEICHERHEIZUNG":   NACHTSPEICHERHEIZUNG,
		"SCHWACHLASTZEITFENSTER": SCHWACHLASTZEITFENSTER,
		"SONSTIGE":               SONSTIGE,
		"HOCHLASTZEITFENSTER":    HOCHLASTZEITFENSTER,
	}

	_ZaehlzeitdefinitionTypValueToName = map[ZaehlzeitdefinitionTyp]string{
		WAERMEPUMPE:            "WAERMEPUMPE",
		NACHTSPEICHERHEIZUNG:   "NACHTSPEICHERHEIZUNG",
		SCHWACHLASTZEITFENSTER: "SCHWACHLASTZEITFENSTER",
		SONSTIGE:               "SONSTIGE",
		HOCHLASTZEITFENSTER:    "HOCHLASTZEITFENSTER",
	}
)

func init() {
	var v ZaehlzeitdefinitionTyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ZaehlzeitdefinitionTypNameToValue = map[string]ZaehlzeitdefinitionTyp{
			interface{}(WAERMEPUMPE).(fmt.Stringer).String():            WAERMEPUMPE,
			interface{}(NACHTSPEICHERHEIZUNG).(fmt.Stringer).String():   NACHTSPEICHERHEIZUNG,
			interface{}(SCHWACHLASTZEITFENSTER).(fmt.Stringer).String(): SCHWACHLASTZEITFENSTER,
			interface{}(SONSTIGE).(fmt.Stringer).String():               SONSTIGE,
			interface{}(HOCHLASTZEITFENSTER).(fmt.Stringer).String():    HOCHLASTZEITFENSTER,
		}
	}
}

// MarshalJSON is generated so ZaehlzeitdefinitionTyp satisfies json.Marshaler.
func (r ZaehlzeitdefinitionTyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ZaehlzeitdefinitionTypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ZaehlzeitdefinitionTyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ZaehlzeitdefinitionTyp satisfies json.Unmarshaler.
func (r *ZaehlzeitdefinitionTyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ZaehlzeitdefinitionTyp should be a string, got %s", data)
	}
	v, ok := _ZaehlzeitdefinitionTypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ZaehlzeitdefinitionTyp %q", s)
	}
	*r = v
	return nil
}