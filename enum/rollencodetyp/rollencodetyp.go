package rollencodetyp

// Rollencodetyp describes the "vergebende Stelle" / Code Number Registry.
//
//go:generate stringer --type Rollencodetyp
//go:generate jsonenums --type Rollencodetyp
type Rollencodetyp int

const (
	// BDEW is Bundesverband der Energie- u. Wasserwirtschaft (293)
	BDEW Rollencodetyp = iota + 1
	DVGW               // DVGW is Deutscher Verein des Gas- und Wasserfaches (332)
	GLN                // GLN is Global Location Number (9)
)
