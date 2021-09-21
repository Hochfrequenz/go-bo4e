package geschaeftspartnerrolle

// Geschaeftspartnerrolle are the possible roles of a Geschaeftspartner
//go:generate stringer --type Geschaeftspartnerrolle
//go:generate jsonenums --type Geschaeftspartnerrolle
type Geschaeftspartnerrolle int

const (
	Lieferant     Geschaeftspartnerrolle = iota + 1 // Lieferant
	Dienstleister                                   // Dienstleister
	Kunde                                           //	Kunde
	Interessent                                     //	Interessent
	Marktpartner                                    //	Marktpartner
)
