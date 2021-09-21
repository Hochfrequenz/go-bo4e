package wertermittlungsverfahren

// The Wertermittlungsverfahren describes how a value was "created". It is typically used to describe the character of e.g. a Verbrauch.
//go:generate stringer --type Wertermittlungsverfahren
//go:generate jsonenums --type Wertermittlungsverfahren
type Wertermittlungsverfahren int

const (
	PROGNOSE Wertermittlungsverfahren = iota + 1 // Prognose / prognosis
	MESSUNG                                      // Messung / measurement
)
