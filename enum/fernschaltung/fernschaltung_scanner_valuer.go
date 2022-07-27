package fernschaltung

import (
	"database/sql/driver"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

// Value returns the string representation of r or an error, if no string representation exists.
// It implements the sql.Valuer interface to be useable by sql drivers when storing enums.
func (r Fernschaltung) Value() (driver.Value, error) {
	s, ok := _FernschaltungValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan sets r to the enum value represented by string
func (r *Fernschaltung) Scan(src interface{}) error {
	f := typemapper.TypeFromValue[Fernschaltung]
	v, err := f(src, _FernschaltungNameToValue)
	if err != nil {
		return err
	}

	*r = v
	return nil
}