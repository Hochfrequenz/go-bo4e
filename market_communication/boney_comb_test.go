package market_communication_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
	"github.com/hochfrequenz/go-bo4e/market_communication"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
	"time"
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
	boneyComb := market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{
			&bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.GESCHAEFTSPARTNER,
					VersionStruktur: "1",
				},
				Anrede: anrede.Frau,
				Name1:  "Musterfrau",
				Name2:  "Erika",
			},
			&bo.Zaehler{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.ZAEHLER,
					VersionStruktur: "1",
				},
				Sparte:             sparte.STROM,
				Zaehlernummer:      "1ASD23",
				Zaehlerauspraegung: zaehlerauspraegung.Einrichtungszaehler,
				Zaehlertyp:         zaehlertyp.Drehkolbenzaehler,
				Tarifart:           tarifart.Eintarif,
			},
			&bo.Energiemenge{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.ENERGIEMENGE,
					VersionStruktur: "1",
				},
				LokationsId:  "54321012345",
				LokationsTyp: lokationstyp.MaLo,
				Verbrauch: []com.Verbrauch{
					{
						Startdatum:               time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
						Enddatum:                 time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
						Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
						Obiskennzahl:             "1-0:1.8.1",
						Wert:                     decimal.NewFromFloat(17),
						Einheit:                  mengeneinheit.KWH,
					},
				},
			},
		},
		Transaktionsdaten: map[string]string{
			"Foo": "Bar",
		},
	}
	validate := validator.New()
	validationErr := validate.Struct(boneyComb)
	then.AssertThat(s.T(), validationErr, is.Nil())
	serializedBoneyComb, err := json.Marshal(boneyComb)
	//jsonString := string(serializedBoneyComb)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedBoneyComb, is.Not(is.Nil()))
	var deserializedBoneyComb market_communication.BOneyComb
	err = json.Unmarshal(serializedBoneyComb, &deserializedBoneyComb)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedBoneyComb.Stammdaten[0], is.EqualTo(boneyComb.Stammdaten[0]))
	then.AssertThat(s.T(), deserializedBoneyComb.Stammdaten[1], is.EqualTo(boneyComb.Stammdaten[1]))
	then.AssertThat(s.T(), deserializedBoneyComb.Stammdaten[2], is.EqualTo(boneyComb.Stammdaten[2]))
	then.AssertThat(s.T(), deserializedBoneyComb.Transaktionsdaten, is.EqualTo(boneyComb.Transaktionsdaten))
	then.AssertThat(s.T(), deserializedBoneyComb, is.EqualTo(boneyComb))
}

func (s *Suite) Test_Empty_BOneyComb_Is_Invalid() {
	validate := validator.New()
	var emptyBoneyComb = market_communication.BOneyComb{}
	err := validate.Struct(emptyBoneyComb)
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) Test_Empty_BOneyComb_With_Empty_Stammdaten_Is_Serializable() {
	var boneyCombWithEmptyStammdaten = market_communication.BOneyComb{
		Stammdaten:        []bo.BusinessObject{},
		Transaktionsdaten: map[string]string{"foo": "bar"},
	}
	serializedBoneyComb, err := json.Marshal(boneyCombWithEmptyStammdaten)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedBoneyComb, is.Not(is.Nil()))
	var deserializedBoneyComb market_communication.BOneyComb
	err = json.Unmarshal(serializedBoneyComb, &deserializedBoneyComb)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedBoneyComb.Stammdaten, is.EqualTo(boneyCombWithEmptyStammdaten.Stammdaten))
	then.AssertThat(s.T(), deserializedBoneyComb.Transaktionsdaten, is.EqualTo(boneyCombWithEmptyStammdaten.Transaktionsdaten))
	then.AssertThat(s.T(), deserializedBoneyComb, is.EqualTo(boneyCombWithEmptyStammdaten))
}

func (s *Suite) Test_GetTransactionData_Returns_Nil_For_Nil_Transaktionsdaten() {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(s.T(), emptyBoneyComb.Transaktionsdaten, is.Nil())
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData("foo"), is.Nil())
}

func (s *Suite) Test_GetTransactionData_Returns_Nil_For_Not_Found_Key_And_Value_Otherwise() {
	var emptyBoneyComb = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"foo": "bar",
		},
	}
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData("asd"), is.Nil())
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData(""), is.Nil())
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData("foo"), is.Not(is.Nil()))
	then.AssertThat(s.T(), *emptyBoneyComb.GetTransactionData("foo"), is.EqualTo("bar"))
}

// Test_BOneyComb_Deserialization loops over the test_boney_combs directory and tries to deserialize all the json files there as boneycomb
func (s *Suite) Test_BOneyComb_Deserialization() {
	const dirName = "test_boney_combs"
	jsonFiles, err := ioutil.ReadDir(dirName)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), len(jsonFiles), is.Not(is.EqualTo(0)))

	for _, file := range jsonFiles {
		fileContent, readErr := ioutil.ReadFile(dirName + "/" + file.Name())
		then.AssertThat(s.T(), readErr, is.Nil())
		then.AssertThat(s.T(), fileContent, is.Not(is.Nil()))
		var boneyComb market_communication.BOneyComb
		err = json.Unmarshal(fileContent, &boneyComb)
		then.AssertThat(s.T(), err, is.Nil())
		then.AssertThat(s.T(), boneyComb, is.Not(is.Nil()))
	}
}
