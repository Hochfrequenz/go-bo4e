// Code auto-generated; DO NOT EDIT.
package lokationstyp

import (
	"database/sql/driver"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

// Value returns the string representation of r or an error, if no string representation exists.
// It implements the sql.Valuer interface to be useable by sql drivers when storing enums.
func (r Lokationstyp) Value() (driver.Value, error) {
	s, ok := _LokationstypValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan sets r to the enum value represented by src.
// It implements the sql.Scanner interface to be useable by sql drivers when reading from database.
func (r *Lokationstyp) Scan(src interface{}) error {
	f := typemapper.TypeFromValue[Lokationstyp]
	v, err := f(src, _LokationstypNameToValue)
	if err != nil {
		return err
	}

	*r = v
	return nil
}
