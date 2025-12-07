package qualitaet

// Qualitaet Qualitätsstufen eines Business-Objektes, v.a. aus Sicht Marktkommunikation
//
//go:generate stringer --type Qualitaet
//go:generate jsonenums --type Qualitaet
type Qualitaet int

const (
	VOLLSTAENDIG Qualitaet = iota + 1
	INFORMATIV
	IM_SYSTEM_VORHANDEN
	ERWARTET
	VORLAEUFIG
	UNVOLLSTAENDIG
)
