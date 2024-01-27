package bo_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/hochfrequenz/go-bo4e/internal"

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
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
	"github.com/hochfrequenz/go-bo4e/enum/titel"
)

// Test_Marktteilnehmer_Deserialization deserializes an Marktteilnehmer json
func Test_Marktteilnehmer_Deserialization(t *testing.T) {
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
		},
		Ansprechpartner: &bo.Ansprechpartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp: botyp.ANSPRECHPARTNER,
			},
			Anrede:             anrede.INDIVIDUELL,
			IndividuelleAnrede: "Eure Herzoglichkeit",
			Titel:              titel.PROF_DR,
			Vorname:            "Testor",
			Nachname:           "von Test",
			Kommentar:          "beste",
		},
	}
	serializedMarktteilnehmer, err := json.Marshal(mt)
	jsonString := string(serializedMarktteilnehmer)
	then.AssertThat(t, strings.Contains(jsonString, "DVGW"), is.True()) // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "LF"), is.True())   // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedMarktteilnehmer, is.Not(is.NilArray[byte]()))
	var deserializedMarktteilnehmer bo.Marktteilnehmer
	err = json.Unmarshal(serializedMarktteilnehmer, &deserializedMarktteilnehmer)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedMarktteilnehmer, is.EqualTo(mt))
}

// Test_Failed_Marktteilnehmer_Validation verifies that the validators of a Marktteilnehmer work
func Test_Failed_Marktteilnehmer_Validation(t *testing.T) {
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
	VerifyFailedValidations(t, validate, invalidMarktteilnehmer)
}

// Test_Successful_Marktteilnehmer_Validation verifies that a valid BO is validated without errors
func Test_Successful_Marktteilnehmer_Validation(t *testing.T) {
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
			},
		},
	}
	VerifySuccessfulValidations(t, validate, validMarktteilnehmers)
}

func Test_Empty_Markteilnehmer_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.MARKTTEILNEHMER)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Marktteilnehmer{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.MARKTTEILNEHMER))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Marktteilnehmer_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.MARKTTEILNEHMER))
}
