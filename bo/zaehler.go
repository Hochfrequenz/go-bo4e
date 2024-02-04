package bo

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/fernschaltung"
	"github.com/hochfrequenz/go-bo4e/enum/geraetemerkmal"
	"github.com/hochfrequenz/go-bo4e/enum/messwerterfassung"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
	"github.com/shopspring/decimal"
)

// Zaehler ist ein Modell für die Abbildung der Informationen zu einem Zähler
type Zaehler struct {
	Geschaeftsobjekt
	Zaehlernummer      string                                 `json:"zaehlernummer,omitempty" validate:"required,alphanum"` // Zaehlernummer ist eine Nummerierung des Zaehlers, vergeben durch den Messstellenbetreiber
	Sparte             sparte.Sparte                          `json:"sparte,omitempty" validate:"required"`                 // Sparte ist eine Unterscheidungsmöglichkeit für die Sparte
	Zaehlerauspraegung *zaehlerauspraegung.Zaehlerauspraegung `json:"zaehlerauspraegung,omitempty"`                         // Zaehlerauspraegung ist eine Spezifikation die Richtung des Zählers betreffend
	Zaehlertyp         *zaehlertyp.Zaehlertyp                 `json:"zaehlertyp,omitempty" validate:"required"`             // Zaehlertyp erlaubt eine Typisierung des Zählers
	Tarifart           *tarifart.Tarifart                     `json:"tarifart,omitempty"`                                   // Tarifart erlaubt eine Spezifikation bezüglich unterstützter Tarifarten
	Zaehlerkonstante   decimal.NullDecimal                    `json:"zaehlerkonstante,omitempty"`                           // Zaehlerkonstante ist die Zählerkonstante auf dem Zähler
	EichungBis         time.Time                              `json:"eichungBis,omitempty"`                                 // EichungBis ist das exklusive Enddatum bis zu dem der Zähler geeicht ist
	LetzteEichung      time.Time                              `json:"letzteEichung,omitempty"`                              // LetzteEichung ist das Datum, an dem die letzte Eichprüfung des Zählers stattfand
	Zaehlwerke         []com.Zaehlwerk                        `json:"zaehlwerke,omitempty" validate:"required,min=1"`       // Zaehlwerke sind die Zählwerke des Zählers
	Zaehlerhersteller  *Geschaeftspartner                     `json:"zaehlerhersteller,omitempty"`                          // Zaehlerhersteller ist der Hersteller des Zählers
	Zaehlergroesse     *geraetemerkmal.Geraetemerkmal         `json:"zaehlergroesse,omitempty"`                             // Zaehlergroesse
	// the following fields are not part of the official bo4e standard (yet)

	Fernschaltung     *fernschaltung.Fernschaltung         `json:"fernschaltung,omitempty"`     // Fernschaltung is set to VORHANDEN if there is a fernschaltung
	Messwerterfassung *messwerterfassung.Messwerterfassung `json:"messwerterfassung,omitempty"` // Messwerterfassung describes if meter readings have to happen manually
}

func (_ Zaehler) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
