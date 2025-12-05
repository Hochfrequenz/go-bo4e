package vertragstatus

// Vertragstatus is an enum
//
//go:generate stringer --type Vertragstatus
//go:generate jsonenums --type Vertragstatus
type Vertragstatus int

const (
	IN_ARBEIT Vertragstatus = iota + 1
	UEBERMITTELT
	ANGENOMMEN
	AKTIV
	ABGELEHNT
	WIDERRUFEN
	STORNIERT
	GEKUENDIGT
	BEENDET
)
