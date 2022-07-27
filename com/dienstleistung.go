package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/dienstleistungstyp"
)

// A Dienstleistung is a billable Service
type Dienstleistung struct {
	DienstleistungsTyp dienstleistungstyp.Dienstleistungstyp `json:"dienstleistungstyp,omitempty" validate:"required"`       // DienstleistungsTyp describes the type of the service
	Bezeichnung        string                                `json:"bezeichnung,omitempty" validate:"alphaunicode,required"` // Bezeichnung is a description
}
