package zeitreihentyp

// Zeitreihentyp sind die Codes der Summenzeitreihentypen
// vgl. https://www.edi-energy.de/index.php?id=38&tx_bdew_bdew%5Buid%5D=695&tx_bdew_bdew%5Baction%5D=download&tx_bdew_bdew%5Bcontroller%5D=Dokument&cHash=67782e05d8b0f75fbe3a0e1801d07ed0
// Note that this enum is not official BO4E standard (yet)!
//
//go:generate stringer --type Zeitreihentyp
//go:generate jsonenums --type Zeitreihentyp
type Zeitreihentyp int

const (
	EGS     Zeitreihentyp = iota + 1 // EGS ist eine Einspeisegangsumme
	LGS                              // LGS ist eine Lastgangsumme
	NZR                              // NZR ist eine Netzzeitreihe
	SES                              // SES ist eine Standardeinspeiseprofilsumme
	SLS                              // SLS ist eine Standardlastsumme
	TES                              // TES ist eine tagesparameter-abhängige Einspeiseprofilsumme
	TLS                              // TLS ist eine tagesparameter-abhängige Lastprofilsumme
	SLS_TLS                          // SLS_TLS ist die gemeinsame Messung aus SLS und TLS
	SES_TES                          // SES_TES ist die gemeinsame Messung aus SES und TES
)
