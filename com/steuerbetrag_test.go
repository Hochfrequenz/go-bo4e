package com

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

// TestFailedSteuerbetragValidation asserts that the validation fails for invalid Steuerbetrag
func (s *Suite) Test_Failed_SteuerbetragValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(SteuerbetragStructLevelValidation, Steuerbetrag{})
	invalidSteuerbetrag := map[string][]interface{}{
		"required": {
			Steuerbetrag{
				Steuerkennzeichen: 0,
				Basiswert:         decimal.NewFromFloat(0),
				Steuerwert:        decimal.NewFromFloat(0),
				Waehrung:          0,
			},
		},
		"Steuerwert=Basiswert*Steuerkennzeichen": {
			Steuerbetrag{
				Steuerkennzeichen: steuerkennzeichen.Ust19,
				Basiswert:         decimal.NewFromFloat(100),
				Steuerwert:        decimal.NewFromFloat(7), // expected 19
				Waehrung:          0,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidSteuerbetrag)
}

// TestSuccessfulSteuerbetragValidation asserts that the validation does not fail for a valid Steuerbetrag
func (s *Suite) Test_Successful_SteuerbetragValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(SteuerbetragStructLevelValidation, Steuerbetrag{})
	validSteuerbetraege := []interface{}{
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Ust7,
			Waehrung:          waehrungscode.ALL,
		},
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Ust7,
			Basiswert:         decimal.NewFromFloat(100),
			Steuerwert:        decimal.NewFromFloat(7),
			Waehrung:          waehrungscode.ALL,
		},
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Ust19,
			Basiswert:         decimal.NewFromFloat(50),
			Steuerwert:        decimal.NewFromFloat(9.5),
			Waehrung:          waehrungscode.EUR,
		},
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Vst0,
			Basiswert:         decimal.NewFromFloat(100),
			Steuerwert:        decimal.NewFromFloat(0),
			Waehrung:          waehrungscode.EUR,
		},
	}
	VerfiySuccessfulValidations(s, validate, validSteuerbetraege)
}
