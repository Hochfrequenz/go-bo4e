package sperrauftragsverhinderungsgrund

//go:generate stringer --type Sperrauftragsverhinderungsgrund
//go:generate jsonenums --type Sperrauftragsverhinderungsgrund
type Sperrauftragsverhinderungsgrund int

const (
	RECHTLICHER_GRUND_FEHLT Sperrauftragsverhinderungsgrund = iota + 1
	AKTIVE_ZUTRITTSVERWEIGERUNG
	PASSIVE_ZUTRITTSVERWEIGERUNG
	ANDERER_VERHINDERUNGSGRUND
	TATSAECHLICHER_VERHINDERUNGSGRUND
	TECHNISCHER_VERHINDERUNGSGRUND
	ANSCHLUSSNUTZER_WURDE_NICHT_ANGETROFFEN // ANSCHLUSSNUTZER_WURDE_NICHT_ANGETROFFEN means: Es gab keine Anzeichen dafür, dass dieser anwesend war (EBD 0472 A08). Details werden im Freitext bo4e.bo.Auftrag.Bemerkungen vermerkt.
)
