package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
)

// Test_Messlokation_Deserialization tests serialization and deserialization of Messlokation
func (s *Suite) Test_Messlokation_Deserialization() {
	var melo = bo.Messlokation{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.MESSLOKATION,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
			ExtensionData:     unmappeddatamarshaller.ExtensionData{},
		},
		MesslokationsId:              "DE0000011111222223333344444555556",
		Sparte:                       sparte.STROM,
		NetzebeneMessung:             internal.Ptr(netzebene.MD),
		MessgebietNr:                 "",
		Geraete:                      nil,
		Messdienstleistung:           nil,
		GrundzustaendigerMsbCodeNr:   "",
		GrundzustaendigerMsbImCodeNr: "",
		Messadresse: &com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Strasse:      "Nördliche Münchner Straße",
			Hausnummer:   "27A",
			Landescode:   landescode.DE,
		},
		Messlokationszaehler: []bo.Zaehler{
			{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.ZAEHLER,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Zaehlernummer:      "0815",
				Sparte:             sparte.STROM,
				Zaehlerauspraegung: internal.Ptr(zaehlerauspraegung.EINRICHTUNGSZAEHLER),
				Zaehlertyp:         zaehlertyp.DREHSTROMZAEHLER,
				Tarifart:           internal.Ptr(tarifart.EINTARIF),
				Zaehlerkonstante:   decimal.NullDecimal{},
				Zaehlwerke: []com.Zaehlwerk{{
					ZaehlwerkId:   "1",
					Bezeichnung:   "",
					Richtung:      energierichtung.AUSSP,
					ObisKennzahl:  "1-0:1.8.0",
					Wandlerfaktor: decimal.NewFromFloat(1),
					Einheit:       mengeneinheit.KWH,
				}},
			},
		},
	}
	serializedMelo, err := json.Marshal(melo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMelo, is.Not(is.NilArray[byte]()))
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "MD"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "STROM"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "EINTARIF"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "EINRICHTUNGSZAEHLER"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "EINTARIF"), is.True())
	var deserializedMelo bo.Messlokation
	err = json.Unmarshal(serializedMelo, &deserializedMelo)
	then.AssertThat(s.T(), err, is.Nil())

	areEqual, err := internal.CompareAsJson(melo, deserializedMelo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), areEqual, is.True())
}

// Test_Failed_MesslokationValidation verify that the validators of Messlokation work
func (s *Suite) Test_Failed_MesslokationValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(bo.XorStructLevelMesslokationValidation, bo.Messlokation{})
	invalidMesslokationMap := map[string][]interface{}{
		"required": {
			bo.Messlokation{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				MesslokationsId:              "",
				Sparte:                       0,
				NetzebeneMessung:             nil,
				MessgebietNr:                 "",
				Geraete:                      nil,
				Messdienstleistung:           nil,
				GrundzustaendigerMsbCodeNr:   "",
				GrundzustaendigerMsbImCodeNr: "",
				Messadresse:                  nil,
				Geoadresse:                   nil,
				Katasterinformation:          nil,
			},
		},
		"required_without_all": {
			bo.Messlokation{},
		},
		"onlyOneAddressType": {
			bo.Messlokation{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				MesslokationsId:              "",
				Sparte:                       0,
				NetzebeneMessung:             nil,
				MessgebietNr:                 "",
				Geraete:                      nil,
				Messdienstleistung:           nil,
				GrundzustaendigerMsbCodeNr:   "",
				GrundzustaendigerMsbImCodeNr: "",
				Messadresse: &com.Adresse{
					Postleitzahl: "44056",
					Ort:          "Neustadt",
					Strasse:      "Neue Straße",
					Hausnummer:   "17",
					Landescode:   landescode.DE,
				},
				Geoadresse: &com.Geokoordinaten{
					Breitengrad: newDecimalFromString("12.34"),
					Laengengrad: newDecimalFromString("45.67"),
				},
				Katasterinformation: nil,
			},
		},
		"len": {
			bo.Messlokation{
				MesslokationsId: "not33",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMesslokationMap)
}

//nolint:dupl // This can only be simplified if we use generics. anything else seems overly complicated but maybe it's just me
func (s *Suite) Test_Messlokation_DeSerialization_With_Unkonwn_Fields() {
	// this is a json that contains fields/keys that are not part of the official bo4e standard.
	// our expectation is that they are deserialized into the "ExtensionData" field.
	// They should also be serialized from there / not be lost during a marshalling/unmarshalling round trip
	meloJsonWithUnknownFields := `{
      "messlokationsId": "DE00026930926E0000000000100763575",
      "sparte": "STROM",
      "netzebeneMessung": "NSP",
      "abrechnungmessstellenbetriebnna": true,
      "marktrollen": [
        {
          "code": "9906214000003",
          "marktrolle": "MSB"
        }
      ],
      "boTyp": "MESSLOKATION",
      "versionStruktur": "1"
    }`
	melo := bo.Messlokation{}

	// unmarshalling tests
	err := json.Unmarshal([]byte(meloJsonWithUnknownFields), &melo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), melo.Geschaeftsobjekt.ExtensionData, is.Not(is.EqualTo[unmappeddatamarshaller.ExtensionData](nil)))
	then.AssertThat(s.T(), melo.ExtensionData["marktrollen"], is.Not(is.EqualTo[any](nil)))       // messlokation->marktrollen is NOT part of the bo4e standard ==> present in extension data
	then.AssertThat(s.T(), melo.ExtensionData["messlokationsId"], is.EqualTo[any](nil))           // messlokation->messlokationsId is part of the bo4e standard ==> not present in extension data
	then.AssertThat(s.T(), melo.MesslokationsId, is.EqualTo("DE00026930926E0000000000100763575")) // but where it should be
	// the other fields should be fine, too, without explicit tests; Add them if you feel like it doesn't work

	// marshaling tests
	serializedMeLoBytes, errSerializing := json.Marshal(melo)
	then.AssertThat(s.T(), errSerializing, is.Nil())
	serializedMeLo := string(serializedMeLoBytes)
	then.AssertThat(s.T(), strings.Contains(serializedMeLo, "marktrollen"), is.True())     // unmapped fields should be part of the serialized melo
	then.AssertThat(s.T(), strings.Contains(serializedMeLo, "messlokationsId"), is.True()) // mapped fields should be part of the serialized melo
}

// Test_Successful_Messlokation_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Messlokation_Validation() {
	validate := validator.New()
	validMelos := []bo.BusinessObject{
		bo.Messlokation{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.MESSLOKATION,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			MesslokationsId:              "DE0000011111222223333344444555556",
			Sparte:                       sparte.STROM,
			NetzebeneMessung:             internal.Ptr(netzebene.MD),
			MessgebietNr:                 "",
			Geraete:                      nil,
			Messdienstleistung:           nil,
			GrundzustaendigerMsbCodeNr:   "",
			GrundzustaendigerMsbImCodeNr: "",
			Messadresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördliche Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMelos)
}

func (s *Suite) Test_Empty_Messlokation_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.MESSLOKATION)
	then.AssertThat(s.T(), object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Messlokation{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.MESSLOKATION))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Messlokation_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.MESSLOKATION))
}
