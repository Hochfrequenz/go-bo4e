package com

import "github.com/hochfrequenz/go-bo4e/enum/geraetetyp"

// A Hardware is a billable device
type Hardware struct {
	GeraeteTyp  *geraetetyp.Geraetetyp `json:"geraetetyp,omitempty"`                          // GeraeteTyp type of the hardware
	Bezeichnung *string                `json:"bezeichnung,omitempty" validate:"alphaunicode"` // Bezeichnung is a description

	// Geraeteeigenschaften gibt die Eigenschaften des Geräts an.
	Geraeteeigenschaften *Geraeteeigenschaften `json:"geraeteeigenschaften,omitempty"`

	// Geraetenummer ist die Gerätenummer des Wandlers.
	Geraetenummer *string `json:"geraetenummer,omitempty"`

	// Geraetereferenz ist die Referenz auf die Gerätenummer des Zählers.
	Geraetereferenz *string `json:"geraetereferenz,omitempty"`
}
