package nnrechnungsart

// NNRechnungsart ist die Art der bo.Netznutzungsrechnung (Abbildung verschiedener in der INVOIC angegebenen Rechnungsarten)
//
//go:generate stringer --type NNRechnungsart
//go:generate jsonenums --type NNRechnungsart
type NNRechnungsart int

const (
	// HANDELSRECHNUNG is a commercial invoid / Handelsrechnung
	HANDELSRECHNUNG   NNRechnungsart = iota + 1
	SELBSTAUSGESTELLT                // SELBSTAUSGESTELLT means selbst ausgestellte Rechnung, z.B. für Einspeiserechnungen.
)
