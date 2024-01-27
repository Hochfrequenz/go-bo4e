package com_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"github.com/shopspring/decimal"
)

// TestZeitreihenwertDeserialization deserializes a Zeitreihenwert json
func Test_Zeitreihenwert_Deserialization(t *testing.T) {
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
	then.AssertThat(t, strings.Contains(jsonString, "ENERGIEMENGESUMMIERT"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedZeitreihenwert, is.Not(is.NilArray[byte]()))
	var deserializedZeitreihenwert com.Zeitreihenwert
	err = json.Unmarshal(serializedZeitreihenwert, &deserializedZeitreihenwert)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZeitreihenwert, is.EqualTo(zeitreihenwert))
}

// Test_Zeitreihenwert_Failed_Validation verifies that the validation fails for invalid Zeitreihenwert s
func Test_Zeitreihenwert_Failed_Validation(t *testing.T) {
	validate := validator.New()
	invalidZeitreihenwertkompakts := map[string][]interface{}{
		"required": {
			com.Zeitreihenwert{
				Zeitreihenwertkompakt: com.Zeitreihenwertkompakt{
					Wert:         decimal.NewFromFloat(17),
					Status:       0,
					Statuszusatz: 0,
				},
				DatumUhrzeitVon: time.Time{},
				DatumUhrzeitBis: time.Time{},
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidZeitreihenwertkompakts)
}

// Test_Successful_Zeitreihenwert_Validation asserts that the validation does not fail for a valid Zeitreihenwert
func Test_Successful_Zeitreihenwert_Validation(t *testing.T) {
	validate := validator.New()
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
	VerifySuccessfulValidations(t, validate, validAddresses)
}

func Test_Serialized_Empty_Zeitreihenwert_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Zeitreihenwert{})
}
