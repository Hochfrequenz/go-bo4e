package market_communication_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/hochfrequenz/go-bo4e/enum/anfragekategorie"

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
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/hochfrequenz/go-bo4e/market_communication"
	"github.com/shopspring/decimal"
)

func Test_BOneyComb_DeSerialization(t *testing.T) {
	anfrage := market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{},
		Transaktionsdaten: map[string]string{
			"Bar": "Bar",
		},
	}
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
			&bo.Anfrage{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.ANFRAGE,
					VersionStruktur: "1",
				},
				LokationsId:             "015111111",
				Lokationstyp:            lokationstyp.MALO,
				OBISKennzahl:            internal.Ptr("1-0:1.8.1"),
				ZeitraumMesswertanfrage: nil,
				Anfragekategorie:        anfragekategorie.ENTSPERRUNG,
				Anfragetyp:              nil,
			},
		},
		Transaktionsdaten: map[string]string{
			"Foo": "Bar",
		},
		Anfrage: &anfrage,
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
	then.AssertThat(t, *deserializedBoneyComb.Anfrage, is.EqualTo(anfrage))
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
	const dirName = "test_boney_combs"
	jsonFiles, err := os.ReadDir(dirName)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, len(jsonFiles), is.Not(is.EqualTo(0)))

	for _, file := range jsonFiles {
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

func processJSONFilesInDir(dir string, processFunc func(path string) error) error {
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".json" {
			return processFunc(path)
		}

		return nil
	})
}

// submoduleIsCheckedOut returns true if the 'example_market_communication_bo4e_transactions' submodule is checked out
func submoduleIsCheckedOut(dir string) (bool, error) {
	var hasJSONFiles bool
	err := processJSONFilesInDir(dir, func(path string) error {
		hasJSONFiles = true
		// Return an error to stop further traversal
		return filepath.SkipDir
	})

	if err != nil {
		if err == filepath.SkipDir {
			return hasJSONFiles, nil
		}
		return false, err
	}

	return hasJSONFiles, nil
}

// getJsonFilePathsFromSubmodule returns a list of JSON file paths from the specified directory
func getJsonFilePathsFromSubmodule(dir string) ([]string, error) {
	var jsonFiles []string
	err := processJSONFilesInDir(dir, func(path string) error {
		jsonFiles = append(jsonFiles, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return jsonFiles, nil
}

var formatAllowList = []string{"COMDIS", "INVOIC", "MSCONS", "PRICAT", "QUOTES", "REMADV", "UTILMD"}
var formatVersionAllowList = []string{"FV2210", "FV2310", "FV2404", "FV2504"}

func matchesAllowLists(fileName string) bool {
	for _, format := range formatAllowList {
		if !strings.Contains(fileName, format) {
			continue
		}
		for _, formatVersion := range formatVersionAllowList {
			if strings.Contains(fileName, formatVersion) {
				return true
			}
		}
	}
	return false
}

func Test_Submodule_Example_Messages(t *testing.T) {
	const relativeSubmodulePath = "example_market_communication_bo4e_transactions"
	submoduleIsCheckedOut, submoduleCheckoutErr := submoduleIsCheckedOut(relativeSubmodulePath)
	then.AssertThat(t, submoduleCheckoutErr, is.Nil())
	if !submoduleIsCheckedOut {
		// happens if you don't have access to the submodule => exit silently
		return
	}
	jsonFiles, err := getJsonFilePathsFromSubmodule(relativeSubmodulePath)
	then.AssertThat(t, err, is.Nil())
	for _, fileName := range jsonFiles {
		isBoneyCombFile := strings.HasSuffix(fileName, ".bo.json")
		if !isBoneyCombFile || !matchesAllowLists(fileName) {
			continue
		}
		fileContent, readErr := os.ReadFile(filepath.FromSlash(fileName))
		then.AssertThat(t, readErr, is.Nil())
		then.AssertThat(t, fileContent, is.Not(is.NilArray[byte]()))
		var boneyComb market_communication.BOneyComb
		err = json.Unmarshal(fileContent, &boneyComb)
		if err != nil {
			fmt.Printf("Failed to unmarshal file content from %s: %v\n", fileName, err)
		}
		then.AssertThat(t, err, is.Nil())
	}

}
