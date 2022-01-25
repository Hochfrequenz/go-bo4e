package mengeneinheit_test

import (
	"encoding/json"
	"fmt"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

// SetupSuite sets up the tests
func (s *Suite) SetupSuite() {
}

func (s *Suite) Test_Stueck_ToString() {
	then.AssertThat(s.T(), fmt.Sprintf("%v", mengeneinheit.STUECK), is.EqualTo("STUECK"))
}

type structWithMengeneinheit struct {
	Foo mengeneinheit.Mengeneinheit `json:"foo"`
}

func (s *Suite) Test_ANZAHL_Is_Unmarshaled_To_Stueck() {
	jsonAnzahl := []byte(`{"foo":"ANZAHL"}`)
	jsonStueck := []byte(`{"foo":"STUECK"}`)
	var structAnzahl structWithMengeneinheit
	err := json.Unmarshal(jsonAnzahl, &structAnzahl)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), structAnzahl.Foo, is.EqualTo(mengeneinheit.STUECK))
	// if this test fails, add the following line to mengeneinheit_jsonenums.go line 63:
	// "ANZAHL": STUECK, // I manually added this fallback: Whenever one tries to unmarshal "ANZAHL" from a json it will return the enum value STUECK

	var structStueck structWithMengeneinheit
	err = json.Unmarshal(jsonStueck, &structStueck)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), structStueck.Foo, is.EqualTo(mengeneinheit.STUECK))
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
