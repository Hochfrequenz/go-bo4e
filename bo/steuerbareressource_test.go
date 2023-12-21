package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkanalsleistungsbeschreibung"
	"reflect"
)

var stbressource = bo.Steuerbareressource{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.STEUERBARERESSOURCE,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
	},
	SteuerbareRessourceId:             "1234567890",
	SteuerkanalsLeistungsbeschreibung: steuerkanalsleistungsbeschreibung.AN_AUS,
	ZugeordnetMSBCodeNr:               "1234567890",
}

// Test_Steuerbareressource_Deserialization deserializes an Steuerbareressource json
func (s *Suite) Test_Steuerbareressource_Deserialization() {
	serializedStbressource, err := json.Marshal(stbressource)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedStbressource, is.Not(is.NilArray[byte]()))
	var deserializedStbressource bo.Steuerbareressource
	err = json.Unmarshal(serializedStbressource, &deserializedStbressource)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedStbressource, is.EqualTo(stbressource))

}

// Test_Failed_SteuerbarreRessource_Validation verifies that the validators of a Steuerbareressource work
func (s *Suite) Test_Failed_SteuerbarreRessource_Validation() {
	validate := validator.New()
	invalidStbressource := map[string][]interface{}{
		"required": {
			bo.Steuerbareressource{},
		},
		"empty_value": {
			bo.Steuerbareressource{SteuerbareRessourceId: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			bo.Steuerbareressource{SteuerkanalsLeistungsbeschreibung: steuerkanalsleistungsbeschreibung.GESTUFT},
			bo.Steuerbareressource{ZugeordnetMSBCodeNr: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		},
		"min": {
			bo.Steuerbareressource{SteuerbareRessourceId: "a"},
			bo.Steuerbareressource{SteuerkanalsLeistungsbeschreibung: steuerkanalsleistungsbeschreibung.GESTUFT},
			bo.Steuerbareressource{ZugeordnetMSBCodeNr: "a"},
		},
		"alphanum": {
			bo.Steuerbareressource{SteuerbareRessourceId: "!123a"},
		},
		"numeric": {
			bo.Steuerbareressource{SteuerbareRessourceId: "asd"},
			bo.Steuerbareressource{ZugeordnetMSBCodeNr: "asd"},
		},
	}
	VerfiyFailedValidations(s, validate, invalidStbressource)
}

// Test_Successful_SteuerbarreRessource_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_SteuerbarreRessource_Validation() {
	validate := validator.New()
	validSteuerbarreRessource := []bo.BusinessObject{
		stbressource,
	}
	VerfiySuccessfulValidations(s, validate, validSteuerbarreRessource)
}

func (s *Suite) Test_Empty_SteuerbarreRessource_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.STEUERBARERESSOURCE)
	then.AssertThat(s.T(), object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Steuerbareressource{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.STEUERBARERESSOURCE))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_SteuerbarreRessource_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.STEUERBARERESSOURCE))
}
