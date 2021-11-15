package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/shopspring/decimal"
	"strings"
)

// Test_Preis_Deserialization deserializes an Preis json
func (s *Suite) Test_Preis_Deserialization() {
	var preis = com.Preis{
		Wert:       decimal.NewFromFloat(17.89),
		Einheit:    waehrungseinheit.EUR,
		Bezugswert: mengeneinheit.MONAT,
		Status:     preisstatus.ENDGUELTIG,
	}
	serializedPreis, err := json.Marshal(preis)
	jsonString := string(serializedPreis)
	then.AssertThat(s.T(), strings.Contains(jsonString, "MONAT"), is.True())      // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "ENDGUELTIG"), is.True()) // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "EUR"), is.True())        // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedPreis, is.Not(is.Nil()))
	var deserializedPreis com.Preis
	err = json.Unmarshal(serializedPreis, &deserializedPreis)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedPreis, is.EqualTo(preis))
}

//  Test_Successful_Preis_Validation asserts that the validation does not fail for a valid Preis
func (s *Suite) Test_Successful_Preis_Validation() {
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
	VerfiySuccessfulValidations(s, validate, validPrices)
}

//  TestPreisFailedValidation verifies that invalid Preis values are considered invalid
func (s *Suite) Test_Preis_FailedValidation() {
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
	VerfiyFailedValidations(s, validate, invalidPrices)
}

func (s *Suite) Test_Serialized_Empty_Preis_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Preis{})
}
