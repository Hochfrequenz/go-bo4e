package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/profilart"
	"github.com/hochfrequenz/go-bo4e/enum/profilverfahren"
)

// Lastprofil defines a Lastprofil (no shit, sherlock)
// Note that this component is not official BO4E standard (yet)!
type Lastprofil struct {
	Bezeichnung    string                          `json:"bezeichnung,omitempty" validate:"omitempty,alphanum" example:"H0"` // Bezeichnung des Profils, durch DVGW bzw. den Netzbetreiber vergeben
	Verfahren      profilverfahren.Profilverfahren `json:"verfahren,omitempty"`                                              // Verfahren des Profils (analytisch oder synthetisch)
	Einspeisung    *bool                           `json:"einspeisung,omitempty"`                                            // Einspeisung ist true, falls es sich um Einspeisung handelt
	Tagesparameter *Tagesparameter                 `json:"tagesparameter,omitempty"`                                         // Tagesparameter contains the Klimazone or Temperaturmessstelle
	Profilart      profilart.Profilart             `json:"profilart,omitempty"`                                              // Profilart describes the kind of this profile
	Herausgeber    string                          `json:"herausgeber,omitempty" validate:"omitempty,oneof=BDEW HAENDLER" example:"BDEW"`                             // Herausgeber is the publisher of the Tagesparameter. As of 2022-03-29 its value can be either 'BDEW' or 'HAENDLER'.
	// todo: Herausgeber could be an enum, probably: https://github.com/Hochfrequenz/BO4E-dotnet/pull/179#discussion_r837289474
	// if you fix it, also fix it in com.Tagesparameter, please
}
