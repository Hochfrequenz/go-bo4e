package zaehlertyp

// Zaehlertyp contains different types of meters for power and gas
//go:generate stringer --type Zaehlertyp
//go:generate jsonenums --type Zaehlertyp
type Zaehlertyp int

const (
	// Drehstromzaehler is a polyphase meter
	Drehstromzaehler      Zaehlertyp = iota + 1
	Balgengaszaehler                 // Balgengaszähler
	Drehkolbenzaehler                // Drehkolbengaszähler
	SmartMeter                       // Smart Meter Zähler
	Leistungszaehler                 // leistungsmessender Zähler
	Maximumgaehler                   // Maximumzähler
	Turbinenradgaszaehler            // Turbinenradgaszähler
	Ultraschallgaszaehler            // Ultraschallgaszähler
	Wechselstromzaehler              // Wechselstromzähler
)
