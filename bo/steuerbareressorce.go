package bo

import (
	"github.com/hochfrequenz/go-bo4e/enum/steuerkanalsleistungsbeschreibung"
)

type Steuerbareressource struct {
	Geschaeftsobjekt
	SteuerbareRessourceId             string                                                               `json:"steuerbareRessourceId" validate:"required"` // Identifikationsnummer einer SteuerbareRessource
	SteuerkanalsLeistungsbeschreibung *steuerkanalsleistungsbeschreibung.Steuerkanalsleistungsbeschreibung `json:"*steuerkanalsLeistungsbeschreibung"`        // Leistungsbeschreibung des Steuerkanals
	ZugeordnetMSBCodeNr               string                                                               `json:"zugeordnetMSBCodeNr"`                       // Angabe des Messstellenbetreibers, der der Steuerbaren Ressource zugeordnet ist
}

func (_ Steuerbareressource) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
