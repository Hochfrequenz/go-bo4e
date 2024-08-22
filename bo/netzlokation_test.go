package bo_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
)

// Test_Netzlokation_Deserialization deserializes an Netzlokation json
func Test_Netzlokation_Deserialization(t *testing.T) {
	var nelo = bo.Netzlokation{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp: botyp.NETZLOKATION,
		},
		NetzlokationsId: "E3XS7NXYTL6",
		Sparte:          sparte.STROM,
		Netzanschlussleistung: &com.Menge{
			Wert:    decimal.NewFromFloat(125),
			Einheit: internal.Ptr(mengeneinheit.KW),
		},
		GrundzustaendigerMSBCodeNr: internal.Ptr("codenummer"),
		Steuerkanal:                internal.Ptr(false),
		ObisKennzahl:               internal.Ptr("Obis"),
		Verwendungszweck: &com.Verwendungszweck{
			Zweck:      nil,
			Marktrolle: marktrolle.MSB,
		},
		// Konfigurationsprodukte is missing the Marktteilnehmer because adding it causes a serious issue: "could not import github.com/hochfrequenz/go-bo4e/bo (-: import cycle not allowed in test) (typecheck)".
		Konfigurationsprodukte: []com.Konfigurationsprodukt{{
			Produktcode:               "produktcode",
			Leistungskurvendefinition: "leistkurvcode",
			Schaltzeitdefinition:      "schaltzeitdef",
		}},
		EigenschaftMSBLokation:     internal.Ptr(marktrolle.MSB),
		LokationsbuendelObjektcode: internal.Ptr("lokbücode"),
	}
	serializedNetzlokation, err := json.Marshal(nelo)
	jsonString := string(serializedNetzlokation)
	then.AssertThat(t, strings.Contains(jsonString, "STROM"), is.True()) // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "MSB"), is.True())   // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedNetzlokation, is.Not(is.NilArray[byte]()))
	var deserializedNetzlokation bo.Netzlokation
	err = json.Unmarshal(serializedNetzlokation, &deserializedNetzlokation)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedNetzlokation, is.EqualTo(nelo))
}

func Test_Empty_Netzlokation_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.NETZLOKATION)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Netzlokation{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.NETZLOKATION))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Netzlokation_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.NETZLOKATION))
}

func Test_Get_NeloId_Checksum(t *testing.T) {
	actual := bo.GetNeLoIdCheckSum("E113735592")
	then.AssertThat(t, actual, is.EqualTo(1))
}

func Test_Get_NeloId_Doesnt_Panic(t *testing.T) {
	actual := bo.GetNeLoIdCheckSum("E5345G7F6F")
	then.AssertThat(t, actual, is.EqualTo(0))
}
