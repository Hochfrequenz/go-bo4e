package com

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

// TestFailedBetragValidation asserts that the validation fails, if not both Betrag.Wert and Betrag.Waehrung are provided
func (s *Suite) Test_Failed_BetragValidation() {
	validate := validator.New()
	invalidBetragMap := map[string][]interface{}{
		"required": {
			Betrag{
				Wert:     decimal.NewFromFloat(0),
				Waehrung: 0,
			},
			Betrag{
				Wert:     decimal.NewFromFloat(1),
				Waehrung: 0,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidBetragMap)
}

// TestSuccessfulBetragValidation asserts that the validation does not fail for a valid Betrag
func (s *Suite) Test_Successful_BetragValidation() {
	validate := validator.New()
	validBetraege := []interface{}{
		Betrag{
			Wert:     decimal.NewFromFloat(18.36),
			Waehrung: waehrungscode.EUR,
		},
		Betrag{
			Wert:     decimal.NewFromFloat(0), // wert 0 is allowed
			Waehrung: waehrungscode.ANG,
		},
	}
	VerfiySuccessfulValidations(s, validate, validBetraege)
}
