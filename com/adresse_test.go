package com

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/stretchr/testify/require"
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
	_ = Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   Landescode.DE,
	}
	then.Assert
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
