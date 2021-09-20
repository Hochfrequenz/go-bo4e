package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"strings"
)

// Test_Menge_Deserialization deserializes an Menge json
func (s *Suite) Test_Menge_Deserialization() {
	var menge = Menge{
		Wert:    42,
		Einheit: mengeneinheit.KUBIKMETER,
	}
	serializedMenge, err := json.Marshal(menge)
	jsonString := string(serializedMenge)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KUBIKMETER"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMenge, is.Not(is.Nil()))
	var deserializedVerbrauch Menge
	err = json.Unmarshal(serializedMenge, &deserializedVerbrauch)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVerbrauch, is.EqualTo(menge))
}

//  Test_Successful_Menge_Validation asserts that the validation does not fail for a valid Menge
func (s *Suite) Test_Successful_Menge_Validation() {
	validate := validator.New()
	validMenges := []interface{}{
		Menge{
			Wert:    42,
			Einheit: mengeneinheit.KUBIKMETER,
		},
		Menge{
			Wert:    0, // 0 is a valid value
			Einheit: mengeneinheit.W,
		},
	}
	VerfiySuccessfulValidations(s, validate, validMenges)
}

//  TestMengeFailedValidation verifies that invalid verbrauch values are considered invalid
func (s *Suite) TestMengeFailedValidation() {
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			Menge{
				Wert: 42,
			},
			Menge{},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVerbrauchMap)
}
