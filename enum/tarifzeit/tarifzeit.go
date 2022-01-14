package tarifzeit

// Tarifzeit wird zur Kennzeichnung verschiedener Tarifzeiten, beispielsweise zur Bepreisung oder zur Verbrauchsermittlung verwendet.
//go:generate stringer --type Tarifzeit
//go:generate jsonenums --type Tarifzeit
type Tarifzeit int

const (
	// TZ_STANDARD = single tariff
	TZ_STANDARD Tarifzeit = iota + 1
	TZ_HT                 // 	Tarifzeit für Hochtarif bei Mehrtarif-Konfigurationen
	TZ_NT                 // 	Tarifzeit für Niedrigtarif bei Mehrtarif-Konfigurationen
)
