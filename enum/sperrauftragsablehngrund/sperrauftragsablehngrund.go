package sperrauftragsablehngrund

// Sperrauftragsablehngrund is an enum
//
//go:generate stringer --type Sperrauftragsablehngrund
//go:generate jsonenums --type Sperrauftragsablehngrund
type Sperrauftragsablehngrund int

const (
	DUPLIKAT Sperrauftragsablehngrund = iota + 1
	FALSCHER_MSB
	FALSCHE_SPANNUNGSEBENE
	WEITERE_MALO_BETROFFEN
	ANDERER_ABLEHNGRUND
	FRISTVERLETZUNG_TERMINGEBUNDEN
	ANDERER_FEHLER
	LIEGT_BEREITS_VOR
	ANDERER_ZUKUENFTIGER_LIEFERANT
	BESTAETIGTER_LIEFERBEGINN
)
