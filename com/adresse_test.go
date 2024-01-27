package com_test

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/hochfrequenz/go-bo4e/internal"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
)

// Test_Deserialization deserializes an address json
func (s *Suite) Test_Address_Deserialization() {
	var adresse = com.Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Ortsteil:     "Geiselgasteig",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   internal.Ptr(landescode.DE),
	}
	serializedAdresse, err := json.Marshal(adresse)
	jsonString := string(serializedAdresse)
	then.AssertThat(s.T(), strings.Contains(jsonString, "DE"), is.True())                              // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, strconv.Itoa(int(landescode.DE))), is.False()) // no raw int enum value
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedAdresse, is.Not(is.NilArray[byte]()))
	var deserializedAdresse com.Adresse
	err = json.Unmarshal(serializedAdresse, &deserializedAdresse)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedAdresse, is.EqualTo(adresse))
}

// Test_StrasseXorPostfach_Validation verifies that the Strasse XOR Postfach validation works
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
				Landescode:   internal.Ptr(landescode.DE),
			},
			com.Adresse{
				// no strasse given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Hausnummer:   "27A",
				Landescode:   internal.Ptr(landescode.DE),
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
				Landescode:   internal.Ptr(landescode.DE),
			},
			com.Adresse{
				// neither postfach not strasse+hausnummer given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Landescode:   internal.Ptr(landescode.DE),
			},
			com.Adresse{
				// no strasse given
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Hausnummer:   "27A",
				Landescode:   internal.Ptr(landescode.DE),
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidAdressMaps)
}

// Test_Successful_Validation asserts that the validation does not fail for a valid Adresse
func (s *Suite) Test_Successful_Adresse_Validation() {
	validate := validator.New()
	validAddresses := []interface{}{
		com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Hausnummer:   "27A",
			Strasse:      "Nördliche Münchner Straße",
			Landescode:   internal.Ptr(landescode.DE),
		},
	}
	VerfiySuccessfulValidations(s, validate, validAddresses)
}

func Test_Serialized_Empty_Adresse_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Adresse{})
}
