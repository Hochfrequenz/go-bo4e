package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// Test_Deserialization deserializes an Verbrauch json
func (s *Suite) TestVerbrauchDeserialization() {
	var verbrauch = Verbrauch{
		Startdatum:               time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Enddatum:                 time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     decimal.NewFromFloat(17),
		Einheit:                  mengeneinheit.KWH,
	}
	serializedVerbrauch, err := json.Marshal(verbrauch)
	jsonString := string(serializedVerbrauch)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KWH"), is.True())                      // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "\"2021-08-01T00:00:00Z\""), is.True()) // ISO8601 ftw
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedVerbrauch, is.Not(is.Nil()))
	var deserializedVerbrauch Verbrauch
	err = json.Unmarshal(serializedVerbrauch, &deserializedVerbrauch)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVerbrauch, is.EqualTo(verbrauch))
}

//  Test_Successful_Verbrauch_Validation asserts that the validation does not fail for a valid Verbrauch
func (s *Suite) Test_Successful_Verbrauch_Validation() {
	validate := validator.New()
	validVerbrauch := []interface{}{
		Verbrauch{
			Startdatum:               time.Now(),
			Enddatum:                 time.Now().Add(time.Minute * 15),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Obiskennzahl:             "1-0:1.8.1",
			Wert:                     decimal.NewFromFloat(17),
			Einheit:                  mengeneinheit.KWH,
		},
	}
	VerfiySuccessfulValidations(s, validate, validVerbrauch)
}

//  TestVerbrauchFailedValidation verifies that invalid verbrauch values are considered invalid
func (s *Suite) TestVerbrauchFailedValidation() {
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			Verbrauch{
				Startdatum:               time.Now(),
				Enddatum:                 time.Now().Add(time.Minute * 15),
				Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
				Obiskennzahl:             "1-0:1.8.1",
				Wert:                     decimal.NewFromFloat(0),
				Einheit:                  0,
			},
		},
		"gtfield": {
			Verbrauch{
				Startdatum:               time.Now(),
				Enddatum:                 time.Now().Add(time.Minute * -15),
				Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
				Obiskennzahl:             "1-0:1.8.1",
				Wert:                     decimal.NewFromFloat(17),
				Einheit:                  mengeneinheit.KWH,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVerbrauchMap)
}
