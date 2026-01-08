package com

import "github.com/hochfrequenz/go-bo4e/enum/lokationstyp"

// LokationsTypZuordnung ordnet einer Lokations-ID einen Lokationstyp zu.
type LokationsTypZuordnung struct {
	// Lokationstyp ist der Typ der Lokation.
	Lokationstyp *lokationstyp.Lokationstyp `json:"lokationstyp,omitempty"`

	// LokationsId ist der Wert, der einem Lokationstyp zugeordnet werden soll.
	LokationsId *string `json:"lokationsId,omitempty"`
}
