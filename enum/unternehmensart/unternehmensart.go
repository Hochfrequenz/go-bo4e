package unternehmensart

// Unternehmensart kategorisiert Unternehmen nach Ihrer Marktrolle (basierend auf dem PARTIN-Datenmodell)
//
//go:generate stringer --type Unternehmensart
//go:generate jsonenums --type Unternehmensart
type Unternehmensart int

const (
	LIEFERANT                  Unternehmensart = iota + 1 // LIEFERANT is a supplier
	NETZBETREIBER                                         // NETZBETREIBER is a grid operator
	MESSSTELLENBETREIBER                                  // MESSSTELLENBETREIBER is a metering point operator
	UEBERTRAGUNGSNETZBETREIBER                            // UEBERTRAGUNGSNETZBETREIBER is a transmission system operator
	BILANZKOORDINATOR                                     // BILANZKOORDINATOR is a balancing coordinator
	BILANZKREISVERANTWORTLICHER                           // BILANZKREISVERANTWORTLICHER is a balance responsible party
	ENERGIESERVICEANBIETER                                // ENERGIESERVICEANBIETER is an energy service provider
)
