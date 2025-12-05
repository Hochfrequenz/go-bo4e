package preismodell

// Preismodell is an enum
//
//go:generate stringer --type Preismodell
//go:generate jsonenums --type Preismodell
type Preismodell int

const (
	TRANCHE Preismodell = iota + 1
)
