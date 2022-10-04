package com

type Rueckmeldungsposition struct {
	Positionsnummer       *string               `json:"positionsnummer,omitempty"`
	Abweichungspositionen []Abweichungsposition `json:"abweichungspositionen"`
}
