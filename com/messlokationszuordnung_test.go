package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
	"strings"
	"time"
)

// Test_Messlokationszuordnung_Deserialization deserializes an MesslokationszuordnungMesslokationszuordnung json
func (s *Suite) Test_Messlokationszuordnung_Deserialization() {
	var messlokationszuordnung = com.Messlokationszuordnung{
		MesslokationsId: "DE0123456789012345678901234567890",
		Arithmetik:      arithmetischeoperation.ADDITION,
		GueltigSeit:     time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
		GueltigBis:      time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}
	then.AssertThat(s.T(), messlokationszuordnung, is.Not(is.Nil()))

	serializedMesslokationszuordnung, err := json.Marshal(messlokationszuordnung)
	then.AssertThat(s.T(), serializedMesslokationszuordnung, is.Not(is.Nil()))

	jsonString := string(serializedMesslokationszuordnung)
	then.AssertThat(s.T(), strings.Contains(jsonString, "ADDITION"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())

	var deserializedMesslokationszuordnung com.Messlokationszuordnung
	err = json.Unmarshal(serializedMesslokationszuordnung, &deserializedMesslokationszuordnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMesslokationszuordnung, is.EqualTo(messlokationszuordnung))
}

// Test_Failed_Validation asserts that the validation fails for invalid Messlokationszuordnung
func (s *Suite) Test_Failed_MesslokationszuordnungValidation() {
	validate := validator.New()
	invalidMesslokationszuordnung := map[string][]interface{}{
		"required": {
			com.Messlokationszuordnung{},
		},
		"len": {
			com.Messlokationszuordnung{
				MesslokationsId: "not33",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMesslokationszuordnung)
}

//  Test_Successful_Validation asserts that the validation does not fail for a valid Messlokationszuordnung
func (s *Suite) Test_Successful_Messlokationszuordnung_Validation() {
	validate := validator.New()
	validMesslokationszuordnung := []interface{}{
		com.Messlokationszuordnung{
			MesslokationsId: "DE0123456789012345678901234567890",
			Arithmetik:      arithmetischeoperation.ADDITION,
		},
		com.Messlokationszuordnung{
			MesslokationsId: "DE0123456789012345678901234567890",
			Arithmetik:      arithmetischeoperation.ADDITION,
			GueltigSeit:     time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
			GueltigBis:      time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	VerfiySuccessfulValidations(s, validate, validMesslokationszuordnung)
}
