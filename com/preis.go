package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
)

// Preis is a price
type Preis struct {
	Wert       float32                           `json:"wert" validate:"required"`       // Gibt die nominale Höhe des Preises an
	Einheit    waehrungseinheit.Waehrungseinheit `json:"einheit" validate:"required"`    // Währungseinheit für den Preis
	Bezugswert mengeneinheit.Mengeneinheit       `json:"bezugswert" validate:"required"` // Angabe, für welche Bezugsgröße der Preis gilt. Z.B. kWh
	Status     preisstatus.Preisstatus           `json:"status,omitempty"`               // Gibt den Status des veröffentlichten Preises an, Details siehe ENUM Preisstatus
}
