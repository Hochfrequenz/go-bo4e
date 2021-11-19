package com

import "time"

// Vertragsteil wird dazu verwendet, eine vertragliche Leistung in Bezug zu einer Lokation (Markt- oder Messlokation) festzulegen.
type Vertragsteil struct {
	Vertragsteilbeginn       time.Time `json:"vertragsteilbeginn,omitempty" validate:"required"`                          // Vertragsteilbeginn ist der inklusive Start der Gültigkeit des Vertragsteils
	Vertragsteilende         time.Time `json:"vertragsteilende,omitempty" validate:"required,gtfield=Vertragsteilbeginn"` // Vertragsteilende ist das exklusive Ende der Gültigkeit des Vertragsteils
	Lokation                 string    `json:"lokation,omitempty" validate:"omitempty,alphanum,min=11,max=33"`            // Lokation ist der Identifier für diejenigen Markt- oder Messlokation, die zu diesem Vertragsteil gehören. Verträge für mehrere Lokationen werden mit mehreren Vertragsteilen abgebildet.
	VertraglichFixierteMenge *Menge    `json:"vertraglichFixierteMenge,omitempty"`                                        // VertraglichFixierteMenge ist die für die Lokation festgeschriebene Abnahmemenge
	MinimaleAbnahmemenge     *Menge    `json:"minimaleAbnahmemenge,omitempty"`                                            // MinimaleAbnahmemenge ist die, für die Lokation festgelegte Mindestabnahmemenge
	MaximaleAbnahmemenge     *Menge    `json:"maximaleAbnahmemenge,omitempty"`                                            // MaximaleAbnahmemenge ist die, für die Lokation festgelegte maximale Abnahmemenge
}
