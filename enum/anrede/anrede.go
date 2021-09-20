package anrede

//go:generate stringer --type Anrede
//go:generate jsonenums --type Anrede
// Anrede ist eine Auflistung möglicher abzurechnender Dienstleistungen.
type Anrede int

const (
	Herr             Anrede = iota + 1 // Herr
	Frau                               // Frau
	Divers                             // divers
	Eheleute                           // Eheleute
	Firma                              // Firma
	Wohngemeinschaft                   // WG
	Individuell                        // Individuell festgelegt
)
