package netznutzungsabrechnungsvariante

// Netznutzungsabrechnungsvariante beschreibt, wie die Netznutzung abgerechnet wird
// Note that this is not official BO4E standard (yet).
//
//go:generate stringer --type Netznutzungsabrechnungsvariante
//go:generate jsonenums --type Netznutzungsabrechnungsvariante
type Netznutzungsabrechnungsvariante int

const (
	// ARBEITSPREIS_GRUNDPREIS meint Arbeitspreis oder Grundpreis (EDIFACT Z14)
	ARBEITSPREIS_GRUNDPREIS Netznutzungsabrechnungsvariante = iota + 1
	// ARBEITSPREIS_LEISTUNGSPREIS meint Arbeitspreis oder Leistungspreis (EDFIACT Z15)
	ARBEITSPREIS_LEISTUNGSPREIS
)
