package com

import "github.com/hochfrequenz/go-bo4e/enum/profilverfahren"

// Lastprofil defines a Lastprofil (no shit, sherlock)
// Note that this component is not official BO4E standard (yet)!
type Lastprofil struct {
	Bezeichnung    string                          `json:"bezeichnung" validate:"omitempty,alphanum" example:"H0"` // Bezeichnung des Profils, durch DVGW bzw. den Netzbetreiber vergeben
	Verfahren      profilverfahren.Profilverfahren `json:"verfahren,omitempty"`                                    // Verfahren des Profils (analytisch oder synthetisch)
	Einspeisung    *bool                           `json:"einspeisung,omitempty" validate:"omitempty"`             // Einspeisung ist true, falls es sich um Einspeisung handelt
	Tagesparameter *Tagesparameter                 `json:"tagesparameter,omitempty"`                               // Tagesparameter contains the Klimazone or Temperaturmessstelle
}
