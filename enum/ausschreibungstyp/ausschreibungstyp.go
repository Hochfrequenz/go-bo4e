package ausschreibungstyp

// Ausschreibungstyp Aufzählung für die Typisierung von Ausschreibungen.
//
//go:generate stringer --type Ausschreibungstyp
//go:generate jsonenums --type Ausschreibungstyp
type Ausschreibungstyp int

const (
	OEFFENTLICHRECHTLICH Ausschreibungstyp = iota + 1
	EUROPAWEIT
)
