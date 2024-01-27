package com_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/shopspring/decimal"
)

// TestZaehlwerkDeserialization deserializes a Zaehlwerk json
func (s *Suite) Test_Zaehlwerk_Deserialization() {
	// Warning: This test makes use of global variables of a package (3rd-party). It must not be parallelized!

	decimal.MarshalJSONWithoutQuotes = true

	var zaehlwerk = com.Zaehlwerk{
		ZaehlwerkId:    "1",
		Bezeichnung:    "bestes Zählwerk",
		Richtung:       energierichtung.AUSSP,
		ObisKennzahl:   "1-0:1.8.0",
		Wandlerfaktor:  decimal.NewFromFloat(1.2),
		Einheit:        mengeneinheit.KWH,
		Zaehlerstaende: nil,
	}
	serializedZaehlwerk, err := json.Marshal(zaehlwerk)
	jsonString := string(serializedZaehlwerk)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KWH"), is.True())                 // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "AUSSP"), is.True())               // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "wandlerfaktor\":1.2"), is.True()) // no quotes around the decimal
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZaehlwerk, is.Not(is.NilArray[byte]()))
	var deserializedZaehlwerk com.Zaehlwerk
	err = json.Unmarshal(serializedZaehlwerk, &deserializedZaehlwerk)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZaehlwerk.ExtensionData.CompareTo(zaehlwerk.ExtensionData), is.True())
	zaehlwerk.ExtensionData = deserializedZaehlwerk.ExtensionData
	then.AssertThat(s.T(), deserializedZaehlwerk, is.EqualTo(zaehlwerk))
}

// Test_Zaehlwerk_Decimal_As_Number_Serialization serializes a Zaehlwerk, once with the wandlerfaktor as string (default), once as number
func (s *Suite) Test_Zaehlwerk_Decimal_As_Number_Serialization() {
	// Warning: This test makes use of global variables of a package (3rd-party). It must not be parallelized!

	var zaehlwerk = com.Zaehlwerk{
		Wandlerfaktor: decimal.NewFromFloat(1.2),
	}

	decimal.MarshalJSONWithoutQuotes = false
	serializedZaehlwerk, err := json.Marshal(zaehlwerk)
	then.AssertThat(s.T(), err, is.Nil())
	jsonString := string(serializedZaehlwerk)
	then.AssertThat(s.T(), strings.Contains(jsonString, "wandlerfaktor\":\"1.2\""), is.True()) // decimal as string by default: https://github.com/shopspring/decimal/issues/21

	decimal.MarshalJSONWithoutQuotes = true // https://github.com/shopspring/decimal/blob/fa3b22f4d484d626ee81919285cf3d22ad3a4000/decimal.go#L47
	serializedZaehlwerk, err = json.Marshal(zaehlwerk)
	then.AssertThat(s.T(), err, is.Nil())
	jsonString = string(serializedZaehlwerk)
	then.AssertThat(s.T(), strings.Contains(jsonString, "wandlerfaktor\":1.2"), is.True()) // decimal as number
}

// Test_Zaehlwerk_Failed_Validation verifies that the validation fails for invalid Zaehlwerk
func (s *Suite) Test_Zaehlwerk_Failed_Validation() {
	validate := validator.New()
	invalidVertragsteile := map[string][]interface{}{
		"required": {
			com.Zaehlwerk{
				ZaehlwerkId:    "",
				Bezeichnung:    "unvollständiges Zaehlwerk",
				ObisKennzahl:   "",
				Zaehlerstaende: nil,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertragsteile)
}

// Test_Successful_Zaehlwerk_Validation asserts that the validation does not fail for a valid Zaehlwerk
func (s *Suite) Test_Successful_Zaehlwerk_Validation() {
	validate := validator.New()
	validZaehlwerke := []interface{}{
		com.Zaehlwerk{
			ZaehlwerkId:    "1",
			Bezeichnung:    "ZW, Juhee",
			Richtung:       energierichtung.AUSSP,
			ObisKennzahl:   "1-0:1.8.0",
			Wandlerfaktor:  decimal.NewFromFloat(1.0), // Mit diesem Faktor wird eine Zählerstandsdifferenz multipliziert, um zum eigentlichen Verbrauch im Zeitraum zu kommen. => must not be empty
			Einheit:        mengeneinheit.KWH,
			Zaehlerstaende: nil,
		},
	}
	VerfiySuccessfulValidations(s, validate, validZaehlwerke)
}

func Test_Serialized_Empty_Zaehlwerk_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Zaehlwerk{})
}
