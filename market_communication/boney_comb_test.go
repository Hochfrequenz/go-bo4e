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
					BoTyp:           botyp.Geschaeftspartner,
					VersionStruktur: "1",
				},
				Anrede: anrede.Frau,
				Name1:  "Musterfrau",
				Name2:  "Erika",
			},
			&bo.Zaehler{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.Zaehler,
					VersionStruktur: "1",
				},
				Sparte:             sparte.Strom,
				Zaehlernummer:      "1ASD23",
				Zaehlerauspraegung: zaehlerauspraegung.Einrichtungszaehler,
				Zaehlertyp:         zaehlertyp.Drehkolbenzaehler,
				Tarifart:           tarifart.Eintarif,
			},
			&bo.Energiemenge{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.Energiemenge,
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
		Transaktionsdaten: map[string]interface{}{
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

func (s *Suite) Test_BOneyComb_DeSerialization_With_Invalid_BoTyp() {
	jsonWithInvalidBoTyp := "{\"Stammdaten\":[{\"boTyp\":\"foo\"}], \"Transaktionsdaten\":{}}" // foo is not a valid boTyp => deserialization should fail
	var deserializedBoneyComb market_communication.BOneyComb
	err := json.Unmarshal([]byte(jsonWithInvalidBoTyp), &deserializedBoneyComb)
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) Test_BOneyComb_DeSerialization_With_Unimplemented_BoTyp() {
	jsonWithUnimplementedBoTyp := "{\"Stammdaten\":[{\"boTyp\":\"PreisblattUmlagen\"}], \"Transaktionsdaten\":{}}" // PreisblattUmlagen is not (yet) an implemented boTyp => deserialization should fail
	var deserializedBoneyComb market_communication.BOneyComb
	err := json.Unmarshal([]byte(jsonWithUnimplementedBoTyp), &deserializedBoneyComb)
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
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
		Transaktionsdaten: map[string]interface{}{"foo": "bar"},
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
