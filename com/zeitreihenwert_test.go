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
	"log"
	"strings"
	"time"
)

// TestZeitreihenwertDeserialization deserializes a Zeitreihenwert json
func (s *Suite) Test_Zeitreihenwert_Deserialization() {
	var zeitreihenwert = com.Zeitreihenwert{
		Zeitreihenwertkompakt: com.Zeitreihenwertkompakt{
			Wert:         decimal.NewFromFloat(17.43),
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
	then.AssertThat(s.T(), serializedZeitreihenwert, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Zeitreihenwert
	err = json.Unmarshal(serializedZeitreihenwert, &deserializedZeitreihenwert)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZeitreihenwert, is.EqualTo(zeitreihenwert))
}

// Test_Zeitreihenwert_Failed_Validation verifies that the validation fails for invalid Zeitreihenwert s
func (s *Suite) Test_Zeitreihenwert_Failed_Validation() {
	validate := validator.New()
	if err := validate.RegisterValidation("timenotzero", valTimeNotZero); err != nil { //add custom validator tag for time not zero and handle error of registration
		//generate error if registration of validation failed
		log.Fatalln("Failed to register custom validation for checking time != 0:", err)
	}
	invalidZeitreihenwertkompakts := map[string][]interface{}{
		"timenotzero": {
			com.Zeitreihenwert{
				Zeitreihenwertkompakt: com.Zeitreihenwertkompakt{
					Wert:         decimal.NewFromFloat(17),
					Status:       0,
					Statuszusatz: 0,
				},
				DatumUhrzeitVon: time.Time{},
				DatumUhrzeitBis: time.Time{},
			},
			//custom timenotzero and not required because time is initialized with 0 value -> todo: required validator for time.Time
			//todo: implement required dive for Zeitreihenwertkompakt?
		},
	}
	VerfiyFailedValidations(s, validate, invalidZeitreihenwertkompakts)
}

// Test_Successful_Zeitreihenwert_Validation asserts that the validation does not fail for a valid Zeitreihenwert
func (s *Suite) Test_Successful_Zeitreihenwert_Validation() {
	validate := validator.New()
	if err := validate.RegisterValidation("timenotzero", valTimeNotZero); err != nil { //add custom validator tag for time not zero and handle error of registration
		//generate error if registration of validation failed
		log.Fatalln("Failed to register custom validation for checking time != 0:", err)
	}
	validAddresses := []interface{}{
		com.Zeitreihenwert{
			Zeitreihenwertkompakt: com.Zeitreihenwertkompakt{
				Wert:         decimal.NewFromFloat(17.43),
				Status:       messwertstatus.ENERGIEMENGESUMMIERT,
				Statuszusatz: messwertstatuszusatz.Z81_MESSEINRICHTUNGGESTOERT_DEFEKT,
			},
			DatumUhrzeitVon: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			DatumUhrzeitBis: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		},
	}
	VerfiySuccessfulValidations(s, validate, validAddresses)
}

func (s *Suite) Test_Serialized_Empty_Zeitreihenwert_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Zeitreihenwert{})
}
