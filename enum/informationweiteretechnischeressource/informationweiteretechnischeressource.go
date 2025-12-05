package informationweiteretechnischeressource

// InformationWeitereTechnischeRessource is an enum
//
//go:generate stringer --type InformationWeitereTechnischeRessource
//go:generate jsonenums --type InformationWeitereTechnischeRessource
type InformationWeitereTechnischeRessource int

const (
	WEITERE_EINRICHTUNG_VORHANDEN InformationWeitereTechnischeRessource = iota + 1
	KEINE_WEITERE_EINRICHTUNG_VORHANDEN
)
