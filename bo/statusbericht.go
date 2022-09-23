package bo

import (
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"time"

	"github.com/hochfrequenz/go-bo4e/com"

	"github.com/hochfrequenz/go-bo4e/enum/berichtstatus"
)

// Statusbericht kann verwendet werden um CONTRL oder APREAK Nachrichten auszutauschen
type Statusbericht struct {
	Geschaeftsobjekt
	Status          berichtstatus.BerichtStatus `json:"status,omitempty"  validate:"required"` // Status des Berichtes (Fehlerhaft, Erfolgreich)
	Pruefgegenstand *string                     `json:"pruefgegenstand,omitempty"`             //  Das geprüfte Dokument, z.B. die Referenz auf die EDIFACT-Nachricht die geprüft / beanstandet wurde
	DatumPruefung   time.Time                   `json:"datumPruefung"`                         // Pruefdatum (wann wurde der Pruefgegenstand geprüft)
	Fehler          *com.Fehler                 `json:"fehler,omitempty"`                      // Liste der Fehler
}

func (s Statusbericht) GetDefaultJsonTags() []string {
	panic("implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
func (_ Statusbericht) GetBoTyp() botyp.BOTyp {
	// this is useful for generics to work
	return botyp.STATUSBERICHT
}
