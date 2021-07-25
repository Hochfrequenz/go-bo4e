package netzebene

//go:generate stringer --type Netzebene

type Netzebene int

const (
	NSP          Netzebene = iota // Niederspannung
	MSP                           // Mittelspannung
	HSP                           // Hochspannung
	HSS                           // Hoechstspannung
	MSP_NSP_UMSP                  // MS/NS Umspannung
	HSP_MSP_UMSP                  // HS/MS Umspannung
	HSS_HSP_UMSP                  // HOES/HS Umspannung
	HD                            // Hochdruck
	MD                            // Mitteldruck
	N                             // Niederdruck
)
