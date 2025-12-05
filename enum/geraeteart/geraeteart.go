package geraeteart

// Geraeteart is an enum
//
//go:generate stringer --type Geraeteart
//go:generate jsonenums --type Geraeteart
type Geraeteart int

const (
	WANDLER Geraeteart = iota + 1
	KOMMUNIKATIONSEINRICHTUNG
	TECHNISCHE_STEUEREINRICHTUNG
	MENGENUMWERTER
	SMARTMETER_GATEWAY
	STEUERBOX
	ZAEHLEINRICHTUNG
)
