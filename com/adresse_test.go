package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"strings"
)

// Test_Deserialization deserializes an address json
func (s *Suite) Test_Address_Deserialization() {
	var adresse = com.Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   landescode.DE,
	}
	serializedAdresse, err := json.Marshal(adresse)
	jsonString := string(serializedAdresse)
	then.AssertThat(s.T(), strings.Contains(jsonString, "DE"), is.True())  // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "61"), is.False()) // no "61" for DE
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedAdresse, is.Not(is.Nil()))
	var deserializedAdresse com.Adresse
	err = json.Unmarshal(serializedAdresse, &deserializedAdresse)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedAdresse, is.EqualTo(adresse))
}

//  Test_StrasseXorPostfach_Validation verifies that the Strasse XOR Postfach validation works
func (s *Suite) Test_Strasse_XorPostfachValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(com.AdresseStructLevelValidation, com.Adresse{})
	invalidAdressMaps := map[string][]interface{}{
		"StrasseRequiresHausnummer": {
			com.Adresse{
				// no hausnummer given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Landescode:   landescode.DE,
			},
			com.Adresse{
				// no strasse given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
		"StrasseXORPostfach": {
			com.Adresse{
				// both postfach and strasse+hausnummer given
				Postfach:     "Foo-Fach",
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
			com.Adresse{
				// neither postfach not strasse+hausnummer given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Landescode:   landescode.DE,
			},
			com.Adresse{
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
		com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Hausnummer:   "27A",
			Strasse:      "Nördliche Münchner Straße",
			Landescode:   landescode.DE,
		},
	}
	VerfiySuccessfulValidations(s, validate, validAddresses)
}

func (s *Suite) Test_Serialized_Empty_Adresse_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Adresse{})
}

func (s *Suite) assert_Does_Not_Serialize_Default_Enums(com interface{}) {
	jsonBytes, err := json.Marshal(com)
	then.AssertThat(s.T(), err, is.Nil())
	jsonString := string(jsonBytes)
	then.AssertThat(s.T(), strings.Contains(jsonString, "(0)"), is.False())
}
