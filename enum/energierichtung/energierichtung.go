package energierichtung

// Energierichtung describes if the energy is supplied out of or feeded into the grid
//go:generate stringer --type Energierichtung
//go:generate jsonenums --type Energierichtung
type Energierichtung int

const (
	Aussp  Energierichtung = iota + 1 // supply
	Einsp                               // feeding
)
