package bo_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/definitionennotwendigkeit"
	"github.com/hochfrequenz/go-bo4e/enum/ermittlungleistungsmaximum"
	"github.com/hochfrequenz/go-bo4e/enum/haeufigkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/schwachlastfaehigkeit"
	"github.com/hochfrequenz/go-bo4e/enum/uebermittelbarkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlzeitdefinitiontyp"
	"reflect"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/internal"
)

var zaehlzeitDefinition = bo.Zaehlzeitdefinition{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.ZAEHLZEITDEFINITION,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
	},
	Version:       time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
	Notwendigkeit: internal.Ptr(definitionennotwendigkeit.DEFINITIONEN_WERDEN_VERWENDET),

	Zaehlzeiten: []com.Zaehlzeit{{
		Code:                       internal.Ptr("HL1"),
		Haeufigkeit:                internal.Ptr(haeufigkeitzaehlzeit.JAEHRLICH),
		Uebermittelbarkeit:         internal.Ptr(uebermittelbarkeitzaehlzeit.ELEKTRONISCH),
		ErmittlungLeistungsmaximum: internal.Ptr(ermittlungleistungsmaximum.VERWENDUNG_HOCHLASTFENSTER),
		IstBestellbar:              internal.Ptr(true),
		Typ:                        internal.Ptr(zaehlzeitdefinitiontyp.HOCHLASTZEITFENSTER),
		BeschreibungTyp:            internal.Ptr("EineBeschreibung"),
	}},

	Zaehlzeitregister: []com.Zaehlzeitregister{{
		ZaehlzeitDefinition: internal.Ptr("HL1"),
		Register:            internal.Ptr("HT"),
		Schwachlastfaehig:   internal.Ptr(schwachlastfaehigkeit.NICHT_SCHWACHLASTFAEHIG),
	}, {
		ZaehlzeitDefinition: internal.Ptr("HL1"),
		Register:            internal.Ptr("NT"),
		Schwachlastfaehig:   internal.Ptr(schwachlastfaehigkeit.SCHWACHLASTFAEHIG),
	}},
}

// Test_Zaehlzeitdefinition_Deserialization deserializes an Zaehlzeitdefinition json
func Test_Zaehlzeitdefinition_Deserialization(t *testing.T) {
	serializedZaehlzeitdefinition, err := json.Marshal(zaehlzeitDefinition)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedZaehlzeitdefinition, is.Not(is.NilArray[byte]()))
	var deserializedZaehlzeitdefinition bo.Zaehlzeitdefinition
	err = json.Unmarshal(serializedZaehlzeitdefinition, &deserializedZaehlzeitdefinition)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZaehlzeitdefinition, is.EqualTo(zaehlzeitDefinition))
}

// Test_Failed_Zaehlzeitdefinition_Validation verifies that the validators of a Zaehlzeitdefinition work
func Test_Failed_Zaehlzeitdefinition_Validation(t *testing.T) {
	validate := validator.New()
	invalidZaehlzeitdefinition := map[string][]interface{}{
		"required": {
			bo.Zaehlzeitdefinition{},
		},
	}
	VerifyFailedValidations(t, validate, invalidZaehlzeitdefinition)
}

// Test_Successful_Zaehlzeitdefinition_Validation verifies that a valid BO is validated without errors
func Test_Successful_Zaehlzeitdefinition_Validation(t *testing.T) {
	validate := validator.New()
	validTranche := []bo.BusinessObject{
		zaehlzeitDefinition,
	}
	VerifySuccessfulValidations(t, validate, validTranche)
}

func Test_Empty_Zaehlzeitdefinition_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.ZAEHLZEITDEFINITION)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Zaehlzeitdefinition{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.ZAEHLZEITDEFINITION))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Zaehlzeitdefinition_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.ZAEHLZEITDEFINITION))
}
