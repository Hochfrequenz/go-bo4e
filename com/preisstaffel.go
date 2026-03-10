package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/shopspring/decimal"
)

// Preisstaffel is part of a Preisposition and includes a price for a distinct set of parameters
type Preisstaffel struct {
	Einheitspreis    decimal.Decimal `json:"einheitspreis" validate:"required"` // Preis pro abgerechneter Mengeneinheit
	StaffelgrenzeVon decimal.Decimal `json:"staffelgrenzeVon,omitempty"`        // Unterer Wert, ab dem die Staffel gilt (inklusiv)
	StaffelgrenzeBis decimal.Decimal `json:"staffelgrenzeBis,omitempty"`        // Oberer Wert, bis zu dem die Staffel gilt (exklusiv) TODO Inklusivität/Exklusivität noch nicht global abgestimmt!
	// sigmoidparameter TODO

	// Mengeneinheit is the unit for the annual quantity range of the Staffel
	// https://github.com/Hochfrequenz/BO4E-dotnet/blob/0.54.0/BO4E/COM/Preisstaffel.cs#L44
	Mengeneinheit *mengeneinheit.Mengeneinheit `json:"mengeneinheit,omitempty"`
}
