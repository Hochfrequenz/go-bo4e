package tarifkalkulationsmethode

// Tarifkalkulationsmethode Aulistung der verschiedenen Berechnungsmethoden für ein Preisblatt.
//
//go:generate stringer --type Tarifkalkulationsmethode
//go:generate jsonenums --type Tarifkalkulationsmethode
type Tarifkalkulationsmethode int

const (
	STAFFELN Tarifkalkulationsmethode = iota + 1
	ZONEN
	BESTABRECHNUNG_STAFFEL
	PAKETPREIS
)
