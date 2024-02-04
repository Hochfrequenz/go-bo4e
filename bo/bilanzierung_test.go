package bo_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/abwicklungsmodell"
	"github.com/hochfrequenz/go-bo4e/enum/aggregationsverantwortung"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/fallgruppenzuordnung"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/profiltyp"
	"github.com/hochfrequenz/go-bo4e/enum/profilverfahren"
	"github.com/hochfrequenz/go-bo4e/enum/prognosegrundlage"
	"github.com/hochfrequenz/go-bo4e/enum/zeitreihentyp"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/hochfrequenz/go-bo4e/internal/testcase"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
	"github.com/shopspring/decimal"
)

var aFalse = false
var aTrue = true
var seventeen = 17
var validBilanzierung = bo.Bilanzierung{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.BILANZIERUNG,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
		ExtensionData:     unmappeddatamarshaller.ExtensionData{},
	},
	Lastprofile: []com.Lastprofil{{
		Bezeichnung: "mein lieblingslastprofil",
		Verfahren:   profilverfahren.SYNTHETISCH,
		Tagesparameter: &com.Tagesparameter{
			Klimazone:            "ASD",
			Temperaturmessstelle: "Neustadt OST",
			Dienstanbieter:       "DWD",
		},
		Einspeisung: &aTrue,
	}},
	Bilanzierungsbeginn: internal.Ptr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	Bilanzierungsende:   internal.Ptr(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
	Bilanzkreis:         internal.Ptr("11XVE-N-GHM----Q"),
	Jahresverbrauchsprognose: &com.Menge{
		Wert:    newDecimalFromString("1500"),
		Einheit: internal.Ptr(mengeneinheit.KWH),
	},
	Kundenwert: &com.Menge{
		Wert:    newDecimalFromString("0.17"),
		Einheit: internal.Ptr(mengeneinheit.MWH),
	},
	Verbrauchsaufteilung:      decimal.NewNullDecimal(newDecimalFromString("0.17")),
	Zeitreihentyp:             internal.Ptr(zeitreihentyp.LGS),
	Aggregationsverantwortung: internal.Ptr(aggregationsverantwortung.VNB),
	Prognosegrundlage:         internal.Ptr(prognosegrundlage.WERTE),
	DetailsPrognosegrundlage: []profiltyp.Profiltyp{
		profiltyp.SLP_SEP,
	},
	WahlrechtPrognosegrundlage: &aFalse,
	Fallgruppenzuordnung:       internal.Ptr(fallgruppenzuordnung.GABI_RLMoT),
	Prioritaet:                 &seventeen,
	MarktlokationsId:           internal.Ptr("51238696781"),
	Abwicklungsmodell:          internal.Ptr(abwicklungsmodell.MODELL_1),
}

// Test_Bilanzierung_Deserialization deserializes an Bilanzierung json
func Test_Bilanzierung_Deserialization(t *testing.T) {
	serializedBilanzierung, err := json.Marshal(validBilanzierung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedBilanzierung, is.Not(is.NilArray[byte]()))
	jsonString := string(serializedBilanzierung)
	then.AssertThat(t, strings.Contains(jsonString, "SLP_SEP"), is.True()) // stringified enum
	var deserializedBilanzierung bo.Bilanzierung
	err = json.Unmarshal(serializedBilanzierung, &deserializedBilanzierung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedBilanzierung, is.EqualTo(validBilanzierung))
}

func TestBilanzierungDeserializeExplicitNulls(t *testing.T) {
	testcases := testcase.Map[testcase.JSONDeserializationSucceeds[bo.Bilanzierung]]{
		"marktlokationsId": `{
			"boTyp":"BILANZIERUNG",
			"versionStruktur":"1",
			"marktlokationsId": null
		}`,
		"zeitreihentyp": `{
			"boTyp":"BILANZIERUNG",
			"versionStruktur":"1",
			"zeitreihentyp": null
		}`,
		"aggregationsverantwortung": `{
			"boTyp":"BILANZIERUNG",
			"versionStruktur":"1",
			"aggregationsverantwortung": null
		}`,
		"prognosegrundlage": `{
			"boTyp":"BILANZIERUNG",
			"versionStruktur":"1",
			"prognosegrundlage":null
		}`,
		"fallgruppenzuordnung": `{
			"boTyp":"BILANZIERUNG",
			"versionStruktur":"1",
			"fallgruppenzuordnung":null
		}`,
	}

	testcases.Run(t)
}

func Test_Bilanzierung_Deserializes_Unknown_Fields(t *testing.T) {
	jsonString := `{"jahresverbrauchsprognose":{"wert":2345},"kundenwert":{"wert":38},"bilanzierungsbeginn":"2021-12-31T23:00:00+00:00","bilanzkreis":"THE0BFL002410004","prognosegrundlage":"PROFILE","detailsPrognosegrundlage":["SLP_SEP"],"boTyp":"BILANZIERUNG","versionStruktur":"1","Klassentyp":"Z12","Verfahren":"SYNTHETISCH","Profil":"D13","Lastprofil_Codeliste":"293","Temperaturmessstelle_Klassentyp":"Z99","Temperaturmessstelle_ID":"107290","Temperaturmessstelle_Anbieter":"ZT3","Temperaturmessstelle_Codeliste":"293"}`
	var deserializedBilanzierung bo.Bilanzierung
	err := json.Unmarshal([]byte(jsonString), &deserializedBilanzierung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, *deserializedBilanzierung.Bilanzkreis, is.EqualTo("THE0BFL002410004"))                  // a "normal" property/field of Bilanzierung
	then.AssertThat(t, deserializedBilanzierung.ExtensionData["Lastprofil_Codeliste"], is.EqualTo[any]("293")) // an extension data key
	jsonStringBytes, serializationErr := json.Marshal(deserializedBilanzierung)
	then.AssertThat(t, serializationErr, is.Nil())
	jsonString = string(jsonStringBytes)
	then.AssertThat(t, strings.Contains(jsonString, "\"Lastprofil_Codeliste\""), is.True())
}

// Test_Failed_Bilanzierung_Validation verifies that the validators of a Bilanzierung BO work
func Test_Failed_Bilanzierung_Validation(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("eic", bo.EICFieldLevelValidation)
	then.AssertThat(t, err, is.Nil())
	err = validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(t, err, is.Nil())
	invalidBilanzierungs := map[string][]interface{}{
		"required": {
			bo.Bilanzierung{},
		},
		"len": {
			bo.Bilanzierung{Bilanzkreis: internal.Ptr("not 16 chars")},
		},
		"eic": {
			bo.Bilanzierung{Bilanzkreis: internal.Ptr("FOO BAR XYZ ASD ")}, // 16 chars but does not match regex
			bo.Bilanzierung{Bilanzkreis: internal.Ptr("1234567890123456")}, // 16 chars but does not match regex
			bo.Bilanzierung{Bilanzkreis: internal.Ptr("11XVE-N-GHM----R")}, // wrong check sum
		},
		"gtefield": { // on ende
			bo.Bilanzierung{
				Bilanzierungsbeginn: internal.Ptr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				Bilanzierungsende:   internal.Ptr(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		"ltefield": { // on beginn
			bo.Bilanzierung{
				Bilanzierungsbeginn: internal.Ptr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				Bilanzierungsende:   internal.Ptr(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		"maloid": { // on beginn
			bo.Bilanzierung{
				MarktlokationsId: internal.Ptr("asdasd"), // not numeric, not len 11
			},
			bo.Bilanzierung{
				MarktlokationsId: internal.Ptr("12345678901"), // wrong checksum
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidBilanzierungs)
}

// Test_Successful_Bilanzierung_Validation verifies that a valid BO is validated without errors
func Test_Successful_Bilanzierung_Validation(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("eic", bo.EICFieldLevelValidation)
	then.AssertThat(t, err, is.Nil())
	err = validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(t, err, is.Nil())
	validMarktteilnehmers := []bo.BusinessObject{
		validBilanzierung,
		bilanzierungWithNilAbwicklungsmodell(validBilanzierung),
	}
	VerifySuccessfulValidations(t, validate, validMarktteilnehmers)
}

func Test_Empty_Bilanzierung_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.BILANZIERUNG)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Bilanzierung{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.BILANZIERUNG))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Bilanzierung_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.BILANZIERUNG))
}

// bilanzierungWithNilAbwicklungsmodell returns a copy of bilanzierung with Abwicklungsmodell set to nil.
func bilanzierungWithNilAbwicklungsmodell(bilanzierung bo.Bilanzierung) bo.Bilanzierung {
	bilanzierung.Abwicklungsmodell = nil
	return bilanzierung
}
