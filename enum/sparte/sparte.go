package sparte

import (
	"database/sql/driver"
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
	v, err := r.MarshalJSON()
	s := string(v)
	s = s[1:len(s)-1]
	return s,err
}

// Scan - Implement sql scanner interface to read the json representation from the DB
func (r *Sparte) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		// set the value of the pointer to 0 as default
		*r = 0
		return nil
	}
	s := *value.(*string)
	s = "\"" + s + "\""
	return r.UnmarshalJSON([]byte(s))
}