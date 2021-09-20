package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"strings"
)

// Test_Geschaeftspartner_Deserialization deserializes an Geschaeftspartner json
func (s *Suite) Test_Geschaeftspartner_Deserialization() {
	var gp = Geschaeftspartner{
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
	}
	serializedGp, err := json.Marshal(gp)
	jsonString := string(serializedGp)
	then.AssertThat(s.T(), strings.Contains(jsonString, "Divers"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedGp, is.Not(is.Nil()))
	var deserializedGp Geschaeftspartner
	err = json.Unmarshal(serializedGp, &deserializedGp)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedGp, is.EqualTo(gp))
}

// TestFailedGeschaeftspartnerValidation verifies that the validators of a Geschaeftspartner work
func (s *Suite) TestFailedGeschaeftspartnerValidation() {
	var adresse = com.Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   landescode.DE,
	}
	validate := validator.New()
	invalidVertrags := map[string][]interface{}{
		"required": {
			Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Anrede:                  0,
				Name1:                   "",
				Name2:                   "",
				Name3:                   "",
				Gewerbekennzeichnung:    false,
				HrNummer:                "",
				Amtsgericht:             "",
				Kontaktweg:              0,
				UmsatzsteuerId:          "",
				GlaeubigerId:            "",
				EMailAdresse:            "",
				Website:                 "",
				Geschaeftspartnerrollen: nil,
				Partneradresse:          com.Adresse{},
			},
			Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Anrede:                  0,
				Name1:                   "",
				Name2:                   "",
				Name3:                   "",
				Gewerbekennzeichnung:    false,
				HrNummer:                "",
				Amtsgericht:             "",
				Kontaktweg:              0,
				UmsatzsteuerId:          "",
				GlaeubigerId:            "",
				EMailAdresse:            "",
				Website:                 "",
				Geschaeftspartnerrollen: nil,
				Partneradresse:          adresse,
			},
		},
		"min": {
			Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:           botyp.Geschaeftspartner,
					VersionStruktur: "1",
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{},
			},
		},
		"email": {
			Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:           botyp.Geschaeftspartner,
					VersionStruktur: "1",
				},
				EMailAdresse: "notanemail",
			},
		},
		"url": {
			Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:           botyp.Geschaeftspartner,
					VersionStruktur: "1",
				},
				Website: "not a website",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertrags)
}

//  Test_Successful_Geschaeftspartner_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Geschaeftspartner_Validation() {

	validate := validator.New()
	validGeschaeftspartners := []interface{}{
		Geschaeftspartner{
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
	}
	VerfiySuccessfulValidations(s, validate, validGeschaeftspartners)
}