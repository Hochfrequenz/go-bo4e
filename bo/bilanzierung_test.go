package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	aggregationsverantwortung "github.com/hochfrequenz/go-bo4e/enum/aggregationsverwantwortung"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/fallgruppenzuordnung"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/profiltyp"
	"github.com/hochfrequenz/go-bo4e/enum/profilverfahren"
	"github.com/hochfrequenz/go-bo4e/enum/prognosegrundlage"
	"github.com/hochfrequenz/go-bo4e/enum/zeitreihentyp"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"time"
)

var aFalse = false
var aTrue = true
var seventeen = 17
var validBilanzierung = bo.Bilanzierung{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.BILANZIERUNG,
		VersionStruktur: "1.1",
		ExterneReferenzen: nil,
		ExtensionData:     map[string]interface{}{},
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
	Bilanzierungsbeginn: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	Bilanzierungsende:   time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
	Bilanzkreis:         "11XVE-N-GHM----Q",
	Jahresverbrauchsprognose: &com.Menge{
		Wert:    newDecimalFromString("1500"),
		Einheit: mengeneinheit.KWH,
	},
	Kundenwert: &com.Menge{
		Wert:    newDecimalFromString("0.17"),
		Einheit: mengeneinheit.MWH,
	},
	Verbrauchsaufteilung:      decimal.NewNullDecimal(newDecimalFromString("0.17")),
	Zeitreihentyp:             zeitreihentyp.LGS,
	Aggregationsverantwortung: aggregationsverantwortung.VNB,
	Prognosegrundlage:         prognosegrundlage.WERTE,
	DetailsPrognosegrundlage: []profiltyp.Profiltyp{
		profiltyp.SLP_SEP,
	},
	WahlrechtPrognosegrundlage: &aFalse,
	Fallgruppenzuordnung:       fallgruppenzuordnung.GABI_RLMoT,
	Prioritaet:                 &seventeen,
	MarktlokationsId:           "51238696781",
}

// Test_Bilanzierung_Deserialization deserializes an Bilanzierung json
func (s *Suite) Test_Bilanzierung_Deserialization() {
	serializedBilanzierung, err := json.Marshal(validBilanzierung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedBilanzierung, is.Not(is.Nil()))
	jsonString := string(serializedBilanzierung)
	then.AssertThat(s.T(), strings.Contains(jsonString, "SLP_SEP"), is.True()) // stringified enum
	var deserializedBilanzierung bo.Bilanzierung
	err = json.Unmarshal(serializedBilanzierung, &deserializedBilanzierung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedBilanzierung, is.EqualTo(validBilanzierung))
}

func (s *Suite) Test_Bilanzierung_Deserializes_Unknown_Fields(){
	jsonString := `{"jahresverbrauchsprognose":{"wert":2345},"kundenwert":{"wert":38},"bilanzierungsbeginn":"2021-12-31T23:00:00+00:00","bilanzkreis":"THE0BFL002410004","prognosegrundlage":"PROFILE","detailsPrognosegrundlage":["SLP_SEP"],"boTyp":"BILANZIERUNG","versionStruktur":"1","Klassentyp":"Z12","Verfahren":"SYNTHETISCH","Profil":"D13","Lastprofil_Codeliste":"293","Temperaturmessstelle_Klassentyp":"Z99","Temperaturmessstelle_ID":"107290","Temperaturmessstelle_Anbieter":"ZT3","Temperaturmessstelle_Codeliste":"293"}`
	var deserializedBilanzierung bo.Bilanzierung
	err := json.Unmarshal([]byte(jsonString), &deserializedBilanzierung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedBilanzierung.Bilanzkreis, is.EqualTo("THE0BFL002410004")) // a "normal" property/field of Bilanzierung
	then.AssertThat(s.T(), deserializedBilanzierung.ExtensionData["Lastprofil_Codeliste"], is.EqualTo("293")) // an extension data key
	jsonStringBytes, serializationErr := json.Marshal(deserializedBilanzierung)
	then.AssertThat(s.T(), serializationErr, is.Nil())
	jsonString = string(jsonStringBytes)
	then.AssertThat(s.T(), strings.Contains(jsonString, "\"Lastprofil_Codeliste\""), is.True())
}

// Test_Failed_Bilanzierung_Validation verifies that the validators of a Bilanzierung BO work
func (s *Suite) Test_Failed_Bilanzierung_Validation() {
	validate := validator.New()
	err := validate.RegisterValidation("eic", bo.EICFieldLevelValidation)
	then.AssertThat(s.T(), err, is.Nil())
	err = validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(s.T(), err, is.Nil())
	invalidBilanzierungs := map[string][]interface{}{
		"required": {
			bo.Bilanzierung{},
		},
		"len": {
			bo.Bilanzierung{Bilanzkreis: "not 16 chars"},
		},
		"eic": {
			bo.Bilanzierung{Bilanzkreis: "FOO BAR XYZ ASD "}, // 16 chars but does not match regex
			bo.Bilanzierung{Bilanzkreis: "1234567890123456"}, // 16 chars but does not match regex
			bo.Bilanzierung{Bilanzkreis: "11XVE-N-GHM----R"}, // wrong check sum
		},
		"gtefield": { // on ende
			bo.Bilanzierung{
				Bilanzierungsbeginn: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Bilanzierungsende:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		"ltefield": { // on beginn
			bo.Bilanzierung{
				Bilanzierungsbeginn: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Bilanzierungsende:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		"maloid": { // on beginn
			bo.Bilanzierung{
				MarktlokationsId: "asdasd", // not numeric, not len 11
			},
			bo.Bilanzierung{
				MarktlokationsId: "12345678901", // wrong checksum
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidBilanzierungs)
}

//  Test_Successful_Bilanzierung_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Bilanzierung_Validation() {
	validate := validator.New()
	err := validate.RegisterValidation("eic", bo.EICFieldLevelValidation)
	then.AssertThat(s.T(), err, is.Nil())
	err = validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(s.T(), err, is.Nil())
	validMarktteilnehmers := []bo.BusinessObject{
		validBilanzierung,
	}
	VerfiySuccessfulValidations(s, validate, validMarktteilnehmers)
}

func (s *Suite) Test_Empty_Bilanzierung_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.BILANZIERUNG)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Bilanzierung{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.BILANZIERUNG))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Bilanzierung_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.BILANZIERUNG))
}
