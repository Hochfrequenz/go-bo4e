package bo

import (
	"encoding/json"
	"errors"
	"fmt"
)

// BusinessObjectSlice is a slice that contains 0-n BusinessObject s
type BusinessObjectSlice []BusinessObject

var ErrorUnimplementedBusinessObject = errors.New("unimplemented business object")

func (boSlice *BusinessObjectSlice) UnmarshalJSON(data []byte) error {
	// https://stackoverflow.com/a/69557652/10009545
	var array []json.RawMessage
	if err := json.Unmarshal(data, &array); err != nil {
		return err
	}

	*boSlice = BusinessObjectSlice{}
	errs := make([]error, 0)

	for i := range array {
		base := Geschaeftsobjekt{}
		data := []byte(array[i])

		if err := json.Unmarshal(data, &base); err != nil {
			errs = append(errs, err)
			continue
		}

		elem := NewBusinessObject(base.GetBoTyp())

		if elem == nil {
			errs = append(errs, fmt.Errorf("%w: BusinessObject of type %v at index %v not implemented (or not mapped yet)", ErrorUnimplementedBusinessObject, base.GetBoTyp(), i))
			continue
		}

		if err := json.Unmarshal(data, elem); err != nil {
			errs = append(errs, err)
			continue
		}

		*boSlice = append(*boSlice, elem)
	}

	return errors.Join(errs...)
}
