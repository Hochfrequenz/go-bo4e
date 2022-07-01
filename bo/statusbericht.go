package bo

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/enum/berichtstatus"
)

// Statusbericht kann verwendet werden um CONTRL oder APREAK Nachrichten auszutauschen
type Statusbericht struct {
	Geschaeftsobjekt
	Status          berichtstatus.BerichtStatus `json:"status"  validate:"required"` // Status des Berichtes (Fehlerhaft, Erfolgreich)
	Pruefgegenstand *string                     `json:"pruefgegenstand,omitempty"`   //  Das geprüfte Dokument, z.B. die Referenz auf die EDIFACT-Nachricht die geprüft / beanstandet wurde
	DatumPruefung   time.Time                   `json:"datumPruefung"`               // Pruefdatum (wann wurde der Pruefgegenstand geprüft)
	Fehler          *string                     `json:"fehler,omitempty"`            // Liste der Fehler
}
