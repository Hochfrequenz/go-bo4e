package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/fehlertyp"
)

//  Die Komponente wird dazu verwendet Fehler innerhalb eines bo.Statusbericht abzubilden
type Fehler struct {
	Typ           fehlertyp.FehlerTyp `json:"typ" validate:"required"` //
	FehlerDetails []FehlerDetail      `json:"fehlerDetails"`           //
}
