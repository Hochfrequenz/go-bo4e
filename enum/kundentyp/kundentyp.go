package kundentyp

//go:generate stringer --type Kundentyp
//go:generate jsonenums --type Kundentyp
type Kundentyp int

const (
	PRIVAT Kundentyp = iota + 1
	LANDWIRT
	SONSTIGE
	HAUSHALT
	DIREKTHEIZUNG
	GEMEINSCHAFT_MFH
	KIRCHE
	KWK
	LADESAEULE
	BELEUCHTUNG_OEFFENTLICH
	BELEUCHTUNG_STRASSE
	SPEICHERHEIZUNG
	UNTERBR_EINRICHTUNG
	WAERMEPUMPE
)
