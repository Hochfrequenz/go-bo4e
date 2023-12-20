package bo

type Steuerbareressource struct {
	BusinessObject
	steuerbareRessourceId             *string                           `json:"steuerbareRessourceId" validate:"required"` // Identifikationsnummer einer SteuerbareRessource
	steuerkanalsLeistungsbeschreibung steuerkanalsLeistungsbeschreibung `json:"steuerkanalsLeistungsbeschreibung"`         // Leistungsbeschreibung des Steuerkanals
	zugeordnetMSBCodeNr               *string                           `json:"zugeordnetMSBCodeNr"`                       // Angabe des Messstellenbetreibers, der der Steuerbaren Ressource zugeordnet ist
}
