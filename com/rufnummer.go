package com

import "github.com/hochfrequenz/go-bo4e/enum/rufnummernart"

// Rufnummer bildet eine Telefonnummer ab
type Rufnummer struct {
	Nummerntyp rufnummernart.Rufnummernart `json:"nummerntyp,omitempty" validate:"required"`                           // Nummerntyp ist die Ausprägung der Nummer
	Rufnummer  string                      `json:"rufnummer,omitempty" validate:"required" example:"4989209090153910"` // Rufnummer ist die konkrete Rufnummer
}
