package profilverfahren

// Profilverfahren describes how profiles are constructed
// Note that this enum is not official BO4E standard (yet)!
//go:generate stringer --type Profilverfahren
//go:generate jsonenums --type Profilverfahren
type Profilverfahren int

const (
	// SYNTHETISCH means SLP
	SYNTHETISCH Profilverfahren = iota + 1
	ANALYTISCH                  // ANALYTISCH means ALP
)
