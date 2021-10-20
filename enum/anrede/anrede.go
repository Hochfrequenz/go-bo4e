package anrede

// Anrede ist eine Auflistung möglicher abzurechnender Dienstleistungen.
//go:generate stringer --type Anrede
//go:generate jsonenums --type Anrede
type Anrede int

const (
	// HERR (male)
	HERR             Anrede = iota + 1
	FRAU                    // FRAU (female)
	DIVERS                  // DIVERS (non-binary)
	EHELEUTE                // EHELEUTE (married couple)
	FIRMA                   // FIRMA (company)
	WOHNGEMEINSCHAFT        // WOHNGEMEINSCHAFT is a share flat/WG
	INDIVIDUELL             // INDIVIDUELL festgelegt
)
