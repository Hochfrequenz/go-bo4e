package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/bilanzierungsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/dienstleistungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
)

// PreisblattMessung ist eine Preisblatt-Variante mit zusätzlichen Merkmalen für Messpreise
type PreisblattMessung struct {
	Preisblatt
	Bilanzierungsmethode    bilanzierungsmethode.Bilanzierungsmethode `json:"bilanzierungsmethode,omitempty"`
	Messebene               netzebene.Netzebene                       `json:"messebene,omitempty"`
	InklusiveDienstleistung []dienstleistungstyp.Dienstleistungstyp   `json:"inklusiveDienstleistung,omitempty"`
	Basisgeraet             com.Geraeteeigenschaften                  `json:"basisgeraet,omitempty"`
	InklusiveGeraete        []com.Geraeteeigenschaften                `json:"inklusiveGeraete,omitempty"`
	Herausgeber             Marktteilnehmer                           `json:"herausgeber"`
}
