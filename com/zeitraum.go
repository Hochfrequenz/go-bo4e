package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
	"time"
)

// Zeitraum ist eine Komponente zur Abbildung von Zeiträumen in Form von Dauern oder der Angabe von Start und Ende verwendet.
type Zeitraum struct {
	Einheit        zeiteinheit.Zeiteinheit `json:"einheit,omitempty" validate:"omitempty,required_with=Dauer"`                                                                       // Einheit is the unit of measurement
	Dauer          decimal.NullDecimal     `json:"dauer,omitempty" validate:"omitempty,required_with=Einheit"`                                                                       // Dauer ist eh duration
	Startdatum     *time.Time              `json:"startdatum,omitempty" validate:"required_without_all=Einheit Dauer" example:"2018-01-28T10:15:00+01"`                              // Startdatum is the day and time of begin
	Enddatum       *time.Time              `json:"enddatum,omitempty" validate:"required_without_all=Einheit Dauer" example:"2018-01-28T10:15:00+01"`                                // Enddatum is the day and time of end
	Startzeitpunkt *time.Time              `json:"startzeitpunkt,omitempty" validate:"required_without_all=Einheit Dauer" example:"2018-01-28T10:15:00+01"`                          // Startzeitpunkt is the inclusive begin
	Endzeitpunkt   *time.Time              `json:"endzeitpunkt,omitempty" validate:"required_with=Startzeitpunkt,omitempty,gtfield=Startzeitpunkt" example:"2018-01-28T10:30:00+01"` // Endzeitpunkt is the exclusive end
}

func (zeitraum Zeitraum) AsPointer() *Zeitraum {
	return &zeitraum
}
