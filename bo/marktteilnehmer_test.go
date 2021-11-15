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
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
	"reflect"
	"strings"
)

// Test_Marktteilnehmer_Deserialization deserializes an Marktteilnehmer json
func (s *Suite) Test_Marktteilnehmer_Deserialization() {
	var mt = bo.Marktteilnehmer{
		Marktrolle:       marktrolle.LF,
		Makoadresse:      "edifact@my-favourite-marketpartner.de",
		Rollencodetyp:    rollencodetyp.DVGW,
		Rollencodenummer: "9903100000006",
		Geschaeftspartner: bo.Geschaeftspartner{
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
			Kontaktweg:           0,
			UmsatzsteuerId:       "umsatzsteuer foo",
			GlaeubigerId:         "glauebiger bar",
			EMailAdresse:         "email@lieschen-mueller.de",
			Website:              "https://lieschen-mueller.de",
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.KUNDE,
				geschaeftspartnerrolle.MARKTPARTNER,
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
	var deserializedMarktteilnehmer bo.Marktteilnehmer
	err = json.Unmarshal(serializedMarktteilnehmer, &deserializedMarktteilnehmer)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMarktteilnehmer, is.EqualTo(mt))
}

// Test_Failed_Marktteilnehmer_Validation verifies that the validators of a Marktteilnehmer work
func (s *Suite) Test_Failed_Marktteilnehmer_Validation() {
	validate := validator.New()
	invalidMarktteilnehmer := map[string][]interface{}{
		"required": {
			bo.Marktteilnehmer{},
		},
		"min": {
			bo.Marktteilnehmer{Rollencodenummer: "012345678901"},
		},
		"max": {
			bo.Marktteilnehmer{Rollencodenummer: "01234567890123"},
		},
		"numeric": {
			bo.Marktteilnehmer{Rollencodenummer: "asd"},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMarktteilnehmer)
}

//  Test_Successful_Marktteilnehmer_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Marktteilnehmer_Validation() {
	validate := validator.New()
	validMarktteilnehmers := []bo.BusinessObject{
		bo.Marktteilnehmer{
			Marktrolle: marktrolle.LF,
			// mako adresse might be empty
			Rollencodetyp:    rollencodetyp.DVGW,
			Rollencodenummer: "9903100000006",
			Geschaeftspartner: bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.GESCHAEFTSPARTNER,
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
					geschaeftspartnerrolle.DIENSTLEISTER,
				},
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMarktteilnehmers)
}

func (s *Suite) Test_Empty_Markteilnehmer_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.MARKTTEILNEHMER)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Marktteilnehmer{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.MARKTTEILNEHMER))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Marktteilnehmer_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.MARKTTEILNEHMER))
}
