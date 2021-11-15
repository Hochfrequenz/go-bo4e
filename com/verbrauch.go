package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"time"
)

// Verbrauch is a a consumption during a specific time frame
type Verbrauch struct {
	Startdatum               time.Time                                         `json:"startdatum" validate:"omitempty"`                        // Startdatum is an inclusive start
	Enddatum                 time.Time                                         `json:"enddatum" validate:"omitempty,gtfield=Startdatum"`       // Enddatum is an exclusive end
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren,omitempty" validate:"required"` // Wertermittlungsverfahren is the type of the consumption
	Obiskennzahl             string                                            `json:"obiskennzahl" validate:"required" example:"1-0:1.8.1"`   // Obiskennzahl is the OBIS Kennzahl // (?<power>(1)-((?:[0-5]?[0-9])|(?:6[0-5])):((?:[1-8]|99))\.((?:6|8|9|29))\.([0-9]{1,2}))||(?<gas>)(7)-((?:[0-5]?[0-9])|(?:6[0-5])):(.{1,2})\.(.{1,2})\.([0-9]{1,2})
	Wert                     decimal.Decimal                                   `json:"wert" validate:"required" example:"17"`                  // Wert is the value (of einheit)
	Einheit                  mengeneinheit.Mengeneinheit                       `json:"einheit,omitempty" validate:"required" example:"KWH"`    // Einheit is the unit (associated to the wert)
}
