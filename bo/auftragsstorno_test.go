package bo_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

var validAuftragsStorno = bo.AuftragsStorno{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.GESCHAEFTSOBJEKT,
		VersionStruktur: "1.1",
	},
	AuftragsId: "AUFTRAG-12345",
}

// Test_AuftragsStorno_Deserialization deserializes an AuftragsStorno json
func Test_AuftragsStorno_Deserialization(t *testing.T) {
	serializedAuftragsStorno, err := json.Marshal(validAuftragsStorno)
	jsonString := string(serializedAuftragsStorno)
	then.AssertThat(t, strings.Contains(jsonString, "AUFTRAG-12345"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedAuftragsStorno, is.Not(is.NilArray[byte]()))
	var deserializedAuftragsStorno bo.AuftragsStorno
	err = json.Unmarshal(serializedAuftragsStorno, &deserializedAuftragsStorno)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedAuftragsStorno, is.EqualTo(validAuftragsStorno))
}

// Test_AuftragsStorno_Fields tests that all fields are correctly serialized
func Test_AuftragsStorno_Fields(t *testing.T) {
	serializedAuftragsStorno, _ := json.Marshal(validAuftragsStorno)
	jsonString := string(serializedAuftragsStorno)
	then.AssertThat(t, strings.Contains(jsonString, "auftragsId"), is.True())
}
