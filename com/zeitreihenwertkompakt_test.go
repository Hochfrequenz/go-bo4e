package com_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"github.com/shopspring/decimal"
)

// Test_Deserialization deserializes a Zeitreihenwertkompakt json
func Test_Zeitreihenwertkompakt_Deserialization(t *testing.T) {
	var zeitreihenwertkompakt = com.Zeitreihenwertkompakt{
		Wert:         decimal.NewFromFloat(17.32),
		Status:       messwertstatus.ERSATZWERT,
		Statuszusatz: messwertstatuszusatz.Z77_SPANNUNGSAUSFALL,
	}
	serializedZeitreihenwert, err := json.Marshal(zeitreihenwertkompakt)
	jsonString := string(serializedZeitreihenwert)
	then.AssertThat(t, strings.Contains(jsonString, "ERSATZWERT"), is.True())           // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "Z77_SPANNUNGSAUSFALL"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedZeitreihenwert, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Zeitreihenwertkompakt
	err = json.Unmarshal(serializedZeitreihenwert, &deserializedZeitreihenwert)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZeitreihenwert, is.EqualTo(zeitreihenwertkompakt))
}

// Test_Zeitreihenwertkompakt_Failed_Validation verifies that the validation fails for invalid Zeitreihenwertkompakt s
func Test_Zeitreihenwertkompakt_Failed_Validation(t *testing.T) {
	validate := validator.New()
	invalidZeitreihenwertkompakts := map[string][]interface{}{
		"required": {
			// Zeitreihenwertkompakt{},
			// todo: find out how to validate empty decimals as required.
		},
	}
	VerifyFailedValidations(t, validate, invalidZeitreihenwertkompakts)
}

// Test_Successful_Validation asserts that the validation does not fail for a valid Zeitreihenwertkompakt
func Test_Successful_Zeitreihenwertkompakt_Validation(t *testing.T) {
	validate := validator.New()
	validAddresses := []interface{}{
		com.Zeitreihenwertkompakt{
			Wert:         decimal.NewFromFloat(0),
			Status:       0,
			Statuszusatz: 0,
		},
	}
	VerifySuccessfulValidations(t, validate, validAddresses)
}

func (s *Suite) Test_Serialized_Empty_Zeitreihenwertkompakt_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Zeitreihenwertkompakt{})
}
