package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"github.com/shopspring/decimal"
	"strings"
)

// Test_Deserialization deserializes a Zeitreihenwertkompakt json
func (s *Suite) Test_Zeitreihenwertkompakt_Deserialization() {
	var zeitreihenwertkompakt = com.Zeitreihenwertkompakt{
		Wert:         decimal.NewFromFloat(17.32),
		Status:       messwertstatus.ERSATZWERT,
		Statuszusatz: messwertstatuszusatz.Z77_SPANNUNGSAUSFALL,
	}
	serializedZeitreihenwert, err := json.Marshal(zeitreihenwertkompakt)
	jsonString := string(serializedZeitreihenwert)
	then.AssertThat(s.T(), strings.Contains(jsonString, "ERSATZWERT"), is.True())           // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "Z77_SPANNUNGSAUSFALL"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZeitreihenwert, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Zeitreihenwertkompakt
	err = json.Unmarshal(serializedZeitreihenwert, &deserializedZeitreihenwert)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZeitreihenwert, is.EqualTo(zeitreihenwertkompakt))
}

// Test_Zeitreihenwertkompakt_Failed_Validation verifies that the validation fails for invalid Zeitreihenwertkompakt s
func (s *Suite) Test_Zeitreihenwertkompakt_Failed_Validation() {
	validate := validator.New()
	invalidZeitreihenwertkompakts := map[string][]interface{}{
		"required": {
			// Zeitreihenwertkompakt{},
			// todo: find out how to validate empty decimals as required.
		},
	}
	VerfiyFailedValidations(s, validate, invalidZeitreihenwertkompakts)
}

// Test_Successful_Validation asserts that the validation does not fail for a valid Zeitreihenwertkompakt
func (s *Suite) Test_Successful_Zeitreihenwertkompakt_Validation() {
	validate := validator.New()
	validAddresses := []interface{}{
		com.Zeitreihenwertkompakt{
			Wert:         decimal.NewFromFloat(0),
			Status:       0,
			Statuszusatz: 0,
		},
	}
	VerfiySuccessfulValidations(s, validate, validAddresses)
}

func (s *Suite) Test_Serialized_Empty_Zeitreihenwertkompakt_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Zeitreihenwertkompakt{})
}
