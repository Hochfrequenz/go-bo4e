package anfragekategorie

//go:generate stringer --type Anfragekategorie
//go:generate jsonenums --type Anfragekategorie
type Anfragekategorie int

const (
	PROZESSDATENBERICHT Anfragekategorie = iota + 1
	GERAETEUEBERNAHME
	WEITERVERPFLICHTUNG_BETRIEB_MELO
	AENDERUNG_MELO
	STAMMDATEN_MALO_ODER_MELO
	BILANZIERTE_MENGE_MEHR_MINDER_MENGEN
	ALLOKATIONSLISTE_MEHR_MINDER_MENGEN
	ENERGIEMENGE_UND_LEISTUNGSMAXIMUM
	ABRECHNUNG_MESSSTELLENBETRIEB_MSB_AN_LF
	AENDERUNG_PROGNOSEGRUNDLAGE_GERAETEKONFIGURATION
	AENDERUNG_GERAETEKONFIGURATION
	REKLAMATION_VON_WERTEN
	LASTGANG_MALO_TRANCHE
	SPERRUNG
	ENTSPERRUNG
)
