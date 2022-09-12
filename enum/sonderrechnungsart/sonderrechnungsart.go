package sonderrechnungsart

// Sonderrechnungsart beschreibt die Art einer enum Rechnungstyp.SONDERRECHNUNG
//
//go:generate stringer --type Sonderrechnungsart
//go:generate jsonenums --type Sonderrechnungsart
type Sonderrechnungsart int

const (
	KONZESSIONSABGABE_TESTAT      Sonderrechnungsart = iota + 1 // KONZESSIONSABGABE_TESTAT is Konzessionsabgabe (Testat)
	INDIVIDUELL_ATYPISCH                                        // Individuelle Vereinbarung für atypische und energieintensive Netznutzung
	INDIVIDUELL_SINGULAER                                       // Individuelle Vereinbarung für singuläre Netznutzung
	KWKG_UMLAGE                                                 // KWKG-Umlage
	OFFSHORE_UMLAGE                                             // Offshore-Netzumlage
	P19_STROM_NEV_UMLAGE                                        // § 19 StromNEV-Umlage
	P18_ABLAV                                                   // §18 AbLaV
	KONZESSIONSABGABE_WECHSEL_RLM                               // Konzessionsabgabe (Wechsel auf Lastgangmessung)
)
