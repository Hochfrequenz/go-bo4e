// This maps the EDIFACT Values to the Rechnungstyp and vice versa
// nolint:dupl // linter thinks the mapping is a duplicate to jsonenums
package rechnungstyp

import (
	"fmt"
)

// EdiValue outputs the EDIFACT representation of the enum (AJT 4465)
func (r Rechnungstyp) EdiValue() (string, error) {
	s, ok := _RechnungstypValueToEdi[r]
	if ok {
		return s, nil
	}
	return "", fmt.Errorf("could not map %s", r)
}

// FromEdi maps the EDIFACT code to the enum
func (r *Rechnungstyp) FromEdi(value string) error {
	if v, ok := _RechnungstypEdiToValue[value]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("could not read %s", value)
}

var (
	_RechnungstypEdiToValue = map[string]Rechnungstyp{
		"ABR": ABSCHLUSSRECHNUNG,
		"ABS": ABSCHLAGSRECHNUNG,
		"JVR": TURNUSRECHNUNG,
		"MVR": MONATSRECHNUNG,
		"WIM": WIMRECHNUNG,
		"ZVR": ZWISCHENRECHNUNG,
		"13I": INTEGRIERTE_13TE_RECHNUNG,
		"13R": ZUSAETZLICHE_13TE_RECHNUNG,
		"MMM": MEHRMINDERMENGENRECHNUNG,
		"MSB": MSBRECHNUNG,
		"NAP": KAPAZITAETSRECHNUNG,
	}

	_RechnungstypValueToEdi = map[Rechnungstyp]string{
		ABSCHLUSSRECHNUNG:          "ABR",
		ABSCHLAGSRECHNUNG:          "ABS",
		TURNUSRECHNUNG:             "JVR",
		MONATSRECHNUNG:             "MVR",
		WIMRECHNUNG:                "WIM",
		ZWISCHENRECHNUNG:           "ZVR",
		INTEGRIERTE_13TE_RECHNUNG:  "13I",
		ZUSAETZLICHE_13TE_RECHNUNG: "13R",
		MEHRMINDERMENGENRECHNUNG:   "MMM",
		MSBRECHNUNG:                "MSB",
		KAPAZITAETSRECHNUNG:        "NAP",
	}
)
