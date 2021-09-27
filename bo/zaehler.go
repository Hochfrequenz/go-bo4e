package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
	"time"
)

// Zaehler ist ein Modell für die Abbildung der Informationen zu einem Zähler
type Zaehler struct {
	BusinessObject
	Zaehlernummer      string                                `json:"zahlernummer" validate:"required,alphanum"` // Nummerierung des Zaehlers, vergeben durch den Messstellenbetreiber
	Sparte             sparte.Sparte                         `json:"sparte" validate:"required"`                // Unterscheidungsmöglichkeiten für die Sparte
	Zaehlerauspraegung zaehlerauspraegung.Zaehlerauspraegung `json:"zaehlerauspraegung" validate:"required"`    // Spezifikation die Richtung des Zählers betreffend
	Zaehlertyp         zaehlertyp.Zaehlertyp                 `json:"zaehlertyp" validate:"required"`            // Typisierung des Zählers
	Tarifart           tarifart.Tarifart                     `json:"tarifart" validate:"required"`              // Spezifikation bezüglich unterstützter Tarifarten
	Zaehlerkonstante   float32                               `json:"zaehlerkonstante,omitempty"`                // Zählerkonstante auf dem Zähler
	EichungBis         time.Time                             `json:"eichung_bis,omitempty"`                     // Bis zu diesem Datum ist der Zähler geeicht
	LetzteEichung      time.Time                             `json:"letzte_eichung,omitempty"`                  // Zu diesem Datum fand die letzte Eichprüfung des Zählers statt
	Zaehlwerke         []com.Zaehlwerk                       `json:"zaehlwerke" validate:"required,min=1"`      // Die Zählwerke des Zählers
	//Zaehlerhersteller  Geschaeftspartner                     `json:"zaehlerhersteller,omitempty"`            // Der Hersteller des Zählers
}
