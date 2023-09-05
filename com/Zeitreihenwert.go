package com

import (
	"time"
)

// Zeitreihenwert ist eine Abbildung eines Zeitreihenwertes bestehend aus Zeitraum, Wert und Statusinformationen.
type Zeitreihenwert struct {
	Zeitreihenwertkompakt
	DatumUhrzeitVon time.Time `json:"datumUhrzeitVon,omitempty" validate:"timenotzero" example:"2018-01-28T10:15:00+01"` // DatumUhrzeitVon is the inclusive begin //add custom validator tag for time not zero
	DatumUhrzeitBis time.Time `json:"datumUhrzeitBis,omitempty" validate:"timenotzero" example:"2018-01-28T10:30:00+01"` // DatumUhrzeitBis is the exclusive end //add custom validator tag for time not zero
}
