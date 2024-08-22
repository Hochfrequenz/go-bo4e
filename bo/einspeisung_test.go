package bo_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/internal"
	"reflect"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/eegvermarktungsform"
	"github.com/hochfrequenz/go-bo4e/enum/fernsteuerbarkeitstatus"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
)

// Test_Einspeisung_Deserialization deserializes an Einspeiser json
func Test_Einspeisung_Deserialization(t *testing.T) {
	var einsp = bo.Einspeisung{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp: botyp.EINSPEISUNG,
		},
		MarktlokationsId:        internal.Ptr("84443210210"),
		TrancheId:               internal.Ptr("20232281644"),
		Verguetungsempfaenger:   internal.Ptr(geschaeftspartnerrolle.LIEFERANT),
		EEGVermarktungsform:     internal.Ptr(eegvermarktungsform.KWKG_VERGUETUNG),
		Landescode:              internal.Ptr(landescode.DE),
		FernsteuerbarkeitStatus: internal.Ptr(fernsteuerbarkeitstatus.NICHT_FERNSTEUERBAR),
	}
	serializedEinspeisung, err := json.Marshal(einsp)
	jsonString := string(serializedEinspeisung)
	then.AssertThat(t, strings.Contains(jsonString, "DE"), is.True())              // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "KWKG_VERGUETUNG"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedEinspeisung, is.Not(is.NilArray[byte]()))
	var deserializedEinspeisung bo.Einspeisung
	err = json.Unmarshal(serializedEinspeisung, &deserializedEinspeisung)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedEinspeisung, is.EqualTo(einsp))
}

func Test_Empty_Einspeisung_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.EINSPEISUNG)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Einspeisung{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.EINSPEISUNG))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Einspeisung_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.EINSPEISUNG))
}
