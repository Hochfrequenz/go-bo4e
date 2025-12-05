package angebotsstatus

// Angebotsstatus is an enum
//
//go:generate stringer --type Angebotsstatus
//go:generate jsonenums --type Angebotsstatus
type Angebotsstatus int

const (
	KONZEPTION Angebotsstatus = iota + 1
	UNVERBINDLICH
	VERBINDLICH
	BEAUFTRAGT
	UNGUELTIG
	ABGELEHNT
	NACHGEFASST
	AUSSTEHEND
	ERLEDIGT
)
