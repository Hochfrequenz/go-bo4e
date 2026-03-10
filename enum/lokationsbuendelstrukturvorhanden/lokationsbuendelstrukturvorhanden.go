package lokationsbuendelstrukturvorhanden

// LokationsbuendelstrukturVorhanden: Gibt an, ob eine Lokationsbündelstruktur vorhanden ist.
//
//go:generate stringer --type LokationsbuendelstrukturVorhanden
//go:generate jsonenums --type LokationsbuendelstrukturVorhanden
type LokationsbuendelstrukturVorhanden int

const (
	// VORHANDEN: Lokationsbündelstruktur vorhanden (EDIFACT Z31)
	VORHANDEN LokationsbuendelstrukturVorhanden = iota + 1
	// NICHT_VORHANDEN: Lokationsbündelstruktur nicht vorhanden (EDIFACT Z39)
	NICHT_VORHANDEN
)
