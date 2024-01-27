package com_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

// TestFailedSteuerbetragValidation asserts that the validation fails for invalid Steuerbetrag
func Test_Failed_SteuerbetragValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(com.SteuerbetragStructLevelValidation, com.Steuerbetrag{})
	invalidSteuerbetrag := map[string][]interface{}{
		"required": {
			com.Steuerbetrag{
				Steuerkennzeichen: 0,
				Basiswert:         decimal.NewFromFloat(0),
				Steuerwert:        decimal.NewFromFloat(0),
				Waehrung:          0,
			},
		},
		"Steuerwert=Basiswert*Steuerkennzeichen": {
			com.Steuerbetrag{
				Steuerkennzeichen: steuerkennzeichen.UST_19,
				Basiswert:         decimal.NewFromFloat(100),
				Steuerwert:        decimal.NewFromFloat(7), // expected 19
				Waehrung:          0,
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidSteuerbetrag)
}

// TestSuccessfulSteuerbetragValidation asserts that the validation does not fail for a valid Steuerbetrag
func Test_Successful_SteuerbetragValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(com.SteuerbetragStructLevelValidation, com.Steuerbetrag{})
	validSteuerbetraege := []interface{}{
		com.Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.UST_7,
			Waehrung:          waehrungscode.ALL,
		},
		com.Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.UST_7,
			Basiswert:         decimal.NewFromFloat(100),
			Steuerwert:        decimal.NewFromFloat(7),
			Waehrung:          waehrungscode.ALL,
		},
		com.Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.UST_19,
			Basiswert:         decimal.NewFromFloat(50),
			Steuerwert:        decimal.NewFromFloat(9.5),
			Waehrung:          waehrungscode.EUR,
		},
		com.Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.VST_0,
			Basiswert:         decimal.NewFromFloat(100),
			Steuerwert:        decimal.NewFromFloat(0),
			Waehrung:          waehrungscode.EUR,
		},
	}
	VerifySuccessfulValidations(t, validate, validSteuerbetraege)
}

func (s *Suite) Test_Serialized_Empty_Steuerbetrag_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Steuerbetrag{})
}
