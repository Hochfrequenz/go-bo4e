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
)
