package preisstatus

// Preisstatus contains Statusinformation für Preise
//go:generate stringer --type Preisstatus
//go:generate jsonenums --type Preisstatus
type Preisstatus int

const (
	// Vorlaufig means preliminary
	Vorlaufig Preisstatus = iota + 1
	// Endgueltig means final
	Endgueltig
)
