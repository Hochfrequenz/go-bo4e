package zaehlerauspraegung

// Zaehlerauspraegung states if a meter measures in one or two directions, i.e. only supply xor feeding or both
//go:generate stringer --type Zaehlerauspraegung
//go:generate jsonenums --type Zaehlerauspraegung
type Zaehlerauspraegung int

const (
	// EINRICHTUNGSZAEHLER is a single direction meter
	EINRICHTUNGSZAEHLER Zaehlerauspraegung = iota + 1
	// ZWEIRICHTUNGSZAEHLER is a bidirectional meter
	ZWEIRICHTUNGSZAEHLER
)
