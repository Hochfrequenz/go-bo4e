package com_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
	"time"
)

func (s *Suite) Test_Failed_VorauszahlungValidation() {
	validate := validator.New()
	invalidVorauszahlungMap := map[string][]interface{}{
		"required": {
			com.Vorauszahlung{
				// no Betrag given
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVorauszahlungMap)
}

func (s *Suite) Test_Successful_VorauszahlungValidation() {
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
	VerfiySuccessfulValidations(s, validate, validBetraege)
}

func (s *Suite) Test_Serialized_Empty_Vorauszahlung_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Vorauszahlung{})
}
