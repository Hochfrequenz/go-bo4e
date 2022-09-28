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
	"os"
	"path/filepath"
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
				Anrede: anrede.FRAU,
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
				Zaehlerauspraegung: zaehlerauspraegung.EINRICHTUNGSZAEHLER,
				Zaehlertyp:         zaehlertyp.DREHKOLBENZAEHLER,
				Tarifart:           tarifart.EINTARIF,
			},
			&bo.Energiemenge{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.ENERGIEMENGE,
					VersionStruktur: "1",
				},
				LokationsId:  "54321012345",
				LokationsTyp: lokationstyp.MALO,
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

func (s *Suite) Test_BOneyComb_Without_Pruefi_Is_Invalid() {
	validate := validator.New()
	validate.RegisterStructValidation(market_communication.PruefidentifikatorInTransaktionsdatenValidation, market_communication.BOneyComb{})
	var emptyBoneyComb = market_communication.BOneyComb{}
	err := validate.Struct(emptyBoneyComb)
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) Test_BOneyComb_With_Wrong_Pruefi_Is_Invalid() {
	validate := validator.New()
	validate.RegisterStructValidation(market_communication.PruefidentifikatorInTransaktionsdatenValidation, market_communication.BOneyComb{})
	var boneyComb = market_communication.BOneyComb{}
	boneyComb.SetPruefidentifikator("asdfg")
	err := validate.Struct(boneyComb)
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) Test_BOneyComb_With_Correct_Pruefi_Is_Valid() {
	validate := validator.New()
	validate.RegisterStructValidation(market_communication.PruefidentifikatorInTransaktionsdatenValidation, market_communication.BOneyComb{})
	var boneyComb = market_communication.BOneyComb{}
	boneyComb.SetPruefidentifikator("11042")
	boneyComb.Stammdaten = []bo.BusinessObject{} // empty slice because nil is invalid
	err := validate.Struct(boneyComb)
	then.AssertThat(s.T(), err, is.Nil())
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

// Test_BOneyComb_Deserialization loops over the test_boney_combs directory and tries to deserialize all the json files there as boneycomb
func (s *Suite) Test_BOneyComb_Deserialization() {
	const dirName = "test_boney_combs"
	jsonFiles, err := os.ReadDir(dirName)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), len(jsonFiles), is.Not(is.EqualTo(0)))

	for _, file := range jsonFiles {
		fileContent, readErr := os.ReadFile(filepath.FromSlash(dirName + "/" + file.Name()))
		then.AssertThat(s.T(), readErr, is.Nil())
		then.AssertThat(s.T(), fileContent, is.Not(is.Nil()))
		var boneyComb market_communication.BOneyComb
		err = json.Unmarshal(fileContent, &boneyComb)
		then.AssertThat(s.T(), err, is.Nil())
		then.AssertThat(s.T(), boneyComb, is.Not(is.Nil()))
		if file.Name() == "20211011.json" {
			firstMaLo, _ := boneyComb.GetSingle(botyp.MARKTLOKATION)
			then.AssertThat(s.T(), firstMaLo.(*bo.Marktlokation).ExtensionData["some untyped prop"], is.EqualTo("hello world"))
			then.AssertThat(s.T(), boneyComb.Links["foo"][1], is.EqualTo("baz"))
		}
		if file.Name() == "20220926.json" {
			firstRechnung, _ := boneyComb.GetSingle(botyp.RECHNUNG)
			expectedBuchungsdatum := time.Date(2022, 6, 01, 13, 37, 0, 0, time.UTC)
			actualBuchungsdatum := firstRechnung.(*bo.Rechnung).Buchungsdatum
			then.AssertThat(s.T(), actualBuchungsdatum.Unix(), is.EqualTo(expectedBuchungsdatum.Unix()))
		}
	}
}
