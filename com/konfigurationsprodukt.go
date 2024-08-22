package com

// Konfigurationsprodukt TODO explain what this is
type Konfigurationsprodukt struct {
	Produktcode               string `json:"produktcode,omitempty"`               // Der Konfigurationsprodukt-Code für das Objekt
	Leistungskurvendefinition string `json:"leistungskurvendefinition,omitempty"` // Code der zugeordneten Leistungskurvendefinition für das Objekt
	Schaltzeitdefinition      string `json:"schaltzeitdefinition,omitempty"`      // Code der zugeordneten Schaltzeitdefinition für das Objekt
	// Unfortunately adding the Marktteilnehmer here causes a serious issue: "could not import github.com/hochfrequenz/go-bo4e/bo (-: import cycle not allowed in test) (typecheck)". I comment it until we find a good solution
	// Marktpartner              bo.Marktteilnehmer `json:"marktpartner,omitempty"`              //Auftraggebender Marktpartner
}
