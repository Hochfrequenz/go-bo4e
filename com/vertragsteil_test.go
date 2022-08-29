package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// TestVertragsteilDeserialization deserializes a Vertragsteil json
func (s *Suite) Test_Vertragsteil_Deserialization() {
	randomStartDate := time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC)
	var vertraqsteil = com.Vertragsteil{
		Vertragsteilbeginn: &randomStartDate,
		Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		Lokation:           "DE0123456789012345678901234567890",
		VertraglichFixierteMenge: &com.Menge{
			Wert:    decimal.NewFromFloat(42),
			Einheit: mengeneinheit.KUBIKMETER,
		},
		MinimaleAbnahmemenge: &com.Menge{
			Wert:    decimal.NewFromFloat(17),
			Einheit: mengeneinheit.MW,
		},
		MaximaleAbnahmemenge: &com.Menge{
			Wert:    decimal.NewFromFloat(-3),
			Einheit: mengeneinheit.MONAT,
		},
	}
	serializedVertragsteil, err := json.Marshal(vertraqsteil)
	jsonString := string(serializedVertragsteil)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KUBIKMETER"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedVertragsteil, is.Not(is.Nil()))
	var deserializedVertragsteil com.Vertragsteil
	err = json.Unmarshal(serializedVertragsteil, &deserializedVertragsteil)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVertragsteil, is.EqualTo(vertraqsteil))
}

//  Test_Vertragsteil_Failed_Validation verifies that the validation fails for invalid Vertragsteil
func (s *Suite) Test_Vertragsteil_Failed_Validation() {
	validate := validator.New()
	randomStartDate := time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC)
	randomStartDate2024 := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	invalidVertragsteile := map[string][]interface{}{
		"required": {
			com.Vertragsteil{
				Vertragsteilbeginn: &randomStartDate,
				Vertragsteilende:   time.Time{},
			},
		},
		"min": {
			com.Vertragsteil{
				Vertragsteilbeginn: &randomStartDate,
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "tooshort",
			},
		},
		"max": {
			com.Vertragsteil{
				Vertragsteilbeginn: &randomStartDate,
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "tooooooooooooooooooooooooooooooolong",
			},
		},
		"alphanum": {
			com.Vertragsteil{
				Vertragsteilbeginn: &randomStartDate,
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "not!alpha&num",
			},
		},
		"gtfield": {
			com.Vertragsteil{
				Vertragsteilbeginn: &randomStartDate2024,
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertragsteile)
}

//  Test_Successful_Vertragskonditionen_Validation asserts that the validation does not fail for a valid Vertragskonditionen
func (s *Suite) Test_Successful_Vertragsteil_Validation() {
	validate := validator.New()
	randomStartDate := time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC)
	emptyDate := time.Time{}
	validVertragsteile := []interface{}{
		com.Vertragsteil{
			Vertragsteilbeginn: &randomStartDate,
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Lokation:           "DE0123456789012345678901234567890",
			VertraglichFixierteMenge: &com.Menge{
				Wert:    decimal.NewFromFloat(42),
				Einheit: mengeneinheit.KUBIKMETER,
			},
			MinimaleAbnahmemenge: &com.Menge{
				Wert:    decimal.NewFromFloat(17),
				Einheit: mengeneinheit.MW,
			},
			MaximaleAbnahmemenge: &com.Menge{
				Wert:    decimal.NewFromFloat(-3),
				Einheit: mengeneinheit.MONAT,
			},
		},
		com.Vertragsteil{
			Vertragsteilbeginn: &randomStartDate,
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		},
		com.Vertragsteil{
			Vertragsteilbeginn: &randomStartDate,
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Lokation:           "543231012345",
		},
		com.Vertragsteil{
			Vertragsteilende:   time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilbeginn: &emptyDate, // empty start date is supported now (was not prior to https://github.com/Hochfrequenz/BO4E-dotnet/commit/88b974e7e30f8d7c9d23aa4284278c5e344142ef )
		},
	}
	VerfiySuccessfulValidations(s, validate, validVertragsteile)
}

func (s *Suite) Test_Serialized_Empty_Vertragsteil_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Vertragsteil{})
}
