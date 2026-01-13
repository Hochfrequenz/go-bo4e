package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/shopspring/decimal"
)

// PhysikalischerWert ist eine Kombination aus einem Wert und einer Einheit.
type PhysikalischerWert struct {
	// Wert ist der numerische Wert.
	Wert decimal.Decimal `json:"wert,omitempty" validate:"required"`

	// Einheit ist die Einheit des Werts.
	Einheit mengeneinheit.Mengeneinheit `json:"einheit,omitempty" validate:"required"`
}
