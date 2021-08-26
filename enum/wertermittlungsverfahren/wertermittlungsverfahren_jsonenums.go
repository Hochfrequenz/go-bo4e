// Code generated by jsonenums --type Wertermittlungsverfahren; DO NOT EDIT.

package lokationstyp

import (
	"encoding/json"
	"fmt"
)

var (
	_WertermittlungsverfahrenNameToValue = map[string]Wertermittlungsverfahren{
		"PROGNOSE": PROGNOSE,
		"MESSUNG":  MESSUNG,
	}

	_WertermittlungsverfahrenValueToName = map[Wertermittlungsverfahren]string{
		PROGNOSE: "PROGNOSE",
		MESSUNG:  "MESSUNG",
	}
)

func init() {
	var v Wertermittlungsverfahren
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_WertermittlungsverfahrenNameToValue = map[string]Wertermittlungsverfahren{
			interface{}(PROGNOSE).(fmt.Stringer).String(): PROGNOSE,
			interface{}(MESSUNG).(fmt.Stringer).String():  MESSUNG,
		}
	}
}

// MarshalJSON is generated so Wertermittlungsverfahren satisfies json.Marshaler.
func (r Wertermittlungsverfahren) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _WertermittlungsverfahrenValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Wertermittlungsverfahren: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Wertermittlungsverfahren satisfies json.Unmarshaler.
func (r *Wertermittlungsverfahren) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Wertermittlungsverfahren should be a string, got %s", data)
	}
	v, ok := _WertermittlungsverfahrenNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Wertermittlungsverfahren %q", s)
	}
	*r = v
	return nil
}
