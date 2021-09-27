package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
)

// Test_Zaehler_Deserialization deserializes an Zaehler json
func (s *Suite) Test_Zaehler_Deserialization() {
	var meter = Zaehler{
		BusinessObject: BusinessObject{
			BoTyp:             botyp.Zaehler,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		Zaehlernummer:      "0815",
		Sparte:             sparte.Strom,
		Zaehlerauspraegung: zaehlerauspraegung.Einrichtungszaehler,
		Zaehlertyp:         zaehlertyp.Drehstromzaehler,
		Tarifart:           tarifart.Eintarif,
		Zaehlerkonstante:   0,
		//EichungBis:         time.Time{},
		//LetzteEichung:      time.Time{},
		Zaehlwerke: []com.Zaehlwerk{{
			ZaehlwerkId:    "1",
			Bezeichnung:    "",
			Richtung:       energierichtung.Aussp,
			ObisKennzahl:   "1-0:1.8.0",
			Wandlerfaktor:  0,
			Einheit:        mengeneinheit.KWH,
			Zaehlerstaende: nil,
		}},
		Zaehlerhersteller: nil,
	}
	serializedMeter, err := json.Marshal(meter)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMeter, is.Not(is.Nil()))
	var deserializedMeter Zaehler
	err = json.Unmarshal(serializedMeter, &deserializedMeter)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMeter, is.EqualTo(meter))
}

// TestFailedZaehlerValidation verifies that the validators of a Zaehler work
func (s *Suite) TestFailedZaehlerValidation() {
	validate := validator.New()
	invalidVertrags := map[string][]interface{}{
		"required": {
			Zaehler{},
		},
		"min": { // min 1 zaehlwerk is required
			Zaehler{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Zaehler,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Zaehlernummer:      "0815",
				Sparte:             sparte.Strom,
				Zaehlerauspraegung: zaehlerauspraegung.Einrichtungszaehler,
				Zaehlertyp:         zaehlertyp.Drehstromzaehler,
				Tarifart:           tarifart.Eintarif,
				Zaehlerkonstante:   17,
				Zaehlwerke: []com.Zaehlwerk{
				},
				Zaehlerhersteller: nil,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertrags)
}

//  Test_Successful_Vertrag_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Zaehler_Validation() {

	validate := validator.New()
	validZaehler := []interface{}{
		Zaehler{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Zaehler,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Zaehlernummer:      "08150",
			Sparte:             sparte.Strom,
			Zaehlerauspraegung: zaehlerauspraegung.Einrichtungszaehler,
			Zaehlertyp:         zaehlertyp.Wechselstromzaehler,
			Tarifart:           tarifart.Eintarif,
			Zaehlerkonstante:   0,
			//EichungBis:         time.Time{},
			//LetzteEichung:      time.Time{},
			Zaehlwerke: []com.Zaehlwerk{{
				ZaehlwerkId:    "1",
				Bezeichnung:    "",
				Richtung:       energierichtung.Aussp,
				ObisKennzahl:   "1-0:1.8.0",
				Wandlerfaktor:  0,
				Einheit:        mengeneinheit.KWH,
				Zaehlerstaende: nil,
			}},
			//Zaehlerhersteller:  Geschaeftspartner{},
		},
	}
	VerfiySuccessfulValidations(s, validate, validZaehler)
}
