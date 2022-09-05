package com_test

import (
	"encoding/json"
	"strings"
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
func (s *Suite) Test_AvisPosition_Deserialization() {
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
	then.AssertThat(s.T(), serializedAvisPosition, is.Not(is.Nil()))
	jsonString := string(serializedAvisPosition)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), strings.Contains(jsonString, "AAAA"), is.True())
	then.AssertThat(s.T(), strings.Contains(jsonString, "ZZZZ"), is.False())
	deserializedAvisPosition := com.AvisPosition{}
	err = json.Unmarshal(serializedAvisPosition, &deserializedAvisPosition)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedAvisPosition, is.EqualTo(avisPosition))
}

//  Test_Successful_Validation asserts that the validation does not fail for a valid AvisPosition
func (s *Suite) Test_Successful_AvisPosition_Validation() {
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
			Abweichung: &com.Abweichung{
				Referenz:                  "BBBB",
				AbweichungsGrund:          &ungleichVertragsbeginn,
				AbweichungsGrundBemerkung: &bemerkung,
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validAvisPosition)
}

func (s *Suite) Test_Serialized_Empty_AvisPosition_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.AvisPosition{})
}
