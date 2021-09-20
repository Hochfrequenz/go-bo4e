package vertragsstatus

//go:generate stringer --type Vertragsstatus
//go:generate jsonenums --type Vertragsstatus
// Vertragsstatus describes different kinds of contracs
type Vertragsstatus int

const (
	InArbeit     Vertragsstatus = iota + 1 // in Arbeit
	Uebermittelt                           // übermittelt
	Angenommen                             // angenommen
	Aktiv                                  // aktiv
	Abgelehnt                              // abgelehnt
	Widerrufen                             // widerrufen
	Storniert                              // storniert
	Gekuendigt                             // gekündigt
	Beendet                                // beendet
)
