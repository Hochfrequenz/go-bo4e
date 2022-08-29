package com_test

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"strings"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
)

func (s *Suite) Test_Zaehlzeit_Deserialization() {
	foo := "foo"
	bar := "bar"
	zaehlzeit := com.Zaehlzeit{
		Zaehlzeitdefinition: &foo,
		Zaehlzeitregister:   &bar,
	}
	serializedZaehlzeit, err := json.Marshal(zaehlzeit)
	then.AssertThat(s.T(), serializedZaehlzeit, is.Not(is.Nil()))
	jsonString := string(serializedZaehlzeit)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), strings.Contains(jsonString, "AAAA"), is.True())
	then.AssertThat(s.T(), strings.Contains(jsonString, "ZZZZ"), is.False())
	deserializedZaehlzeit := com.Zaehlzeit{}
	err = json.Unmarshal(serializedZaehlzeit, &deserializedZaehlzeit)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZaehlzeit, is.EqualTo(zaehlzeit))
}

func (s *Suite) Test_Successful_Zaehlzeit_Validation() {
	validate := validator.New()
	foo := "foo"
	bar := "bar"
	validAbweichung := []interface{}{
		com.Zaehlzeit{
			Zaehlzeitdefinition: &foo,
			Zaehlzeitregister:   &bar,
		},
		com.Zaehlzeit{
			Zaehlzeitregister: &bar,
		},
		com.Zaehlzeit{
			Zaehlzeitdefinition: &foo,
		},
		com.Zaehlzeit{},
	}
	VerfiySuccessfulValidations(s, validate, validAbweichung)
}

func (s *Suite) Test_Serialized_Empty_Zaehlzeit_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Zaehlzeit{})
}
