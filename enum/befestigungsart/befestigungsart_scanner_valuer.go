// Code auto-generated; DO NOT EDIT.
package befestigungsart

import (
	"database/sql/driver"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

// Value returns the string representation of r or an error, if no string representation exists.
// It implements the sql.Valuer interface to be useable by sql drivers when storing enums.
func (r Befestigungsart) Value() (driver.Value, error) {
	s, ok := _BefestigungsartValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan sets r to the enum value represented by src.
// It implements the sql.Scanner interface to be useable by sql drivers when reading from database.
func (r *Befestigungsart) Scan(src interface{}) error {
	f := typemapper.TypeFromValue[Befestigungsart]
	v, err := f(src, _BefestigungsartNameToValue)
	if err != nil {
		return err
	}

	*r = v
	return nil
}
