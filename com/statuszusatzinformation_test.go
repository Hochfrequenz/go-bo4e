package com_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/status"
	"github.com/hochfrequenz/go-bo4e/enum/statusart"
)

// Test_StatusZusatzInformation_Deserialization deserializes an StatusZusatzInformation json
func Test_StatusZusatzInformation_Deserialization(t *testing.T) {
	gasQualitaetArt := statusart.GASQUALITAET
	var sts = com.StatusZusatzInformation{
		Art:    &gasQualitaetArt,
		Status: status.BRENNWERTKORREKTUR,
	}

	serializedSts, err := json.Marshal(sts)
	jsonString := string(serializedSts)
	then.AssertThat(t, strings.Contains(jsonString, "GASQUALITAET"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "BRENNWERTKORREKTUR"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedSts, is.Not(is.NilArray[byte]()))
	var deserializedSts com.StatusZusatzInformation
	err = json.Unmarshal(serializedSts, &deserializedSts)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedSts, is.EqualTo(sts))
}

// Test_Successful_StatusZusatzInformation_Validation asserts that the validation
// does not fail for a valid StatusZusatzInformation
func Test_Successful_StatusZusatzInformation_Validation(t *testing.T) {
	statusartGasqualitaet := statusart.GASQUALITAET
	validate := validator.New()
	validSts := []interface{}{
		com.StatusZusatzInformation{
			Art:    &statusartGasqualitaet,
			Status: status.BRENNWERTKORREKTUR,
		},
		com.StatusZusatzInformation{
			Status: status.BRENNWERTKORREKTUR,
		},
	}
	VerifySuccessfulValidations(t, validate, validSts)
}

// Test_StatusZusatzInformation_FailedValidation verifies that invalid
// StatusZusatzInformation values are considered invalid
func Test_StatusZusatzInformation_FailedValidation(t *testing.T) {
	statusartGasqualitaet := statusart.GASQUALITAET
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			com.StatusZusatzInformation{
				Art: &statusartGasqualitaet,
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidVerbrauchMap)
}

func (s *Suite) Test_Serialized_Empty_StatusZusatzInformation_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.StatusZusatzInformation{})
}
