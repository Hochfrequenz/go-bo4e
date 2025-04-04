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
func Test_Address_Deserialization(t *testing.T) {
	var adresse = com.Adresse{
		Postleitzahl: internal.Ptr("82031"),
		Ort:          internal.Ptr("Grünwald"),
		Ortsteil:     internal.Ptr("Geiselgasteig"),
		Strasse:      internal.Ptr("Nördlicher Münchner Straße"),
		Hausnummer:   internal.Ptr("27A"),
		Landescode:   internal.Ptr(landescode.DE),
	}
	serializedAdresse, err := json.Marshal(adresse)
	jsonString := string(serializedAdresse)
	then.AssertThat(t, strings.Contains(jsonString, "DE"), is.True())                              // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, strconv.Itoa(int(landescode.DE))), is.False()) // no raw int enum value
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedAdresse, is.Not(is.NilArray[byte]()))
	var deserializedAdresse com.Adresse
	err = json.Unmarshal(serializedAdresse, &deserializedAdresse)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedAdresse, is.EqualTo(adresse))
}

// Test_StrasseXorPostfach_Validation verifies that the Strasse XOR Postfach validation works
func Test_Strasse_XorPostfachValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(com.AdresseStructLevelValidation, com.Adresse{})
	invalidAdressMaps := map[string][]interface{}{
		"StrasseRequiresHausnummer": {
			com.Adresse{
				// no hausnummer given
				Postleitzahl: internal.Ptr("82031"),
				Ort:          internal.Ptr("Grünwald"),
				Strasse:      internal.Ptr("Nördlicher Münchner Straße"),
				Landescode:   internal.Ptr(landescode.DE),
			},
			com.Adresse{
				// no strasse given
				Postleitzahl: internal.Ptr("82031"),
				Ort:          internal.Ptr("Grünwald"),
				Hausnummer:   internal.Ptr("27A"),
				Landescode:   internal.Ptr(landescode.DE),
			},
		},
		"StrasseXORPostfach": {
			com.Adresse{
				// both postfach and strasse+hausnummer given
				Postfach:     internal.Ptr("Foo-Fach"),
				Postleitzahl: internal.Ptr("82031"),
				Ort:          internal.Ptr("Grünwald"),
				Strasse:      internal.Ptr("Nördlicher Münchner Straße"),
				Hausnummer:   internal.Ptr("27A"),
				Landescode:   internal.Ptr(landescode.DE),
			},
			com.Adresse{
				// neither postfach not strasse+hausnummer given
				Postleitzahl: internal.Ptr("82031"),
				Ort:          internal.Ptr("Grünwald"),
				Landescode:   internal.Ptr(landescode.DE),
			},
			com.Adresse{
				// no strasse given
				Postleitzahl: internal.Ptr("82031"),
				Ort:          internal.Ptr("Grünwald"),
				Hausnummer:   internal.Ptr("27A"),
				Landescode:   internal.Ptr(landescode.DE),
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidAdressMaps)
}

// Test_Successful_Validation asserts that the validation does not fail for a valid Adresse
func Test_Successful_Adresse_Validation(t *testing.T) {
	validate := validator.New()
	validAddresses := []interface{}{
		com.Adresse{
			Postleitzahl: internal.Ptr("82031"),
			Ort:          internal.Ptr("Grünwald"),
			Hausnummer:   internal.Ptr("27A"),
			Strasse:      internal.Ptr("Nördliche Münchner Straße"),
			Landescode:   internal.Ptr(landescode.DE),
		},
	}
	VerifySuccessfulValidations(t, validate, validAddresses)
}

func Test_Serialized_Empty_Adresse_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Adresse{})
}
