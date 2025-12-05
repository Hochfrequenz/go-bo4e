package tarifmerkmal

//go:generate stringer --type Tarifmerkmal
//go:generate jsonenums --type Tarifmerkmal
type Tarifmerkmal int

const (
	VORKASSE Tarifmerkmal = iota + 1
	PAKET
	KOMBI
	FESTPREIS
	BAUSTROM
	HAUSLICHT
	HEIZSTROM
)
