package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
	"github.com/shopspring/decimal"
	"strings"
)

// Test_Messlokation_Deserialization tests serialization and deserialization of Messlokation
func (s *Suite) Test_Messlokation_Deserialization() {
	var melo = Messlokation{
		Geschaeftsobjekt: Geschaeftsobjekt{
			BoTyp:             botyp.Messlokation,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		MesslokationsId:              "DE0000011111222223333344444555556",
		Sparte:                       sparte.Strom,
		NetzebeneMessung:             netzebene.MD,
		MessgebietNr:                 "",
		Gerate:                       nil,
		Messdienstleistung:           nil,
		GrundzustaendigerMsbCodeNr:   "",
		GrundzustaendigerMsbImCodeNr: "",
		Messadresse: &com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Strasse:      "Nördliche Münchner Straße",
			Hausnummer:   "27A",
			Landescode:   landescode.DE,
		},
		Messlokationszaehler: []Zaehler{
			{
				Geschaeftsobjekt: Geschaeftsobjekt{
					BoTyp:             botyp.Zaehler,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Zaehlernummer:      "0815",
				Sparte:             sparte.Strom,
				Zaehlerauspraegung: zaehlerauspraegung.Einrichtungszaehler,
				Zaehlertyp:         zaehlertyp.Drehstromzaehler,
				Tarifart:           tarifart.Eintarif,
				Zaehlerkonstante:   decimal.NullDecimal{},
				Zaehlwerke: []com.Zaehlwerk{{
					ZaehlwerkId:   "1",
					Bezeichnung:   "",
					Richtung:      energierichtung.Aussp,
					ObisKennzahl:  "1-0:1.8.0",
					Wandlerfaktor: decimal.NewFromFloat(1),
					Einheit:       mengeneinheit.KWH,
				}},
			},
		},
	}
	serializedMelo, err := json.Marshal(melo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMelo, is.Not(is.Nil()))
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "MD"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "Strom"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "Eintarif"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "Einrichtungszaehler"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMelo), "Eintarif"), is.True())
	var deserializedMelo Messlokation
	err = json.Unmarshal(serializedMelo, &deserializedMelo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMelo, is.EqualTo(melo))
}

// TestFailedMesslokationValidation verifies that the validators of Messlokation work
func (s *Suite) Test_Failed_MesslokationValidation() {
	validate := validator.New()
	invalidMesslokationMap := map[string][]interface{}{
		"required": {
			Messlokation{
				Geschaeftsobjekt: Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				MesslokationsId:              "",
				Sparte:                       0,
				NetzebeneMessung:             0,
				MessgebietNr:                 "",
				Gerate:                       nil,
				Messdienstleistung:           nil,
				GrundzustaendigerMsbCodeNr:   "",
				GrundzustaendigerMsbImCodeNr: "",
				Messadresse:                  nil,
			},
		},
		"len": {
			Messlokation{
				MesslokationsId: "not33",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMesslokationMap)
}

//  TestSuccessfulMesslokationValidation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_MesslokationValidation() {
	validate := validator.New()
	validMelos := []interface{}{
		Messlokation{
			Geschaeftsobjekt: Geschaeftsobjekt{
				BoTyp:             botyp.Messlokation,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			MesslokationsId:              "DE0000011111222223333344444555556",
			Sparte:                       sparte.Strom,
			NetzebeneMessung:             netzebene.MD,
			MessgebietNr:                 "",
			Gerate:                       nil,
			Messdienstleistung:           nil,
			GrundzustaendigerMsbCodeNr:   "",
			GrundzustaendigerMsbImCodeNr: "",
			Messadresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördliche Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMelos)
}
