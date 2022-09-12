package schwachlastfaehigkeit

// Schwachlastfaehigkeit ist de facto ein boolean der die Schwachlastfaehigkeit beschreibt
// Note that this is not official BO4E standard (yet).
//
//go:generate stringer --type Schwachlastfaehigkeit
//go:generate jsonenums --type Schwachlastfaehigkeit
type Schwachlastfaehigkeit int

const (
	// SCHWACHLASTFAEHIG heißt "true" (EDIFACT Z60)
	SCHWACHLASTFAEHIG Schwachlastfaehigkeit = iota + 1
	// NICHT_SCHWACHLASTFAEHIG heißt "false" (EDIFACT Z59)
	NICHT_SCHWACHLASTFAEHIG
)
