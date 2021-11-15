package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"github.com/shopspring/decimal"
)

// Zeitreihenwertkompakt ist eine Abbildung eines kompakten Zeitreihenwertes in dem ausschliesslich der Wert und Statusinformationen stehen.
type Zeitreihenwertkompakt struct {
	Wert         decimal.Decimal                           `json:"wert" validate:"required" example:"17"` // Wert ist der der im Zeitintervall gültige Wert.
	Status       messwertstatus.Messwertstatus             `json:"status,omitempty"`                      // Der Status gibt an, wie der Wert zu interpretieren ist, z.B. in Berechnungen.
	Statuszusatz messwertstatuszusatz.Messwertstatuszusatz `json:"statuszusatz,omitempty"`                // Statuszusatz ist eine Zusatzinformation zum Status, beispielsweise ein Grund für einen fehlenden Wert.
}
