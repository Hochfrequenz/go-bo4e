package com

import (
	"github.com/shopspring/decimal"
)

// Preisstaffel is part of a Preisposition and includes a price for a distinct set of parameters
type Preisstaffel struct {
	Einheitspreis    decimal.Decimal `json:"einheitspreis" validate:"required"` // Preis pro abgerechneter Mengeneinheit
	StaffelgrenzeVon decimal.Decimal `json:"staffelgrenzevon,omitempty"`        // Unterer Wert, ab dem die Staffel gilt.
	StaffelgrenzeBis decimal.Decimal `json:"staffelgrenzebis,omitempty"`        // Oberer Wert, bis zu dem die Staffel gilt.
	// sigmoidparameter TODO
}
