package netzebene

// Netzebene ist eine Auflistung möglicher Netzebenen innerhalb der Energiearten Strom und Gas.
//go:generate stringer --type Netzebene
//go:generate jsonenums --type Netzebene
type Netzebene int

const (
	NSP          Netzebene = iota + 1 // NSP is for Niederspannung
	MSP                               // MSP is for Mittelspannung
	HSP                               // HSP is for Hochspannung
	HSS                               // HSS is for Hoechstspannung
	MSP_NSP_UMSP                      // MSP_NSP_UMSP is for MS/NS Umspannung
	HSP_MSP_UMSP                      // HSP_MSP_UMSP is for HS/MS Umspannung
	HSS_HSP_UMSP                      // HSS_HSP_UMSP is for HOES/HS Umspannung
	HD                                // HD is for Hochdruck
	MD                                // MD is for Mitteldruck
	ND                                // ND is for Niederdruck
)
