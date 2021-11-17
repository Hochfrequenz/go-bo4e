package com

// Tagesparameter speichert Informationen zu einer tagesparameter-abhängigen Messstelle. z.B. den Namen einer Klimazone oder die ID der Wetterstation für die Temperaturmessstelle
type Tagesparameter struct {
	Klimazone            string `json:"klimazone" validate:"required,alphanum" example:"7624q"`            // Klimazone is the qualifier of the climate zone
	Temperaturmessstelle string `json:"temperaturmessstelle" validate:"required,alphanum" example:"1234x"` // Klimazone is the qualifier of the temperature measurement station
	Dienstanbieter       string `json:"dienstanbieter" validate:"required_with=Temperaturmessstelle"`      // Dienstanbieter is the service provider for temperature measurements
}
