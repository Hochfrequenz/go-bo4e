// Code generated by jsonenums --type NNRechnungstyp; DO NOT EDIT.

package nnrechnungstyp

import (
	"encoding/json"
	"fmt"
)

var (
	_NNRechnungstypNameToValue = map[string]NNRechnungstyp{
		"ABSCHLUSSRECHNUNG":          ABSCHLUSSRECHNUNG,
		"ABSCHLAGSRECHNUNG":          ABSCHLAGSRECHNUNG,
		"TURNUSRECHNUNG":             TURNUSRECHNUNG,
		"MONATSRECHNUNG":             MONATSRECHNUNG,
		"WIMRECHNUNG":                WIMRECHNUNG,
		"ZWISCHENRECHNUNG":           ZWISCHENRECHNUNG,
		"INTEGRIERTE_13TE_RECHNUNG":  INTEGRIERTE_13TE_RECHNUNG,
		"ZUSAETZLICHE_13TE_RECHNUNG": ZUSAETZLICHE_13TE_RECHNUNG,
		"MEHRMINDERMENGENRECHNUNG":   MEHRMINDERMENGENRECHNUNG,
	}

	_NNRechnungstypValueToName = map[NNRechnungstyp]string{
		ABSCHLUSSRECHNUNG:          "ABSCHLUSSRECHNUNG",
		ABSCHLAGSRECHNUNG:          "ABSCHLAGSRECHNUNG",
		TURNUSRECHNUNG:             "TURNUSRECHNUNG",
		MONATSRECHNUNG:             "MONATSRECHNUNG",
		WIMRECHNUNG:                "WIMRECHNUNG",
		ZWISCHENRECHNUNG:           "ZWISCHENRECHNUNG",
		INTEGRIERTE_13TE_RECHNUNG:  "INTEGRIERTE_13TE_RECHNUNG",
		ZUSAETZLICHE_13TE_RECHNUNG: "ZUSAETZLICHE_13TE_RECHNUNG",
		MEHRMINDERMENGENRECHNUNG:   "MEHRMINDERMENGENRECHNUNG",
	}
)

func init() {
	var v NNRechnungstyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_NNRechnungstypNameToValue = map[string]NNRechnungstyp{
			interface{}(ABSCHLUSSRECHNUNG).(fmt.Stringer).String():          ABSCHLUSSRECHNUNG,
			interface{}(ABSCHLAGSRECHNUNG).(fmt.Stringer).String():          ABSCHLAGSRECHNUNG,
			interface{}(TURNUSRECHNUNG).(fmt.Stringer).String():             TURNUSRECHNUNG,
			interface{}(MONATSRECHNUNG).(fmt.Stringer).String():             MONATSRECHNUNG,
			interface{}(WIMRECHNUNG).(fmt.Stringer).String():                WIMRECHNUNG,
			interface{}(ZWISCHENRECHNUNG).(fmt.Stringer).String():           ZWISCHENRECHNUNG,
			interface{}(INTEGRIERTE_13TE_RECHNUNG).(fmt.Stringer).String():  INTEGRIERTE_13TE_RECHNUNG,
			interface{}(ZUSAETZLICHE_13TE_RECHNUNG).(fmt.Stringer).String(): ZUSAETZLICHE_13TE_RECHNUNG,
			interface{}(MEHRMINDERMENGENRECHNUNG).(fmt.Stringer).String():   MEHRMINDERMENGENRECHNUNG,
		}
	}
}

// MarshalJSON is generated so NNRechnungstyp satisfies json.Marshaler.
func (r NNRechnungstyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _NNRechnungstypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid NNRechnungstyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so NNRechnungstyp satisfies json.Unmarshaler.
func (r *NNRechnungstyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NNRechnungstyp should be a string, got %s", data)
	}
	v, ok := _NNRechnungstypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid NNRechnungstyp %q", s)
	}
	*r = v
	return nil
}
