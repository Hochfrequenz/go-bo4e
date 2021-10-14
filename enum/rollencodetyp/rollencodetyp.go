package rollencodetyp

// Rollencodetyp describes the "vergebende Stelle" / Code Number Registry.
//go:generate stringer --type Rollencodetyp
//go:generate jsonenums --type Rollencodetyp
type Rollencodetyp int

const (
	BDEW Rollencodetyp = iota + 1 // BDEW is Bundesverband der Energie- u. Wasserwirtschaft (293)
	DVGW                          // DVGW is Deutscher Verein des Gas- und Wasserfaches (332)
	GLN                           // GLN is Global Location Number (9)
)
