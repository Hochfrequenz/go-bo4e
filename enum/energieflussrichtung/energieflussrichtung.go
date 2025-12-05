package energieflussrichtung

// Energieflussrichtung is an enum
//
//go:generate stringer --type Energieflussrichtung
//go:generate jsonenums --type Energieflussrichtung
type Energieflussrichtung int

const (
	VERBRAUCH Energieflussrichtung = iota + 1
	ERZEUGUNG
)
