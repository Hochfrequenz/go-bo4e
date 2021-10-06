package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungsart"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Test_Netznutzungsrechnung_Deserialization deserializes an Netznutzungsrechnung json
func (s *Suite) Test_Netznutzungsrechnung_Deserialization() {
	var nnrechnung = Netznutzungsrechnung{
		Rechnung:             CompleteValidRechnung,
		Sparte:               sparte.Gas,
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
	var deserializedRechnung Netznutzungsrechnung
	err = json.Unmarshal(serializedNnrechnung, &deserializedRechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedRechnung, is.EqualTo(nnrechnung))
}

// Test_Failed_Netznutzungsrechnung_Validation verifies that the validators of a Netznutzungsrechnung work
func (s *Suite) Test_Failed_Netznutzungsrechnung_Validation() {
	validate := validator.New()
	invalidNnrs := map[string][]interface{}{
		"required": {
			Netznutzungsrechnung{},
		},
		"max": {
			Netznutzungsrechnung{LokationsId: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			Netznutzungsrechnung{Absendercodenummer: "012345678901234"},
			Netznutzungsrechnung{Empfaengercodenummer: "012345678901234"},
		},
		"min": {
			Netznutzungsrechnung{LokationsId: "a"},
			Netznutzungsrechnung{Absendercodenummer: "012345678901"},
			Netznutzungsrechnung{Empfaengercodenummer: "012345678901"},
		},
		"alphanum": {
			Netznutzungsrechnung{LokationsId: "%!23"},
		},
		"numeric": {
			Netznutzungsrechnung{Absendercodenummer: "asd"},
			Netznutzungsrechnung{Empfaengercodenummer: "asd"},
		},
	}
	VerfiyFailedValidations(s, validate, invalidNnrs)
}

// Test_Successful_Netznutzungsrechnung_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Netznutzungsrechnung_Validation() {
	validate := validator.New()
	validNnrs := []interface{}{
		Netznutzungsrechnung{
			Rechnung:             CompleteValidRechnung,
			Sparte:               sparte.Gas,
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
