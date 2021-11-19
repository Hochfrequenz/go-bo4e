package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/netznutzungsabrechnungsgrundlage"
	"github.com/hochfrequenz/go-bo4e/enum/netznutzungsabrechnungsvariante"
	"github.com/hochfrequenz/go-bo4e/enum/netznutzungsvertragsart"
	"time"
)

// Vertragskonditionen ist eine Komponente zur Abbildung von Vertragskonditionen. Die Komponente wird sowohl im Vertrag als auch im Tarif verwendet.
type Vertragskonditionen struct {
	Beschreibung          string   `json:"beschreibung,omitempty"`          // Beschreibung ist ein Freitext zur Beschreibung der Konditionen, z.B. "Standardkonditionen Gas"
	AnzahlAbschlaege      int      `json:"anzahlAbschlaege,omitempty"`      // AnzahlAbschlaege enthält die Anzahl der vereinbarten Abschläge pro Jahr, z.B. 12
	Vertragslaufzeit      Zeitraum `json:"vertragslaufzeit,omitempty"`      // Vertragslaufzeit enthält den Zeitraum, über den dieser Vertrag läuft
	Kuendigungsfrist      Zeitraum `json:"kuendigungsfrist,omitempty"`      // Kuendigungsfrist ist die Frist, innerhalb derer der Vertrag gekündigt werden kann
	Vertragsverlaengerung Zeitraum `json:"vertragsverlaengerung,omitempty"` // Vertragsverlaengerung beschreibt, dass, falls der Vertrag nicht gekündigt wird, er sich automatisch um die angegebene Zeit verlängert
	Abschlagszyklus       Zeitraum `json:"abschlagszyklus,omitempty"`       // Abschlagszyklus sind die Zyklen, in denen Abschläge erstellt werden. Alternativ kann auch die Anzahl in den Konditionen angeben werden.

	// the fields below are not bo4e standard (yet)

	StartAbrechnungsjahr    time.Time                                       `json:"StartAbrechnungsjahr,omitempty"`                        // StartAbrechnungsjahr ist der Beginn des Abrechnungsjahres
	GeplanteTurnusablesung  *Zeitraum                                       `json:"geplanteTurnusablesung,omitempty" validate:"omitempty"` // GeplanteTurnusablesung beschreibt die geplante Turnusablesung
	TurnusablesungIntervall *int                                            `json:"TurnusablesungIntervall" validate:"omitempty,min=1"`    // Das TurnusablesungIntervall beschreibt die Anzahl Monate die zwischen zwei Ablesungen liegen
	Netznutzungsabrechnung  *Zeitraum                                       `json:"netznutzungsabrechnung,omitempty" validate:"omitempty"` // Netznutzungsabrechnung ist der Zeitraum der bei der Netznutzung abgerechnet wird
	Haushaltskunde          *bool                                           `json:"haushaltskunde,omitempty" validate:"omitempty"`         // Haushaltskunde ist true, wenn es sich um einen privat/Haushaltskunden handelt
	Netznutzungsvertragsart netznutzungsvertragsart.Netznutzungsvertragsart `json:"netznutzungsVertrag,omitempty"`                         // Netznutzungsvertragsart beschreibt, zwischen welchen Vertragspartnern der Vertrag besteht
	// the json tag is "netznutzungsvertrag" for consistency with the dotnet-implementation
	Netznutzungsabrechnungsvariante  netznutzungsabrechnungsvariante.Netznutzungsabrechnungsvariante   `json:"netznutzungsabrechnungsvariante,omitempty"`
	Netznutzungsabrechnungsgrundlage netznutzungsabrechnungsgrundlage.Netznutzungsabrechnungsgrundlage `json:"netznutzungsabrechnungsgrundlage,omitempty"`
}
