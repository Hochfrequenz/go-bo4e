package com_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

// TestFailedExterneReferenzValidation asserts that the validation fails, if not both name and value are provided
func Test_Failed_ExterneReferenzValidation(t *testing.T) {
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
	VerifyFailedValidations(t, validate, invalidReferenzMap)
}

// TestSuccessfulExterneReferenzValidation asserts that the validation does not fail for a valid ExterneReferenz
func Test_Successful_ExterneReferenzValidation(t *testing.T) {
	validate := validator.New()
	validReferences := []interface{}{
		com.ExterneReferenz{
			ExRefName: "foo",
			ExRefWert: "bar",
		},
	}
	VerifySuccessfulValidations(t, validate, validReferences)
}

func Test_Serialized_Empty_Externe_Referenz_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.ExterneReferenz{})
}
