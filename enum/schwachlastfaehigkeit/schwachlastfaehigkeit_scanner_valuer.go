// Code auto-generated; DO NOT EDIT.
package schwachlastfaehigkeit

import (
	"database/sql/driver"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

// Value returns the string representation of r or an error, if no string representation exists.
// It implements the sql.Valuer interface to be useable by sql drivers when storing enums.
func (r Schwachlastfaehigkeit) Value() (driver.Value, error) {
	s, ok := _SchwachlastfaehigkeitValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan sets r to the enum value represented by src.
// It implements the sql.Scanner interface to be useable by sql drivers when reading from database.
func (r *Schwachlastfaehigkeit) Scan(src interface{}) error {
	f := typemapper.TypeFromValue[Schwachlastfaehigkeit]
	v, err := f(src, _SchwachlastfaehigkeitNameToValue)
	if err != nil {
		return err
	}

	*r = v
	return nil
}
