package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
	"strings"
)

// Test_Marktteilnehmer_Deserialization deserializes an Marktteilnehmer json
func (s *Suite) Test_Marktteilnehmer_Deserialization() {
	var mt = Marktteilnehmer{
		Marktrolle:       marktrolle.LF,
		Makoadresse:      "edifact@my-favourite-marketpartner.de",
		Rollencodetyp:    rollencodetyp.DVGW,
		Rollencodenummer: "9903100000006",
		Geschaeftspartner: Geschaeftspartner{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Geschaeftspartner,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Anrede:               anrede.Divers,
			Name1:                "Müller",
			Name2:                "Lieschen",
			Name3:                "",
			Gewerbekennzeichnung: false,
			HrNummer:             "handelsregister foo",
			Amtsgericht:          "amtsgericht bar",
			Kontaktweg:           0,
			UmsatzsteuerId:       "umsatzsteuer foo",
			GlaeubigerId:         "glauebiger bar",
			EMailAdresse:         "email@lieschen-mueller.de",
			Website:              "https://lieschen-mueller.de",
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.Kunde,
				geschaeftspartnerrolle.Marktpartner,
			},
			Partneradresse: com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
	}
	serializedMarktteilnehmer, err := json.Marshal(mt)
	jsonString := string(serializedMarktteilnehmer)
	then.AssertThat(s.T(), strings.Contains(jsonString, "DVGW"), is.True()) // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "LF"), is.True())   // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMarktteilnehmer, is.Not(is.Nil()))
	var deserializedMarktteilnehmer Marktteilnehmer
	err = json.Unmarshal(serializedMarktteilnehmer, &deserializedMarktteilnehmer)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMarktteilnehmer, is.EqualTo(mt))
}

// Test_Failed_Marktteilnehmer_Validation verifies that the validators of a Marktteilnehmer work
func (s *Suite) Test_Failed_Marktteilnehmer_Validation() {
	validate := validator.New()
	invalidMarktteilnehmer := map[string][]interface{}{
		"required": {
			Marktteilnehmer{},
		},
		"min": {
			Marktteilnehmer{Rollencodenummer: "012345678901"},
		},
		"max": {
			Marktteilnehmer{Rollencodenummer: "01234567890123"},
		},
		"numeric": {
			Marktteilnehmer{Rollencodenummer: "asd"},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMarktteilnehmer)
}

//  Test_Successful_Marktteilnehmer_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Marktteilnehmer_Validation() {
	validate := validator.New()
	validMarktteilnehmers := []interface{}{
		Marktteilnehmer{
			Marktrolle: marktrolle.LF,
			// mako adresse might be empty
			Rollencodetyp:    rollencodetyp.DVGW,
			Rollencodenummer: "9903100000006",
			Geschaeftspartner: Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Geschaeftspartner,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Name1:                "Musterfrau",
				Gewerbekennzeichnung: false,
				Partneradresse: com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördlicher Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   landescode.DE,
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.Dienstleister,
				},
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMarktteilnehmers)
}
