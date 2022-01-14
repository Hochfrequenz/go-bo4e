package com

import (
	"github.com/shopspring/decimal"
)

// Preisstaffel is part of a Preisposition and includes a price for a distinct set of parameters
type Preisstaffel struct {
	Einheitspreis    decimal.Decimal `json:"einheitspreis" validate:"required"` // Preis pro abgerechneter Mengeneinheit
	StaffelgrenzeVon decimal.Decimal `json:"staffelgrenzeVon,omitempty"`        // Unterer Wert, ab dem die Staffel gilt (inklusiv)
	StaffelgrenzeBis decimal.Decimal `json:"staffelgrenzeBis,omitempty"`        // Oberer Wert, bis zu dem die Staffel gilt (exklusiv) TODO Inklusivität/Exklusivität noch nicht global abgestimmt!
	// sigmoidparameter TODO
}
