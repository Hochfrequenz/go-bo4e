package market_communication

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

// SetupSuite sets up the tests
func (s *Suite) SetupSuite() {
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_BOneyComb_DeSerialization() {
	boneyComb := BOneyComb{
		Stammdaten: []bo.BusinessObject{
			bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.Geschaeftspartner,
					VersionStruktur: "1",
				},
				Anrede: anrede.Frau,
				Name1:  "Musterfrau",
				Name2:  "Erika",
			},
		},
		Transaktionsdaten: map[string]interface{}{
			"Foo": "Bar",
		},
	}
	serializedBoneyComb, err := json.Marshal(boneyComb)
	//jsonString := string(serializedBoneyComb)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedBoneyComb, is.Not(is.Nil()))
	var deserializedGp BOneyComb
	err = json.Unmarshal(serializedBoneyComb, &deserializedGp)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedGp, is.EqualTo(boneyComb))
}
