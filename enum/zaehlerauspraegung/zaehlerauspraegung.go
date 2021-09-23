package zaehlerauspraegung

// Zaehlerauspraegung states if a meter measures in one or two directions, i.e. only supply xor feeding or both
//go:generate stringer --type Zaehlerauspraegung
//go:generate jsonenums --type Zaehlerauspraegung
type Zaehlerauspraegung int

const (
	Einrichtungszaehler  Zaehlerauspraegung = iota + 1 // power/electricity
	Zweirichtungszaehler                               // gas
)
