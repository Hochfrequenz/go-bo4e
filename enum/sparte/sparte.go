package sparte

import (
	"database/sql/driver"
	"fmt"

	"github.com/hochfrequenz/go-bo4e/internal/dbScanValue"
)

// Sparte contains different divisions of typical utilities.
//go:generate stringer --type Sparte
//go:generate jsonenums --type Sparte
type Sparte int

const (
	// STROM means power/electricity
	STROM      Sparte = iota + 1
	GAS               // GAS means gas
	FERNWAERME        // FERNWAERME means long-distance heat
	NAHWAERME         // NAHWAERME means local heat
	WASSER            // WASSER means water
	ABWASSER          // ABWASSER means waste water
)

// Value sets the value that is written to a database, for that we just use the json representation
func (r Sparte) Value() (driver.Value, error) {
	s, ok := _SparteValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan - Implement sql scanner interface to read the json representation from the DB
func (r *Sparte) Scan(value interface{}) error {
	if s, ok := dbScanValue.GetStringFromDB(value); ok {
		if v, ok := _SparteNameToValue[s]; ok {
			*r = v
			return nil
		}
		return fmt.Errorf("could not read %s", s)
	}
	if i, ok := dbScanValue.GetIntFromDB(value); ok {
		*r = Sparte(i)
		return nil
	}
	if r != nil {
		// We could not read a value, if there is already a value it would not be set to nil
		return fmt.Errorf("could not convert %v to string", value)
	}
	return nil
}
