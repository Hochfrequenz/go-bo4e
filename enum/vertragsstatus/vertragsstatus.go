package vertragsstatus

// Vertragsstatus describes different status of contracts
//go:generate stringer --type Vertragsstatus
//go:generate jsonenums --type Vertragsstatus
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
