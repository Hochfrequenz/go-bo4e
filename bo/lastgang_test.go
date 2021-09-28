package bo

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatus"
	"github.com/hochfrequenz/go-bo4e/enum/messwertstatuszusatz"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
	"time"
)

// TestFailedLastgangValidation verifies that the validators of a Lastgang work
func (s *Suite) TestFailedLastgangValidation() {
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
			Lastgang{
				BusinessObject: BusinessObject{
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
			Lastgang{
				BusinessObject: BusinessObject{
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
			Lastgang{
				BusinessObject: BusinessObject{
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
