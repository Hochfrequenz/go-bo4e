package eegvermarktungsform

// Vermarktungsformen gemäß dem Erneuerbare-Energien-Gesetz (EEG).
//
//go:generate stringer --type EEGVermarktungsform
//go:generate jsonenums --type EEGVermarktungsform
type EEGVermarktungsform int

const (
	// Ausfallvergütung für den Fall, dass andere Vermarktungsmethoden nicht verfügbar sind. Z90
	AUSFALLVERGUETUNG EEGVermarktungsform = iota + 1
	MARKTPRAEMIE                          // Marktprämie für die geförderte Direktvermarktung. Z91
	SONSTIGE                              // Sonstige Vermarktungsformen ohne gesetzliche Vergütung. Z92
	KWKG_VERGUETUNG                       // Vergütung nach dem Kraft-Wärme-Kopplungsgesetz (KWKG). Z94
)
