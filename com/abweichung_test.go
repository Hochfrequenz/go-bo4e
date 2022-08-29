package com_test

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"strings"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/abweichungsgrund"
)

// Test_Deserialization deserializes an abweichung json
func (s *Suite) Test_Abweichung_Deserialization() {
	ungleichVertragsbeginn := abweichungsgrund.ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN
	abweichung := com.Abweichung{
		Referenz:                  "AAAA",
		AbweichungsGrund:          &ungleichVertragsbeginn,
		AbweichungsGrundBemerkung: "BBBB",
	}
	serializedAbweichung, err := json.Marshal(abweichung)
	then.AssertThat(s.T(), serializedAbweichung, is.Not(is.Nil()))
	jsonString := string(serializedAbweichung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), strings.Contains(jsonString, "AAAA"), is.True())
	then.AssertThat(s.T(), strings.Contains(jsonString, "ZZZZ"), is.False())
	deserializedAbweichung := com.Abweichung{}
	err = json.Unmarshal(serializedAbweichung, &deserializedAbweichung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedAbweichung, is.EqualTo(abweichung))
}

//  Test_Successful_Validation asserts that the validation does not fail for a valid Abweichung
func (s *Suite) Test_Successful_Abweichung_Validation() {
	validate := validator.New()
	bilanzierteMengeFehlt := abweichungsgrund.BILANZIERTE_MENGE_FEHLT
	validAbweichung := []interface{}{
		com.Abweichung{
			Referenz:                  "",
			AbweichungsGrund:          &bilanzierteMengeFehlt,
			AbweichungsGrundBemerkung: "",
		},
	}
	VerfiySuccessfulValidations(s, validate, validAbweichung)
}

func (s *Suite) Test_Serialized_Empty_Abweichung_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Abweichung{})
}
