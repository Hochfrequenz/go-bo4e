package com_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/enum/schwachlastfaehigkeit"
	"github.com/hochfrequenz/go-bo4e/internal"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

func Test_Zaehlzeitregister_Deserialization(t *testing.T) {
	zaehlzeitregister := com.Zaehlzeitregister{
		ZaehlzeitDefinition: internal.Ptr("HL1"),
		Register:            internal.Ptr("HT"),
		Schwachlastfaehig:   internal.Ptr(schwachlastfaehigkeit.NICHT_SCHWACHLASTFAEHIG),
	}
	serializedZaehlzeitregister, err := json.Marshal(zaehlzeitregister)
	then.AssertThat(t, serializedZaehlzeitregister, is.Not(is.NilArray[byte]()))
	then.AssertThat(t, err, is.Nil())
	deserializedZaehlzeitregister := com.Zaehlzeitregister{}
	err = json.Unmarshal(serializedZaehlzeitregister, &deserializedZaehlzeitregister)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZaehlzeitregister, is.EqualTo(zaehlzeitregister))
}

func Test_Successful_Zaehlzeitregister_Validation(t *testing.T) {
	validate := validator.New()

	validAbweichung := []interface{}{
		com.Zaehlzeitregister{
			ZaehlzeitDefinition: internal.Ptr("HL1"),
			Register:            internal.Ptr("HT"),
			Schwachlastfaehig:   internal.Ptr(schwachlastfaehigkeit.NICHT_SCHWACHLASTFAEHIG),
		},
		com.Zaehlzeitregister{},
	}
	VerifySuccessfulValidations(t, validate, validAbweichung)
}

func Test_Serialized_Empty_Zaehlzeitregister_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Zaehlzeitregister{})
}
