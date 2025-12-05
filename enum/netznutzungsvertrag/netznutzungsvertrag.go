package netznutzungsvertrag

// Netznutzungsvertrag is an enum
//
//go:generate stringer --type Netznutzungsvertrag
//go:generate jsonenums --type Netznutzungsvertrag
type Netznutzungsvertrag int

const (
	KUNDEN_NB Netznutzungsvertrag = iota + 1
	LIEFERANTEN_NB
)
