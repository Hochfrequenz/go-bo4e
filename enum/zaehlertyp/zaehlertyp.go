package zaehlertyp

// Zaehlertyp contains different types of meters for power and gas
//
//go:generate stringer --type Zaehlertyp
//go:generate jsonenums --type Zaehlertyp
type Zaehlertyp int

const (
	// DREHSTROMZAEHLER is a polyphase meter
	DREHSTROMZAEHLER      Zaehlertyp = iota + 1
	BALGENGASZAEHLER                 // Balgengaszähler
	DREHKOLBENZAEHLER                // Drehkolbengaszähler
	SMARTMETER                       // Smart Meter Zähler
	LEISTUNGSZAEHLER                 // leistungsmessender Zähler
	MAXIMUMZAEHLER                   // Maximumzähler
	TURBINENRADGASZAEHLER            // Turbinenradgaszähler
	ULTRASCHALLGASZAEHLER            // Ultraschallgaszähler
	WECHSELSTROMZAEHLER              // Wechselstromzähler
	WIRBELGASZAEHLER
	MESSDATENREGISTRIERGERAET
	ELEKTRONISCHERHAUSHALTSZAEHLER
	SONDERAUSSTATTUNG
	WASSERZAEHLER
	MODERNEMESSEINRICHTUNG
	NEUEMESSEINRICHTUNGGAS
)
