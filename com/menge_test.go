package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
	"strings"
)

// Test_Menge_Deserialization deserializes an Menge json
func (s *Suite) Test_Menge_Deserialization() {
	var menge = com.Menge{
		Wert:    decimal.NewFromFloat(42),
		Einheit: internal.Ptr(mengeneinheit.KUBIKMETER),
	}
	serializedMenge, err := json.Marshal(menge)
	jsonString := string(serializedMenge)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KUBIKMETER"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMenge, is.Not(is.NilArray[byte]()))
	var deserializedVerbrauch com.Menge
	err = json.Unmarshal(serializedMenge, &deserializedVerbrauch)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVerbrauch, is.EqualTo[com.Menge](menge))
}

// Test_Successful_Menge_Validation asserts that the validation does not fail for a valid Menge
func (s *Suite) Test_Successful_Menge_Validation() {
	validate := validator.New()
	validMenges := []interface{}{
		com.Menge{
			Wert:    decimal.NewFromFloat(42),
			Einheit: internal.Ptr(mengeneinheit.KUBIKMETER),
		},
		com.Menge{
			Wert:    decimal.NewFromFloat(0), // 0 is a valid value
			Einheit: internal.Ptr(mengeneinheit.W),
		},
	}
	VerfiySuccessfulValidations(s, validate, validMenges)
}

// TestMengeFailedValidation verifies that invalid verbrauch values are considered invalid
func (s *Suite) Test_Menge_FailedValidation() {
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			com.Menge{
				Wert: decimal.NewFromFloat(42),
			},
			com.Menge{},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVerbrauchMap)
}

func (s *Suite) Test_Serialized_Empty_Menge_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Menge{})
}
