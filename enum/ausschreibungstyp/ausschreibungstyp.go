package ausschreibungstyp

// Ausschreibungstyp is an enum
//
//go:generate stringer --type Ausschreibungstyp
//go:generate jsonenums --type Ausschreibungstyp
type Ausschreibungstyp int

const (
	OEFFENTLICHRECHTLICH Ausschreibungstyp = iota + 1
	EUROPAWEIT
)
