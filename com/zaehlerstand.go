package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"time"
)

// Zaehlerstand is the value of a meter for a specific point in time (non-standard)
type Zaehlerstand struct {
	Ablesedatum              time.Time                                         `json:"ablesedatum,omitempty" validate:"required"`              // point in time
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren,omitempty" validate:"required"` // type of the consumption
	Wert                     decimal.Decimal                                   `json:"wert,omitempty" validate:"required" example:"17086"`     // value (of einheit)
	Einheit                  mengeneinheit.Mengeneinheit                       `json:"einheit,omitempty" validate:"required" example:"KWH"`    // unit (associated to the wert)
	Brennwert                decimal.NullDecimal                               `json:"brennwert,omitempty"`                                    // used to calculate caloric Energy from gas volume
	Zustandszahl             decimal.NullDecimal                               `json:"zustandszahl,omitempty"`                                 // used to calculate caloric Energy from gas volume
}
