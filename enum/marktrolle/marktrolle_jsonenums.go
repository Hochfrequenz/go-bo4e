// Code generated by jsonenums --type Marktrolle; DO NOT EDIT.

package marktrolle

import (
	"encoding/json"
	"fmt"
)

var (
	_MarktrolleNameToValue = map[string]Marktrolle{
		"NB":              NB,
		"LF":              LF,
		"MSB":             MSB,
		"DL":              DL,
		"BKV":             BKV,
		"BIKO":            BIKO,
		"UENB":            UENB,
		"KUNDE_SELBST_NN": KUNDE_SELBST_NN,
		"MGV":             MGV,
		"EIV":             EIV,
		"RB":              RB,
		"KUNDE":           KUNDE,
		"INTERESSENT":     INTERESSENT,
		"GMSB":            GMSB,
		"AMSB":            AMSB,
	}

	_MarktrolleValueToName = map[Marktrolle]string{
		NB:              "NB",
		LF:              "LF",
		MSB:             "MSB",
		DL:              "DL",
		BKV:             "BKV",
		BIKO:            "BIKO",
		UENB:            "UENB",
		KUNDE_SELBST_NN: "KUNDE_SELBST_NN",
		MGV:             "MGV",
		EIV:             "EIV",
		RB:              "RB",
		KUNDE:           "KUNDE",
		INTERESSENT:     "INTERESSENT",
		GMSB:            "GMSB",
		AMSB:            "AMSB",
	}
)

func init() {
	var v Marktrolle
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MarktrolleNameToValue = map[string]Marktrolle{
			interface{}(NB).(fmt.Stringer).String():              NB,
			interface{}(LF).(fmt.Stringer).String():              LF,
			interface{}(MSB).(fmt.Stringer).String():             MSB,
			interface{}(DL).(fmt.Stringer).String():              DL,
			interface{}(BKV).(fmt.Stringer).String():             BKV,
			interface{}(BIKO).(fmt.Stringer).String():            BIKO,
			interface{}(UENB).(fmt.Stringer).String():            UENB,
			interface{}(KUNDE_SELBST_NN).(fmt.Stringer).String(): KUNDE_SELBST_NN,
			interface{}(MGV).(fmt.Stringer).String():             MGV,
			interface{}(EIV).(fmt.Stringer).String():             EIV,
			interface{}(RB).(fmt.Stringer).String():              RB,
			interface{}(KUNDE).(fmt.Stringer).String():           KUNDE,
			interface{}(INTERESSENT).(fmt.Stringer).String():     INTERESSENT,
			interface{}(GMSB).(fmt.Stringer).String():            GMSB,
			interface{}(AMSB).(fmt.Stringer).String():            AMSB,
		}
	}
}

// MarshalJSON is generated so Marktrolle satisfies json.Marshaler.
func (r Marktrolle) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _MarktrolleValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Marktrolle: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Marktrolle satisfies json.Unmarshaler.
func (r *Marktrolle) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Marktrolle should be a string, got %s", data)
	}
	v, ok := _MarktrolleNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Marktrolle %q", s)
	}
	*r = v
	return nil
}
