package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"strings"
)

// Test_Deserialization deserializes an address json
func (s *Suite) TestAddressDeserialization() {
	var adresse = Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   landescode.DE,
	}
	serializedAdresse, err := json.Marshal(adresse)
	jsonString := string(serializedAdresse)
	then.AssertThat(s.T(), strings.Contains(jsonString, "DE"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedAdresse, is.Not(is.Nil()))
	var deserializedAdresse Adresse
	err = json.Unmarshal(serializedAdresse, &deserializedAdresse)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedAdresse, is.EqualTo(adresse))
}

//  Test_StrasseXorPostfach_Validation verifies that the Strasse XOR Postfach validation works
func (s *Suite) TestStrasseXorPostfachValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(AdresseStructLevelValidation, Adresse{})
	invalidAdressMaps := map[string][]interface{}{
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
	VerfiyFailedValidations(s, validate, invalidAdressMaps)
}

//  Test_Successful_Validation asserts that the validation does not fail for a valid Adresse
func (s *Suite) Test_Successful_Adresse_Validation() {
	validate := validator.New()
	validAddresses := []interface{}{
		Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Hausnummer:   "27A",
			Strasse:      "Nördliche Münchner Straße",
			Landescode:   landescode.DE,
		},
	}
	VerfiySuccessfulValidations(s, validate, validAddresses)
}
