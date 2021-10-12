package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"strings"
	"time"
)

// TestUnterschriftDeserialization deserializes an Unterschrift json
func (s *Suite) Test_Unterschrift_Deserialization() {
	var unterschrift = com.Unterschrift{
		Name:  "Max Mustermann",
		Datum: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Ort:   "Musterstadt",
	}

	serializedUnterschrift, err := json.Marshal(unterschrift)
	jsonString := string(serializedUnterschrift)
	then.AssertThat(s.T(), strings.Contains(jsonString, "Max Mustermann"), is.True())           // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "\"2021-08-01T00:00:00Z\""), is.True()) // ISO8601 ftw
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedUnterschrift, is.Not(is.Nil()))
	var deserializedUnterschrift com.Unterschrift
	err = json.Unmarshal(serializedUnterschrift, &deserializedUnterschrift)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedUnterschrift, is.EqualTo(unterschrift))
}

//  Test_Successful_Unterschrift_Validation asserts that the validation does not fail for a valid Unterschrift
func (s *Suite) Test_Successful_Unterschrift_Validation() {
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
	VerfiySuccessfulValidations(s, validate, validUnterschrift)
}

//  TestUnterschriftFailedValidation verifies that invalid Unterschrift values are considered invalid
func (s *Suite) Test_Unterschrift_FailedValidation() {
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
	VerfiyFailedValidations(s, validate, invalidVerbrauchMap)
}
