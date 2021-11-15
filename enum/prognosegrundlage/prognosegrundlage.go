package prognosegrundlage

// Prognosegrundlage describes the data used to perform a prognosis
// Note that this enum is not official BO4E standard (yet)!
//go:generate stringer --type Prognosegrundlage
//go:generate jsonenums --type Prognosegrundlage
type Prognosegrundlage int

const (
	// WERTE means prognosis based on values
	WERTE   Prognosegrundlage = iota + 1
	PROFILE                   // PROFILE means prognosis based on profiles
)
