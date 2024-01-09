package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
	"reflect"
)

var tranche = bo.Tranche{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.TRANCHE,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
	},
	TrancheId:        "20232281644",
	Sparte:           sparte.STROM,
	Aufteilungsmenge: decimal.NewNullDecimal(decimal.NewFromInt(50)),
	ObisKennzahl:     internal.Ptr("1-1:1.9.0"),
}

// Test_Tranche_Deserialization deserializes an Tranche json
func (s *Suite) Test_Tranche_Deserialization() {
	serializedTranche, err := json.Marshal(tranche)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedTranche, is.Not(is.NilArray[byte]()))
	var deserializedTranche bo.Tranche
	err = json.Unmarshal(serializedTranche, &deserializedTranche)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedTranche, is.EqualTo(tranche))

}

// Test_Failed_Tranche_Validation verifies that the validators of a Tranche work
func (s *Suite) Test_Failed_Tranche_Validation() {
	validate := validator.New()
	invalidTranche := map[string][]interface{}{
		"required": {
			bo.Tranche{},
		},
	}
	VerfiyFailedValidations(s, validate, invalidTranche)
}

// Test_Successful_Tranche_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Tranche_Validation() {
	validate := validator.New()
	validTranche := []bo.BusinessObject{
		tranche,
	}
	VerfiySuccessfulValidations(s, validate, validTranche)
}

func (s *Suite) Test_Empty_Tranche_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.TRANCHE)
	then.AssertThat(s.T(), object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Tranche{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.TRANCHE))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Tranche_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.TRANCHE))
}
