package bo

import (
	"encoding/json"
	"fmt"
)

// BusinessObjectSlice is a slice that contains 0-n BusinessObject s
type BusinessObjectSlice []BusinessObject

func (boSlice *BusinessObjectSlice) UnmarshalJSON(data []byte) error {
	// https://stackoverflow.com/a/69557652/10009545
	var array []json.RawMessage
	if err := json.Unmarshal(data, &array); err != nil {
		return err
	}

	*boSlice = make(BusinessObjectSlice, len(array))
	for i := range array {
		base := Geschaeftsobjekt{}
		data := []byte(array[i])
		if err := json.Unmarshal(data, &base); err != nil {
			return err
		}
		elem := NewBusinessObject(base.GetBoTyp())
		if elem == nil {
			return fmt.Errorf("The BusinessObject with type %v is not implemented (or not mapped yet)", base.GetBoTyp())
		}
		if err := json.Unmarshal(data, elem); err != nil {
			return err
		}
		(*boSlice)[i] = elem
	}
	return nil
}
