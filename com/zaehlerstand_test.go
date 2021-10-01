package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// TestZaehlerstandDeserialization deserializes a Zaehlerstand json
func (s *Suite) Test_Zaehlerstand_Deserialization() {
	var zaehlerstand = Zaehlerstand{
		Ablesedatum:              time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Wert:                     decimal.NewFromFloat(847439),
		Einheit:                  mengeneinheit.KWH,
		Zustandszahl: decimal.NullDecimal{
			Decimal: decimal.NewFromFloat(17.23),
			Valid:   true,
		},
	}
	serializedZaehlerstand, err := json.Marshal(zaehlerstand)
	jsonString := string(serializedZaehlerstand)
	then.AssertThat(s.T(), strings.Contains(jsonString, "MESSUNG"), is.True())      // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "KWH"), is.True())          // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "zustandszahl"), is.True()) // is not omitted
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZaehlerstand, is.Not(is.Nil()))
	var deserializedZaehlerstand Zaehlerstand
	err = json.Unmarshal(serializedZaehlerstand, &deserializedZaehlerstand)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZaehlerstand, is.EqualTo(zaehlerstand))
}

//  Test_Zaehlerstand_Failed_Validation verifies that the validation fails for invalid Zaehlerstand s
func (s *Suite) Test_Zaehlerstand_Failed_Validation() {
	validate := validator.New()
	invalidZaehlerstand := map[string][]interface{}{
		"required": {
			Zaehlerstand{
				Ablesedatum: time.Time{},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidZaehlerstand)
}

//  Test_Successful_Zaehlerstand_Validation asserts that the validation does not fail for a valid Zaehlerstand
func (s *Suite) Test_Successful_Zaehlerstand_Validation() {
	validate := validator.New()
	validZaehlerstaende := []interface{}{
		Zaehlerstand{
			Ablesedatum:              time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(484535),
			Einheit:                  mengeneinheit.KWH,
		},
	}
	VerfiySuccessfulValidations(s, validate, validZaehlerstaende)
}
