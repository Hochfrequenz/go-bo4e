package vertragsstatus

// Vertragsstatus describes different status of contracts
//go:generate stringer --type Vertragsstatus
//go:generate jsonenums --type Vertragsstatus
type Vertragsstatus int

const (
	// INARBEIT means work in progress
	INARBEIT     Vertragsstatus = iota + 1
	UEBERMITTELT                // UEBERMITTELT means sent
	ANGENOMMEN                  // ANGENOMMEN means accepted
	AKTIV                       // AKTIV means aktive
	ABGELEHNT                   // ABGELEHNT means rejected
	WIDERRUFEN                  // WIDERRUFEN means revoked
	STORNIERT                   // STORNIERT means cancelled
	GEKUENDIGT                  // GEKUENDIGT means resigned
	BEENDET                     // BEENDET means stopped
)
