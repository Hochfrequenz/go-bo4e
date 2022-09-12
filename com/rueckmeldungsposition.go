package com

type Rueckmeldungsposition struct {
	Positionsnummer       *string               `json:"positionenummer,omitempty"`
	Abweichungspositionen []Abweichungsposition `json:"abweichungspositionen"`
}
