package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/shopspring/decimal"
)

// Preis is a price
type Preis struct {
	Wert       decimal.Decimal                   `json:"wert" validate:"required"`                 // Wert gibt die nominale Höhe des Preises an
	Einheit    waehrungseinheit.Waehrungseinheit `json:"einheit,omitempty" validate:"required"`    // Einheit ist die Währungseinheit für den Preis
	Bezugswert mengeneinheit.Mengeneinheit       `json:"bezugswert,omitempty" validate:"required"` // Bezugswert ist eine Angabe, für welche Bezugsgröße der Preis gilt. Z.B. kWh
	Status     preisstatus.Preisstatus           `json:"status,omitempty"`                         // Status gibt den Status des veröffentlichten Preises an
}
