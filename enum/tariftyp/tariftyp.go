package tariftyp

// Tariftyp is an enum
//
//go:generate stringer --type Tariftyp
//go:generate jsonenums --type Tariftyp
type Tariftyp int

const (
	GRUND_ERSATZVERSORGUNG Tariftyp = iota + 1
	GRUNDVERSORGUNG
	ERSATZVERSORGUNG
	SONDERTARIF
)
