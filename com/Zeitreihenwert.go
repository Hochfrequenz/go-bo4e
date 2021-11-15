package com

import (
	"time"
)

// Zeitreihenwert ist eine Abbildung eines Zeitreihenwertes bestehend aus Zeitraum, Wert und Statusinformationen.
type Zeitreihenwert struct {
	Zeitreihenwertkompakt
	DatumUhrzeitVon time.Time `json:"datumUhrzeitVon" validate:"required" example:"2018-01-28T10:15:00+01"` // DatumUhrzeitVon is the inclusive begin
	DatumUhrzeitBis time.Time `json:"datumUhrzeitBis" validate:"required" example:"2018-01-28T10:30:00+01"` // DatumUhrzeitBis is the exclusive end
}
