package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/shopspring/decimal"
)

// Menge ist die Abbildung einer Menge mit Wert und Einheit.
type Menge struct {
	Wert    decimal.Decimal             `json:"wert" validate:"required"`    //Gibt den absoluten Wert der Menge an
	Einheit mengeneinheit.Mengeneinheit `json:"einheit" validate:"required"` //Gibt die Einheit zum jeweiligen Wert an
}
