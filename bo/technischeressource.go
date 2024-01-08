package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/emobilitaetsart"
	"github.com/hochfrequenz/go-bo4e/enum/erzeugungsart"
	"github.com/hochfrequenz/go-bo4e/enum/speicherart"
	"github.com/hochfrequenz/go-bo4e/enum/technischeressourcenutzung"
	"github.com/hochfrequenz/go-bo4e/enum/technischeressourceverbrauchsart"
	"github.com/hochfrequenz/go-bo4e/enum/waermenutzung"
)

type TechnischeRessource struct {
	Geschaeftsobjekt
	TechnischeRessourceId            *string                                                            `json:"technischeRessourceId" validate:"required"`                                                                             //Identifikationsnummer einer TechnischeRessource
	VorgelagerteMesslokationsId      *string                                                            `json:"vorgelagerteMesslokationsId,omitempty" example:"DE00713739359S0000000000001222221" validate:"alphanum,required,len=33"` // Vorgelagerte Messlokation ID
	ZugeordneteMarktlokationsId      *string                                                            `json:"zugeordneteMarktlokationsId,omitempty" example:"20072281644"`                                                           // ZugeordneteMarktlokationsId Messlokation ID
	ZugeordneteSteuerbareRessourceId *string                                                            `json:"zugeordneteSteuerbareRessourceId" example:"20072281644"`                                                                // Referenz auf die der Technischen Ressource zugeordneten Steuerbaren Ressource
	NennleistungAufnahme             *com.Menge                                                         `json:"nennleistungAufnahme" example:"QTY+Z43:100:KWT"`                                                                        // Nennleistung (Aufnahme)
	NennleistungAbgabe               *com.Menge                                                         `json:"nennleistungAbgabe" example:"QTY+Z44:100:KWT"`                                                                          //Nennleistung (Abgabe)
	Speicherkapazitaet               *com.Menge                                                         `json:"speicherkapazitaet" example:"QTY+Z42:100:KWH"`                                                                          //Speicherkapazität
	TechnischeRessourceNutzung       *technischeressourcenutzung.TechnischeRessourceNutzung             `json:"technischeRessourcenutzung" example:"CCI+Z17"`                                                                          //Art und Nutzung der Technischen Ressource
	Verbrauchsart                    *technischeressourceverbrauchsart.TechnischeRessourceVerbrauchsart `json:"verbrauchsart" example:"CAV+Z64" `                                                                                      //Verbrauchsart der Technischen Ressource
	Waermenutzung                    *waermenutzung.Waermenutzung                                       `json:"waermenutzung" example:"CAV+Z56"`                                                                                       //Wärmenutzung
	EMobilitaetsart                  *emobilitaetsart.EMobilitaetsart                                   `json:"emobilitaetsart" example:"CAV+Z87"`                                                                                     //Art der E-Mobilität  Das Segment dient dazu, im Falle der E-Mobilität eine genauere Angabe über die Art der E-Mobilität zu definieren
	Erzeugungsart                    *erzeugungsart.Erzeugungsart                                       `json:"erzeugungsart" example:"CAV+ZF5"`                                                                                       //Art der Erzeugung der Energie
	Speicherart                      *speicherart.Speicherart                                           `json:"speicherart" example:"CAV+ZF7"`                                                                                         //Art der speicher. Details <see cref="ENUM.Speicherart" />
}

func (_ TechnischeRessource) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
