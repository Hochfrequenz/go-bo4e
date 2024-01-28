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
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
)

var tranche = bo.Tranche{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.TRANCHE,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
	},
	TrancheId:        "20232281644",
	Sparte:           internal.Ptr(sparte.STROM),
	Aufteilungsmenge: decimal.NewNullDecimal(decimal.NewFromInt(50)),
	ObisKennzahl:     internal.Ptr("1-1:1.9.0"),
}

// Test_Tranche_Deserialization deserializes an Tranche json
func Test_Tranche_Deserialization(t *testing.T) {
	serializedTranche, err := json.Marshal(tranche)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedTranche, is.Not(is.NilArray[byte]()))
	var deserializedTranche bo.Tranche
	err = json.Unmarshal(serializedTranche, &deserializedTranche)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedTranche, is.EqualTo(tranche))

}

// Test_Failed_Tranche_Validation verifies that the validators of a Tranche work
func Test_Failed_Tranche_Validation(t *testing.T) {
	validate := validator.New()
	invalidTranche := map[string][]interface{}{
		"required": {
			bo.Tranche{},
		},
	}
	VerifyFailedValidations(t, validate, invalidTranche)
}

// Test_Successful_Tranche_Validation verifies that a valid BO is validated without errors
func Test_Successful_Tranche_Validation(t *testing.T) {
	validate := validator.New()
	validTranche := []bo.BusinessObject{
		tranche,
	}
	VerifySuccessfulValidations(t, validate, validTranche)
}

func Test_Empty_Tranche_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.TRANCHE)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Tranche{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.TRANCHE))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Tranche_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.TRANCHE))
}
