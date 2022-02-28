package avistyp

// Gibt den Typ des Avis an. (REMADV BGM 1001)
//go:generate stringer --type AvisTyp
//go:generate jsonenums --type AvisTyp
type AvisTyp int

const (
	ABGELEHNTE_FORDERUNG AvisTyp = iota + 1 // Abgelehnte Forderung (239)
	ZAHLUNGSAVIS                            // Zahlungsavis (481)
)
