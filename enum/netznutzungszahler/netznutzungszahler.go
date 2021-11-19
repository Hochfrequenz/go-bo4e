package netznutzungszahler

// Netznutzungszahler bildet ab, wer für die Netznutzung zahlt
// Note that this is not official BO4E standard (yet).
//go:generate stringer --type Netznutzungszahler
//go:generate jsonenums --type Netznutzungszahler
type Netznutzungszahler int

const (
	// KUNDE heißt, der Endkunde zahlt (EDIFACT Z10)
	KUNDE     Netznutzungszahler = iota + 1
	LIEFERANT                    // LIEFERANT bedeutet, der Lieferant zahlt (EDIFACT Z11)
)
