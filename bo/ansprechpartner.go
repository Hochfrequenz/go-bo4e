package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/titel"
)

// Ansprechpartner sind ansprechbare Geschäftspartner (Personen)
type Ansprechpartner struct {
	Geschaeftsobjekt
	Anrede anrede.Anrede `json:"anrede,omitempty"` // Anrede ist die mögliche Anrede des Ansprechpartners
	// IndividuelleAnrede ermöglicht im Falle einer nicht standardisierten Anrede (7=anrede.INDIVIDUELL) eine frei definierbare Anrede vorzugeben.
	IndividuelleAnrede string               `json:"individuelleAnrede,omitempty" validate:"required_if=Anrede 7" example:"Sehr geehrte Frau Müller, sehr geehrter Herr Dr. Müller"`
	Titel              titel.Titel          `json:"titel,omitempty"`                                    // Titel ist ein möglicher Titel des Ansprechpartners
	Vorname            string               `json:"vorname,omitempty"`                                  // Vorname ist der Vorname des Ansprechpartners
	Nachname           string               `json:"nachname,omitempty" validate:"required"`             // Nachname ist der Familienname des Ansprechpartners
	EMailAdresse       string               `json:"eMailAdresse,omitempty" validate:"omitempty,email"`  // EMailAdresse des Ansprechpartners
	Kommentar          string               `json:"kommentar,omitempty"`                                // Kommentar sind zusätzlich Freitextinformationen zum Ansprechpartner
	Geschaeftspartner  *Geschaeftspartner   `json:"geschaeftspartner,omitempty" validate:"required"`    // Geschaeftspartner ist der Geschäftspartner für den der Ansprechpartner modelliert wird
	Adresse            *com.Adresse         `json:"adresse,omitempty"`                                  // Adresse des Ansprechpartners, falls diese von der Adresse des Geschaeftspartner s abweicht
	Rufnummern         []com.Rufnummer      `json:"rufnummern,omitempty" validate:"omitempty,dive"`     // Rufnummern ist eine Liste der Telefonnummern, unter denen der Ansprechpartner erreichbar ist
	Zustaendigkeiten   []com.Zustaendigkeit `json:"zustaendigkeit,omitempty" validate:"omitempty,dive"` // Zustaendigkeiten ist eine Liste der Abteilungen und Zuständigkeiten des Ansprechpartner
	// the json tags are singular "zustaendigkeit"/"rufnummer" instead of plural "zustaendigkeiten"/"rufnummern" for compatability reasons
	// https://github.com/Hochfrequenz/bo4e-modification-proposals/blob/master/markdown/ansprechpartner_rufnummern.md

}

func (_ Ansprechpartner) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
