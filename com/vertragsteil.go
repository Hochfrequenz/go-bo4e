package com

import "time"

// Vertragsteil wird dazu verwendet, eine vertragliche Leistung in Bezug zu einer Lokation (Markt- oder Messlokation) festzulegen.
type Vertragsteil struct {
	Vertragsteilbeginn       time.Time `json:"vertragsteilbeginn" validate:"required"`                          // Start der Gültigkeit des Vertragsteils.
	Vertragsteilende         time.Time `json:"Vertragsteilende" validate:"required,gtfield=Vertragsteilbeginn"` // Ende der Gültigkeit des Vertragsteils.
	Lokation                 string    `json:"lokation" validate:"omitempty,alphanum,min=11,max=33"`            // Der Identifier für diejenigen Markt- oder Messlokation, die zu diesem Vertragsteil gehören. Verträge für mehrere Lokationen werden mit mehreren Vertragsteilen abgebildet.
	VertraglichFixierteMenge *Menge    `json:"vertraglichFixierteMenge"`                                        // Für die Lokation festgeschriebene Abnahmemenge
	MinimaleAbnahmemenge     *Menge    `json:"minimaleAbnahmemenge"`                                            // Für die Lokation festgelegte Mindestabnahmemenge
	MaximaleAbnahmemenge     *Menge    `json:"maximaleAbnahmemenge"`                                            // Für die Lokation festgelegte maximale Abnahmemenge
}
