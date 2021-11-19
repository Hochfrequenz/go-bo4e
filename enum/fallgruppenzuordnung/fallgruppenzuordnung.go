package fallgruppenzuordnung

// Fallgruppenzuordnung as described by edi@energy
// Note that this enum is not official BO4E standard (yet)!
//go:generate stringer --type Fallgruppenzuordnung
//go:generate jsonenums --type Fallgruppenzuordnung
type Fallgruppenzuordnung int

const (
	// GABI_RLMmT is an abbreviation for "RLM mit Tagesband"
	GABI_RLMmT  Fallgruppenzuordnung = iota + 1
	GABI_RLMoT                       // GABI_RLMoT is an abbreviation for "RLM ohne Tagesband"
	GABI_RLMNEV                      // GABI_RLMNEV is an abbreviation for "RLM im Nominierungsersatzverfahren"
)
