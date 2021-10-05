package com

import (
	"time"
)

// Unterschrift ist die Modellierung einer Unterschrift, z.B. für Verträge, Angebote etc.
type Unterschrift struct {
	Name  string    `json:"name,omitempty" validate:"required"`   // Name des Unterschreibers
	Datum time.Time `json:"datum,omitempty" validate:"omitempty"` // Datum der Unterschrift
	Ort   string    `json:"ort,omitempty" validate:"omitempty"`   // Ort, an dem die Unterschrift geleistet wird
}
