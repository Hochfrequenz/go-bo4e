package com_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
)

// TestZaehlerstandDeserialization deserializes a Zaehlerstand json
func Test_Zaehlerstand_Deserialization(t *testing.T) {
	var zaehlerstand = com.Zaehlerstand{
		Ablesedatum:              time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Wert:                     decimal.NewFromFloat(847439),
		Einheit:                  mengeneinheit.KWH,
		Zustandszahl:             decimal.NewNullDecimal(decimal.NewFromFloat(17.23)),
	}
	serializedZaehlerstand, err := json.Marshal(zaehlerstand)
	jsonString := string(serializedZaehlerstand)
	then.AssertThat(t, strings.Contains(jsonString, "MESSUNG"), is.True())      // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "KWH"), is.True())          // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "zustandszahl"), is.True()) // is not omitted
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedZaehlerstand, is.Not(is.NilArray[byte]()))
	var deserializedZaehlerstand com.Zaehlerstand
	err = json.Unmarshal(serializedZaehlerstand, &deserializedZaehlerstand)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZaehlerstand, is.EqualTo(zaehlerstand))
}

// Test_Zaehlerstand_Failed_Validation verifies that the validation fails for invalid Zaehlerstand s
func Test_Zaehlerstand_Failed_Validation(t *testing.T) {
	validate := validator.New()
	invalidZaehlerstand := map[string][]interface{}{
		"required": {
			com.Zaehlerstand{
				Ablesedatum: time.Time{},
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidZaehlerstand)
}

// Test_Successful_Zaehlerstand_Validation asserts that the validation does not fail for a valid Zaehlerstand
func Test_Successful_Zaehlerstand_Validation(t *testing.T) {
	validate := validator.New()
	validZaehlerstaende := []interface{}{
		com.Zaehlerstand{
			Ablesedatum:              time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(484535),
			Einheit:                  mengeneinheit.KWH,
		},
	}
	VerifySuccessfulValidations(t, validate, validZaehlerstaende)
}

func Test_Serialized_Empty_Zaehlerstand_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Zaehlerstand{})
}
