package bo_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/hochfrequenz/go-bo4e/internal/testcase"
	"github.com/shopspring/decimal"
)

// TestFailedEnergiemengeValidation verifies that the validators of Energiemenge work
func Test_Failed_EnergiemengeValidation(t *testing.T) {
	var verbrauch = com.Verbrauch{
		Startdatum:               time.Now(),
		Enddatum:                 time.Now().Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     decimal.NewFromFloat(17),
		Einheit:                  mengeneinheit.KWH,
	}
	validate := validator.New()
	invalidEnergiemengeMap := map[string][]interface{}{
		"required": {
			bo.Energiemenge{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				LokationsId:  "",
				LokationsTyp: internal.Ptr(lokationstyp.Lokationstyp(0)),
				Verbrauch:    nil,
			},
		},
		"min": {
			bo.Energiemenge{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.ENERGIEMENGE,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				LokationsId:  "not11len",
				LokationsTyp: internal.Ptr(lokationstyp.MELO),
				Verbrauch:    []com.Verbrauch{verbrauch},
			},
		},
		"alphanum": {
			bo.Energiemenge{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.ENERGIEMENGE,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				LokationsId:  "not alpha numeric",
				LokationsTyp: internal.Ptr(lokationstyp.MELO),
				Verbrauch:    []com.Verbrauch{verbrauch},
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidEnergiemengeMap)
}

// TestSuccessfulMesslokationValidation verifies that a valid BO is validated without errors
func Test_Successful_EnergiemengeValidation(t *testing.T) {
	validate := validator.New()
	var verbrauch = com.Verbrauch{
		Startdatum:               time.Now(),
		Enddatum:                 time.Now().Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     decimal.NewFromFloat(17),
		Einheit:                  mengeneinheit.KWH,
	}
	validEnergiemengen := []bo.BusinessObject{
		bo.Energiemenge{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.ENERGIEMENGE,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			LokationsId:  "54321098765",
			LokationsTyp: internal.Ptr(lokationstyp.MELO),
			Verbrauch:    []com.Verbrauch{verbrauch},
		},
	}
	VerifySuccessfulValidations(t, validate, validEnergiemengen)
}

func TestEnergiemengeDeserializeExplicitNulls(t *testing.T) {
	testcases := testcase.Map[testcase.JSONDeserializationSucceeds[bo.Energiemenge]]{
		"lokationstyp": `{
			"boTyp": "ENERGIEMENGE",
			"versionStruktur":"1",
			"lokationsTyp": null
		}`,
	}

	testcases.Run(t)
}

func Test_Empty_Energiemenge_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.ENERGIEMENGE)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Energiemenge{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.ENERGIEMENGE))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Empty_Something_Is_Creatable_Using_BoTyp(t *testing.T) {
	// remove this test as soon as the TARIFPREISBLATT is implemented. just to cover the nil case
	object := bo.NewBusinessObject(botyp.TARIFPREISBLATT)
	then.AssertThat(t, object, is.EqualTo[bo.BusinessObject](nil))
}

func Test_Serialized_Empty_Energiemenge_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.ENERGIEMENGE))
}
