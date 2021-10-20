package preisstatus

// Preisstatus contains Statusinformation für Preise
//go:generate stringer --type Preisstatus
//go:generate jsonenums --type Preisstatus
type Preisstatus int

const (
	// VORLAUEFIG means preliminary
	VORLAUEFIG Preisstatus = iota + 1
	// ENDGUELTIG means final
	ENDGUELTIG
)
