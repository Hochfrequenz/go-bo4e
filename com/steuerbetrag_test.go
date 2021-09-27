package com

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
)

// TestFailedSteuerbetragValidation asserts that the validation fails for invalid Steuerbetrag
func (s *Suite) TestFailedSteuerbetragValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(SteuerbetragStructLevelValidation, Steuerbetrag{})
	invalidSteuerbetrag := map[string][]interface{}{
		"required": {
			Steuerbetrag{
				Steuerkennzeichen: 0,
				Basiswert:         0,
				Steuerwert:        0,
				Waehrung:          0,
			},
		},
		"Steuerwert=Basiswert*Steuerkennzeichen": {
			Steuerbetrag{
				Steuerkennzeichen: steuerkennzeichen.Ust19,
				Basiswert:         100,
				Steuerwert:        7, // expected 19
				Waehrung:          0,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidSteuerbetrag)
}

// TestSuccessfulSteuerbetragValidation asserts that the validation does not fail for a valid Steuerbetrag
func (s *Suite) TestSuccessfulSteuerbetragValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(SteuerbetragStructLevelValidation, Steuerbetrag{})
	validSteuerbetraege := []interface{}{
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Ust7,
			Waehrung:          waehrungscode.ALL,
		},
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Ust7,
			Basiswert:         100,
			Steuerwert:        7,
			Waehrung:          waehrungscode.ALL,
		},
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Ust19,
			Basiswert:         50,
			Steuerwert:        9.5,
			Waehrung:          waehrungscode.EUR,
		},
		Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Vst0,
			Basiswert:         100,
			Steuerwert:        0,
			Waehrung:          waehrungscode.EUR,
		},
	}
	VerfiySuccessfulValidations(s, validate, validSteuerbetraege)
}
