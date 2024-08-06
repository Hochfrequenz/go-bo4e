package market_communication_test

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
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
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/hochfrequenz/go-bo4e/market_communication"
	"github.com/shopspring/decimal"
)

func Test_BOneyComb_DeSerialization(t *testing.T) {
	boneyComb := market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{
			&bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.GESCHAEFTSPARTNER,
					VersionStruktur: "1",
				},
				Anrede: internal.Ptr(anrede.FRAU),
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
				Zaehlerauspraegung: internal.Ptr(zaehlerauspraegung.EINRICHTUNGSZAEHLER),
				Zaehlertyp:         internal.Ptr(zaehlertyp.DREHKOLBENZAEHLER),
				Tarifart:           internal.Ptr(tarifart.EINTARIF),
			},
			&bo.Energiemenge{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.ENERGIEMENGE,
					VersionStruktur: "1",
				},
				LokationsId:  "54321012345",
				LokationsTyp: internal.Ptr(lokationstyp.MALO),
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
	then.AssertThat(t, validationErr, is.Nil())
	serializedBoneyComb, err := json.Marshal(boneyComb)
	//jsonString := string(serializedBoneyComb)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedBoneyComb, is.Not(is.NilArray[byte]()))
	var deserializedBoneyComb market_communication.BOneyComb
	err = json.Unmarshal(serializedBoneyComb, &deserializedBoneyComb)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedBoneyComb.Stammdaten[0], is.EqualTo(boneyComb.Stammdaten[0]))
	then.AssertThat(t, deserializedBoneyComb.Stammdaten[1], is.EqualTo(boneyComb.Stammdaten[1]))
	then.AssertThat(t, deserializedBoneyComb.Stammdaten[2], is.EqualTo(boneyComb.Stammdaten[2]))
	then.AssertThat(t, deserializedBoneyComb.Transaktionsdaten, is.EqualTo(boneyComb.Transaktionsdaten))
	then.AssertThat(t, deserializedBoneyComb, is.EqualTo(boneyComb))
}

func Test_Empty_BOneyComb_Is_Invalid(t *testing.T) {
	validate := validator.New()
	var emptyBoneyComb = market_communication.BOneyComb{}
	err := validate.Struct(emptyBoneyComb)
	then.AssertThat(t, err, is.Not(is.Nil()))
}

func Test_BOneyComb_Without_Pruefi_Is_Invalid(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(market_communication.PruefidentifikatorInTransaktionsdatenValidation, market_communication.BOneyComb{})
	var emptyBoneyComb = market_communication.BOneyComb{}
	err := validate.Struct(emptyBoneyComb)
	then.AssertThat(t, err, is.Not(is.Nil()))
}

func Test_BOneyComb_With_Wrong_Pruefi_Is_Invalid(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(market_communication.PruefidentifikatorInTransaktionsdatenValidation, market_communication.BOneyComb{})
	var boneyComb = market_communication.BOneyComb{}
	boneyComb.SetPruefidentifikator("asdfg")
	err := validate.Struct(boneyComb)
	then.AssertThat(t, err, is.Not(is.Nil()))
}

func Test_BOneyComb_With_Correct_Pruefi_Is_Valid(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(market_communication.PruefidentifikatorInTransaktionsdatenValidation, market_communication.BOneyComb{})
	var boneyComb = market_communication.BOneyComb{}
	boneyComb.SetPruefidentifikator("11042")
	boneyComb.Stammdaten = []bo.BusinessObject{} // empty slice because nil is invalid
	err := validate.Struct(boneyComb)
	then.AssertThat(t, err, is.Nil())
}

func Test_Empty_BOneyComb_With_Empty_Stammdaten_Is_Serializable(t *testing.T) {
	var boneyCombWithEmptyStammdaten = market_communication.BOneyComb{
		Stammdaten:        []bo.BusinessObject{},
		Transaktionsdaten: map[string]string{"foo": "bar"},
	}
	serializedBoneyComb, err := json.Marshal(boneyCombWithEmptyStammdaten)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedBoneyComb, is.Not(is.NilArray[byte]()))
	var deserializedBoneyComb market_communication.BOneyComb
	err = json.Unmarshal(serializedBoneyComb, &deserializedBoneyComb)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedBoneyComb.Stammdaten, is.EqualTo(boneyCombWithEmptyStammdaten.Stammdaten))
	then.AssertThat(t, deserializedBoneyComb.Transaktionsdaten, is.EqualTo(boneyCombWithEmptyStammdaten.Transaktionsdaten))
	then.AssertThat(t, deserializedBoneyComb, is.EqualTo(boneyCombWithEmptyStammdaten))
}

// Test_BOneyComb_Deserialization loops over the test_boney_combs directory and tries to deserialize all the json files there as boneycomb
func Test_BOneyComb_Deserialization(t *testing.T) {
	const dirName = "../../example_market_communication_bo4e_transactions/UTILMD/FV2310" // assuming you have checked out the example-repo side by side with go-bo4e
	jsonFiles, err := os.ReadDir(dirName)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, len(jsonFiles), is.Not(is.EqualTo(0)))

	for _, file := range jsonFiles {
		if !strings.HasSuffix(file.Name(), ".bo.json") {
			continue
		}
		fileContent, readErr := os.ReadFile(filepath.FromSlash(dirName + "/" + file.Name()))
		then.AssertThat(t, readErr, is.Nil())
		then.AssertThat(t, fileContent, is.Not(is.NilArray[byte]()))
		var boneyComb market_communication.BOneyComb
		err = json.Unmarshal(fileContent, &boneyComb)
		then.AssertThat(t, err, is.Nil())
		//then.AssertThat(t, boneyComb, is.Not(is.Nil())) // boneyComb is not nullable
		if file.Name() == "20211011.json" {
			firstMaLo, _ := boneyComb.GetSingle(botyp.MARKTLOKATION)
			then.AssertThat(t, firstMaLo.(*bo.Marktlokation).ExtensionData["some untyped prop"], is.EqualTo[any]("hello world"))
			then.AssertThat(t, boneyComb.Links["foo"][1], is.EqualTo("baz"))
		}
		if file.Name() == "20220926.json" {
			firstRechnung, _ := boneyComb.GetSingle(botyp.RECHNUNG)
			expectedBuchungsdatum := time.Date(2022, 6, 01, 13, 37, 0, 0, time.UTC)
			actualBuchungsdatum := firstRechnung.(*bo.Rechnung).Buchungsdatum
			then.AssertThat(t, actualBuchungsdatum.Unix(), is.EqualTo(expectedBuchungsdatum.Unix()))
		}
	}
}
