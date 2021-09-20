package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"time"
)

// Verbrauch is a a consumption during a specific time frame
type Verbrauch struct {
	Startdatum               time.Time                                         `json:"startdatum" validate:"omitempty"`                      // inclusive start
	Enddatum                 time.Time                                         `json:"enddatum" validate:"omitempty,gtfield=Startdatum"`     // exclusive end
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren" validate:"required"`         // type of the consumption
	Obiskennzahl             string                                            `json:"obiskennzahl" validate:"required" example:"1-0:1.8.1"` // OBIS Kennzahl // (?<power>(1)-((?:[0-5]?[0-9])|(?:6[0-5])):((?:[1-8]|99))\.((?:6|8|9|29))\.([0-9]{1,2}))||(?<gas>)(7)-((?:[0-5]?[0-9])|(?:6[0-5])):(.{1,2})\.(.{1,2})\.([0-9]{1,2})
	Wert                     float32                                           `json:"wert" validate:"required" example:"17"`                // value (of einheit)
	Einheit                  mengeneinheit.Mengeneinheit                       `json:"einheit" validate:"required" example:"KWH"`            // unit (associated to the wert)
}
