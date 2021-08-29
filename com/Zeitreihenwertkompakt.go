package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
)

// Zeitreihenwertkompakt ist eine Abbildung eines kompakten Zeitreihenwertes in dem ausschliesslich der Wert und Statusinformationen stehen.
type Zeitreihenwertkompakt struct {
	Wert         float32                                   `json:"wert" validate:"required" example:"17"` // 	Der im Zeitintervall gültige Wert.
	Status       messwertstatus.Messwertstatus             `json:"status,omitempty"`                      // Der Status gibt an, wie der Wert zu interpretieren ist, z.B. in Berechnungen.
	Statuszusatz messwertstatuszusatz.Messwertstatuszusatz `json:"statuszusatz,omitempty"`                // Eine Zusatzinformation zum Status, beispielsweise ein Grund für einen fehlenden Wert.
}
