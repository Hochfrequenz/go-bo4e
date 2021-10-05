package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

// Betrag wird dazu verwendet, Summenbeträge - beispielsweise in Angeboten und Rechnungen - als Geldbeträge abzubilden. Die Einheit ist dabei immer die Hauptwährung also Euro, Dollar etc
type Betrag struct {
	Wert     decimal.Decimal             `json:"wert" example:"18.36" validate:"required"`             // amount
	Waehrung waehrungscode.Waehrungscode `json:"waehrung,omitempty" example:"EUR" validate:"required"` // currency
}
