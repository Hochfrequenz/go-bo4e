package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/shopspring/decimal"
	"log"
	"strings"
	"time"
)

// TestVertragsteilDeserialization deserializes a Vertragsteil json
func (s *Suite) Test_Vertragsteil_Deserialization() {
	var vertraqsteil = com.Vertragsteil{
		Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
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
	then.AssertThat(s.T(), serializedVertragsteil, is.Not(is.NilArray[byte]()))
	var deserializedVertragsteil com.Vertragsteil
	err = json.Unmarshal(serializedVertragsteil, &deserializedVertragsteil)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVertragsteil, is.EqualTo(vertraqsteil))
}

// Test_Vertragsteil_Failed_Validation verifies that the validation fails for invalid Vertragsteil
func (s *Suite) Test_Vertragsteil_Failed_Validation() {
	validate := validator.New()
	if err := validate.RegisterValidation("timenotzero", valTimeNotZero); err != nil { //add custom validator tag for time not zero and handle error of registration
		//generate error if registration of validation failed
		log.Fatalln("Failed to register custom validation for checking time != 0:", err)
	}
	invalidVertragsteile := map[string][]interface{}{
		"timenotzero": { //todo: required validator for time.Time
			com.Vertragsteil{
				Vertragsteilbeginn: time.Time{},
				Vertragsteilende:   time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			},
			com.Vertragsteil{
				Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Vertragsteilende:   time.Time{},
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
	VerfiyFailedValidations(s, validate, invalidVertragsteile)
}

// Test_Successful_Vertragskonditionen_Validation asserts that the validation does not fail for a valid Vertragskonditionen
func (s *Suite) Test_Successful_Vertragsteil_Validation() {
	validate := validator.New()
	if err := validate.RegisterValidation("timenotzero", valTimeNotZero); err != nil { //add custom validator tag for time not zero and handle error of registration
		//generate error if registration of validation failed
		log.Fatalln("Failed to register custom validation for checking time != 0:", err)
	}
	validVertragsteile := []interface{}{
		com.Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
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
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		},
		com.Vertragsteil{
			Vertragsteilbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Vertragsteilende:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Lokation:           "543231012345",
		},
	}
	VerfiySuccessfulValidations(s, validate, validVertragsteile)
}

func (s *Suite) Test_Serialized_Empty_Vertragsteil_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Vertragsteil{})
}
