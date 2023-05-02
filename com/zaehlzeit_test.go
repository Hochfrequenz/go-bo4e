package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

const foo = "foo"
const bar = "bar"

func (s *Suite) Test_Zaehlzeit_Deserialization() {
	_foo := foo
	_bar := bar
	zaehlzeit := com.Zaehlzeit{
		Zaehlzeitdefinition: &_foo,
		Zaehlzeitregister:   &_bar,
	}
	serializedZaehlzeit, err := json.Marshal(zaehlzeit)
	then.AssertThat(s.T(), serializedZaehlzeit, is.Not(is.NilArray[byte]()))
	then.AssertThat(s.T(), err, is.Nil())
	deserializedZaehlzeit := com.Zaehlzeit{}
	err = json.Unmarshal(serializedZaehlzeit, &deserializedZaehlzeit)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZaehlzeit, is.EqualTo(zaehlzeit))
}

func (s *Suite) Test_Successful_Zaehlzeit_Validation() {
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
	VerfiySuccessfulValidations(s, validate, validAbweichung)
}

func (s *Suite) Test_Serialized_Empty_Zaehlzeit_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Zaehlzeit{})
}
