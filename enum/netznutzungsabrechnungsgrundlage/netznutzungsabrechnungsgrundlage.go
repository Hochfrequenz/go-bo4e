package netznutzungsabrechnungsgrundlage

// Netznutzungsabrechnungsgrundlage beschreibt, auf welcher Grundlage die Netznutzungsabrechnung basiert
// Note that this is not official BO4E standard (yet).
//go:generate stringer --type Netznutzungsabrechnungsgrundlage
//go:generate jsonenums --type Netznutzungsabrechnungsgrundlage
type Netznutzungsabrechnungsgrundlage int

const (
	// LIEFERSCHEIN heißt Netznutzungsabrechnung auf Grundlage Lieferschein (Z12)
	LIEFERSCHEIN Netznutzungsabrechnungsgrundlage = iota + 1
	// ABWEICHENDE_GRUNDLAGE heißt, dass die Abrechnung auf einer anderen, vertraglich mit Anschlussnutzer vereinbarten Grundlage geschieht (Z13)
	ABWEICHENDE_GRUNDLAGE
)
