package fernschaltung

// Fernschaltung describes whether a meter has a remote control
// Note that this enum is not official BO4E standard (yet)!
//
//go:generate stringer --type Fernschaltung
//go:generate jsonenums --type Fernschaltung
type Fernschaltung int

const (
	// VORHANDEN means there is remote control (EDIFACT Z06)
	VORHANDEN       Fernschaltung = iota + 1
	NICHT_VORHANDEN               // NICHT_VORHANDEN means there is not remote control (EDIFACT Z07)
)
