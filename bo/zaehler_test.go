package bo_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

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
)

// Test_Zaehler_Deserialization deserializes an Zaehler json
func Test_Zaehler_Deserialization(t *testing.T) {
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
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedMeter, is.Not(is.NilArray[byte]()))
	then.AssertThat(t, strings.Contains(string(serializedMeter), "EINTARIF"), is.True())
	then.AssertThat(t, strings.Contains(string(serializedMeter), "EINRICHTUNGSZAEHLER"), is.True())
	then.AssertThat(t, strings.Contains(string(serializedMeter), "EINTARIF"), is.True())
	var deserializedMeter bo.Zaehler
	err = json.Unmarshal(serializedMeter, &deserializedMeter)
	then.AssertThat(t, err, is.Nil())

	areEqual, err := internal.CompareAsJson(meter, deserializedMeter)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, areEqual, is.True())
}

// TestFailedZaehlerValidation verifies that the validators of a Zaehler work
func Test_Failed_ZaehlerValidation(t *testing.T) {
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
	VerifyFailedValidations(t, validate, invalidVertrags)
}

// Test_Successful_Zaehler_Validation verifies that a valid BO is validated without errors
func Test_Successful_Zaehler_Validation(t *testing.T) {
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
					Landescode:   internal.Ptr(landescode.DE),
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.DIENSTLEISTER,
				},
			},
		},
	}
	VerifySuccessfulValidations(t, validate, validZaehler)
}

func Test_Empty_Zaehler_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.ZAEHLER)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Zaehler{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.ZAEHLER))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Zaehler_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.ZAEHLER))
}
