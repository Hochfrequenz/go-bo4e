package aufabschlagsziel

// AufAbschlagsziel Der Preis, auf den sich ein Auf- oder Abschlag bezieht.
//
//go:generate stringer --type AufAbschlagsziel
//go:generate jsonenums --type AufAbschlagsziel
type AufAbschlagsziel int

const (
	ARBEITSPREIS_HT AufAbschlagsziel = iota + 1
	ARBEITSPREIS_NT
	ARBEITSPREIS_HT_NT
	GRUNDPREIS
	GESAMTPREIS
)
