package com_test

import (
	"encoding/json"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

const foo = "foo"
const bar = "bar"

func Test_Zaehlzeit_Deserialization(t *testing.T) {
	_foo := foo
	_bar := bar
	zaehlzeit := com.Zaehlzeit{
		Zaehlzeitdefinition: &_foo,
		Zaehlzeitregister:   &_bar,
	}
	serializedZaehlzeit, err := json.Marshal(zaehlzeit)
	then.AssertThat(t, serializedZaehlzeit, is.Not(is.NilArray[byte]()))
	then.AssertThat(t, err, is.Nil())
	deserializedZaehlzeit := com.Zaehlzeit{}
	err = json.Unmarshal(serializedZaehlzeit, &deserializedZaehlzeit)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZaehlzeit, is.EqualTo(zaehlzeit))
}

func Test_Successful_Zaehlzeit_Validation(t *testing.T) {
	validate := validator.New()
	_foo := foo
	_bar := bar
	validAbweichung := []interface{}{
		com.Zaehlzeit{
			Zaehlzeitdefinition: &_foo,
			Zaehlzeitregister:   &_bar,
		},
		com.Zaehlzeit{
			Zaehlzeitregister: &_bar,
		},
		com.Zaehlzeit{
			Zaehlzeitdefinition: &_foo,
		},
		com.Zaehlzeit{},
	}
	VerifySuccessfulValidations(t, validate, validAbweichung)
}

func Test_Serialized_Empty_Zaehlzeit_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Zaehlzeit{})
}
