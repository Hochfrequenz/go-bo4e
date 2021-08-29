package netzebene

//go:generate stringer --type Netzebene
//go:generate jsonenums --type Netzebene
// Netzebene ist eine Auflistung möglicher Netzebenen innerhalb der Energiearten Strom und Gas.
type Netzebene int

const (
	NSP          Netzebene = iota + 1 // Niederspannung
	MSP                               // Mittelspannung
	HSP                               // Hochspannung
	HSS                               // Hoechstspannung
	MSP_NSP_UMSP                      // MS/NS Umspannung
	HSP_MSP_UMSP                      // HS/MS Umspannung
	HSS_HSP_UMSP                      // HOES/HS Umspannung
	HD                                // Hochdruck
	MD                                // Mitteldruck
	N                                 // Niederdruck
)
