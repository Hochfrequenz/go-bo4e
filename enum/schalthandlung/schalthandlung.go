package schalthandlung

//go:generate stringer --type Schalthandlung
//go:generate jsonenums --type Schalthandlung
type Schalthandlung int

const (
	LEISTUNG_LOKATION_AN Schalthandlung = iota + 1
	LEISTUNG_LOKATION_AUS
)
