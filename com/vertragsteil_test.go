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
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
)

// TestVertragsteilDeserialization deserializes a Vertragsteil json
func Test_Vertragsteil_Deserialization(t *testing.T) {
	var vertraqsteil = com.Vertragsteil{
		Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		Lokation:           "DE0123456789012345678901234567890",
		VertraglichFixierteMenge: &com.Menge{
			Wert:    decimal.NewFromFloat(42),
			Einheit: internal.Ptr(mengeneinheit.KUBIKMETER),
		},
		MinimaleAbnahmemenge: &com.Menge{
			Wert:    decimal.NewFromFloat(17),
			Einheit: internal.Ptr(mengeneinheit.MW),
		},
		MaximaleAbnahmemenge: &com.Menge{
			Wert:    decimal.NewFromFloat(-3),
			Einheit: internal.Ptr(mengeneinheit.MONAT),
		},
	}
	serializedVertragsteil, err := json.Marshal(vertraqsteil)
	jsonString := string(serializedVertragsteil)
	then.AssertThat(t, strings.Contains(jsonString, "KUBIKMETER"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedVertragsteil, is.Not(is.NilArray[byte]()))
	var deserializedVertragsteil com.Vertragsteil
	err = json.Unmarshal(serializedVertragsteil, &deserializedVertragsteil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedVertragsteil, is.EqualTo(vertraqsteil))
}

// Test_Vertragsteil_Failed_Validation verifies that the validation fails for invalid Vertragsteil
func Test_Vertragsteil_Failed_Validation(t *testing.T) {
	validate := validator.New()
	invalidVertragsteile := map[string][]interface{}{
		"required": {
			com.Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Time{},
			},
			com.Vertragsteil{
				Vertragsteilende:   time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilbeginn: time.Time{},
			},
		},
		"min": {
			com.Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "tooshort",
			},
		},
		"max": {
			com.Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "tooooooooooooooooooooooooooooooolong",
			},
		},
		"alphanum": {
			com.Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "not!alpha&num",
			},
		},
		"gtfield": {
			com.Vertragsteil{
				Vertragsteilbeginn: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidVertragsteile)
}

// Test_Successful_Vertragskonditionen_Validation asserts that the validation does not fail for a valid Vertragskonditionen
func Test_Successful_Vertragsteil_Validation(t *testing.T) {
	validate := validator.New()
	validVertragsteile := []interface{}{
		com.Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Lokation:           "DE0123456789012345678901234567890",
			VertraglichFixierteMenge: &com.Menge{
				Wert:    decimal.NewFromFloat(42),
				Einheit: internal.Ptr(mengeneinheit.KUBIKMETER),
			},
			MinimaleAbnahmemenge: &com.Menge{
				Wert:    decimal.NewFromFloat(17),
				Einheit: internal.Ptr(mengeneinheit.MW),
			},
			MaximaleAbnahmemenge: &com.Menge{
				Wert:    decimal.NewFromFloat(-3),
				Einheit: internal.Ptr(mengeneinheit.MONAT),
			},
		},
		com.Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		},
		com.Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Lokation:           "543231012345",
		},
	}
	VerifySuccessfulValidations(t, validate, validVertragsteile)
}

func Test_Serialized_Empty_Vertragsteil_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Vertragsteil{})
}
