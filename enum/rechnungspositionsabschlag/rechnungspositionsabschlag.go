package rechnungspositionsabschlag

// Zeitreihentyp sind die Codes der Summenzeitreihentypen
// Note that this enum is not official BO4E standard (yet)!
//
//go:generate stringer --type RechnungspositionsAbschlag
//go:generate jsonenums --type RechnungspositionsAbschlag
type RechnungspositionsAbschlag int

const (
	GEMEINDERABATT           RechnungspositionsAbschlag = iota + 1 // Gemeinderabatt nach Konzessionsabgabenverordnung
	ANPASSUNG_STROM_NEV_19_2                                       // Anpassung nach § 19, Absatz 2 Stromnetzentgeltverordnung
)
