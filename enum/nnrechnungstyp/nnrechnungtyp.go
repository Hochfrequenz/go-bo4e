package nnrechnungstyp

// NNRechnungstyp ist der Typ der bo.Netznutzungsrechnung (Abbildung verschiedener in der INVOIC angegebenen Rechnungstypen)
//go:generate stringer --type NNRechnungstyp
//go:generate jsonenums --type NNRechnungstyp
type NNRechnungstyp int

const (
	ABSCHLUSSRECHNUNG          NNRechnungstyp = iota + 1 // ABSCHLUSSRECHNUNG is a Abschlussrechnung
	ABSCHLAGSRECHNUNG                                    // ABSCHLAGSRECHNUNG is a Abschlagsrechnung
	TURNUSRECHNUNG                                       // TURNUSRECHNUNG is a Turnusrechnung
	MONATSRECHNUNG                                       // MONATSRECHNUNG is a Monatsrechnung
	WIMRECHNUNG                                          // WIMRECHNUNG is a Rechnung für WiM
	ZWISCHENRECHNUNG                                     // ZWISCHENRECHNUNG is a Zwischenrechnung
	INTEGRIERTE_13TE_RECHNUNG                            // INTEGRIERTE_13TE_RECHNUNG is a Integrierte 13. Rechnung
	ZUSAETZLICHE_13TE_RECHNUNG                           // ZUSAETZLICHE_13TE_RECHNUNG is a Zusätzliche 13. Rechnung
	MEHRMINDERMENGENRECHNUNG                             // MEHRMINDERMENGENRECHNUNG is a Mehr/Mindermengenabrechnung
)
