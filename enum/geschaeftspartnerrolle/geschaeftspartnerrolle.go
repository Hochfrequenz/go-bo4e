package geschaeftspartnerrolle

// Geschaeftspartnerrolle are the possible roles of a bo.Geschaeftspartner
//
//go:generate stringer --type Geschaeftspartnerrolle
//go:generate jsonenums --type Geschaeftspartnerrolle
type Geschaeftspartnerrolle int

const (
	// LIEFERANT
	LIEFERANT     Geschaeftspartnerrolle = iota + 1
	DIENSTLEISTER                        // DIENSTLEISTER
	KUNDE                                //	KUNDE
	INTERESSENT                          //	INTERESSENT
	MARKTPARTNER                         //	MARKTPARTNER
)
