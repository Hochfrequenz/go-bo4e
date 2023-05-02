package bo_test

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/rufnummernart"
	"github.com/hochfrequenz/go-bo4e/enum/themengebiet"
	"github.com/hochfrequenz/go-bo4e/enum/titel"
)

var validAp = bo.Ansprechpartner{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.ANSPRECHPARTNER,
		VersionStruktur: "1.1",
	},
	Anrede:             anrede.INDIVIDUELL,
	IndividuelleAnrede: "Master of MSCONS (Er/Ihn)",
	Titel:              titel.PROF,
	Vorname:            "Winfried",
	Nachname:           "Müller",
	EMailAdresse:       "send_mails@test.com",
	Kommentar:          "liest seinen Kindern QTY-Segmente zum Einschlafen vor",
	Geschaeftspartner:  &validGp,
	Adresse: &com.Adresse{
		Postleitzahl: "10289",
		Ort:          "Berlin",
		Strasse:      "Ableserstraße",
		Hausnummer:   "17",
		Landescode:   landescode.DE,
	},
	Rufnummern: []com.Rufnummer{{
		Nummerntyp: rufnummernart.FAX_DURCHWAHL,
		Rufnummer:  "030/10239039029302930",
	}},
	Zustaendigkeiten: []com.Zustaendigkeit{{
		Jobtitel:     "Master of MSCONS (Agile)",
		Abteilung:    "MSCONS Special Command Operation Netznutzungs-Spezialisten",
		Themengebiet: themengebiet.MSCONS,
	}},
}

// Test_Ansprechpartner_Deserialization deserializes an Ansprechpartner json
func (s *Suite) Test_Ansprechpartner_Deserialization() {
	serializedAp, err := json.Marshal(validAp)
	jsonString := string(serializedAp)
	then.AssertThat(s.T(), strings.Contains(jsonString, "PROF"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedAp, is.Not(is.NilArray[byte]()))
	var deserializedAp bo.Ansprechpartner
	err = json.Unmarshal(serializedAp, &deserializedAp)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedAp, is.EqualTo(validAp))
}

// Test_Failed_Ansprechpartner_Validation verifies that the validators of an Ansprechpartner work
func (s *Suite) Test_Failed_Ansprechpartner_Validation() {
	validate := validator.New()
	invalidAps := map[string][]interface{}{
		"required": {
			bo.Ansprechpartner{},
		},
		"required_if": {
			bo.Ansprechpartner{
				Anrede:             anrede.INDIVIDUELL,
				IndividuelleAnrede: "", // <-- must be set
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidAps)
}

// Test_Successful_Ansprechpartner_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Ansprechpartner_Validation() {
	validate := validator.New()
	validAnsprechpartners := []bo.BusinessObject{
		validAp,
	}
	VerfiySuccessfulValidations(s, validate, validAnsprechpartners)
}

func (s *Suite) Test_Empty_Ansprechpartner_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.ANSPRECHPARTNER)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Ansprechpartner{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.ANSPRECHPARTNER))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Ansprechpartner_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.ANSPRECHPARTNER))
}
