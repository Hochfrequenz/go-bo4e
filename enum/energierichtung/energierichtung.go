package energierichtung

// Energierichtung describes if the energy is supplied out of or fed into the grid
//go:generate stringer --type Energierichtung
//go:generate jsonenums --type Energierichtung
type Energierichtung int

const (
	AUSSP Energierichtung = iota + 1 // AUSSP is short for "Ausspeisung" / supply
	EINSP                            // EINSP is short for "Einspeisung" / feeding
)
