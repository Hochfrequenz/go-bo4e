package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

// SetupSuite sets up the tests
func (s *Suite) SetupSuite() {
}

// Test_Deserialization deserializes an address json
func (s *Suite) Test_Deserialization() {
	var adresse = Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   landescode.DE,
	}
	serializedAdresse, err := json.Marshal(adresse)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedAdresse, is.Not(is.Nil()))
	var deserializedAdresse Adresse
	err = json.Unmarshal(serializedAdresse, &deserializedAdresse)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedAdresse, is.EqualTo(adresse))
}

//  Test_StrasseXorPostfach_Validation verifies that the Strasse XOR Postfach validation works
func (s *Suite) Test_StrasseXorPostfach_Validation() {
	validate := validator.New()
	validate.RegisterStructValidation(AdresseStructLevelValidation, Adresse{})
	invalidAdressMaps := map[string][]Adresse{
		"StrasseRequiresHausnummer": {
			Adresse{
				// no hausnummer given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Landescode:   landescode.DE,
			},
			Adresse{
				// no strasse given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
		"StrasseXORPostfach": {
			Adresse{
				// both postfach and strasse+hausnummer given
				Postfach:     "Foo-Fach",
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
			Adresse{
				// neither postfach not strasse+hausnummer given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Landescode:   landescode.DE,
			},
			Adresse{
				// no strasse given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
	}
	for validationTag, invalidAddresses := range invalidAdressMaps {
		for _, invalidAddress := range invalidAddresses {
			err := validate.Struct(invalidAddress)
			then.AssertThat(s.T(), err, is.Not(is.Nil()))
			tagFound := false
			for _, validationError := range err.(validator.ValidationErrors) {
				then.AssertThat(s.T(), validationError, is.Not(is.Nil()))
				// sometimes theres more than one tag/validation error
				if validationError.Tag() == validationTag {
					tagFound = true
					break
				}
			}
			then.AssertThat(s.T(), tagFound, is.True())
		}
	}

}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
