package energierichtung

// Energierichtung describes if the energy is supplied out of or fed into the grid
//
//go:generate stringer --type Energierichtung
//go:generate jsonenums --type Energierichtung
type Energierichtung int

const (
	// AUSSP is short for "Ausspeisung" / supply
	AUSSP  Energierichtung = iota + 1
	EINSP                  // EINSP is short for "Einspeisung" / feeding
	RUHEND                 // RUHEND e Lokation (in einer Kundenanlage)

	KUNDENANLAGE // KUNDENANLAGE kennzeichnet eine Marktlokation als Kundenanlage gemäß EnWG
)
