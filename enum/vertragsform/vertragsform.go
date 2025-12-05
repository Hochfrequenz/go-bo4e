package vertragsform

// Vertragsform is an enum
//
//go:generate stringer --type Vertragsform
//go:generate jsonenums --type Vertragsform
type Vertragsform int

const (
	ONLINE Vertragsform = iota + 1
	DIREKT
	FAX
)
