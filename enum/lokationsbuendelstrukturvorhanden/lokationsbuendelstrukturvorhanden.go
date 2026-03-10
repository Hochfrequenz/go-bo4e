package lokationsbuendelstrukturvorhanden

// LokationsbuendelstrukturVorhanden gibt an, ob eine Lokationsbündelstruktur vorhanden ist
//
//go:generate stringer --type LokationsbuendelstrukturVorhanden
//go:generate jsonenums --type LokationsbuendelstrukturVorhanden
type LokationsbuendelstrukturVorhanden int

const (
	VORHANDEN       LokationsbuendelstrukturVorhanden = iota + 1 // VORHANDEN means a Lokationsbündelstruktur is present (EDIFACT Z31)
	NICHT_VORHANDEN                                              // NICHT_VORHANDEN means a Lokationsbündelstruktur is not present (EDIFACT Z39)
)
