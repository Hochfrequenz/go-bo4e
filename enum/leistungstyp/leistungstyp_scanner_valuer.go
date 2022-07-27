package leistungstyp

import (
	"database/sql/driver"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

// Value returns the string representation of r or an error, if no string representation exists.
// It implements the sql.Valuer interface to be useable by sql drivers when storing enums.
func (r Leistungstyp) Value() (driver.Value, error) {
	s, ok := _LeistungstypValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan sets r to the enum value represented by string
func (r *Leistungstyp) Scan(src interface{}) error {
	f := typemapper.TypeFromValue[Leistungstyp]
	v, err := f(src, _LeistungstypNameToValue)
	if err != nil {
		return err
	}

	*r = v
	return nil
}