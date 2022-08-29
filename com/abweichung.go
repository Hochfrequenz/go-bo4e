package com

import "github.com/hochfrequenz/go-bo4e/enum/abweichungsgrund"

type Abweichung struct {
	Referenz                  string                             `json:"referenz,omitempty"`
	AbweichungsGrund          *abweichungsgrund.AbweichungsGrund `json:"abweichungsgrund,omitempty" validate:"omitempty"`
	AbweichungsGrundBemerkung string                             `json:"abweichungsgrundBemerkung,omitempty"`
}
