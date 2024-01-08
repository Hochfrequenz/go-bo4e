package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/emobilitaetsart"
	"github.com/hochfrequenz/go-bo4e/enum/erzeugungsart"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/speicherart"
	"github.com/hochfrequenz/go-bo4e/enum/technischeressourcenutzung"
	"github.com/hochfrequenz/go-bo4e/enum/technischeressourceverbrauchsart"
	"github.com/hochfrequenz/go-bo4e/enum/waermenutzung"
	"github.com/hochfrequenz/go-bo4e/internal"
	"reflect"
)

var tressource = bo.TechnischeRessource{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.TECHNISCHERESSOURCE,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
	},
	TechnischeRessourceId:            "C816417ST77",
	VorgelagerteMesslokationsId:      "DE00713739359S0000000000001222221",
	ZugeordneteMarktlokationsId:      "20072281644",
	ZugeordneteSteuerbareRessourceId: "20072281644",
	NennleistungAufnahme: com.Menge{
		Wert:    newDecimalFromString("100"),
		Einheit: internal.Ptr(mengeneinheit.KW),
	},
	NennleistungAbgabe: com.Menge{
		Wert:    newDecimalFromString("100"),
		Einheit: internal.Ptr(mengeneinheit.KW),
	},
	Speicherkapazitaet: com.Menge{
		Wert:    newDecimalFromString("100"),
		Einheit: internal.Ptr(mengeneinheit.KWH),
	},
	TechnischeRessourceNutzung: technischeressourcenutzung.SPEICHER,
	Verbrauchsart:              technischeressourceverbrauchsart.E_MOBILITAET,
	Waermenutzung:              waermenutzung.DIREKTHEIZUNG,
	EMobilitaetsart:            emobilitaetsart.LADEPARK,
	Erzeugungsart:              erzeugungsart.SOLAR,
	Speicherart:                speicherart.PUMPSPEICHER,
}

// Test_TechnischeRessource_Deserialization deserializes an TechnischeRessource json
func (s *Suite) Test_TechnischeRessource_Deserialization() {
	serializedTressource, err := json.Marshal(tressource)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedTressource, is.Not(is.NilArray[byte]()))
	var deserializedTressource bo.TechnischeRessource
	err = json.Unmarshal(serializedTressource, &deserializedTressource)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedTressource, is.EqualTo(tressource))

}

// Test_Failed_TechnischeRessource_Validation verifies that the validators of a TechnischeRessource work
func (s *Suite) Test_Failed_TechnischeRessource_Validation() {
	validate := validator.New()
	invalidTressource := map[string][]interface{}{
		"required": {
			bo.TechnischeRessource{},
		},
	}
	VerfiyFailedValidations(s, validate, invalidTressource)
}

// Test_Successful_TechnischeRessource_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_TechnischeRessource_Validation() {
	validate := validator.New()
	validTechnischeRessource := []bo.BusinessObject{
		tressource,
	}
	VerfiySuccessfulValidations(s, validate, validTechnischeRessource)
}

func (s *Suite) Test_Empty_TechnischeRessource_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.TECHNISCHERESSOURCE)
	then.AssertThat(s.T(), object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.TechnischeRessource{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.TECHNISCHERESSOURCE))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_TechnischeRessource_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.TECHNISCHERESSOURCE))
}
