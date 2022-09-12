package titel

// Titel is a title
//
//go:generate stringer --type Titel
//go:generate jsonenums --type Titel
type Titel int

const (
	// DR is a PhD
	DR      Titel = iota + 1
	PROF          // PROF is a professor
	PROF_DR       // PROF_DR is a professor with PhD title
)
