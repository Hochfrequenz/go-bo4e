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
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
)

// TestZeitraumDeserialization deserializes a Zeitraum json
func Test_Zeitraum_Deserialization(t *testing.T) {
	var zeitraum = com.Zeitraum{
		Einheit:        zeiteinheit.MINUTE,
		Dauer:          decimal.NewNullDecimal(decimal.NewFromFloat(15)),
		Startzeitpunkt: internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
		Startdatum:     internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
		Enddatum:       internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15)),
		Endzeitpunkt:   internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15)),
	}
	serializedZeitraum, err := json.Marshal(zeitraum)
	jsonString := string(serializedZeitraum)
	then.AssertThat(t, strings.Contains(jsonString, "MINUTE"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedZeitraum, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Zeitraum
	err = json.Unmarshal(serializedZeitraum, &deserializedZeitreihenwert)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZeitreihenwert, is.EqualTo(zeitraum))
}

// TestZeitraumDeserializationWithoutEinheit
func Test_Zeitraum_DeserializationWithoutEinheit(t *testing.T) {
	var zeitraum = com.Zeitraum{
		Startzeitpunkt: internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
		Endzeitpunkt:   internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15)),
	}
	serializedZeitraum, err := json.Marshal(zeitraum)
	jsonString := string(serializedZeitraum)
	then.AssertThat(t, strings.Contains(jsonString, "zeiteinheit"), is.False())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedZeitraum, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Zeitraum
	err = json.Unmarshal(serializedZeitraum, &deserializedZeitreihenwert)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZeitreihenwert, is.EqualTo(zeitraum))
}

// Test_Zeitraum_Failed_Validation verifies that the validation fails for invalid Zeitraum s
func Test_Zeitraum_Failed_Validation(t *testing.T) {
	validate := validator.New()
	invalidZeitraums := map[string][]interface{}{
		"required_with": {
			com.Zeitraum{
				Startzeitpunkt: internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		"gtfield": {
			com.Zeitraum{
				Startzeitpunkt: internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
				Endzeitpunkt:   internal.Ptr(time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidZeitraums)
}

// Test_Successful_Zeitreihenwert_Validation asserts that the validation does not fail for a valid Zeitreihenwert
func Test_Successful_Zeitraum_Validation(t *testing.T) {
	validate := validator.New()
	validZeitraums := []interface{}{
		com.Zeitraum{
			Einheit:        zeiteinheit.Zeiteinheit(0),
			Dauer:          decimal.NewNullDecimal(decimal.NewFromFloat(15)),
			Startzeitpunkt: internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
			Startdatum:     internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
			Enddatum:       internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15)),
			Endzeitpunkt:   internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15)),
		},
		com.Zeitraum{
			Einheit: zeiteinheit.MINUTE,
			Dauer:   decimal.NewNullDecimal(decimal.NewFromFloat(15)),
		},
		com.Zeitraum{
			Startzeitpunkt: internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
			Endzeitpunkt:   internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15)),
			Startdatum:     internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)),
			Enddatum:       internal.Ptr(time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15)),
		},
	}
	VerifySuccessfulValidations(t, validate, validZeitraums)
}

func Test_Serialized_Empty_Zeitraum_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Zeitraum{})
}
