package reklamationsgrund

import (
	"database/sql/driver"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/typemapper"
)

// Value returns the string representation of r or an error, if no string representation exists.
// It implements the sql.Valuer interface to be useable by sql drivers when storing enums.
func (r Reklamationsgrund) Value() (driver.Value, error) {
	s, ok := _ReklamationsgrundValueToName[r]
	if ok {
		return s, nil
	}
	return nil, fmt.Errorf("could not stringify %s", r)
}

// Scan sets r to the enum value represented by string
func (r *Reklamationsgrund) Scan(src interface{}) error {
	f := typemapper.TypeFromValue[Reklamationsgrund]
	v, err := f(src, _ReklamationsgrundNameToValue)
	if err != nil {
		return err
	}

	*r = v
	return nil
}