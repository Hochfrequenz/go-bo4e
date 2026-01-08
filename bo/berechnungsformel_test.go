package bo_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
	"github.com/hochfrequenz/go-bo4e/enum/berechnungsformelnotwendigkeit"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/verwendungszweck"
	"github.com/hochfrequenz/go-bo4e/internal"
)

var validBerechnungsformel = bo.Berechnungsformel{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.GESCHAEFTSOBJEKT,
		VersionStruktur: "1.1",
	},
	Beginndatum:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	Notwendigkeit:    berechnungsformelnotwendigkeit.BERECHNUNGSFORMEL_NOTWENDIG,
	RechenschrittId:  internal.Ptr(12345),
	Verwendungszweck: internal.Ptr(verwendungszweck.NETZNUTZUNGSABRECHNUNG),
	Rechenschritt: &com.Rechenschritt{
		RechenschrittBestandteilId: 1,
		Operation:                  arithmetischeoperation.ADDITION,
		MesslokationId:             internal.Ptr("DE0001234567890123456789012345678"),
	},
}

// Test_Berechnungsformel_Deserialization deserializes a Berechnungsformel json
func Test_Berechnungsformel_Deserialization(t *testing.T) {
	serializedBerechnungsformel, err := json.Marshal(validBerechnungsformel)
	jsonString := string(serializedBerechnungsformel)
	then.AssertThat(t, strings.Contains(jsonString, "rechenschrittId"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedBerechnungsformel, is.Not(is.NilArray[byte]()))
	var deserializedBerechnungsformel bo.Berechnungsformel
	err = json.Unmarshal(serializedBerechnungsformel, &deserializedBerechnungsformel)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedBerechnungsformel.Notwendigkeit, is.EqualTo(validBerechnungsformel.Notwendigkeit))
}

// Test_Berechnungsformel_Fields tests that all fields are correctly serialized
func Test_Berechnungsformel_Fields(t *testing.T) {
	serializedBerechnungsformel, _ := json.Marshal(validBerechnungsformel)
	jsonString := string(serializedBerechnungsformel)
	then.AssertThat(t, strings.Contains(jsonString, "beginndatum"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "notwendigkeit"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "rechenschrittId"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "verwendungszweck"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "rechenschritt"), is.True())
}
