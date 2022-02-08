package com

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/enum/ersatzwertbildungsverfahren"
	"github.com/hochfrequenz/go-bo4e/enum/gasqualitaet"
	"github.com/hochfrequenz/go-bo4e/enum/grundlageenergiemenge"
	"github.com/hochfrequenz/go-bo4e/enum/korrekturgrund"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/plausibilisierungshinweis"
	"github.com/hochfrequenz/go-bo4e/enum/tarif"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
)

// Verbrauch is a consumption during a specific time frame
type Verbrauch struct {
	// Startdatum is an inclusive start
	Startdatum time.Time `json:"startdatum,omitempty" validate:"omitempty"`

	// Enddatum is an exclusive end
	Enddatum time.Time `json:"enddatum,omitempty" validate:"omitempty,gtfield=Startdatum"`

	// Wertermittlungsverfahren is the type of the consumption
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren,omitempty" validate:"required"`

	// Obiskennzahl is the OBIS Kennzahl
	// (?<power>(1)-((?:[0-5]?[0-9])|(?:6[0-5])):((?:[1-8]|99))\.((?:6|8|9|29))\.([0-9]{1,2}))||(?<gas>)(7)-((?:[0-5]?[0-9])|(?:6[0-5])):(.{1,2})\.(.{1,2})\.([0-9]{1,2})
	Obiskennzahl string `json:"obiskennzahl,omitempty" validate:"required" example:"1-0:1.8.1"`

	// Wert is the value (of einheit)
	Wert decimal.Decimal `json:"wert,omitempty" validate:"required" example:"17"`

	// Einheit is the unit (associated to the wert)
	Einheit mengeneinheit.Mengeneinheit `json:"einheit,omitempty" validate:"required" example:"KWH"`

	// Gasqualitaet optionally describes the gas quality in case of gas rearrangements
	Gasqualitaet *gasqualitaet.Gasqualitaet `json:"gasqualitaet,omitempty" example:"H_GAS"`

	// Plausibilisierungshinweis optionally describes the plausibility of a consumption
	Plausibilisierungshinweis *plausibilisierungshinweis.Plausibilisierungshinweis `json:"plausibilisierungshinweis,omitempty" example:"Z83_KUNDENSELBSTABLESUNG"`

	// GrundlageEnergiemenge optionally describes hints to the foundation of a consumption
	GrundlageEnergiemenge *grundlageenergiemenge.GrundlageEnergiemenge `json:"grundlageenergiemenge,omitempty" example:"ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT"`

	// Korrekturgrund optionally describes reasons on why a correction has to be made
	Korrekturgrund *korrekturgrund.Korrekturgrund `json:"korrekturgrund,omitempty" example:"KEIN_ZUGANG"`

	// ErsatzwertBildungsverfahren optionally describes the method used to generate replacement values
	ErsatzwertBildungsverfahren *ersatzwertbildungsverfahren.Ersatzwertbildungsverfahren `json:"ersatzwertbildungsverfahren,omitempty" example:"Z88_VERGLEICHSMESSUNGGEEICHT"`

	// Tarif optionally describes the tarif used, if there is any
	Tarif *tarif.Tarif `json:"tarif,omitempty" example:"T1_TARIF1"`
}
