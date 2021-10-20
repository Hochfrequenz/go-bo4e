package anrede

// Anrede ist eine Auflistung möglicher abzurechnender Dienstleistungen.
//go:generate stringer --type Anrede
//go:generate jsonenums --type Anrede
type Anrede int

const (
	HERR             Anrede = iota + 1 // HERR (male)
	FRAU                               // FRAU (female)
	DIVERS                             // DIVERS (non-binary)
	EHELEUTE                           // EHELEUTE (married couple)
	FIRMA                              // FIRMA (company)
	WOHNGEMEINSCHAFT                   // WOHNGEMEINSCHAFT is a share flat/WG
	INDIVIDUELL                        // INDIVIDUELL festgelegt
)
