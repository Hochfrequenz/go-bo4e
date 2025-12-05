package betriebszustand

// Betriebszustand is an enum
//
//go:generate stringer --type Betriebszustand
//go:generate jsonenums --type Betriebszustand
type Betriebszustand int

const (
	GESPERRT_NICHT_ENTSPERREN Betriebszustand = iota + 1
	GESPERRT
	REGELBETRIEB
)
