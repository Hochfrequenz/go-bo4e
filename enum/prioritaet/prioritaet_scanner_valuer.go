// Code auto-generated; DO NOT EDIT.
package prioritaet

import (
	"database/sql/driver"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

// Value returns the string representation of r or an error, if no string representation exists.
// It implements the sql.Valuer interface to be useable by sql drivers when storing enums.
func (r Prioritaet) Value() (driver.Value, error) {
	s, ok := _PrioritaetValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan sets r to the enum value represented by src.
// It implements the sql.Scanner interface to be useable by sql drivers when reading from database.
func (r *Prioritaet) Scan(src interface{}) error {
	f := typemapper.TypeFromValue[Prioritaet]
	v, err := f(src, _PrioritaetNameToValue)
	if err != nil {
		return err
	}

	*r = v
	return nil
}
