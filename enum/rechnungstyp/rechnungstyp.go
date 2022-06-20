package rechnungstyp

import (
	"database/sql/driver"
	"fmt"

	"github.com/hochfrequenz/go-bo4e/internal/dbScanValue"
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
	v, err := dbScanValue.GetValueFromDB[Rechnungstyp](value, _RechnungstypNameToValue)
	if err != nil {
		return err
	}
	*r = v
	return nil
}
