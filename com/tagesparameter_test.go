package com_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

// Test_Failed_Tagesparameter_Validation asserts that the validation fails for invalid Tagesparameter
func Test_Failed_Tagesparameter_Validation(t *testing.T) {
	validate := validator.New()
	invalidTagesparameters := map[string][]interface{}{
		"required_with": {
			com.Tagesparameter{
				Klimazone:            "",
				Temperaturmessstelle: "asd",
				Dienstanbieter:       "", // <-- required if temperaturmessstelle is provided
			},
		},
		"oneof": {
			com.Tagesparameter{
				Herausgeber: "Foo",
			},
		},
		"alphanum": {
			com.Tagesparameter{
				Klimazone: "😍",
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidTagesparameters)
}

// Test_Successful_Tagesparameter_Validation asserts that the validation does not fail for a valid Tagesparameter
func Test_Successful_Tagesparameter_Validation(t *testing.T) {
	validate := validator.New()
	validTagesparameter := []interface{}{
		com.Tagesparameter{}, // by default nothing is required
		com.Tagesparameter{
			Klimazone:            "asd",
			Temperaturmessstelle: "foo",
			Dienstanbieter:       "bar",
		},
	}
	VerifySuccessfulValidations(t, validate, validTagesparameter)
}
