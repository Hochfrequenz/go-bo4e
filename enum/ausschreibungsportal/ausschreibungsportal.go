package ausschreibungsportal

// Ausschreibungsportal is an enum
//
//go:generate stringer --type Ausschreibungsportal
//go:generate jsonenums --type Ausschreibungsportal
type Ausschreibungsportal int

const (
	ENPORTAL Ausschreibungsportal = iota + 1
	ENERGIE_AGENTUR
	BMWI
	ENERGIE_HANDELSPLATZ
	BUND
	VERA_ONLINE
	ISPEX
	ENERGIEMARKTPLATZ
	EVERGABE
	DTAD
)
