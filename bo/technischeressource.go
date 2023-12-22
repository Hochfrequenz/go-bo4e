package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
)

type TechnischeRessource struct {
	Geschaeftsobjekt
	TechnischeRessourceId            *string   `json:"technischeRessourceId" validate:"required"`                                                                             //Identifikationsnummer einer TechnischeRessource
	VorgelagerteMesslokationsId      string    `json:"vorgelagerteMesslokationsId,omitempty" example:"DE00713739359S0000000000001222221" validate:"alphanum,required,len=33"` // Vorgelagerte Messlokation ID
	ZugeordneteMarktlokationsId      string    `json:"zugeordneteMarktlokationsId,omitempty" example:"20072281644" validate:"required"`                                       // ZugeordneteMarktlokationsId Messlokation ID
	ZugeordneteSteuerbareRessourceId string    `json:"zugeordneteSteuerbareRessourceId" example:"20072281644" validate:"required"`                                            // Referenz auf die der Technischen Ressource zugeordneten Steuerbaren Ressource
	NennleistungAufnahme             com.Menge `json:"nennleistungAufnahme" example:"QTY+Z43:100:KWT" validate:"required"`                                                    // Nennleistung (Aufnahme)
	NennleistungAbgabe               com.Menge `json:"nennleistungAbgabe" example:"QTY+Z44:100:KWT" validate:"required"`                                                      //Nennleistung (Abgabe)
	Speicherkapazitaet               com.Menge `json:"speicherkapazitaet" example:"QTY+Z42:100:KWH" validate:"required"`                                                      //Speicherkapazität
	TechnischeRessourceNutzung
}
