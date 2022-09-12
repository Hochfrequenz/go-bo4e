package com

type Abweichungsposition struct {
	AbweichungsgrundCode      *string `json:"abweichungsgrundCode,omitempty" validate:"omitempty"`
	AbweichungsgrundCodeliste *string `json:"abweichungsgrundCodeliste,omitempty" validate:"required_with=AbweichungsgrundCode"`
	AbweichungsGrundBemerkung *string `json:"abweichungsgrundBemerkung,omitempty"`
	ZugehoerigeRechnung       *string `json:"zugehoerigeRechnung,omitempty"`
	ZugehoerigeBestellung     *string `json:"zugehoerigeBestellung,omitempty"`
}
