package steuerkennzeichen

// Steuerkennzeichen dient der Kennzeichnung verschiedener Steuersätze und Verfahren.
//go:generate stringer --type Steuerkennzeichen
//go:generate jsonenums --type Steuerkennzeichen
type Steuerkennzeichen int

const (
	// UST0 means Keine Umsatzsteuer, bzw. nicht steuerbar.
	UST_0  Steuerkennzeichen = iota + 1
	UST_19                   // UST19 means Umsatzsteuer 19%
	US_T7                    // UST7 means Umsatzsteuer 7%
	VST_0                    // VST0 means Keine Vorsteuer, bzw. nicht steuerbar.
	VST_19                   // VST19 means Vorsteuer 19%
	VST_7                    // VST7 means Vorsteuer 7%
	RCV                      // RCV means Reverse Charge Verfahren (Umkehrung der Steuerpflicht)
)
