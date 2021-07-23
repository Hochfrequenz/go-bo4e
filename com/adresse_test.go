package com

import (
	"../enum"
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
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
		Landescode:   enum.DE,
	}
	serializedAdresse, err := json.Marshal(adresse)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedAdresse, is.Not(is.Nil()))
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
