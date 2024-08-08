package wahlrechtprognosegrundlage

// WahlrechtPrognosegrundlage describes if and how the right to choose the basis for the prognosis is given
//
//go:generate stringer --type WahlrechtPrognosegrundlage
//go:generate jsonenums --type WahlrechtPrognosegrundlage

type WahlrechtPrognosegrundlage int

const (
	DURCH_LF WahlrechtPrognosegrundlage = iota + 1
	DURCH_LF_NICHT_GEGEBEN
	NICHT_WEGEN_GROSSEN_VERBRAUCHS
	NICHT_WEGEN_EIGENVERBRAUCH
	NICHT_WEGEN_TAGES_VERBRAUCH
	NICHT_WEGEN_ENWG
)
