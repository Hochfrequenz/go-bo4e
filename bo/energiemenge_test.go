package bo_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
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
				BusinessObject: bo.BusinessObject{
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
				BusinessObject: bo.BusinessObject{
					BoTyp:             botyp.Energiemenge,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				LokationsId:  "not11len",
				LokationsTyp: lokationstyp.MeLo,
				Verbrauch:    []com.Verbrauch{verbrauch},
			},
		},
		"alphanum": {
			bo.Energiemenge{
				BusinessObject: bo.BusinessObject{
					BoTyp:             botyp.Energiemenge,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				LokationsId:  "not alpha numeric",
				LokationsTyp: lokationstyp.MeLo,
				Verbrauch:    []com.Verbrauch{verbrauch},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidEnergiemengeMap)
}

//  TestSuccessfulMesslokationValidation verifies that a valid BO is validated without errors
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
	validEnergiemengen := []interface{}{
		bo.Energiemenge{
			BusinessObject: bo.BusinessObject{
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
