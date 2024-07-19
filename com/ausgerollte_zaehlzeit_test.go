package com_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/internal"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

func Test_AusgerollteZaehlzeit_Deserialization(t *testing.T) {
	ausgerollteZaehlzeit := com.AusgerollteZaehlzeit{
		Aenderungszeitpunkt: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Register:            internal.Ptr("HT"),
	}
	serializedAusgerollteZaehlzeit, err := json.Marshal(ausgerollteZaehlzeit)
	then.AssertThat(t, serializedAusgerollteZaehlzeit, is.Not(is.NilArray[byte]()))
	then.AssertThat(t, err, is.Nil())
	deserializedAusgerollteZaehlzeit := com.AusgerollteZaehlzeit{}
	err = json.Unmarshal(serializedAusgerollteZaehlzeit, &deserializedAusgerollteZaehlzeit)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedAusgerollteZaehlzeit, is.EqualTo(ausgerollteZaehlzeit))
}

func Test_Successful_AusgerollteZaehlzeit_Validation(t *testing.T) {
	validate := validator.New()

	validAbweichung := []interface{}{
		com.AusgerollteZaehlzeit{
			Aenderungszeitpunkt: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Register:            internal.Ptr("HT"),
		},
		com.AusgerollteZaehlzeit{},
	}
	VerifySuccessfulValidations(t, validate, validAbweichung)
}

func Test_Serialized_Empty_AusgerollteZaehlzeit_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.AusgerollteZaehlzeit{})
}
