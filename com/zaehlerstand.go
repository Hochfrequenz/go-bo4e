package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"time"
)

// Zaehlerstand is the value of a meter for a specific point in time (non-standard)
type Zaehlerstand struct {
	Ablesedatum              time.Time                                         `json:"ablesedatum" validate:"required"`                        // Ablesedatum is a point in time
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren,omitempty" validate:"required"` // Wertermittlungsverfahren is the type of the consumption
	Wert                     decimal.Decimal                                   `json:"wert" validate:"required" example:"17086"`               // Wert is the value (of Einheit)
	Einheit                  mengeneinheit.Mengeneinheit                       `json:"einheit,omitempty" validate:"required" example:"KWH"`    // Einheit ist the unit (associated to the wert)
	Brennwert                decimal.NullDecimal                               `json:"brennwert,omitempty"`                                    // Brennwert is used to calculate caloric Energy from gas volume
	Zustandszahl             decimal.NullDecimal                               `json:"zustandszahl,omitempty"`                                 // Zustandszahl is used to calculate caloric Energy from gas volume
}
