package unternehmensart

// Unternehmensart kategorisiert Unternehmen nach ihrer Marktrolle (basierend auf dem PARTIN-Datenmodell)
//
//go:generate stringer --type Unternehmensart
//go:generate jsonenums --type Unternehmensart
type Unternehmensart int

const (
	// LIEFERANT: Energielieferant
	LIEFERANT Unternehmensart = iota + 1
	// NETZBETREIBER: Verteilnetzbetreiber
	NETZBETREIBER
	// MESSSTELLENBETREIBER: Messstellenbetreiber
	MESSSTELLENBETREIBER
	// UEBERTRAGUNGSNETZBETREIBER: Übertragungsnetzbetreiber
	UEBERTRAGUNGSNETZBETREIBER
	// BILANZKOORDINATOR: Bilanzkoordinator
	BILANZKOORDINATOR
	// BILANZKREISVERANTWORTLICHER: Bilanzkreisverantwortlicher
	BILANZKREISVERANTWORTLICHER
	// ENERGIESERVICEANBIETER: Energiedienstleistungsunternehmen
	ENERGIESERVICEANBIETER
)
