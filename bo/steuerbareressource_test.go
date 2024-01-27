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
)

var stbressource = bo.SteuerbareRessource{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.STEUERBARERESSOURCE,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
	},
	SteuerbareRessourceId:             "1234567890",
	SteuerkanalsLeistungsbeschreibung: nil,
	ZugeordnetMSBCodeNr:               nil,
}

// Test_Steuerbareressource_Deserialization deserializes an SteuerbareRessource json
func Test_Steuerbareressource_Deserialization(t *testing.T) {
	serializedStbressource, err := json.Marshal(stbressource)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedStbressource, is.Not(is.NilArray[byte]()))
	var deserializedStbressource bo.SteuerbareRessource
	err = json.Unmarshal(serializedStbressource, &deserializedStbressource)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedStbressource, is.EqualTo(stbressource))

}

// Test_Failed_SteuerbarreRessource_Validation verifies that the validators of a SteuerbareRessource work
func Test_Failed_SteuerbarreRessource_Validation(t *testing.T) {
	validate := validator.New()
	invalidStbressource := map[string][]interface{}{
		"required": {
			bo.SteuerbareRessource{},
		},
	}
	VerifyFailedValidations(t, validate, invalidStbressource)
}

// Test_Successful_SteuerbarreRessource_Validation verifies that a valid BO is validated without errors
func Test_Successful_SteuerbarreRessource_Validation(t *testing.T) {
	validate := validator.New()
	validSteuerbarreRessource := []bo.BusinessObject{
		stbressource,
	}
	VerifySuccessfulValidations(t, validate, validSteuerbarreRessource)
}

func Test_Empty_SteuerbarreRessource_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.STEUERBARERESSOURCE)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.SteuerbareRessource{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.STEUERBARERESSOURCE))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_SteuerbarreRessource_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.STEUERBARERESSOURCE))
}
