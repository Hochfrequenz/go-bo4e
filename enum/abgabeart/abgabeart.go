package abgabeart

//go:generate stringer --type AbgabeArt
//go:generate jsonenums --type AbgabeArt
type AbgabeArt int

const (
	// KAS für alle konzessionsvertraglichen Sonderregelungen, die nicht in die Systematik der KAV eingegliedert
	KAS AbgabeArt = iota + 1
	//SA: Sondervertragskunden  1 kV, Preis nach § 2 (3) (für Strom 0,11 ct/kWh und für Gas 0,03 ct/kWh)
	SA // SA =
	//SAS: Kennzeichnung, dass ein abweichender Preis für Sondervertragskunden vorliegt
	SAS
	//TA: Tarifkunden, für Strom § 2. (2) 1b HT bzw.ET(hohe KA) und für Gas § 2 (2) 2b
	TA
	//TAS: Kennzeichnung, dass ein abweichender Preis für Tarifkunden vorliegt
	TAS
	//TK: für Gas nach KAV § 2 (2) 2a bei ausschließlicher Nutzung zum Kochen und Warmwassererzeugung
	TK
	//TKS: Kennzeichnung, wenn nach KAV § 2 (2) 2a ein anderen Preis zu verwenden ist
	TKS
	//TS: für Strom mit Schwachlast § 2. (2) 1a NT(niedrige KA, 0,61 ct/kWh)
	TS
	//TSS: Kennzeichnung, dass ein abweichender Preis für Schwachlast angewendet wird
	TSS
)
