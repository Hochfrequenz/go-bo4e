package preistyp

// Preistyp Aufschlüsselung der Preistypen in Tarifen.
//
//go:generate stringer --type Preistyp
//go:generate jsonenums --type Preistyp
type Preistyp int

const (
	ARBEITSPREIS_EINTARIF Preistyp = iota + 1
	ARBEITSPREIS_HT
	ARBEITSPREIS_NT
	LEISTUNGSPREIS
	MESSPREIS
	ENTGELT_ABLESUNG
	ENTGELT_ABRECHNUNG
	ENTGELT_MSB
	PROVISION
)
