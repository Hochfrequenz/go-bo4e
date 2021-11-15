package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/dienstleistung"
)

// A Dienstleistung is a billable Service
type Dienstleistung struct {
	DienstleistungsTyp dienstleistung.Dienstleistungstyp `json:"dienstleistungstyp,omitempty" validate:"required"` // DienstleistungsTyp describes the type of the service
	Bezeichnung        string                            `json:"bezeichnung" validate:"alphaunicode,required"`     // Bezeichnung is a description
}
