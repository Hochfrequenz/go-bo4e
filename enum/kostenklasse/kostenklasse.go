package kostenklasse

// Kostenklasse is an enum
//
//go:generate stringer --type Kostenklasse
//go:generate jsonenums --type Kostenklasse
type Kostenklasse int

const (
	ZERO Kostenklasse = iota + 1
)
