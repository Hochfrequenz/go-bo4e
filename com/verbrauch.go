package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"time"
)

// Verbrauch is a a consumption during a specific time frame
type Verbrauch struct {
	Startdatum               time.Time                                         `json:"startdatum,omitempty" validate:"omitempty"`                      // inclusive start
	Enddatum                 time.Time                                         `json:"enddatum,omitempty" validate:"omitempty,gtfield=Startdatum"`     // exclusive end
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren,omitempty" validate:"required"`         // type of the consumption
	Obiskennzahl             string                                            `json:"obiskennzahl,omitempty" validate:"required" example:"1-0:1.8.1"` // OBIS Kennzahl // (?<power>(1)-((?:[0-5]?[0-9])|(?:6[0-5])):((?:[1-8]|99))\.((?:6|8|9|29))\.([0-9]{1,2}))||(?<gas>)(7)-((?:[0-5]?[0-9])|(?:6[0-5])):(.{1,2})\.(.{1,2})\.([0-9]{1,2})
	Wert                     decimal.Decimal                                   `json:"wert,omitempty" validate:"required" example:"17"`                // value (of einheit)
	Einheit                  mengeneinheit.Mengeneinheit                       `json:"einheit,omitempty" validate:"required" example:"KWH"`            // unit (associated to the wert)
}
