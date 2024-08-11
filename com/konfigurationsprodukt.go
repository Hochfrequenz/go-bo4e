package com

import (
	"github.com/hochfrequenz/go-bo4e/bo"
)

// Konfigurationsprodukt TODO explain what this is
type Konfigurationsprodukt struct {
	Produktcode               string             `json:"produktcode,omitempty"`               // Der Konfigurationsprodukt-Code für das Objekt
	Leistungskurvendefinition string             `json:"leistungskurvendefinition,omitempty"` // Code der zugeordneten Leistungskurvendefinition für das Objekt
	Schaltzeitdefinition      string             `json:"schaltzeitdefinition,omitempty"`      // Code der zugeordneten Schaltzeitdefinition für das Objekt
	Marktpartner              bo.Marktteilnehmer `json:"marktpartner,omitempty"`              //Auftraggebender Marktpartner
}
