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
	"github.com/hochfrequenz/go-bo4e/enum/abweichungsgrund"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

// Test_Deserialization deserializes an avisposition json
func Test_AvisPosition_Deserialization(t *testing.T) {
	avisPosition := com.AvisPosition{
		RechnungsNummer: "AAAA",
		RechnungsDatum:  time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Storno:          false,
		GesamtBrutto: com.Betrag{
			Wert:     decimal.New(0, 0),
			Waehrung: waehrungscode.EUR,
		},
		ZuZahlen: com.Betrag{
			Wert:     decimal.New(0, 0),
			Waehrung: waehrungscode.EUR,
		},
	}
	serializedAvisPosition, err := json.Marshal(avisPosition)
	then.AssertThat(t, serializedAvisPosition, is.Not(is.NilArray[byte]()))
	jsonString := string(serializedAvisPosition)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, strings.Contains(jsonString, "AAAA"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "ZZZZ"), is.False())
	deserializedAvisPosition := com.AvisPosition{}
	err = json.Unmarshal(serializedAvisPosition, &deserializedAvisPosition)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedAvisPosition, is.EqualTo(avisPosition))
}

// Test_Successful_Validation asserts that the validation does not fail for a valid AvisPosition
func Test_Successful_AvisPosition_Validation(t *testing.T) {
	bemerkung := "CCCC"
	validate := validator.New()
	ungleichVertragsbeginn := abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN
	validAvisPosition := []interface{}{
		com.AvisPosition{
			RechnungsNummer: "AAAA",
			RechnungsDatum:  time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Storno:          false,
			GesamtBrutto: com.Betrag{
				Wert:     decimal.New(0, 0),
				Waehrung: waehrungscode.EUR,
			},
			ZuZahlen: com.Betrag{
				Wert:     decimal.New(0, 0),
				Waehrung: waehrungscode.EUR,
			},
			Referenz: "BBBB",

			Abweichungen: []com.Abweichung{{
				AbweichungsGrund:          &ungleichVertragsbeginn,
				AbweichungsGrundBemerkung: &bemerkung,
			},
			},
		},
	}
	VerifySuccessfulValidations(t, validate, validAvisPosition)
}

func Test_Serialized_Empty_AvisPosition_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.AvisPosition{})
}
