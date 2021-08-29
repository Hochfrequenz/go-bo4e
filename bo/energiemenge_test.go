package bo

import (
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"time"
)

// TestFailedEnergiemengeValidation verifies that the validators of Energiemenge work
func (s *Suite) TestFailedEnergiemengeValidation() {
	var verbrauch = com.Verbrauch{
		Startdatum:               time.Now(),
		Enddatum:                 time.Now().Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     17,
		Einheit:                  mengeneinheit.KWH,
	}
	validate := validator.New()
	invalidEnergiemengeMap := map[string][]interface{}{
		"required": {
			Energiemenge{
				BusinessObject: BusinessObject{
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
			Energiemenge{
				BusinessObject: BusinessObject{
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
			Energiemenge{
				BusinessObject: BusinessObject{
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
func (s *Suite) TestSuccessfulEnergiemengeValidation() {
	validate := validator.New()
	var verbrauch = com.Verbrauch{
		Startdatum:               time.Now(),
		Enddatum:                 time.Now().Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     17,
		Einheit:                  mengeneinheit.KWH,
	}
	validEnergiemengen := []interface{}{
		Energiemenge{
			BusinessObject: BusinessObject{
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
