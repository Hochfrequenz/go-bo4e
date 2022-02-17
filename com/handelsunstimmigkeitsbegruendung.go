package com

import "github.com/hochfrequenz/go-bo4e/enum/handelsunstimmigkeitsgrund"

// Handelsunstimmigkeitsbegruendung reasons about the correctness of a "Rechnung" or "Lieferschein"
type Handelsunstimmigkeitsbegruendung struct {
	// Referenzen describes references to older messages.
	Referenzen []string `json:"referenzen,omitempty"`

	// Grund describes the reaseon for a "Handelsunstimmigkeit".
	Grund handelsunstimmigkeitsgrund.Handelsunstimmigkeitsgrund `json:"grund,omitempty" validate:"required"`
}
