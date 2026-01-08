package com

import "github.com/hochfrequenz/go-bo4e/enum/geraetetyp"

// A Hardware is a billable device
type Hardware struct {
	GeraeteTyp  *geraetetyp.Geraetetyp `json:"geraetetyp,omitempty"`                          // GeraeteTyp type of the hardware
	Bezeichnung *string                `json:"bezeichnung,omitempty" validate:"alphaunicode"` // Bezeichnung is a description
}
