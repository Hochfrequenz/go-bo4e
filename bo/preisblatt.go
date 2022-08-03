package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Das allgemeine Modell zur Abbildung von Preisen. Davon abgeleitet können, über die Zuordnung identifizierender
// Merkmale, spezielle Preisblatt-Varianten modelliert werden.
// Die jeweiligen Sätze von Merkmalen sind in der Grafik ergänzt worden und stellen jeweils eine Ausprägung für die
// verschiedenen Anwendungsfälle der Preisblätter dar.
type Preisblatt struct {
	Geschaeftsobjekt
	Bezeichnung     string                  `json:"bezeichnung" validate:"required"`           // Rechnungstitel ist die Bezeichnung für die vorliegende Rechnung.
	Sparte          sparte.Sparte           `json:"sparte,omitempty" validate:"required"`      // Preisblatt gilt für angegebene Sparte.
	Preisstatus     preisstatus.Preisstatus `json:"preisstatus,omitempty" validate:"required"` // Merkmal, das anzeigt, ob es sich um vorläufige oder endgültige Preise handelt
	Herausgeber     Marktteilnehmer         `json:"herausgeber" validate:"required"`           // Der Marktteilnehmer, der die Preise veröffentlicht hat. Details zum Marktteilnehmer
	Gueltigkeit     com.Zeitraum            `json:"gueltigkeit" validate:"required"`           // Der Zeitraum für den der Preis festgelegt ist
	Preispositionen []com.Preisposition     `json:"preispositionen" validate:"required"`       // Die einzelnen Positionen, die mit dem Preisblatt abgerechnet werden können. Z.B. Arbeitspreis, Grundpreis etc
}

func (_ Preisblatt) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
