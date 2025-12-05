package geraeteart

// Geraeteart Auflistung möglicher Geraetearten. This is more broadly defined as a <see cref="Geraetetyp"/>, so a Zaehleinrichtung as Gerateart could be a elektronischer Haushaltszähler as a Gerätetyp.
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
