package com_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/tarifstufe"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/hochfrequenz/go-bo4e/internal"
)

// Test_Deserialization deserializes an Verbrauch json
func Test_Verbrauch_Deserialization(t *testing.T) {
	var verbrauch = com.Verbrauch{
		Startdatum:               time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Enddatum:                 time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     decimal.NewFromFloat(17),
		Einheit:                  mengeneinheit.KWH,
		Nutzungszeitpunkt:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		Tarifstufe:               tarifstufe.TARIFSTUFE_3,
	}
	serializedVerbrauch, err := json.Marshal(verbrauch)
	jsonString := string(serializedVerbrauch)
	then.AssertThat(t, strings.Contains(jsonString, "KWH"), is.True())                      // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "\"2021-08-01T00:00:00Z\""), is.True()) // ISO8601 ftw
	then.AssertThat(t, strings.Contains(jsonString, "\"2023-01-01T00:00:00Z\""), is.True()) // ISO8601 ftw
	then.AssertThat(t, strings.Contains(jsonString, "\"0001-01-01T00:00:00Z\""), is.False())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedVerbrauch, is.Not(is.NilArray[byte]()))
	var deserializedVerbrauch com.Verbrauch
	err = json.Unmarshal(serializedVerbrauch, &deserializedVerbrauch)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedVerbrauch, is.EqualTo(verbrauch))
}

// Test_Successful_Verbrauch_Validation asserts that the validation does not fail for a valid Verbrauch
func Test_Successful_Verbrauch_Validation(t *testing.T) {
	validate := validator.New()
	validVerbrauch := []interface{}{
		com.Verbrauch{
			Startdatum:               time.Now(),
			Enddatum:                 time.Now().Add(time.Minute * 15),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Obiskennzahl:             "1-0:1.8.1",
			Wert:                     decimal.NewFromFloat(17),
			Einheit:                  mengeneinheit.KWH,
			Nutzungszeitpunkt:        time.Now(),
			Ausfuehrungszeitpunkt:    internal.Ptr(time.Now()),
		},
	}
	VerifySuccessfulValidations(t, validate, validVerbrauch)
}

// TestVerbrauchFailedValidation verifies that invalid verbrauch values are considered invalid
func Test_Verbrauch_FailedValidation(t *testing.T) {
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			com.Verbrauch{
				Startdatum:               time.Now(),
				Enddatum:                 time.Now().Add(time.Minute * 15),
				Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
				Obiskennzahl:             "1-0:1.8.1",
				Wert:                     decimal.NewFromFloat(0),
				Einheit:                  0,
			},
		},
		"gtfield": {
			com.Verbrauch{
				Startdatum:               time.Now(),
				Enddatum:                 time.Now().Add(time.Minute * -15),
				Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
				Obiskennzahl:             "1-0:1.8.1",
				Wert:                     decimal.NewFromFloat(17),
				Einheit:                  mengeneinheit.KWH,
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidVerbrauchMap)
}

func Test_Serialized_Empty_Verbrauch_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Verbrauch{})
}
