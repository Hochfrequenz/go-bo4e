package rechnungstyp

import (
	"database/sql/driver"
	"fmt"
)

// Rechnungstyp ist eine Abbildung verschiedener Rechnungstypen zur Kennzeichnung von bo.Rechnung en
// Rechnungstyp entspricht dem Rechnungstypen einer INVOIC und weicht damit ebenso wie die BO4E-dotnet Implementierung
// aktuell vom Standard ab
//go:generate stringer --type Rechnungstyp
//go:generate jsonenums --type Rechnungstyp
type Rechnungstyp int

const (
	ABSCHLAGSRECHNUNG Rechnungstyp = iota + 1
	TURNUSRECHNUNG
	MONATSRECHNUNG
	WIMRECHNUNG
	ZWISCHENRECHNUNG
	INTEGRIERTE_13TE_RECHNUNG
	ZUSAETZLICHE_13TE_RECHNUNG
	MEHRMINDERMENGENRECHNUNG
	ABSCHLUSSRECHNUNG
	MSBRECHNUNG
	KAPAZITAETSRECHNUNG
)

// Value sets the value that is written to a database, for that we just use the json representation
func (r Rechnungstyp) Value() (driver.Value, error) {
	s, ok := _RechnungstypValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan - Implement sql scanner interface to read the json representation from the DB
func (r *Rechnungstyp) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil || value.(*string) == nil {
		// set the value of the pointer to 0 as default
		*r = 0
		return nil
	}
	s := *value.(*string)
	if v, ok := _RechnungstypNameToValue[s]; ok {
		*r = v
		return nil
	}
	if v, ok := _RechnungstypEdiToValue[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("could not read %s", s)
}
