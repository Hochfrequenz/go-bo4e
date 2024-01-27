package com_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

// Test_Failed_Lastprofil_Validation asserts that the validation fails for invalid Lastprofil
func Test_Failed_Lastprofil_Validation(t *testing.T) {
	validate := validator.New()
	invalidLastprofile := map[string][]interface{}{
		"alphanum": {
			com.Lastprofil{
				Bezeichnung: "-*ö",
			},
		},
		"oneof": {
			com.Lastprofil{
				Herausgeber: "weder BDEW noch haendler",
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidLastprofile)
}

// Test_Successful_Lastprofil_Validation asserts that the validation does not fail for a valid Lastprofil
func Test_Successful_Lastprofil_Validation(t *testing.T) {
	validate := validator.New()
	validLastprofile := []interface{}{
		com.Lastprofil{}, // by default nothing is required
		com.Lastprofil{
			Bezeichnung: "Foo",
			Herausgeber: "BDEW",
		},
	}
	VerifySuccessfulValidations(t, validate, validLastprofile)
}
