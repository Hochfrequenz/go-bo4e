package nnrechnungstyp

// NNRechnungstyp ist der Typ der bo.Netznutzungsrechnung (Abbildung verschiedener in der INVOIC angegebenen Rechnungstypen)
//
//go:generate stringer --type NNRechnungstyp
//go:generate jsonenums --type NNRechnungstyp
type NNRechnungstyp int

const (
	// ABSCHLUSSRECHNUNG is a Abschlussrechnung
	ABSCHLUSSRECHNUNG          NNRechnungstyp = iota + 1
	ABSCHLAGSRECHNUNG                         // ABSCHLAGSRECHNUNG is a Abschlagsrechnung
	TURNUSRECHNUNG                            // TURNUSRECHNUNG is a Turnusrechnung
	MONATSRECHNUNG                            // MONATSRECHNUNG is a Monatsrechnung
	WIMRECHNUNG                               // WIMRECHNUNG is a Rechnung für WiM
	ZWISCHENRECHNUNG                          // ZWISCHENRECHNUNG is a Zwischenrechnung
	INTEGRIERTE_13TE_RECHNUNG                 // INTEGRIERTE_13TE_RECHNUNG is a Integrierte 13. Rechnung
	ZUSAETZLICHE_13TE_RECHNUNG                // ZUSAETZLICHE_13TE_RECHNUNG is a Zusätzliche 13. Rechnung
	MEHRMINDERMENGENRECHNUNG                  // MEHRMINDERMENGENRECHNUNG is a Mehr/Mindermengenabrechnung
	MSBRECHNUNG                               // MSBRECHNUNG is the invoice for the meter server operation (Messstellenbetrieb/MSB)
	_13TE_RECHNUNG                            // _13TE_RECHNUNG is the 13th invoice (13. Rechnung)
)
