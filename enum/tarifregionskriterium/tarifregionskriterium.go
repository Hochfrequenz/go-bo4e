package tarifregionskriterium

// Tarifregionskriterium Mit diesen Kriterien können regionale Bereiche definiert werden.
//
//go:generate stringer --type Tarifregionskriterium
//go:generate jsonenums --type Tarifregionskriterium
type Tarifregionskriterium int

const (
	NETZ_NUMMER Tarifregionskriterium = iota + 1
	POSTLEITZAHL
	ORT
	GRUNDVERSORGER_NUMMER
)
