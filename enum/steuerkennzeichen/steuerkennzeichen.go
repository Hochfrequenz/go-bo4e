package steuerkennzeichen

// Steuerkennzeichen dient der Kennzeichnung verschiedener Steuersätze und Verfahren.
//go:generate stringer --type Steuerkennzeichen
//go:generate jsonenums --type Steuerkennzeichen
type Steuerkennzeichen int

const (
	Ust0  Steuerkennzeichen = iota + 1 // Keine Umsatzsteuer, bzw. nicht steuerbar.
	Ust19                              // Umsatzsteuer 19%
	Ust7                               // Umsatzsteuer 7%
	Vst0                               // Keine Vorsteuer, bzw. nicht steuerbar.
	Vst19                              // Vorsteuer 19%
	Vst7                               // Vorsteuer 7%
	RCV                                // Reverse Charge Verfahren (Umkehrung der Steuerpflicht)
)
