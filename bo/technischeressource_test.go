package bo_test

import (
	"encoding/json"
	"reflect"
	"testing"

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
)

var tressource = bo.TechnischeRessource{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:             botyp.TECHNISCHERESSOURCE,
		VersionStruktur:   "1.1",
		ExterneReferenzen: nil,
	},
	TechnischeRessourceId:            internal.Ptr("C816417ST77"),
	VorgelagerteMesslokationsId:      internal.Ptr("DE00713739359S0000000000001222221"),
	ZugeordneteMarktlokationsId:      internal.Ptr("20072281644"),
	ZugeordneteSteuerbareRessourceId: internal.Ptr("20072281644"),
	NennleistungAufnahme: &com.Menge{
		Wert:    newDecimalFromString("100"),
		Einheit: internal.Ptr(mengeneinheit.KW),
	},
	NennleistungAbgabe: &com.Menge{
		Wert:    newDecimalFromString("100"),
		Einheit: internal.Ptr(mengeneinheit.KW),
	},
	Speicherkapazitaet: &com.Menge{
		Wert:    newDecimalFromString("100"),
		Einheit: internal.Ptr(mengeneinheit.KWH),
	},
	TechnischeRessourceNutzung: internal.Ptr(technischeressourcenutzung.SPEICHER),
	Verbrauchsart:              internal.Ptr(technischeressourceverbrauchsart.E_MOBILITAET),
	Waermenutzung:              internal.Ptr(waermenutzung.DIREKTHEIZUNG),
	EMobilitaetsart:            internal.Ptr(emobilitaetsart.LADEPARK),
	Erzeugungsart:              internal.Ptr(erzeugungsart.SOLAR),
	Speicherart:                internal.Ptr(speicherart.PUMPSPEICHER),
}

// Test_TechnischeRessource_Deserialization deserializes an TechnischeRessource json
func Test_TechnischeRessource_Deserialization(t *testing.T) {
	serializedTressource, err := json.Marshal(tressource)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedTressource, is.Not(is.NilArray[byte]()))
	var deserializedTressource bo.TechnischeRessource
	err = json.Unmarshal(serializedTressource, &deserializedTressource)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedTressource, is.EqualTo(tressource))

}

// Test_Failed_TechnischeRessource_Validation verifies that the validators of a TechnischeRessource work
func Test_Failed_TechnischeRessource_Validation(t *testing.T) {
	validate := validator.New()
	invalidTressource := map[string][]interface{}{
		"required": {
			bo.TechnischeRessource{},
		},
	}
	VerifyFailedValidations(t, validate, invalidTressource)
}

// Test_Successful_TechnischeRessource_Validation verifies that a valid BO is validated without errors
func Test_Successful_TechnischeRessource_Validation(t *testing.T) {
	validate := validator.New()
	validTechnischeRessource := []bo.BusinessObject{
		tressource,
	}
	VerifySuccessfulValidations(t, validate, validTechnischeRessource)
}

func Test_Empty_TechnischeRessource_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.TECHNISCHERESSOURCE)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.TechnischeRessource{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.TECHNISCHERESSOURCE))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_TechnischeRessource_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.TECHNISCHERESSOURCE))
}

func Test_Get_TRId_Checksum(t *testing.T) {
	actual := bo.GetTRIdCheckSum("D113735592")
	then.AssertThat(t, actual, is.EqualTo(2))
}

func Test_Get_TRId_Doesnt_Panic(t *testing.T) {
	actual := bo.GetTRIdCheckSum("D5345G7F7F")
	then.AssertThat(t, actual, is.EqualTo(0))
}
