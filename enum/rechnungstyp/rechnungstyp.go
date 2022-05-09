package rechnungstyp

// Rechnungstyp ist eine Abbildung verschiedener Rechnungstypen zur Kennzeichnung von bo.Rechnung en
// Rechnungstyp entspricht dem Rechnungstypen einer INVOIC und weicht damit ebenso wie die BO4E-dotnet Implementierung
// aktuell vom Standard ab
//go:generate stringer --type Rechnungstyp
//go:generate jsonenums --type Rechnungstyp
type Rechnungstyp int

const (
	ABSCHLAGSRECHNUNG Rechnungstyp = iota + 1
	TURNUSRECHNUNG
	MONATSRECHNUNG
	WIMRECHNUNG
	ZWISCHENRECHNUNG
	INTEGRIERTE_13TE_RECHNUNG
	ZUSAETZLICHE_13TE_RECHNUNG
	MEHRMINDERMENGENRECHNUNG
	ABSCHLUSSRECHNUNG
	MSBRECHNUNG
)
