package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/ersatzwertbildungsverfahren"
	"github.com/hochfrequenz/go-bo4e/enum/gasqualitaet"
	"github.com/hochfrequenz/go-bo4e/enum/grundlageenergiemenge"
	"github.com/hochfrequenz/go-bo4e/enum/korrekturgrund"
	"github.com/hochfrequenz/go-bo4e/enum/plausibilisierungshinweis"
	"github.com/hochfrequenz/go-bo4e/enum/tarif"
)

// Statuszusatz defines extended information to a Verbrauch
type Statuszusatz struct {
	// Gasqualitaet describes the gas quality in case of gas rearrangements
	Gasqualitaet *gasqualitaet.Gasqualitaet `json:"gasqualitaet,omitempty" example:"H_GAS"`

	// Plausibilisierungshinweis describes the plausibility of a consumption
	Plausibilisierungshinweis *plausibilisierungshinweis.Plausibilisierungshinweis `json:"plausibilisierungshinweis,omitempty" example:"Z83_KUNDENSELBSTABLESUNG"`

	// GrundlageEnergiemenge describes hints to the foundation of a consumption
	GrundlageEnergiemenge *grundlageenergiemenge.GrundlageEnergiemenge `json:"grundlageenergiemenge,omitempty" example:"ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT"`

	// Korrekturgrund describes reasons on why a correction has to be made
	Korrekturgrund *korrekturgrund.Korrekturgrund `json:"korrekturgrund,omitempty" example:"KEIN_ZUGANG"`

	// ErsatzwertBildungsverfahren describes the method used to generate replacement values
	Ersatzwertbildungsverfahren *ersatzwertbildungsverfahren.Ersatzwertbildungsverfahren `json:"ersatzwertbildungsverfahren,omitempty" example:"INTERPOLATION"`

	// Tarif describes the tarif used, if there is any
	Tarif *tarif.Tarif `json:"tarif,omitempty" example:"T1_TARIF1"`
}
