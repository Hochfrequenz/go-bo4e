// Code generated by jsonenums --type Netzebene; DO NOT EDIT.

package netzebene

import (
	"encoding/json"
	"fmt"
)

var (
	_NetzebeneNameToValue = map[string]Netzebene{
		"NSP":          NSP,
		"MSP":          MSP,
		"HSP":          HSP,
		"HSS":          HSS,
		"MSP_NSP_UMSP": MSP_NSP_UMSP,
		"HSP_MSP_UMSP": HSP_MSP_UMSP,
		"HSS_HSP_UMSP": HSS_HSP_UMSP,
		"HD":           HD,
		"MD":           MD,
		"ND":           ND,
	}

	_NetzebeneValueToName = map[Netzebene]string{
		NSP:          "NSP",
		MSP:          "MSP",
		HSP:          "HSP",
		HSS:          "HSS",
		MSP_NSP_UMSP: "MSP_NSP_UMSP",
		HSP_MSP_UMSP: "HSP_MSP_UMSP",
		HSS_HSP_UMSP: "HSS_HSP_UMSP",
		HD:           "HD",
		MD:           "MD",
		ND:           "ND",
	}
)

func init() {
	var v Netzebene
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_NetzebeneNameToValue = map[string]Netzebene{
			interface{}(NSP).(fmt.Stringer).String():          NSP,
			interface{}(MSP).(fmt.Stringer).String():          MSP,
			interface{}(HSP).(fmt.Stringer).String():          HSP,
			interface{}(HSS).(fmt.Stringer).String():          HSS,
			interface{}(MSP_NSP_UMSP).(fmt.Stringer).String(): MSP_NSP_UMSP,
			interface{}(HSP_MSP_UMSP).(fmt.Stringer).String(): HSP_MSP_UMSP,
			interface{}(HSS_HSP_UMSP).(fmt.Stringer).String(): HSS_HSP_UMSP,
			interface{}(HD).(fmt.Stringer).String():           HD,
			interface{}(MD).(fmt.Stringer).String():           MD,
			interface{}(ND).(fmt.Stringer).String():           ND,
		}
	}
}

// MarshalJSON is generated so Netzebene satisfies json.Marshaler.
func (r Netzebene) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _NetzebeneValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Netzebene: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Netzebene satisfies json.Unmarshaler.
func (r *Netzebene) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Netzebene should be a string, got %s", data)
	}
	v, ok := _NetzebeneNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Netzebene %q", s)
	}
	*r = v
	return nil
}
