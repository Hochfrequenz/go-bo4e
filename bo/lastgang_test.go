package bo_test

import (
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
	"reflect"
	"time"
)

// TestFailedLastgangValidation verifies that the validators of a Lastgang work
func (s *Suite) Test_Failed_LastgangValidation() {
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
					BoTyp:             botyp.Lastgang,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Sparte:       sparte.Strom,
				Version:      "1",
				LokationsId:  "asd",
				LokationsTyp: lokationstyp.MaLo,
				Messgroesse:  mengeneinheit.KW,
				Obiskennzahl: "1-1:0.8.1",
				Werte:        nil,
			},
		},
		"alphanum": {
			bo.Lastgang{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.Lastgang,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Sparte:       sparte.Strom,
				Version:      "1",
				LokationsId:  "not alpha num",
				LokationsTyp: lokationstyp.MaLo,
				Messgroesse:  mengeneinheit.KW,
				Obiskennzahl: "1-1:0.8.1",
				Werte:        []com.Zeitreihenwert{zeitreihenwert},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidLastgangMap)
}

//  Test_Successful_Lastgang_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Lastgang_Validation() {
	validate := validator.New()
	var verbrauch = com.Verbrauch{
		Startdatum:               time.Now(),
		Enddatum:                 time.Now().Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     decimal.NewFromFloat(17),
		Einheit:                  mengeneinheit.KWH,
	}
	validEnergiemengen := []interface{}{
		bo.Energiemenge{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.Energiemenge,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			LokationsId:  "54321098765",
			LokationsTyp: lokationstyp.MeLo,
			Verbrauch:    []com.Verbrauch{verbrauch},
		},
	}
	VerfiySuccessfulValidations(s, validate, validEnergiemengen)
}

func (s *Suite) Test_Empty_Lastgang_Is_Creatable_Using_BoTyp() {
	object := bo.GetNewBusinessObject(botyp.Lastgang)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Lastgang{})))
}
