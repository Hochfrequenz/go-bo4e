package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/fehlercode"
)

//  Die Komponente wird dazu verwendet Fehler innerhalb eines bo.Statusbericht abzubilden
type FehlerDetail struct {
	Code         fehlercode.FehlerCode `json:"code" validate:"required"` //
	Beschreibung *string               `json:"beschreibung,omitempty"`   //
	Ursache      *FehlerUrsache        `json:"ursache,omitempty"`
}
