package bo_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

func Test_Empty_Lokationszuordnung_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.LOKATIONSZUORDNUNG)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Lokationszuordnung{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.LOKATIONSZUORDNUNG))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Lokationszuordnung_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.LOKATIONSZUORDNUNG))
}

func TestUnmarshalLokationsZuordnungExtensionData(t *testing.T) {
	raw := `
	{
		"boTyp": "LOKATIONSZUORDNUNG",
		"versionStruktur": "1.1",
		"extensionProperty": "This is a custom value."
	}
	`

	var lokationsZuordnung bo.Lokationszuordnung

	if err := json.Unmarshal([]byte(raw), &lokationsZuordnung); err != nil {
		t.Fatalf("could not unmarshal lokationszuordnung: %+v", err)
	}

	if lokationsZuordnung.ExtensionData == nil {
		t.Fatalf("lokationsZuordnung.ExtensionData must not be nil")
	}

	extensionProperty, ok := lokationsZuordnung.ExtensionData["extensionProperty"]
	if !ok {
		t.Fatalf(`lokationsZuordnung.ExtensionData["extensionProperty"] not set`)
	}

	customValue, ok := extensionProperty.(string)
	if !ok {
		t.Fatalf(`lokationsZuordnung.ExtensionData["extensionProperty"] is %T instead of string`, extensionProperty)
	}

	then.AssertThat(t, customValue, is.EqualTo("This is a custom value."))
}
