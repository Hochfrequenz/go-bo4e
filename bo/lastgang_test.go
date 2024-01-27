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
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
)

// TestFailedLastgangValidation verifies that the validators of a Lastgang work
func Test_Failed_LastgangValidation(t *testing.T) {
	var zeitreihenwert = com.Zeitreihenwert{
		Zeitreihenwertkompakt: com.Zeitreihenwertkompakt{
			Wert:         decimal.NewFromFloat(17.43),
			Status:       messwertstatus.ENERGIEMENGESUMMIERT,
			Statuszusatz: messwertstatuszusatz.Z81_MESSEINRICHTUNGGESTOERT_DEFEKT,
		},
		DatumUhrzeitVon: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		DatumUhrzeitBis: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
	}
	validate := validator.New()
	invalidLastgangMap := map[string][]interface{}{
		"required": {
			bo.Lastgang{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Sparte:       0,
				Version:      "",
				LokationsId:  "",
				LokationsTyp: 0,
				Messgroesse:  0,
				Obiskennzahl: "",
				Werte:        nil,
			},
		},
		"min": {
			bo.Lastgang{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.LASTGANG,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Sparte:       sparte.STROM,
				Version:      "1",
				LokationsId:  "asd",
				LokationsTyp: lokationstyp.MALO,
				Messgroesse:  mengeneinheit.KW,
				Obiskennzahl: "1-1:0.8.1",
				Werte:        nil,
			},
		},
		"alphanum": {
			bo.Lastgang{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.LASTGANG,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Sparte:       sparte.STROM,
				Version:      "1",
				LokationsId:  "not alpha num",
				LokationsTyp: lokationstyp.MALO,
				Messgroesse:  mengeneinheit.KW,
				Obiskennzahl: "1-1:0.8.1",
				Werte:        []com.Zeitreihenwert{zeitreihenwert},
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidLastgangMap)
}

// Test_Successful_Lastgang_Validation verifies that a valid BO is validated without errors
func Test_Successful_Lastgang_Validation(t *testing.T) {
	validate := validator.New()
	var verbrauch = com.Verbrauch{
		Startdatum:               time.Now(),
		Enddatum:                 time.Now().Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     decimal.NewFromFloat(17),
		Einheit:                  mengeneinheit.KWH,
	}
	validLastgang := []bo.BusinessObject{
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
	VerifySuccessfulValidations(t, validate, validLastgang)
}

func Test_Empty_Lastgang_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.LASTGANG)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Lastgang{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.LASTGANG))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Lastgang_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.LASTGANG))
}
