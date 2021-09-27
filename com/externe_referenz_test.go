package com

import (
	"github.com/go-playground/validator/v10"
)

// TestFailedExterneReferenzValidation asserts that the validation fails, if not both name and value are provided
func (s *Suite) TestFailedExterneReferenzValidation() {
	validate := validator.New()
	invalidReferenzMap := map[string][]interface{}{
		"required": {
			ExterneReferenz{
				ExRefName: "",
				ExRefWert: "",
			},
			ExterneReferenz{
				ExRefName: "foo",
				ExRefWert: "",
			},
			ExterneReferenz{
				ExRefName: "",
				ExRefWert: "bar",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidReferenzMap)
}

// TestSuccessfulExterneReferenzValidation asserts that the validation does not fail for a valid ExterneReferenz
func (s *Suite) TestSuccessfulExterneReferenzValidation() {
	validate := validator.New()
	validReferences := []interface{}{
		ExterneReferenz{
			ExRefName: "foo",
			ExRefWert: "bar",
		},
	}
	VerfiySuccessfulValidations(s, validate, validReferences)
}
