package com

import "github.com/hochfrequenz/go-bo4e/enum/geraetetyp"

// A Hardware is a billable device
type Hardware struct {
	GeraeteTyp  geraetetyp.Geraetetyp `json:"geraeteTyp,omitempty" validate:"required"`     // type of the hardware
	Bezeichnung string                `json:"bezeichnung" validate:"alphaunicode,required"` // description
}
