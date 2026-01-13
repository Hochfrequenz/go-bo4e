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
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal"
)

var validAuftrag = bo.Auftrag{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.GESCHAEFTSOBJEKT,
		VersionStruktur: "1.1",
	},
	Ausfuehrungsdatum:    internal.Ptr(time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)),
	Fertigstellungsdatum: internal.Ptr(time.Date(2024, 1, 16, 14, 0, 0, 0, time.UTC)),
	Sparte:               internal.Ptr(sparte.STROM),
	Lieferanschrift: &com.Adresse{
		Strasse:      "Musterstraße",
		Hausnummer:   "1",
		Postleitzahl: "12345",
		Ort:          "Musterstadt",
	},
	MarktlokationsId: internal.Ptr("12345678901"),
	Bemerkungen:      []string{"Test-Bemerkung"},
}

// Test_Auftrag_Deserialization deserializes an Auftrag json
func Test_Auftrag_Deserialization(t *testing.T) {
	serializedAuftrag, err := json.Marshal(validAuftrag)
	jsonString := string(serializedAuftrag)
	then.AssertThat(t, strings.Contains(jsonString, "12345678901"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedAuftrag, is.Not(is.NilArray[byte]()))
	var deserializedAuftrag bo.Auftrag
	err = json.Unmarshal(serializedAuftrag, &deserializedAuftrag)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedAuftrag.MarktlokationsId, is.EqualTo(validAuftrag.MarktlokationsId))
}

// Test_Auftrag_Fields tests that all fields are correctly serialized
func Test_Auftrag_Fields(t *testing.T) {
	serializedAuftrag, _ := json.Marshal(validAuftrag)
	jsonString := string(serializedAuftrag)
	then.AssertThat(t, strings.Contains(jsonString, "ausfuehrungsdatum"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "fertigstellungsdatum"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "sparte"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "lieferanschrift"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "marktlokationsId"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "bemerkungen"), is.True())
}
