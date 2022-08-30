package bemessungsgroesse

// Bemessungsgroesse
//
//go:generate stringer --type Bemessungsgroesse
//go:generate jsonenums --type Bemessungsgroesse
type Bemessungsgroesse int

const (
	// WIRKARBEIT_EL elektrische Wirkarbeit
	WIRKARBEIT_EL     Bemessungsgroesse = iota + 1
	LEISTUNG_EL                         //	elektrische Leistung
	BLINDARBEIT_KAP                     //	Blindarbeit kapazitiv
	BLINDARBEIT_IND                     // Blindarbeit induktiv
	BLINDLEISTUNG_KAP                   //	Blindleistung kapazitiv
	BLINDLEISTUNG_IND                   //	Blindleistung induktiv
	WIRKARBEIT_TH                       //	thermische Wirkarbeit
	LEISTUNG_TH                         //	thermische Leistung
	VOLUMEN                             //	Volumen
	VOLUMENSTROM                        //	Volumenstrom (Volumen/Zeit)
	BENUTZUNGSDAUER                     //	Benutzungsdauer (Arbeit/Leistung)
	ANZAHL                              //	Darstellung einer Stückzahl
)
