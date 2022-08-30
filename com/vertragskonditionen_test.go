package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

var validVertragskonditionen = com.Vertragskonditionen{
	Beschreibung:     "hallo",
	AnzahlAbschlaege: 17,
	Vertragslaufzeit: &com.Zeitraum{
		Einheit: zeiteinheit.MINUTE,
		Dauer:   decimal.NewNullDecimal(decimal.NewFromFloat(15)),
	},
	Kuendigungsfrist: &com.Zeitraum{
		Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Endzeitpunkt:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		Startdatum:     time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Enddatum:       time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
	},
	Vertragsverlaengerung: &com.Zeitraum{
		Startzeitpunkt: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Endzeitpunkt:   time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		Startdatum:     time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Enddatum:       time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
	},
	Abschlagszyklus: &com.Zeitraum{
		Einheit: zeiteinheit.JAHR,
		Dauer:   decimal.NewNullDecimal(decimal.NewFromFloat(7)),
	},
}

// TestVertragskonditionenDeserialization deserializes a Vertragskonditionen json
func (s *Suite) Test_Vertragskonditionen_Deserialization() {
	serializedVertragskonditionen, err := json.Marshal(validVertragskonditionen)
	jsonString := string(serializedVertragskonditionen)
	then.AssertThat(s.T(), strings.Contains(jsonString, "MINUTE"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedVertragskonditionen, is.Not(is.Nil()))
	var deserializedVertragskonditionen com.Vertragskonditionen
	err = json.Unmarshal(serializedVertragskonditionen, &deserializedVertragskonditionen)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVertragskonditionen, is.EqualTo(validVertragskonditionen))
}

// Test_Vertragskonditionen_Failed_Validation verifies that the validation fails for invalid Vertragskonditionen s
func (s *Suite) Test_Vertragskonditionen_Failed_Validation() {
	validate := validator.New()
	invalidZeitraums := map[string][]interface{}{
		"required_with": {
			com.Vertragskonditionen{
				// is only invalid if a zeitraum is invalid
				Vertragslaufzeit: &com.Zeitraum{
					Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidZeitraums)
}

// Test_Successful_Vertragskonditionen_Validation asserts that the validation does not fail for a valid Vertragskonditionen
func (s *Suite) Test_Successful_Vertragkonditionen_Validation() {
	validate := validator.New()
	validVertragskonditionens := []interface{}{
		validVertragskonditionen,
	}
	VerfiySuccessfulValidations(s, validate, validVertragskonditionens)
}

func (s *Suite) Test_Serialized_Empty_Vertragskonditionen_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Vertragskonditionen{})
}
