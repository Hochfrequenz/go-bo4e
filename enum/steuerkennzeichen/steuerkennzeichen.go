package steuerkennzeichen

// Steuerkennzeichen dient der Kennzeichnung verschiedener Steuersätze und Verfahren.
//go:generate stringer --type Steuerkennzeichen
//go:generate jsonenums --type Steuerkennzeichen
type Steuerkennzeichen int

const (
	UST0  Steuerkennzeichen = iota + 1 // UST0 means Keine Umsatzsteuer, bzw. nicht steuerbar.
	UST19                              // UST19 means Umsatzsteuer 19%
	UST7                               // UST7 means Umsatzsteuer 7%
	VST0                               // VST0 means Keine Vorsteuer, bzw. nicht steuerbar.
	VST19                              // VST19 means Vorsteuer 19%
	VST7                               // VST7 means Vorsteuer 7%
	RCV                                // RCV means Reverse Charge Verfahren (Umkehrung der Steuerpflicht)
)
