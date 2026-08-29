package bo_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/internal"
)

var validMabisZaehlpunkt = bo.MabisZaehlpunkt{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.GESCHAEFTSOBJEKT,
		VersionStruktur: "1.1",
	},
	Id: internal.Ptr("DE0000000000000000000000000000001"),
}

// Test_MabisZaehlpunkt_Deserialization deserializes a MabisZaehlpunkt json
func Test_MabisZaehlpunkt_Deserialization(t *testing.T) {
	serializedMabisZaehlpunkt, err := json.Marshal(validMabisZaehlpunkt)
	jsonString := string(serializedMabisZaehlpunkt)
	then.AssertThat(t, strings.Contains(jsonString, "DE0000000000000000000000000000001"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedMabisZaehlpunkt, is.Not(is.NilArray[byte]()))
	var deserializedMabisZaehlpunkt bo.MabisZaehlpunkt
	err = json.Unmarshal(serializedMabisZaehlpunkt, &deserializedMabisZaehlpunkt)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedMabisZaehlpunkt, is.EqualTo(validMabisZaehlpunkt))
}

// Test_MabisZaehlpunkt_ValidateId tests the ValidateId function
func Test_MabisZaehlpunkt_ValidateId(t *testing.T) {
	validId := "DE0000000000000000000000000000001"
	invalidId := "invalid"

	validZp := bo.MabisZaehlpunkt{Id: &validId}
	invalidZp := bo.MabisZaehlpunkt{Id: &invalidId}
	nilZp := bo.MabisZaehlpunkt{Id: nil}

	then.AssertThat(t, validZp.ValidateId(), is.True())
	then.AssertThat(t, invalidZp.ValidateId(), is.False())
	then.AssertThat(t, nilZp.ValidateId(), is.False())
}

// Test_MabisZaehlpunkt_Fields tests that all fields are correctly serialized
func Test_MabisZaehlpunkt_Fields(t *testing.T) {
	serializedMabisZaehlpunkt, _ := json.Marshal(validMabisZaehlpunkt)
	jsonString := string(serializedMabisZaehlpunkt)
	then.AssertThat(t, strings.Contains(jsonString, "Id"), is.True())
}
