package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
)

// Betrag wird dazu verwendet, Summenbeträge - beispielsweise in Angeboten und Rechnungen - als Geldbeträge abzubilden. Die Einheit ist dabei immer die Hauptwährung also Euro, Dollar etc
type Betrag struct {
	Wert     float32                     `json:"wert" example:"18.36"`                       // amount; this is not required (meaning 0 is valid)
	Waehrung waehrungscode.Waehrungscode `json:"waehrung" example:"EUR" validate:"required"` // currency
}
