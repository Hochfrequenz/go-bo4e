package gueltigkeitstyp

// Gueltigkeitstyp is an enum
//
//go:generate stringer --type Gueltigkeitstyp
//go:generate jsonenums --type Gueltigkeitstyp
type Gueltigkeitstyp int

const (
	NICHT_IN Gueltigkeitstyp = iota + 1
)
