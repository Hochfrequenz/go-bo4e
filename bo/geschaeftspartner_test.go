package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/kontaktart"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"reflect"
	"strings"
)

// Test_Geschaeftspartner_Deserialization deserializes an Geschaeftspartner json
func (s *Suite) Test_Geschaeftspartner_Deserialization() {
	var gp = bo.Geschaeftspartner{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.GESCHAEFTSPARTNER,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		Anrede:               anrede.DIVERS,
		Name1:                "Müller",
		Name2:                "Lieschen",
		Name3:                "",
		Gewerbekennzeichnung: false,
		HrNummer:             "handelsregister foo",
		Amtsgericht:          "amtsgericht bar",
		Kontaktwege: []kontaktart.Kontaktart{
			kontaktart.E_MAIL,
		},
		UmsatzsteuerId: "umsatzsteuer foo",
		GlaeubigerId:   "glauebiger bar",
		EMailAdresse:   "email@lieschen-mueller.de",
		Website:        "https://lieschen-mueller.de",
		Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
			geschaeftspartnerrolle.KUNDE,
			geschaeftspartnerrolle.MARKTPARTNER,
		},
		Partneradresse: &com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Strasse:      "Nördlicher Münchner Straße",
			Hausnummer:   "27A",
			Landescode:   landescode.DE,
		},
	}
	serializedGp, err := json.Marshal(gp)
	jsonString := string(serializedGp)
	then.AssertThat(s.T(), strings.Contains(jsonString, "DIVERS"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedGp, is.Not(is.Nil()))
	var deserializedGp bo.Geschaeftspartner
	err = json.Unmarshal(serializedGp, &deserializedGp)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedGp, is.EqualTo(gp))
}

// TestFailedGeschaeftspartnerValidation verifies that the validators of a Geschaeftspartner work
func (s *Suite) Test_Failed_GeschaeftspartnerValidation() {
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
			bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
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
				Kontaktwege:             []kontaktart.Kontaktart{},
				UmsatzsteuerId:          "",
				GlaeubigerId:            "",
				EMailAdresse:            "",
				Website:                 "",
				Geschaeftspartnerrollen: nil,
				Partneradresse:          nil,
			},
			bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
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
				Kontaktwege:             []kontaktart.Kontaktart{},
				UmsatzsteuerId:          "",
				GlaeubigerId:            "",
				EMailAdresse:            "",
				Website:                 "",
				Geschaeftspartnerrollen: nil,
				Partneradresse:          &adresse,
			},
		},
		"min": {
			bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.GESCHAEFTSPARTNER,
					VersionStruktur: "1",
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{},
			},
		},
		"email": {
			bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.GESCHAEFTSPARTNER,
					VersionStruktur: "1",
				},
				EMailAdresse: "notanemail",
			},
		},
		"url": {
			bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:           botyp.GESCHAEFTSPARTNER,
					VersionStruktur: "1",
				},
				Website: "not a website",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVertrags)
}

var validGp = bo.Geschaeftspartner{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.GESCHAEFTSPARTNER,
		VersionStruktur:   "1",
		ExterneReferenzen: nil,
	},
	Name1:                "Musterfrau",
	Gewerbekennzeichnung: false,
	Partneradresse: &com.Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   landescode.DE,
	},
	Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
		geschaeftspartnerrolle.DIENSTLEISTER,
	},
}

// Test_Successful_Geschaeftspartner_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Geschaeftspartner_Validation() {

	validate := validator.New()
	validGeschaeftspartners := []bo.BusinessObject{
		validGp,
	}
	VerfiySuccessfulValidations(s, validate, validGeschaeftspartners)
}

func (s *Suite) Test_Empty_Geschaeftspartner_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.GESCHAEFTSPARTNER)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Geschaeftspartner{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.GESCHAEFTSPARTNER))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Geschaeftspartner_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.GESCHAEFTSPARTNER))
}
