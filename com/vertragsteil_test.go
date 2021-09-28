package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// TestVertragsteilDeserialization deserializes a Vertragsteil json
func (s *Suite) TestVertragsteilDeserialization() {
	var vertraqsteil = Vertragsteil{
		Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		Lokation:           "DE0123456789012345678901234567890",
		VertraglichFixierteMenge: &Menge{
			Wert:    decimal.NewFromFloat(42),
			Einheit: mengeneinheit.KUBIKMETER,
		},
		MinimaleAbnahmemenge: &Menge{
			Wert:    decimal.NewFromFloat(17),
			Einheit: mengeneinheit.MW,
		},
		MaximaleAbnahmemenge: &Menge{
			Wert:    decimal.NewFromFloat(-3),
			Einheit: mengeneinheit.Monat,
		},
	}
	serializedVertragsteil, err := json.Marshal(vertraqsteil)
	jsonString := string(serializedVertragsteil)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KUBIKMETER"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedVertragsteil, is.Not(is.Nil()))
	var deserializedVertragsteil Vertragsteil
	err = json.Unmarshal(serializedVertragsteil, &deserializedVertragsteil)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVertragsteil, is.EqualTo(vertraqsteil))
}

//  Test_Vertragsteil_Failed_Validation verifies that the validation fails for invalid Vertragsteil
func (s *Suite) Test_Vertragsteil_Failed_Validation() {
	validate := validator.New()
	invalidVertragsteile := map[string][]interface{}{
		"required": {
			Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Time{},
			},
			Vertragsteil{
				Vertragsteilende:   time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilbeginn: time.Time{},
			},
		},
		"min": {
			Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "tooshort",
			},
		},
		"max": {
			Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "tooooooooooooooooooooooooooooooolong",
			},
		},
		"alphanum": {
			Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Lokation:           "not!alpha&num",
			},
		},
		"gtfield": {
			Vertragsteil{
				Vertragsteilbeginn: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertragsteile)
}

//  Test_Successful_Vertragskonditionen_Validation asserts that the validation does not fail for a valid Vertragskonditionen
func (s *Suite) Test_Successful_Vertragsteil_Validation() {
	validate := validator.New()
	validVertragsteile := []interface{}{
		Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Lokation:           "DE0123456789012345678901234567890",
			VertraglichFixierteMenge: &Menge{
				Wert:    decimal.NewFromFloat(42),
				Einheit: mengeneinheit.KUBIKMETER,
			},
			MinimaleAbnahmemenge: &Menge{
				Wert:    decimal.NewFromFloat(17),
				Einheit: mengeneinheit.MW,
			},
			MaximaleAbnahmemenge: &Menge{
				Wert:    decimal.NewFromFloat(-3),
				Einheit: mengeneinheit.Monat,
			},
		},
		Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		},
		Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Lokation:           "543231012345",
		},
	}
	VerfiySuccessfulValidations(s, validate, validVertragsteile)
}
