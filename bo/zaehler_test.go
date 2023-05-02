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
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
)

// Test_Zaehler_Deserialization deserializes an Zaehler json
func (s *Suite) Test_Zaehler_Deserialization() {
	var meter = bo.Zaehler{
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
		//EichungBis:         time.Time{},
		//LetzteEichung:      time.Time{},
		Zaehlwerke: []com.Zaehlwerk{{
			ZaehlwerkId:    "1",
			Bezeichnung:    "",
			Richtung:       energierichtung.AUSSP,
			ObisKennzahl:   "1-0:1.8.0",
			Wandlerfaktor:  newDecimalFromString("1"),
			Einheit:        mengeneinheit.KWH,
			Zaehlerstaende: nil,
		}},
		Zaehlerhersteller: nil,
	}
	serializedMeter, err := json.Marshal(meter)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMeter, is.Not(is.NilArray[byte]()))
	then.AssertThat(s.T(), strings.Contains(string(serializedMeter), "EINTARIF"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMeter), "EINRICHTUNGSZAEHLER"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMeter), "EINTARIF"), is.True())
	var deserializedMeter bo.Zaehler
	err = json.Unmarshal(serializedMeter, &deserializedMeter)
	then.AssertThat(s.T(), err, is.Nil())

	areEqual, err := internal.CompareAsJson(meter, deserializedMeter)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), areEqual, is.True())
}

// TestFailedZaehlerValidation verifies that the validators of a Zaehler work
func (s *Suite) Test_Failed_ZaehlerValidation() {
	validate := validator.New()
	invalidVertrags := map[string][]interface{}{
		"required": {
			bo.Zaehler{},
		},
		"min": { // min 1 zaehlwerk is required
			bo.Zaehler{
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
				Zaehlerkonstante:   decimal.NewNullDecimal(decimal.NewFromFloat(17)),
				Zaehlwerke:         []com.Zaehlwerk{},
				Zaehlerhersteller:  nil,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertrags)
}

// Test_Successful_Zaehler_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Zaehler_Validation() {
	validate := validator.New()
	validZaehler := []bo.BusinessObject{
		bo.Zaehler{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.ZAEHLER,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Zaehlernummer:      "08150",
			Sparte:             sparte.STROM,
			Zaehlerauspraegung: internal.Ptr(zaehlerauspraegung.EINRICHTUNGSZAEHLER),
			Zaehlertyp:         zaehlertyp.WECHSELSTROMZAEHLER,
			Tarifart:           internal.Ptr(tarifart.EINTARIF),
			Zaehlerkonstante:   decimal.NullDecimal{Valid: false, Decimal: decimal.NewFromFloat(0)},
			//EichungBis:         time.Time{},
			//LetzteEichung:      time.Time{},
			Zaehlwerke: []com.Zaehlwerk{{
				ZaehlwerkId:    "1",
				Bezeichnung:    "",
				Richtung:       energierichtung.AUSSP,
				ObisKennzahl:   "1-0:1.8.0",
				Wandlerfaktor:  decimal.NewFromFloat(0),
				Einheit:        mengeneinheit.KWH,
				Zaehlerstaende: nil,
			}},
			Zaehlerhersteller: &bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.GESCHAEFTSPARTNER,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Name1:                "Musterfrau",
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
		},
	}
	VerfiySuccessfulValidations(s, validate, validZaehler)
}

func (s *Suite) Test_Empty_Zaehler_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.ZAEHLER)
	then.AssertThat(s.T(), object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Zaehler{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.ZAEHLER))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Zaehler_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.ZAEHLER))
}
