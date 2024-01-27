package bo_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

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
	"github.com/hochfrequenz/go-bo4e/internal"
)

// Test_Geschaeftspartner_Deserialization deserializes an Geschaeftspartner json
func Test_Geschaeftspartner_Deserialization(t *testing.T) {
	var gp = bo.Geschaeftspartner{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.GESCHAEFTSPARTNER,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		Anrede:               internal.Ptr(anrede.DIVERS),
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
			Landescode:   internal.Ptr(landescode.DE),
		},
	}
	serializedGp, err := json.Marshal(gp)
	jsonString := string(serializedGp)
	then.AssertThat(t, strings.Contains(jsonString, "DIVERS"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedGp, is.Not(is.NilArray[byte]()))
	var deserializedGp bo.Geschaeftspartner
	err = json.Unmarshal(serializedGp, &deserializedGp)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedGp, is.EqualTo(gp))
}

// TestFailedGeschaeftspartnerValidation verifies that the validators of a Geschaeftspartner work
func Test_Failed_GeschaeftspartnerValidation(t *testing.T) {
	var adresse = com.Adresse{
		Postleitzahl: "82031",
		Ort:          "Grünwald",
		Strasse:      "Nördlicher Münchner Straße",
		Hausnummer:   "27A",
		Landescode:   internal.Ptr(landescode.DE),
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
				Anrede:                  nil,
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
				Anrede:                  nil,
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
	VerifyFailedValidations(t, validate, invalidVertrags)
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
		Landescode:   internal.Ptr(landescode.DE),
	},
	Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
		geschaeftspartnerrolle.DIENSTLEISTER,
	},
}

// Test_Successful_Geschaeftspartner_Validation verifies that a valid BO is validated without errors
func Test_Successful_Geschaeftspartner_Validation(t *testing.T) {

	validate := validator.New()
	validGeschaeftspartners := []bo.BusinessObject{
		validGp,
	}
	VerifySuccessfulValidations(t, validate, validGeschaeftspartners)
}

func Test_Empty_Geschaeftspartner_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.GESCHAEFTSPARTNER)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Geschaeftspartner{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.GESCHAEFTSPARTNER))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Geschaeftspartner_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.GESCHAEFTSPARTNER))
}
