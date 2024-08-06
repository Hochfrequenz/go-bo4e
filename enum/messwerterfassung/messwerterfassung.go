package messwerterfassung

// Messwerterfassung describes how meter readings are recorded
// Note that this enum is not official BO4E standard (yet)!
//
//go:generate stringer --type Messwerterfassung
//go:generate jsonenums --type Messwerterfassung
type Messwerterfassung int

const (
	// FERNAUSLESBAR describes automatic meter readings (EDIFACT AMR)
	FERNAUSLESBAR Messwerterfassung = iota + 1
	// MANUELL_AUSGELESENE describes manual meter readings (EDIFACT MMR)
	MANUELL_AUSGELESENE
)
