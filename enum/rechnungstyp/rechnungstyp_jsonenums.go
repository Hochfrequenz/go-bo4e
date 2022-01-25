// Code generated by jsonenums --type Rechnungstyp; DO NOT EDIT.

package rechnungstyp

import (
	"encoding/json"
	"fmt"
)

var (
	_RechnungstypNameToValue = map[string]Rechnungstyp{
		"ENDKUNDENRECHNUNG":              ENDKUNDENRECHNUNG,
		"NETZNUTZUNGSRECHNUNG":           NETZNUTZUNGSRECHNUNG,
		"MEHRMINDERMENGENRECHNUNG":       MEHRMINDERMENGENRECHNUNG,
		"MESSSTELLENBETRIEBSRECHNUNG":    MESSSTELLENBETRIEBSRECHNUNG,
		"BESCHAFFUNGSRECHNUNG":           BESCHAFFUNGSRECHNUNG,
		"AUSGLEICHSENERGIERECHNUNG":      AUSGLEICHSENERGIERECHNUNG,
		"ABSCHLAGSRECHNUNG":              ABSCHLAGSRECHNUNG,
		"WIMRECHNUNG":                    WIMRECHNUNG,
		"SELBSTAUSGESTELLTERECHNUNGMEMI": SELBSTAUSGESTELLTERECHNUNGMEMI,
		"MSBRECHNUNG":                    MSBRECHNUNG,
		"TURNUSRECHNUNG":                 TURNUSRECHNUNG,
	}

	_RechnungstypValueToName = map[Rechnungstyp]string{
		ENDKUNDENRECHNUNG:              "ENDKUNDENRECHNUNG",
		NETZNUTZUNGSRECHNUNG:           "NETZNUTZUNGSRECHNUNG",
		MEHRMINDERMENGENRECHNUNG:       "MEHRMINDERMENGENRECHNUNG",
		MESSSTELLENBETRIEBSRECHNUNG:    "MESSSTELLENBETRIEBSRECHNUNG",
		BESCHAFFUNGSRECHNUNG:           "BESCHAFFUNGSRECHNUNG",
		AUSGLEICHSENERGIERECHNUNG:      "AUSGLEICHSENERGIERECHNUNG",
		ABSCHLAGSRECHNUNG:              "ABSCHLAGSRECHNUNG",
		WIMRECHNUNG:                    "WIMRECHNUNG",
		SELBSTAUSGESTELLTERECHNUNGMEMI: "SELBSTAUSGESTELLTERECHNUNGMEMI",
		MSBRECHNUNG:                    "MSBRECHNUNG",
		TURNUSRECHNUNG:                 "TURNUSRECHNUNG",
	}
)

func init() {
	var v Rechnungstyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_RechnungstypNameToValue = map[string]Rechnungstyp{
			interface{}(ENDKUNDENRECHNUNG).(fmt.Stringer).String():              ENDKUNDENRECHNUNG,
			interface{}(NETZNUTZUNGSRECHNUNG).(fmt.Stringer).String():           NETZNUTZUNGSRECHNUNG,
			interface{}(MEHRMINDERMENGENRECHNUNG).(fmt.Stringer).String():       MEHRMINDERMENGENRECHNUNG,
			interface{}(MESSSTELLENBETRIEBSRECHNUNG).(fmt.Stringer).String():    MESSSTELLENBETRIEBSRECHNUNG,
			interface{}(BESCHAFFUNGSRECHNUNG).(fmt.Stringer).String():           BESCHAFFUNGSRECHNUNG,
			interface{}(AUSGLEICHSENERGIERECHNUNG).(fmt.Stringer).String():      AUSGLEICHSENERGIERECHNUNG,
			interface{}(ABSCHLAGSRECHNUNG).(fmt.Stringer).String():              ABSCHLAGSRECHNUNG,
			interface{}(WIMRECHNUNG).(fmt.Stringer).String():                    WIMRECHNUNG,
			interface{}(SELBSTAUSGESTELLTERECHNUNGMEMI).(fmt.Stringer).String(): SELBSTAUSGESTELLTERECHNUNGMEMI,
			interface{}(MSBRECHNUNG).(fmt.Stringer).String():                    MSBRECHNUNG,
			interface{}(TURNUSRECHNUNG).(fmt.Stringer).String():                 TURNUSRECHNUNG,
		}
	}
}

// MarshalJSON is generated so Rechnungstyp satisfies json.Marshaler.
func (r Rechnungstyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _RechnungstypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Rechnungstyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Rechnungstyp satisfies json.Unmarshaler.
func (r *Rechnungstyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Rechnungstyp should be a string, got %s", data)
	}
	v, ok := _RechnungstypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Rechnungstyp %q", s)
	}
	*r = v
	return nil
}
