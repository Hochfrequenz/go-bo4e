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
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/shopspring/decimal"
)

// Test_Preis_Deserialization deserializes an Preis json
func Test_Preis_Deserialization(t *testing.T) {
	var preis = com.Preis{
		Wert:       decimal.NewFromFloat(17.89),
		Einheit:    waehrungseinheit.EUR,
		Bezugswert: mengeneinheit.MONAT,
		Status:     preisstatus.ENDGUELTIG,
	}
	serializedPreis, err := json.Marshal(preis)
	jsonString := string(serializedPreis)
	then.AssertThat(t, strings.Contains(jsonString, "MONAT"), is.True())      // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "ENDGUELTIG"), is.True()) // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "EUR"), is.True())        // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedPreis, is.Not(is.NilArray[byte]()))
	var deserializedPreis com.Preis
	err = json.Unmarshal(serializedPreis, &deserializedPreis)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedPreis, is.EqualTo(preis))
}

// Test_Successful_Preis_Validation asserts that the validation does not fail for a valid Preis
func Test_Successful_Preis_Validation(t *testing.T) {
	validate := validator.New()
	validPrices := []interface{}{
		com.Preis{
			Wert:       decimal.NewFromFloat(17.89),
			Einheit:    waehrungseinheit.EUR,
			Bezugswert: mengeneinheit.MONAT,
			Status:     preisstatus.ENDGUELTIG,
		},
		com.Preis{
			Wert:       decimal.NewFromFloat(17.89),
			Einheit:    waehrungseinheit.EUR,
			Bezugswert: mengeneinheit.MONAT,
			// status is not obligatory
		},
	}
	VerifySuccessfulValidations(t, validate, validPrices)
}

// TestPreisFailedValidation verifies that invalid Preis values are considered invalid
func Test_Preis_FailedValidation(t *testing.T) {
	validate := validator.New()
	invalidPrices := map[string][]interface{}{
		"required": {
			com.Preis{},
			com.Preis{
				Wert:       decimal.NewFromFloat(1),
				Einheit:    0,
				Bezugswert: 0,
				Status:     0,
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidPrices)
}

func Test_Serialized_Empty_Preis_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Preis{})
}
