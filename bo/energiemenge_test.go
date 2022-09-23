package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"time"
)

// TestFailedEnergiemengeValidation verifies that the validators of Energiemenge work
func (s *Suite) Test_Failed_EnergiemengeValidation() {
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
				LokationsTyp: 0,
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
				LokationsTyp: lokationstyp.MELO,
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
				LokationsTyp: lokationstyp.MELO,
				Verbrauch:    []com.Verbrauch{verbrauch},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidEnergiemengeMap)
}

// TestSuccessfulMesslokationValidation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_EnergiemengeValidation() {
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
			LokationsTyp: lokationstyp.MELO,
			Verbrauch:    []com.Verbrauch{verbrauch},
		},
	}
	VerfiySuccessfulValidations(s, validate, validEnergiemengen)
}

func (s *Suite) Test_Empty_Energiemenge_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.ENERGIEMENGE)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Energiemenge{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.ENERGIEMENGE))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Empty_Something_Is_Creatable_Using_BoTyp() {
	// remove this test as soon as the TARIFPREISBLATT is implemented. just to cover the nil case
	object := bo.NewBusinessObject(botyp.TARIFPREISBLATT)
	then.AssertThat(s.T(), object, is.Nil())
}

func (s *Suite) Test_Serialized_Empty_Energiemenge_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.ENERGIEMENGE))
}

func (s *Suite) assert_Does_Not_Serialize_Default_Enums(bo bo.BusinessObject) {
	jsonBytes, err := json.Marshal(bo)
	then.AssertThat(s.T(), err, is.Nil())
	jsonString := string(jsonBytes)
	then.AssertThat(s.T(), strings.Contains(jsonString, "(0)"), is.False())
}

func (s *Suite) Test_Energiemenge_Implements_GetBoTyp() {
	em := bo.Energiemenge{}
	then.AssertThat(s.T(), em.GetBoTyp(), is.EqualTo(botyp.ENERGIEMENGE))
}
