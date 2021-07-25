package com

import (
	"github.com/go-playground/validator"
)

//  Test_Failed_Validation asserts that the validation fails, if not both name and value are provided
func (s *Suite) Test_Failed_Validation() {
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

//  Test_Successful_Validation asserts that the validation does not fail for a valid ExterneReferenz
func (s *Suite) Test_Successful_ExterneReferenz_Validation() {
	validate := validator.New()
	validReferences := []interface{}{
		ExterneReferenz{
			ExRefName: "foo",
			ExRefWert: "bar",
		},
	}
	VerfiySuccessfulValidations(s, validate, validReferences)
}
