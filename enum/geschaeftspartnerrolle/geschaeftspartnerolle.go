package geschaeftspartnerrolle

//go:generate stringer --type Geschaeftspartnerrolle
//go:generate jsonenums --type Geschaeftspartnerrolle
// Geschaeftspartnerrolle are the possible roles of a Geschaeftspartner
type Geschaeftspartnerrolle int

const (
	Lieferant     Geschaeftspartnerrolle = iota + 1 // Lieferant
	Dienstleister                                   // Dienstleister
	Kunde                                           //	Kunde
	Interessent                                     //	Interessent
	Marktpartner                                    //	Marktpartner
)
