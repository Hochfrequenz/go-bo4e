package technischeressourcenutzung

// Art und Nutzung der Technischen Ressource
//
//go:generate stringer --type TechnischeRessourceNutzung
//go:generate jsonenums --type TechnischeRessourceNutzung

type TechnischeRessourceNutzung int

const (
	STROMVERBRAUCHSART TechnischeRessourceNutzung = iota + 1 //Stromverbrauchsart
	STROMERZEUGUNGSART                                       //Stromerzeugungsart
	SPEICHER                                                 //Speicher
)
