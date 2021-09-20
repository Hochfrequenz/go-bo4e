package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"time"
)

// Zeitraum ist eine Komponente zur Abbildung von Zeiträumen in Form von Dauern oder der Angabe von Start und Ende verwendet.
type Zeitraum struct {
	Einheit        zeiteinheit.Zeiteinheit `json:"zeiteinheit,omitempty" validate:"omitempty,required_with=Dauer"`                                                         // unit of measurement
	Dauer          float32                 `json:"dauer,omitempty" validate:"omitempty,required_with=Einheit"`                                                             // duration
	Startzeitpunkt time.Time               `json:"startzeitpunkt,omitempty" validate:"required_without_all=Einheit Dauer" example:"2018-01-28T10:15:00+01"`                // inclusive begin
	Endzeitpunkt   time.Time               `json:"endzeitpunkt,omitempty" validate:"required_with=Startzeitpunkt,gtfield=Startzeitpunkt" example:"2018-01-28T10:30:00+01"` // exclusive end
}
