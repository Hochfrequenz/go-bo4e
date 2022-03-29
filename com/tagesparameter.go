package com

// Tagesparameter speichert Informationen zu einer tagesparameter-abhängigen Messstelle. z.B. den Namen einer Klimazone oder die ID der Wetterstation für die Temperaturmessstelle
type Tagesparameter struct {
	Klimazone            string `json:"klimazone,omitempty" validate:"omitempty,alphanum" example:"7624q"`             // Klimazone is the qualifier of the climate zone
	Temperaturmessstelle string `json:"temperaturmessstelle,omitempty" validate:"omitempty,alphanum" example:"1234x"`  // Klimazone is the qualifier of the temperature measurement station
	Dienstanbieter       string `json:"dienstanbieter,omitempty" validate:"required_with=Temperaturmessstelle"`        // Dienstanbieter is the service provider for temperature measurements
	Herausgeber          string `json:"herausgeber,omitempty" validate:"omitempty,oneof=BDEW HAENDLER" example:"BDEW"` // Herausgeber is the publisher of the Tagesparameter. As of 2022-03-29 its value can be either 'BDEW' or 'HAENDLER'.
	// todo: Herausgeber could be an enum, probably: https://github.com/Hochfrequenz/BO4E-dotnet/pull/179#discussion_r837289474
	// if you fix it, also fix it in com.Lastprofil, please
}
