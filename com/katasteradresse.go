package com

// Katasteradresse is a cadastre location
type Katasteradresse struct {
	// ToDo: why is this `gemarkung_flur` and not `gemarkungFlur`
	GemarkungFlur string `json:"gemarkung_flur,omitempty" validate:"required"` // GemarkungFlur is the district in which something is located
	Flurstueck    string `json:"flurstueck,omitempty" validate:"required"`     // Flurstueck describes the parcel
}
