package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"time"
)

// Zaehlerstand is the value of a meter for a specific point in time
type Zaehlerstand struct {
	Ablesedatum              time.Time                                         `json:"ablesedatum" validate:"required"`              // point in time
	Wertermittlungsverfahren wertermittlungsverfahren.Wertermittlungsverfahren `json:"wertermittlungsverfahren" validate:"required"` // type of the consumption
	Wert                     float32                                           `json:"wert" validate:"required" example:"17086"`     // value (of einheit)
	Einheit                  mengeneinheit.Mengeneinheit                       `json:"einheit" validate:"required" example:"KWH"`    // unit (associated to the wert)
	Brennwert                float32                                           `json:"brennwert,omitempty"`                          // used to calculate caloric Energy from gas volume
	Zustandszahl             float32                                           `json:"zustandszahl,omitempty"`                       // used to calculate caloric Energy from gas volume
}

// Zaehlerstaende is a table type of Zaehlerstand with defined Descending Sorting of the Ablesedatum for sort.Sort(...)
type Zaehlerstaende []Zaehlerstand

func (z Zaehlerstaende) Len() int {
	return len(z)
}

func (z Zaehlerstaende) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}

func (z Zaehlerstaende) Less(i, j int) bool {
	return z[i].Ablesedatum.After(z[j].Ablesedatum)
}
