package com_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

// TestFailedExterneReferenzValidation asserts that the validation fails, if not both name and value are provided
func (s *Suite) Test_Failed_ExterneReferenzValidation() {
	validate := validator.New()
	invalidReferenzMap := map[string][]interface{}{
		"required": {
			com.ExterneReferenz{
				ExRefName: "",
				ExRefWert: "",
			},
			com.ExterneReferenz{
				ExRefName: "foo",
				ExRefWert: "",
			},
			com.ExterneReferenz{
				ExRefName: "",
				ExRefWert: "bar",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidReferenzMap)
}

// TestSuccessfulExterneReferenzValidation asserts that the validation does not fail for a valid ExterneReferenz
func (s *Suite) Test_Successful_ExterneReferenzValidation() {
	validate := validator.New()
	validReferences := []interface{}{
		com.ExterneReferenz{
			ExRefName: "foo",
			ExRefWert: "bar",
		},
	}
	VerfiySuccessfulValidations(s, validate, validReferences)
}
