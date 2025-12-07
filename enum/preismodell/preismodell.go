package preismodell

// Preismodell Bezeichnung der Preismodelle in Ausschreibungen für die Energielieferung.
//
//go:generate stringer --type Preismodell
//go:generate jsonenums --type Preismodell
type Preismodell int

const (
	TRANCHE Preismodell = iota + 1
)
