package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/shopspring/decimal"
)

// Preis is a price
type Preis struct {
	Wert       decimal.Decimal                   `json:"wert" validate:"required"`                 // Gibt die nominale Höhe des Preises an
	Einheit    waehrungseinheit.Waehrungseinheit `json:"einheit,omitempty" validate:"required"`    // Währungseinheit für den Preis
	Bezugswert mengeneinheit.Mengeneinheit       `json:"bezugswert,omitempty" validate:"required"` // Angabe, für welche Bezugsgröße der Preis gilt. Z.B. kWh
	Status     preisstatus.Preisstatus           `json:"status,omitempty"`                         // Gibt den Status des veröffentlichten Preises an, Details siehe ENUM Preisstatus
}
