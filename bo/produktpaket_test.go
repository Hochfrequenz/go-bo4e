package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/produktcode"
	"reflect"
	"testing"
)

var produktpaket = bo.Produktpaket{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.PRODUKTPAKET,
		VersionStruktur: "1.1",
	},
	PaketId: 1,
	Konfigurationen: []com.Produktkonfiguration{
		{
			Code:              produktcode.MESSTECHNISCHE_EINORDNUNG,
			Eigenschaft:       produktcode.INTELLIGENTES_MESSSYSTEM,
			Zusatzeigenschaft: "",
		},
		{
			Code:              produktcode.BILANZKREIS,
			Eigenschaft:       produktcode.BILANZKREIS,
			Zusatzeigenschaft: "4000",
		},
	},
	Prioritaet:           1,
	MussVollstaendigSein: false,
}

// Test_Produktpaket_Deserialization deserializes an Produktpaket json
func Test_Produktpaket_Deserialization(t *testing.T) {
	serializedProduktpaket, err := json.Marshal(produktpaket)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedProduktpaket, is.Not(is.NilArray[byte]()))
	var deserializedProduktpaket bo.Produktpaket
	err = json.Unmarshal(serializedProduktpaket, &deserializedProduktpaket)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedProduktpaket, is.EqualTo(produktpaket))

}

// Test_Failed_Produktpaket_Validation verifies that the validators of a Produktpaket work
func Test_Failed_Produktpaket_Validation(t *testing.T) {
	validate := validator.New()
	invalidProduktpaket := map[string][]interface{}{
		"required": {
			bo.Produktpaket{},
		},
	}
	VerifyFailedValidations(t, validate, invalidProduktpaket)
}

// Test_Successful_Produktpaket_Validation verifies that a valid BO is validated without errors
func Test_Successful_Produktpaket_Validation(t *testing.T) {
	validate := validator.New()
	validProduktpaket := []bo.BusinessObject{
		produktpaket,
	}
	VerifySuccessfulValidations(t, validate, validProduktpaket)
}

func Test_Empty_Produktpaket_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.PRODUKTPAKET)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Produktpaket{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.PRODUKTPAKET))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Produktpaket_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.PRODUKTPAKET))
}
