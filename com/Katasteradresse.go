package com

// Katasteradresse is a cadastre location
type Katasteradresse struct {
	// ToDo: why is this `gemarkung_flur` and not `gemarkungFlur`
	GemarkungFlur string `json:"gemarkung_flur" validate:"required"` // the district in which something is located
	Flurstueck    string `json:"flurstueck" validate:"required"`     // parcel
}
