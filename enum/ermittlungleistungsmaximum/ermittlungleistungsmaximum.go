package ermittlungleistungsmaximum

//go:generate stringer --type ErmittlungLeistungsmaximum
//go:generate jsonenums --type ErmittlungLeistungsmaximum
type ErmittlungLeistungsmaximum int

const (
	VERWENDUNG_HOCHLASTFENSTER ErmittlungLeistungsmaximum = iota + 1
	KEINE_VERWENDUNG_HOCHLASTFENSTER
)
