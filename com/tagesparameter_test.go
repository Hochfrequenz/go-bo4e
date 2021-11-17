package com_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

// Test_Failed_Tagesparameter_Validation asserts that the validation fails for invalid Tagesparameter
func (s *Suite) Test_Failed_Tagesparameter_Validation() {
	validate := validator.New()
	invalidTagesparameters := map[string][]interface{}{
		"required": {
			com.Tagesparameter{},
		},
		"required_with": {
			com.Tagesparameter{
				Klimazone:            "",
				Temperaturmessstelle: "asd",
				Dienstanbieter:       "", // <-- required if temperaturmessstelle is provided
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidTagesparameters)
}

// Test_Successful_Tagesparameter_Validation asserts that the validation does not fail for a valid Tagesparameter
func (s *Suite) Test_Successful_Tagesparameter_Validation() {
	validate := validator.New()
	validTagesparameter := []interface{}{
		com.Tagesparameter{
			Klimazone:            "asd",
			Temperaturmessstelle: "foo",
			Dienstanbieter:       "bar",
		},
	}
	VerfiySuccessfulValidations(s, validate, validTagesparameter)
}
