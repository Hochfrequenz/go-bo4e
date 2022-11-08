package com

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/abweichungsgrund"
)

type Abweichung struct {
	AbweichungsGrund          *abweichungsgrund.AbweichungsGrund `json:"abweichungsgrund,omitempty" validate:"omitempty"`
	AbweichungsGrundBemerkung *string                            `json:"abweichungsgrundBemerkung,omitempty"`
	AbweichungsgrundCode      *string                            `json:"abweichungsgrundCode,omitempty" validate:"omitempty"`
	AbweichungsgrundCodeliste *string                            `json:"abweichungsgrundCodeliste,omitempty" validate:"omitempty"`
}

// AbweichungStructLevelValidation does a cross check on a Abweichung object
func AbweichungStructLevelValidation(sl validator.StructLevel) {
	// ToDo: use required_without/required_if instead of own validator
	// see https://github.com/go-playground/validator/v10/search?q=required_without
	abweichung := sl.Current().Interface().(Abweichung)
	if abweichung.AbweichungsgrundCode != nil && abweichung.AbweichungsgrundCodeliste == nil {
		sl.ReportError(abweichung.AbweichungsgrundCode, "AbweichungsgrundCode", "AbweichungsgrundCodeliste", "AbweichungsgrundCodeComplete", "")
	}
}
