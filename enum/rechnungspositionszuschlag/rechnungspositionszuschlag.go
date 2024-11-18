package rechnungspositionszuschlag

// Zeitreihentyp sind die Codes der Summenzeitreihentypen
// Note that this enum is not official BO4E standard (yet)!
//
//go:generate stringer --type RechnungspositionsZuschlag
//go:generate jsonenums --type RechnungspositionsZuschlag
type RechnungspositionsZuschlag int

const (
	UMSPANNUNGSZUSCHLAG RechnungspositionsZuschlag = iota + 1 //
	ALLEIN_GENUTZTE_BETRIEBSMITTEL_STROM_NEV
	ANPASSUNG_STROM_NEV_19_2
	PAUSCHALE_NETZENTGELTREDUZIERUNG_ENWG_14A
)
