package com_test

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

func Test_Failed_VorauszahlungValidation(t *testing.T) {
	validate := validator.New()
	invalidVorauszahlungMap := map[string][]interface{}{
		"required": {
			com.Vorauszahlung{
				// no Betrag given
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidVorauszahlungMap)
}

func Test_Successful_VorauszahlungValidation(t *testing.T) {
	validate := validator.New()
	arbitraryReferenz := "foo"
	arbitraryDate := time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC)
	validBetraege := []interface{}{
		com.Vorauszahlung{
			Betrag: com.Betrag{
				Wert:     decimal.NewFromFloat(18.36),
				Waehrung: waehrungscode.EUR,
			},
		},
		com.Vorauszahlung{
			Betrag: com.Betrag{
				Wert:     decimal.NewFromFloat(18.36),
				Waehrung: waehrungscode.EUR,
			},
			Referenz:      &arbitraryReferenz,
			ReferenzDatum: &arbitraryDate,
		},
	}
	VerifySuccessfulValidations(t, validate, validBetraege)
}

func Test_Serialized_Empty_Vorauszahlung_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Vorauszahlung{})
}
