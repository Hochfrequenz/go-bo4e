package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
	"github.com/hochfrequenz/go-bo4e/enum/bilanzierungsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/gebiettyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/sperrstatus"
	"github.com/hochfrequenz/go-bo4e/enum/verbrauchsart"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
)

// Test_Marktlokation_Deserialization tests serialization and deserialization of Marktlokation
func (s *Suite) Test_Marktlokation_Deserialization() {
	f := false
	gesperrt := sperrstatus.GESPERRT
	var malo = bo.Marktlokation{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.MARKTLOKATION,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
			ExtensionData:     unmappeddatamarshaller.ExtensionData{},
		},
		MarktlokationsId:     "51238696781",
		Sparte:               sparte.STROM,
		Energierichtung:      energierichtung.AUSSP,
		Bilanzierungsmethode: bilanzierungsmethode.RLM,
		Verbrauchsart:        verbrauchsart.KL,
		Unterbrechbar:        &f,
		Netzebene:            netzebene.MSP,
		Netzbetreibercodenr:  "0815",
		Gebiettyp:            gebiettyp.GRUNDVERSORGUNGSGEBIET,
		Netzgebietnr:         "1234",
		Bilanzierungsgebiet:  "Foo",
		Grundversorgercodenr: "5678",
		Endkunde: &bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.GESCHAEFTSPARTNER,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Name1:                "Kruemel",
			Gewerbekennzeichnung: false,
			Partneradresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.DIENSTLEISTER,
			},
		},
		Lokationsadresse: &com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Strasse:      "Nördliche Münchner Straße",
			Hausnummer:   "27A",
			Landescode:   landescode.DE,
		},
		ZugehoerigeMesslokationen: []com.Messlokationszuordnung{
			{
				MesslokationsId: "DE0123456789012345678901234567890",
				Arithmetik:      arithmetischeoperation.ADDITION,
			},
		},
		Netznutzungsabrechnungsdaten: []com.Netznutzungsabrechnungsdaten{
			{
				Gemeinderabatt: decimal.NewNullDecimal(decimal.NewFromFloat(12.34)),
				Zuschlag:       decimal.NewNullDecimal(decimal.NewFromFloat(56.78)),
				Abschlag:       decimal.NewNullDecimal(decimal.NewFromFloat(9.10)),
			},
		},
		Sperrstatus: &gesperrt,
	}
	serializedMalo, err := json.Marshal(malo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMalo, is.Not(is.NilArray[byte]()))
	stringSerializedMalo := string(serializedMalo)
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "STROM"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "AUSSP"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "RLM"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "KL"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "MSP"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "GRUNDVERSORGUNGSGEBIET"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "DE"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "ADDITION"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "\"unterbrechbar\":false"), is.True())
	var deserializedMalo bo.Marktlokation
	err = json.Unmarshal(serializedMalo, &deserializedMalo)
	then.AssertThat(s.T(), err, is.Nil())

	expectedJson, _ := json.Marshal(malo)
	actualJson, _ := json.Marshal(deserializedMalo)
	then.AssertThat(s.T(), expectedJson, is.EqualTo(actualJson))
}

//nolint:dupl // This can only be simplified if we use generics. anything else seems overly complicated but maybe it's just me
func (s *Suite) Test_Marktlokation_DeSerialization_With_Unkonwn_Fields() {
	// this is a json that contains fields/keys that are not part of the official bo4e standard.
	// our expectation is that they are deserialized into the "ExtensionData" field.
	// They should also be serialized from there / not be lost during a marshalling/unmarshalling round trip
	maloJsonWithUnknownFields := `{
      "marktlokationsId": "10024073272",
      "sparte": "STROM",
      "energierichtung": "AUSSP",
      "bilanzierungsmethode": "RLM",
      "netzebene": "NSP",
      "bilanzierungsgebiet": "11YN10000762-01E",
      "lokationsadresse": {
        "postleitzahl": "30926",
        "ort": "Seelze",
        "strasse": "Im Sande",
        "hausnummer": "33",
        "landescode": "DE",
        "ortsteil": "Letter"
      },
      "marktrollen": [
        {
          "marktrolle": "LF"
        },
        {
          "code": "4033872000058",
          "marktrolle": "UENB"
        }
      ],
      "regelzone": "10YDE-EON------1",
      "zeitreihentyp": "SLS",
      "zaehlwerke": [
        {
          "verwendungszwecke": [
            {
              "marktrolle": "LF",
              "zweck": null
            }
          ],
          "richtung": "AUSSP",
          "obisKennzahl": "1-1:1.9.0",
          "wandlerfaktor": 0,
          "schwachlastfaehig": "NICHT_SCHWACHLASTFAEHIG",
          "konzessionsabgabe": {
            "satz": "TA",
            "kosten": 0
          }
        }
      ],
      "messtechnischeEinordnung": "KME_MME",
      "boTyp": "MARKTLOKATION",
      "versionStruktur": "1"
    }`
	malo := bo.Marktlokation{}

	// unmarshalling tests
	err := json.Unmarshal([]byte(maloJsonWithUnknownFields), &malo)
	then.AssertThat(s.T(), err, is.Nil())
	//then.AssertThat(s.T(), malo.ExtensionData, is.Not(is.Nil()))             // ExtensionData is not nullable
	then.AssertThat(s.T(), malo.Zaehlwerke, is.Not(is.NilArray[com.Zaehlwerk]()))        // marktloktion->zaehlwerke is NOT part of the bo4e standard ==> present in extension data
	then.AssertThat(s.T(), malo.ExtensionData["marktlokationsId"], is.EqualTo[any](nil)) // marktlokation->marklokationsId is part of the bo4e standard ==> not present in extension data
	then.AssertThat(s.T(), malo.MarktlokationsId, is.EqualTo("10024073272"))             // but where it should be
	// the other fields should be fine, too, without explicit tests; Add them if you feel like it doesn't work

	// marshaling tests
	serializedMaloBytes, errSerializing := json.Marshal(malo)
	then.AssertThat(s.T(), errSerializing, is.Nil())
	serializedMaLo := string(serializedMaloBytes)
	then.AssertThat(s.T(), strings.Contains(serializedMaLo, "zaehlwerke"), is.True())       // unmapped fields should be part of the serialized malo
	then.AssertThat(s.T(), strings.Contains(serializedMaLo, "marktlokationsId"), is.True()) // mapped fields should be part of the serialized malo
}

// TestFailedMarktlokationValidation verifies that the validators of Marktlokation work
func (s *Suite) Test_Failed_MarktlokationValidation() {
	validate := validator.New()
	registerError := validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(s.T(), registerError, is.Nil())
	validate.RegisterStructValidation(bo.XorStructLevelValidation, bo.Marktlokation{})

	invalidMarktlokationMap := map[string][]interface{}{
		"required": {
			bo.Marktlokation{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				MarktlokationsId:     "",
				Sparte:               0,
				Energierichtung:      0,
				Bilanzierungsmethode: 0,
				Netzebene:            0,
			},
		},
		"onlyOneAddressType": {
			bo.Marktlokation{
				Lokationsadresse: &com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördliche Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   landescode.DE,
				},
				Katasterinformation: &com.Katasteradresse{
					GemarkungFlur: "Foo",
					Flurstueck:    "Bar",
				},
			},
		},
		"maloid": {
			bo.Marktlokation{
				MarktlokationsId: "12345678918",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMarktlokationMap)
}

// Test_Successful_Marktlokation_Validation verifies that a valid Marktlokation is validated without errors
func (s *Suite) Test_Successful_Marktlokation_Validation() {
	validate := validator.New()
	registerError := validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(s.T(), registerError, is.Nil())
	validate.RegisterStructValidation(bo.XorStructLevelValidation, bo.Marktlokation{})
	validMalos := []bo.BusinessObject{
		// Minimal attributes
		bo.Marktlokation{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.MARKTLOKATION,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			MarktlokationsId:     "12345678913",
			Sparte:               sparte.STROM,
			Energierichtung:      energierichtung.EINSP,
			Bilanzierungsmethode: bilanzierungsmethode.SLP,
			Netzebene:            netzebene.NSP,
			Lokationsadresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördliche Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMalos)
}

func (s *Suite) Test_Get_MaloId_Checksum() {
	actual := bo.GetMaLoIdCheckSum("5123869678")
	then.AssertThat(s.T(), actual, is.EqualTo(1))
}

func (s *Suite) Test_Empty_Marktlokation_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.MARKTLOKATION)
	then.AssertThat(s.T(), object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Marktlokation{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.MARKTLOKATION))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Marktlokation_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.MARKTLOKATION))
}
