package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungsart"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"reflect"
)

// Test_Netznutzungsrechnung_Deserialization deserializes an Netznutzungsrechnung json
func (s *Suite) Test_Netznutzungsrechnung_Deserialization() {
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
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedNnrechnung, is.Not(is.Nil()))
	var deserializedRechnung bo.Netznutzungsrechnung
	err = json.Unmarshal(serializedNnrechnung, &deserializedRechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedRechnung, is.EqualTo(nnrechnung))
}

// Test_Failed_Netznutzungsrechnung_Validation verifies that the validators of a Netznutzungsrechnung work
func (s *Suite) Test_Failed_Netznutzungsrechnung_Validation() {
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
	VerfiyFailedValidations(s, validate, invalidNnrs)
}

// Test_Successful_Netznutzungsrechnung_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Netznutzungsrechnung_Validation() {
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
	VerfiySuccessfulValidations(s, validate, validNnrs)
}

func (s *Suite) Test_Empty_NNR_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.NETZNUTZUNGSRECHNUNG)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Netznutzungsrechnung{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.NETZNUTZUNGSRECHNUNG))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}
