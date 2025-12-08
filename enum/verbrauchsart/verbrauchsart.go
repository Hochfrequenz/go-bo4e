package verbrauchsart

// Verbrauchsart describes the consumption type of a market location.
//
//go:generate stringer --type Verbrauchsart
//go:generate jsonenums --type Verbrauchsart
type Verbrauchsart int

const (
	// KL stands for Kraft/Licht
	KL   Verbrauchsart = iota + 1
	KLW                // KLW stands for Kraft/Licht/Wärme
	KLWS               // KLWS stands for Kraft/Licht/Wärme/Speicherheizung
	W                  // W stands for Wärme
	WS                 // WS stands for Wärme/Speicherheizung
	WK                 // Z65: Wärme/Kälte
	EM                 // ZE5: E-Mobilität
	STRB               // ZA8: Straßenbeleuchtung
	STW                // ZB3: Steuerung Wärmeabgabe
)
