package rechnungstyp

// Rechnungstyp ist eine Abbildung verschiedener Rechnungstypen zur Kennzeichnung von bo.Rechnung en
//go:generate stringer --type Rechnungstyp
//go:generate jsonenums --type Rechnungstyp
type Rechnungstyp int

const (
	// Eine Rechnung vom Lieferanten an einen Endkunden über die Lieferung von Energie
	ENDKUNDENRECHNUNG              Rechnungstyp = iota + 1
	NETZNUTZUNGSRECHNUNG                        // Eine Rechnung vom Netzbetreiber an den Netznutzer. (i.d.R. der Lieferant) über die Netznutzung
	MEHRMINDERMENGENRECHNUNG                    // Eine Rechnung vom Netzbetreiber an den Netznutzer. (i.d.R. der Lieferant) zur Abrechnung von Mengen-Differenzen zwischen Bilanzierung und Messung
	MESSSTELLENBETRIEBSRECHNUNG                 // Rechnung eines Messstellenbetreibers an den Messkunden
	BESCHAFFUNGSRECHNUNG                        // Rechnungen zwischen einem Händler und Einkäufer von Energie
	AUSGLEICHSENERGIERECHNUNG                   // Rechnung an den Verursacher von Ausgleichsenergie
	ABSCHLAGSRECHNUNG                           // Rechnung über einen Abschlag (non-standard)
	WIMRECHNUNG                                 // Rechnung im Rahmen der Wechselprozesse im Messwesen (non-standard)
	SELBSTAUSGESTELLTERECHNUNGMEMI              // Rechnung für eine Mehrmenge vom VNB an den Lieferanten (non-standard)
)
