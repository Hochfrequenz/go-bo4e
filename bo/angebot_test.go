package bo_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

var validAngebot = bo.Angebot{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.ANGEBOT,
		VersionStruktur: "1.1",
	},
	Angebotsnummer:              "2345678",
	Anfragereferenz:             internal.Ptr("Z432342425"),
	Angebotsdatum:               time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	Sparte:                      sparte.GAS,
	Bindefrist:                  nil,
	Angebotgeber:                &validGp,
	Angebotnehmer:               &validGp,
	UnterzeichnerAngebotsnehmer: &validAp,
	UnterzeichnerAngebotsgeber:  &validAp,
	Varianten: []com.Angebotsvariante{{
		Beschreibung: "testbeschreibung",
	}},
}

// Test_Angebot_Deserialization deserializes an Angebot json
func Test_Angebot_Deserialization(t *testing.T) {
	serializedAngebot, err := json.Marshal(validAngebot)
	jsonString := string(serializedAngebot)
	then.AssertThat(t, strings.Contains(jsonString, "GAS"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedAngebot, is.Not(is.NilArray[byte]()))
	var deserializedAngebot bo.Angebot
	err = json.Unmarshal(serializedAngebot, &deserializedAngebot)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedAngebot, is.EqualTo(validAngebot))
}

// Test_Successful_Angebot_Validation verifies that a valid BO is validated without errors
func Test_Successful_Angebot_Validation(t *testing.T) {
	validate := validator.New()
	validAngebote := []bo.BusinessObject{
		validAngebot,
	}
	VerifySuccessfulValidations(t, validate, validAngebote)
}

func Test_Empty_Angebot_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.ANGEBOT)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Angebot{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.ANGEBOT))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Angebot_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.ANGEBOT))
}
