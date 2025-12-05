package blindarbeitszahler

// Blindarbeitszahler is an enum
//
//go:generate stringer --type Blindarbeitszahler
//go:generate jsonenums --type Blindarbeitszahler
type Blindarbeitszahler int

const (
	ANSCHLUSSNEHMER Blindarbeitszahler = iota + 1
	LIEFERANT
)
