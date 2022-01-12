package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Rechnung ist ein Modell für die Abbildung von Rechnungen im Kontext der Energiewirtschaft. Ausgehend von diesem Basismodell werden weitere spezifische Formen abgeleitet.
type Preisblatt struct {
	Geschaeftsobjekt
	Bezeichnung     string                  `json:"bezeichnung" validate:"required"`     // Rechnungstitel ist die Bezeichnung für die vorliegende Rechnung.
	Sparte          sparte.Sparte           `json:"sparte" validate:"required"`          // Preisblatt gilt für angegebene Sparte. Details siehe ENUM Sparte. TODO Klärung dieses Feld zu entfernen, da Herausgeber, sowie Artikel bereits Spartenscharf sind
	Preisstatus     preisstatus.Preisstatus `json:"preisstatus" validate:"required"`     // Merkmal, das anzeigtr, ob es sich um vorläufige oder endgültige Preise handelt. Details siehe ENUM Preisstatus
	Herausgeber     Marktteilnehmer         `json:"herausgeber" validate:"required"`     // Der Marktteilnehmer, der die Preise veröffentlicht hat. Details zum Marktteilnehmer siehe BO Marktteilnehmer
	Gueltigkeit     com.Zeitraum            `json:"herausgeber" validate:"required"`     // Der Zeitraum für den der Preis festgelegt ist. Siehe COM Zeitraum
	Preispositionen []com.Preisposition     `json:"preispositionen" validate:"required"` // Die einzelnen Positionen, die mit dem Preisblatt abgerechnet werden können. Z.B. Arbeitspreis, Grundpreis etc. Details siehe COM Preisposition
}
