package befestigungsart

// Befestigungsart Befestigungsart
//
//go:generate stringer --type Befestigungsart
//go:generate jsonenums --type Befestigungsart
type Befestigungsart int

const (
	STECKTECHNIK Befestigungsart = iota + 1
	DREIPUNKT
	HUTSCHIENE
	EINSTUTZEN
	ZWEISTUTZEN
)
