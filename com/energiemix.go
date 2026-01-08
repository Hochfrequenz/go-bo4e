package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/oekolabel"
	"github.com/hochfrequenz/go-bo4e/enum/oekozertifikat"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/shopspring/decimal"
)

// Energiemix beschreibt die Zusammensetzung der gelieferten Energie aus den verschiedenen Primärenergieformen.
type Energiemix struct {
	// Energiemixnummer ist die eindeutige Nummer zur Identifizierung des Energiemixes.
	Energiemixnummer *int `json:"energiemixnummer,omitempty"`

	// Energieart gibt an, ob es sich um Strom, Gas etc. handelt.
	Energieart *sparte.Sparte `json:"energieart,omitempty"`

	// Bezeichnung ist die Bezeichnung des Energiemix.
	Bezeichnung *string `json:"bezeichnung,omitempty"`

	// Bemerkung ist eine Bemerkung zum Energiemix.
	Bemerkung *string `json:"bemerkung,omitempty"`

	// Gueltigkeitsjahr ist das Jahr, für das der Energiemix gilt.
	Gueltigkeitsjahr *int `json:"gueltigkeitsjahr,omitempty"`

	// CO2Emission ist die Höhe des erzeugten CO2-Ausstoßes in g/kWh.
	CO2Emission *decimal.Decimal `json:"cO2Emission,omitempty"`

	// Atommuell ist die Höhe des erzeugten Atommülls in g/kWh.
	Atommuell *decimal.Decimal `json:"atommuell,omitempty"`

	// Oekozertifikat sind Zertifikate für den Energiemix.
	Oekozertifikat []oekozertifikat.Oekozertifikat `json:"oekozertifikat,omitempty"`

	// Oekolabel sind Ökolabels für den Energiemix.
	Oekolabel []oekolabel.Oekolabel `json:"oekolabel,omitempty"`

	// OekoTopTen gibt an, ob der Versorger zu den Öko Top Ten gehört.
	OekoTopTen *bool `json:"oekoTopTen,omitempty"`

	// Website ist die Internetseite, auf der die Strommixdaten veröffentlicht sind.
	Website *string `json:"website,omitempty"`

	// Anteil enthält die Anteile der jeweiligen Erzeugungsart.
	Anteil []Energieherkunft `json:"anteil,omitempty"`
}
