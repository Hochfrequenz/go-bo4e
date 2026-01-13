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
	"github.com/hochfrequenz/go-bo4e/enum/geraeteart"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal"
)

var validWechsel = bo.Wechsel{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.GESCHAEFTSOBJEKT,
		VersionStruktur: "1.1",
	},
	Sparte: sparte.STROM,
	Geraete: []com.Geraet{
		{
			Geraetenummer: internal.Ptr("GERAET-001"),
			Geraeteart:    internal.Ptr(geraeteart.ZAEHLEINRICHTUNG),
		},
	},
	Wechseldatum: internal.Ptr(time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)),
	Vollstaendig: internal.Ptr(true),
}

// Test_Wechsel_Deserialization deserializes a Wechsel json
func Test_Wechsel_Deserialization(t *testing.T) {
	serializedWechsel, err := json.Marshal(validWechsel)
	jsonString := string(serializedWechsel)
	then.AssertThat(t, strings.Contains(jsonString, "GERAET-001"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedWechsel, is.Not(is.NilArray[byte]()))
	var deserializedWechsel bo.Wechsel
	err = json.Unmarshal(serializedWechsel, &deserializedWechsel)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedWechsel.Sparte, is.EqualTo(validWechsel.Sparte))
}

// Test_Wechsel_Fields tests that all fields are correctly serialized
func Test_Wechsel_Fields(t *testing.T) {
	serializedWechsel, _ := json.Marshal(validWechsel)
	jsonString := string(serializedWechsel)
	then.AssertThat(t, strings.Contains(jsonString, "sparte"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "geraete"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "wechseldatum"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "vollstaendig"), is.True())
}
