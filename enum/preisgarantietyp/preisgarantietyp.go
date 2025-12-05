package preisgarantietyp

// Preisgarantietyp Aufzählung der Möglichkeiten für die Vergabe von Preisgarantien
//
//go:generate stringer --type Preisgarantietyp
//go:generate jsonenums --type Preisgarantietyp
type Preisgarantietyp int

const (
	ALLE_PREISBESTANDTEILE_NETTO Preisgarantietyp = iota + 1
	PREISBESTANDTEILE_OHNE_ABGABEN
	NUR_ENERGIEPREIS
)
