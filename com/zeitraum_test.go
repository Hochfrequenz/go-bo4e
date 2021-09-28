package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// TestZeitraumDeserialization deserializes a Zeitraum json
func (s *Suite) Test_Zeitraum_Deserialization() {
	var zeitraum = Zeitraum{
		Einheit: zeiteinheit.Minute,
		Dauer: decimal.NullDecimal{
			Decimal: decimal.NewFromFloat(15),
			Valid:   true,
		},
		Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Endzeitpunkt:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
	}
	serializedZeitraum, err := json.Marshal(zeitraum)
	jsonString := string(serializedZeitraum)
	then.AssertThat(s.T(), strings.Contains(jsonString, "Minute"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZeitraum, is.Not(is.Nil()))
	var deserializedZeitreihenwert Zeitraum
	err = json.Unmarshal(serializedZeitraum, &deserializedZeitreihenwert)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZeitreihenwert, is.EqualTo(zeitraum))
}

// TestZeitraumDeserializationWithoutEinheit
func (s *Suite) Test_Zeitraum_DeserializationWithoutEinheit() {
	var zeitraum = Zeitraum{
		Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Endzeitpunkt:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
	}
	serializedZeitraum, err := json.Marshal(zeitraum)
	jsonString := string(serializedZeitraum)
	then.AssertThat(s.T(), strings.Contains(jsonString, "zeiteinheit"), is.False())
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZeitraum, is.Not(is.Nil()))
	var deserializedZeitreihenwert Zeitraum
	err = json.Unmarshal(serializedZeitraum, &deserializedZeitreihenwert)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZeitreihenwert, is.EqualTo(zeitraum))
}

//  Test_Zeitraum_Failed_Validation verifies that the validation fails for invalid Zeitraum s
func (s *Suite) Test_Zeitraum_Failed_Validation() {
	validate := validator.New()
	invalidZeitraums := map[string][]interface{}{
		"required_with": {
			Zeitraum{
				Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		"gtfield": {
			Zeitraum{
				Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
				Endzeitpunkt:   time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidZeitraums)
}

//  Test_Successful_Zeitreihenwert_Validation asserts that the validation does not fail for a valid Zeitreihenwert
func (s *Suite) Test_Successful_Zeitraum_Validation() {
	validate := validator.New()
	validZeitraums := []interface{}{
		Zeitraum{
			Einheit: zeiteinheit.Zeiteinheit(0),
			Dauer: decimal.NullDecimal{
				Decimal: decimal.NewFromFloat(15),
				Valid:   true,
			},
			Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			Endzeitpunkt:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		},
		Zeitraum{
			Einheit: zeiteinheit.Minute,
			Dauer: decimal.NullDecimal{
				Decimal: decimal.NewFromFloat(15),
				Valid:   true,
			},
		},
		Zeitraum{
			Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			Endzeitpunkt:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		},
	}
	VerfiySuccessfulValidations(s, validate, validZeitraums)
}
