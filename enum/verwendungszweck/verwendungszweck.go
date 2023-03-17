package verwendungszweck

// Verwendungungszweck der Werte Marktlokation
//
//go:generate stringer --type Verwendungszweck
//go:generate jsonenums --type Verwendungszweck

type Verwendungszweck int

const (
	NETZNUTZUNGSABRECHNUNG Verwendungszweck = iota + 1
	BILANZKREISABRECHNUNG
	MEHRMINDERMENGENABRECHNUNG
	ENDKUNDENABRECHNUNG
	UEBERMITTLUNG_AN_DAS_HKNR
	ERMITTLUNG_AUSGEGLICHENHEIT_BILANZKREIS
)
