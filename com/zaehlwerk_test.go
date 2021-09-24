package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"strings"
)

// TestVertragsteilDeserialization deserializes a Vertragsteil json
func (s *Suite) TestZaehlwerkDeserialization() {
	var zaehlwerk = Zaehlwerk{
		ZaehlwerkId:    "1",
		Bezeichnung:    "bestes Zählwerk",
		Richtung:       energierichtung.Aussp,
		ObisKennzahl:   "1-0:1.8.0",
		Wandlerfaktor:  0,
		Einheit:        mengeneinheit.KWH,
		Zaehlerstaende: nil,
	}
	serializedZaehlwerk, err := json.Marshal(zaehlwerk)
	jsonString := string(serializedZaehlwerk)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KWH"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedZaehlwerk, is.Not(is.Nil()))
	var deserializedZaehlwerk Zaehlwerk
	err = json.Unmarshal(serializedZaehlwerk, &deserializedZaehlwerk)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedZaehlwerk, is.EqualTo(zaehlwerk))
}

//  Test_Vertragsteil_Failed_Validation verifies that the validation fails for invalid Vertragsteil
func (s *Suite) Test_Zaehlwerk_Failed_Validation() {
	validate := validator.New()
	invalidVertragsteile := map[string][]interface{}{
		"required": {
			Zaehlwerk{
				ZaehlwerkId:    "",
				Bezeichnung:    "unvollständiges Zaehlwerk",
				ObisKennzahl:   "",
				Zaehlerstaende: nil,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertragsteile)
}

//  Test_Successful_Vertragskonditionen_Validation asserts that the validation does not fail for a valid Vertragskonditionen
func (s *Suite) Test_Successful_Zaehlwerk_Validation() {
	validate := validator.New()
	validZaehlwerke := []interface{}{
		Zaehlwerk{
			ZaehlwerkId:    "1",
			Bezeichnung:    "ZW, Juhee",
			Richtung:       energierichtung.Aussp,
			ObisKennzahl:   "1-0:1.8.0",
			Wandlerfaktor:  0,
			Einheit:        mengeneinheit.KWH,
			Zaehlerstaende: nil,
		},
	}
	VerfiySuccessfulValidations(s, validate, validZaehlwerke)
}
