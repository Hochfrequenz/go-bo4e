package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"strings"
	"time"
)

// TestZeitreihenwertDeserialization deserializes a Zeitreihenwert json
func (s *Suite) TestZeitreihenwertDeserialization() {
	var zeitreihenwert = Zeitreihenwert{
		Zeitreihenwertkompakt: Zeitreihenwertkompakt{
			Wert:         17.43,
			Status:       messwertstatus.ENERGIEMENGESUMMIERT,
			Statuszusatz: messwertstatuszusatz.Z81_MESSEINRICHTUNGGESTOERT_DEFEKT,
		},
		DatumUhrzeitVon: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		DatumUhrzeitBis: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
	}
	serializedZeitreihenwert, err := json.Marshal(zeitreihenwert)
	jsonString := string(serializedZeitreihenwert)
	then.AssertThat(s.T(), strings.Contains(jsonString, "ENERGIEMENGESUMMIERT"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZeitreihenwert, is.Not(is.Nil()))
	var deserializedZeitreihenwert Zeitreihenwert
	err = json.Unmarshal(serializedZeitreihenwert, &deserializedZeitreihenwert)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZeitreihenwert, is.EqualTo(zeitreihenwert))
}

//  Test_Zeitreihenwert_Failed_Validation verifies that the validation fails for invalid Zeitreihenwert s
func (s *Suite) Test_Zeitreihenwert_Failed_Validation() {
	validate := validator.New()
	invalidZeitreihenwertkompakts := map[string][]interface{}{
		"required": {
			Zeitreihenwert{
				Zeitreihenwertkompakt: Zeitreihenwertkompakt{
					Wert:         17,
					Status:       0,
					Statuszusatz: 0,
				},
				DatumUhrzeitVon: time.Time{},
				DatumUhrzeitBis: time.Time{},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidZeitreihenwertkompakts)
}

//  Test_Successful_Zeitreihenwert_Validation asserts that the validation does not fail for a valid Zeitreihenwert
func (s *Suite) Test_Successful_Zeitreihenwert_Validation() {
	validate := validator.New()
	validAddresses := []interface{}{
		Zeitreihenwert{
			Zeitreihenwertkompakt: Zeitreihenwertkompakt{
				Wert:         17.43,
				Status:       messwertstatus.ENERGIEMENGESUMMIERT,
				Statuszusatz: messwertstatuszusatz.Z81_MESSEINRICHTUNGGESTOERT_DEFEKT,
			},
			DatumUhrzeitVon: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			DatumUhrzeitBis: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		},
	}
	VerfiySuccessfulValidations(s, validate, validAddresses)
}
