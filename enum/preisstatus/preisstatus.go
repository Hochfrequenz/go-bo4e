package preisstatus

// Preisstatus contains Statusinformation für Preise
//
//go:generate stringer --type Preisstatus
//go:generate jsonenums --type Preisstatus
type Preisstatus int

const (
	// Deprecated: VORLAUEFIG means preliminary (use VORLAEUFIG instead)
	VORLAUEFIG Preisstatus = iota + 1
	// ENDGUELTIG means final
	ENDGUELTIG
	VORLAEUFIG // VORLAEUFIG means preliminary
)
