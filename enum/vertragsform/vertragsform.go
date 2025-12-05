package vertragsform

// Vertragsform Aufzählung der Möglichkeiten zu Vertragsformen in Ausschreibungen.
//
//go:generate stringer --type Vertragsform
//go:generate jsonenums --type Vertragsform
type Vertragsform int

const (
	ONLINE Vertragsform = iota + 1
	DIREKT
	FAX
)
