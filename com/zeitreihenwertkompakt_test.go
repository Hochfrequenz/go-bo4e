package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"strings"
)

// Test_Deserialization deserializes a Zeitreihenwertkompakt json
func (s *Suite) TestZeitreihenwertkompaktDeserialization() {
	var zeitreihenwertkompakt = Zeitreihenwertkompakt{
		Wert:         17.43,
		Status:       messwertstatus.ERSATZWERT,
		Statuszusatz: messwertstatuszusatz.Z77_SPANNUNGSAUSFALL,
	}
	serializedZeitreihenwert, err := json.Marshal(zeitreihenwertkompakt)
	jsonString := string(serializedZeitreihenwert)
	then.AssertThat(s.T(), strings.Contains(jsonString, "ERSATZWERT"), is.True())           // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "Z77_SPANNUNGSAUSFALL"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZeitreihenwert, is.Not(is.Nil()))
	var deserializedZeitreihenwert Zeitreihenwertkompakt
	err = json.Unmarshal(serializedZeitreihenwert, &deserializedZeitreihenwert)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZeitreihenwert, is.EqualTo(zeitreihenwertkompakt))
}

//  Test_Zeitreihenwertkompakt_Failed_Validation verifies that the validation fails for invalid Zeitreihenwertkompakt s
func (s *Suite) Test_Zeitreihenwertkompakt_Failed_Validation() {
	validate := validator.New()
	invalidZeitreihenwertkompakts := map[string][]interface{}{
		"required": {
			Zeitreihenwertkompakt{
				Wert:         0,
				Status:       0,
				Statuszusatz: 0,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidZeitreihenwertkompakts)
}

//  Test_Successful_Validation asserts that the validation does not fail for a valid Zeitreihenwertkompakt
func (s *Suite) Test_Successful_Zeitreihenwertkompakt_Validation() {
	validate := validator.New()
	validAddresses := []interface{}{
		Zeitreihenwertkompakt{
			Wert:         17,
			Status:       0,
			Statuszusatz: 0,
		},
	}
	VerfiySuccessfulValidations(s, validate, validAddresses)
}
