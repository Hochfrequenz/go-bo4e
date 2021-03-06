// Code generated by jsonenums --type ArithmetischeOperation; DO NOT EDIT.

package arithmetischeoperation

import (
	"encoding/json"
	"fmt"
)

var (
	_ArithmetischeOperationNameToValue = map[string]ArithmetischeOperation{
		"ADDITION":       ADDITION,
		"SUBTRAKTION":    SUBTRAKTION,
		"MULTIPLIKATION": MULTIPLIKATION,
		"DIVISION":       DIVISION,
	}

	_ArithmetischeOperationValueToName = map[ArithmetischeOperation]string{
		ADDITION:       "ADDITION",
		SUBTRAKTION:    "SUBTRAKTION",
		MULTIPLIKATION: "MULTIPLIKATION",
		DIVISION:       "DIVISION",
	}
)

func init() {
	var v ArithmetischeOperation
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ArithmetischeOperationNameToValue = map[string]ArithmetischeOperation{
			interface{}(ADDITION).(fmt.Stringer).String():       ADDITION,
			interface{}(SUBTRAKTION).(fmt.Stringer).String():    SUBTRAKTION,
			interface{}(MULTIPLIKATION).(fmt.Stringer).String(): MULTIPLIKATION,
			interface{}(DIVISION).(fmt.Stringer).String():       DIVISION,
		}
	}
}

// MarshalJSON is generated so ArithmetischeOperation satisfies json.Marshaler.
func (r ArithmetischeOperation) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ArithmetischeOperationValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ArithmetischeOperation: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ArithmetischeOperation satisfies json.Unmarshaler.
func (r *ArithmetischeOperation) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ArithmetischeOperation should be a string, got %s", data)
	}
	v, ok := _ArithmetischeOperationNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ArithmetischeOperation %q", s)
	}
	*r = v
	return nil
}
