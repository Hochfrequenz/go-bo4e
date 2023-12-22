package technischeressourcenutzung

//go:generate stringer --type Technischeressourcenutzung
//go:generate jsonenums --type Technischeressourcenutzung

type Technischeressourcenutzung int

const (
	STROMVERBRAUCHSART Technischeressourcenutzung = iota + 1 //Stromverbrauchsart
	STROMERZEUGUNGSART                                       //Stromerzeugungsart
	SPEICHER                                                 //Speicher
)
