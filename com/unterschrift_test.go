package com_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

// TestUnterschriftDeserialization deserializes an Unterschrift json
func Test_Unterschrift_Deserialization(t *testing.T) {
	var unterschrift = com.Unterschrift{
		Name:  "Max Mustermann",
		Datum: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Ort:   "Musterstadt",
	}

	serializedUnterschrift, err := json.Marshal(unterschrift)
	jsonString := string(serializedUnterschrift)
	then.AssertThat(t, strings.Contains(jsonString, "Max Mustermann"), is.True())           // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "\"2021-08-01T00:00:00Z\""), is.True()) // ISO8601 ftw
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedUnterschrift, is.Not(is.NilArray[byte]()))
	var deserializedUnterschrift com.Unterschrift
	err = json.Unmarshal(serializedUnterschrift, &deserializedUnterschrift)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedUnterschrift, is.EqualTo(unterschrift))
}

// Test_Successful_Unterschrift_Validation asserts that the validation does not fail for a valid Unterschrift
func Test_Successful_Unterschrift_Validation(t *testing.T) {
	validate := validator.New()
	validUnterschrift := []interface{}{
		com.Unterschrift{
			Name:  "Max Mustermann",
			Datum: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			Ort:   "Musterstadt",
		},
		com.Unterschrift{
			Name: "Max Mustermann",
		},
	}
	VerifySuccessfulValidations(t, validate, validUnterschrift)
}

// TestUnterschriftFailedValidation verifies that invalid Unterschrift values are considered invalid
func Test_Unterschrift_FailedValidation(t *testing.T) {
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			com.Unterschrift{
				Name:  "",
				Datum: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
				Ort:   "Musterstadt",
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidVerbrauchMap)
}

func Test_Serialized_Empty_Unterschrift_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Unterschrift{})
}
