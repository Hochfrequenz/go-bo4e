package com_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
)

// Test_Menge_Deserialization deserializes an Menge json
func Test_Menge_Deserialization(t *testing.T) {
	var menge = com.Menge{
		Wert:    decimal.NewFromFloat(42),
		Einheit: internal.Ptr(mengeneinheit.KUBIKMETER),
	}
	serializedMenge, err := json.Marshal(menge)
	jsonString := string(serializedMenge)
	then.AssertThat(t, strings.Contains(jsonString, "KUBIKMETER"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedMenge, is.Not(is.NilArray[byte]()))
	var deserializedVerbrauch com.Menge
	err = json.Unmarshal(serializedMenge, &deserializedVerbrauch)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedVerbrauch, is.EqualTo[com.Menge](menge))
}

// Test_Successful_Menge_Validation asserts that the validation does not fail for a valid Menge
func Test_Successful_Menge_Validation(t *testing.T) {
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
	VerifySuccessfulValidations(t, validate, validMenges)
}

// TestMengeFailedValidation verifies that invalid verbrauch values are considered invalid
func Test_Menge_FailedValidation(t *testing.T) {
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			com.Menge{
				Wert: decimal.NewFromFloat(42),
			},
			com.Menge{},
		},
	}
	VerifyFailedValidations(t, validate, invalidVerbrauchMap)
}

func (s *Suite) Test_Serialized_Empty_Menge_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Menge{})
}
