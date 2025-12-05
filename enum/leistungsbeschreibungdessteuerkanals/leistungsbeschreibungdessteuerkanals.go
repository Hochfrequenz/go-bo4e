package leistungsbeschreibungdessteuerkanals

// LeistungsbeschreibungDesSteuerkanals is an enum
//
//go:generate stringer --type LeistungsbeschreibungDesSteuerkanals
//go:generate jsonenums --type LeistungsbeschreibungDesSteuerkanals
type LeistungsbeschreibungDesSteuerkanals int

const (
	AN_AUS LeistungsbeschreibungDesSteuerkanals = iota + 1
	GESTUFT
)
