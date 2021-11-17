// Code generated by jsonenums --type Zeitreihentyp; DO NOT EDIT.

package zeitreihentyp

import (
	"encoding/json"
	"fmt"
)

var (
	_ZeitreihentypNameToValue = map[string]Zeitreihentyp{
		"EGS": EGS,
		"LGS": LGS,
		"NZR": NZR,
		"SES": SES,
		"SLS": SLS,
		"TES": TES,
		"TLS": TLS,
	}

	_ZeitreihentypValueToName = map[Zeitreihentyp]string{
		EGS: "EGS",
		LGS: "LGS",
		NZR: "NZR",
		SES: "SES",
		SLS: "SLS",
		TES: "TES",
		TLS: "TLS",
	}
)

func init() {
	var v Zeitreihentyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ZeitreihentypNameToValue = map[string]Zeitreihentyp{
			interface{}(EGS).(fmt.Stringer).String(): EGS,
			interface{}(LGS).(fmt.Stringer).String(): LGS,
			interface{}(NZR).(fmt.Stringer).String(): NZR,
			interface{}(SES).(fmt.Stringer).String(): SES,
			interface{}(SLS).(fmt.Stringer).String(): SLS,
			interface{}(TES).(fmt.Stringer).String(): TES,
			interface{}(TLS).(fmt.Stringer).String(): TLS,
		}
	}
}

// MarshalJSON is generated so Zeitreihentyp satisfies json.Marshaler.
func (r Zeitreihentyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ZeitreihentypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Zeitreihentyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Zeitreihentyp satisfies json.Unmarshaler.
func (r *Zeitreihentyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Zeitreihentyp should be a string, got %s", data)
	}
	v, ok := _ZeitreihentypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Zeitreihentyp %q", s)
	}
	*r = v
	return nil
}
