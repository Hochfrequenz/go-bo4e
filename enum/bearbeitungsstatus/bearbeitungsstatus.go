package bearbeitungsstatus

// Bearbeitungsstatus is an enum
//
//go:generate stringer --type Bearbeitungsstatus
//go:generate jsonenums --type Bearbeitungsstatus
type Bearbeitungsstatus int

const (
	OFFEN Bearbeitungsstatus = iota + 1
	IN_BEARBEITUNG
	ABGESCHLOSSEN
	STORNIERT
	QUITTIERT
	IGNORIERT
)
