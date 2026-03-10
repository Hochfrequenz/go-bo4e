package unternehmensart

// Unternehmensart kategorisiert Unternehmen nach Ihrer Marktrolle (basierend auf dem PARTIN-Datenmodell)
//
//go:generate stringer --type Unternehmensart
//go:generate jsonenums --type Unternehmensart
type Unternehmensart int

const (
	LIEFERANT Unternehmensart = iota + 1
	NETZBETREIBER
	MESSSTELLENBETREIBER
	UEBERTRAGUNGSNETZBETREIBER
	BILANZKOORDINATOR
	BILANZKREISVERANTWORTLICHER
	ENERGIESERVICEANBIETER
)
