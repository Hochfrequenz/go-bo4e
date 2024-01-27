package bo_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungsart"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Test_Netznutzungsrechnung_Deserialization deserializes an Netznutzungsrechnung json
func Test_Netznutzungsrechnung_Deserialization(t *testing.T) {
	var nnrechnung = bo.Netznutzungsrechnung{
		Rechnung:             serializableRechnung,
		Sparte:               sparte.GAS,
		Absendercodenummer:   "9876543210987",
		Empfaengercodenummer: "0123456789012",
		Nnrechnungsart:       nnrechnungsart.SELBSTAUSGESTELLT,
		Nnrechnungstyp:       nnrechnungstyp.ABSCHLUSSRECHNUNG,
		Original:             true,
		Simuliert:            false,
		LokationsId:          "DE0123456789012345678901234567890",
	}
	serializedNnrechnung, err := json.Marshal(nnrechnung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedNnrechnung, is.Not(is.NilArray[byte]()))
	var deserializedRechnung bo.Netznutzungsrechnung
	err = json.Unmarshal(serializedNnrechnung, &deserializedRechnung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedRechnung, is.EqualTo(nnrechnung))
}

// Test_Failed_Netznutzungsrechnung_Validation verifies that the validators of a Netznutzungsrechnung work
func Test_Failed_Netznutzungsrechnung_Validation(t *testing.T) {
	validate := validator.New()
	invalidNnrs := map[string][]interface{}{
		"required": {
			bo.Netznutzungsrechnung{},
		},
		"max": {
			bo.Netznutzungsrechnung{LokationsId: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			bo.Netznutzungsrechnung{Absendercodenummer: "012345678901234"},
			bo.Netznutzungsrechnung{Empfaengercodenummer: "012345678901234"},
		},
		"min": {
			bo.Netznutzungsrechnung{LokationsId: "a"},
			bo.Netznutzungsrechnung{Absendercodenummer: "012345678901"},
			bo.Netznutzungsrechnung{Empfaengercodenummer: "012345678901"},
		},
		"alphanum": {
			bo.Netznutzungsrechnung{LokationsId: "%!23"},
		},
		"numeric": {
			bo.Netznutzungsrechnung{Absendercodenummer: "asd"},
			bo.Netznutzungsrechnung{Empfaengercodenummer: "asd"},
		},
	}
	VerifyFailedValidations(t, validate, invalidNnrs)
}

// Test_Successful_Netznutzungsrechnung_Validation verifies that a valid BO is validated without errors
func Test_Successful_Netznutzungsrechnung_Validation(t *testing.T) {
	validate := validator.New()
	validNnrs := []bo.BusinessObject{
		bo.Netznutzungsrechnung{
			Rechnung:             completeValidRechnung,
			Sparte:               sparte.GAS,
			Absendercodenummer:   "9876543210987",
			Empfaengercodenummer: "0123456789012",
			Nnrechnungsart:       nnrechnungsart.SELBSTAUSGESTELLT,
			Nnrechnungstyp:       nnrechnungstyp.ABSCHLUSSRECHNUNG,
			Original:             true,
			Simuliert:            false,
			LokationsId:          "DE0123456789012345678901234567890",
		},
	}
	VerifySuccessfulValidations(t, validate, validNnrs)
}

func Test_Empty_NNR_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.NETZNUTZUNGSRECHNUNG)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Netznutzungsrechnung{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.NETZNUTZUNGSRECHNUNG))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Netznutzungsrechnung_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.NETZNUTZUNGSRECHNUNG))
}
