package rechnungstyp

// Rechnungstyp ist eine Abbildung verschiedener Rechnungstypen zur Kennzeichnung von Rechnungen
//go:generate stringer --type Rechnungstyp
//go:generate jsonenums --type Rechnungstyp
type Rechnungstyp int

const (
	Endkundenrechnung           Rechnungstyp = iota + 1 // Eine Rechnung vom Lieferanten an einen Endkunden über die Lieferung von Energie
	Netznutzungsrechnung                                // Eine Rechnung vom Netzbetreiber an den Netznutzer. (i.d.R. der Lieferant) über die Netznutzung
	Mehrmindermengenrechnung                            // Eine Rechnung vom Netzbetreiber an den Netznutzer. (i.d.R. der Lieferant) zur Abrechnung von Mengen-Differenzen zwischen Bilanzierung und Messung
	Messstellenbetriebsrechnung                         // Rechnung eines Messstellenbetreibers an den Messkunden
	Beschaffungsrechnung                                // Rechnungen zwischen einem Händler und Einkäufer von Energie
	Ausgleichsenergierichtung                           // Rechnung an den Verursacher von Ausgleichsenergie
)
