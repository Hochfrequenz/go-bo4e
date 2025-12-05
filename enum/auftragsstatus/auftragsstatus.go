package auftragsstatus

// Auftragsstatus is an enum
//
//go:generate stringer --type Auftragsstatus
//go:generate jsonenums --type Auftragsstatus
type Auftragsstatus int

const (
	GESCHEITERT Auftragsstatus = iota + 1
	ERFOLGREICH
	GEPLANT
	ZUGESTIMMT
	WIDERSPROCHEN
	ABGELEHNT
)
