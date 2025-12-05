package rechnungslegung

// Rechnungslegung is an enum
//
//go:generate stringer --type Rechnungslegung
//go:generate jsonenums --type Rechnungslegung
type Rechnungslegung int

const (
	MONATSRECHN Rechnungslegung = iota + 1
	ABSCHL_MONATSRECHN
	ABSCHL_JAHRESRECHN
	MONATSRECHN_JAHRESRECHN
	VORKASSE
)
