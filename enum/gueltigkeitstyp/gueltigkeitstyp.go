package gueltigkeitstyp

// Gueltigkeitstyp Übersicht der verschiedenen Gültigkeiten zur Umsetzung von Positiv- bzw. Negativlisten.
//
//go:generate stringer --type Gueltigkeitstyp
//go:generate jsonenums --type Gueltigkeitstyp
type Gueltigkeitstyp int

const (
	NICHT_IN Gueltigkeitstyp = iota + 1
)
