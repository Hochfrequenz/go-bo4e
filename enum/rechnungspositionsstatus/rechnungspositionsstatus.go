package rechnungspositionsstatus

// RechnungspositionsStatus is an enum
//
//go:generate stringer --type RechnungspositionsStatus
//go:generate jsonenums --type RechnungspositionsStatus
type RechnungspositionsStatus int

const (
	ROH RechnungspositionsStatus = iota + 1
	ROH_AUSGENOMMEN
	ABRECHENBAR
	ABRECHENBAR_AUSGENOMMEN
	ABGERECHNET
)
