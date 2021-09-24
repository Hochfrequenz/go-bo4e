package zaehlertyp

// Sparte contains different divisions of typical utilities.
//go:generate stringer --type Zaehlertyp
//go:generate jsonenums --type Zaehlertyp
type Zaehlertyp int

const (
	Drehstromzaehler      Zaehlertyp = iota + 1 // Drehstromzähler
	Balgengaszaehler                            //	Balgengaszähler
	Drehkolbenzaehler                           //	Drehkolbengaszähler
	SmartMeter                                  // Smart Meter Zähler
	Leistungszaehler                            //	leistungsmessender Zähler
	Maximumgaehler                              //	Maximumzähler
	Turbinenradgaszaehler                       //	Turbinenradgaszähler
	Ultraschallgaszaehler                       //	Ultraschallgaszähler
	Wechselstromzaehler                         // Wechselstromzähler
)
