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
		},
		MesslokationsId:              "DE0000011111222223333344444555556",
		Sparte:                       sparte.STROM,
		NetzebeneMessung:             netzebene.MD,
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
				Zaehlerauspraegung: zaehlerauspraegung.EINRICHTUNGSZAEHLER,
				Zaehlertyp:         zaehlertyp.DREHSTROMZAEHLER,
				Tarifart:           tarifart.EINTARIF,
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
	then.AssertThat(s.T(), serializedMelo, is.Not(is.Nil()))
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "MD"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "STROM"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "EINTARIF"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "EINRICHTUNGSZAEHLER"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "EINTARIF"), is.True())
	var deserializedMelo bo.Messlokation
	err = json.Unmarshal(serializedMelo, &deserializedMelo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMelo, is.EqualTo(melo))
}

// Test_Failed_MesslokationValidation verifies that the validators of Messlokation work
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
				NetzebeneMessung:             0,
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
				NetzebeneMessung:             0,
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

//  Test_Successful_Messlokation_Validation verifies that a valid BO is validated without errors
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
			NetzebeneMessung:             netzebene.MD,
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
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Messlokation{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.MESSLOKATION))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}
