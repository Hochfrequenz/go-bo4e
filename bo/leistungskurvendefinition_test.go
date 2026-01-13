package bo_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/haeufigkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/uebermittelbarkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/internal"
)

var validLeistungskurvendefinition = bo.Leistungskurvendefinition{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.GESCHAEFTSOBJEKT,
		VersionStruktur: "1.1",
	},
	Ausgerollt:                     internal.Ptr(true),
	Aenderungszeitpunkt:            time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	LeistungskurvendefinitionsCode: internal.Ptr("LKD001"),
	Haeufigkeit:                    internal.Ptr(haeufigkeitzaehlzeit.JAEHRLICH),
	Uebermittelbarkeit:             internal.Ptr(uebermittelbarkeitzaehlzeit.ELEKTRONISCH),
	ObererSchwellwert:              internal.Ptr("100"),
}

// Test_Leistungskurvendefinition_Deserialization deserializes a Leistungskurvendefinition json
func Test_Leistungskurvendefinition_Deserialization(t *testing.T) {
	serializedLeistungskurvendefinition, err := json.Marshal(validLeistungskurvendefinition)
	jsonString := string(serializedLeistungskurvendefinition)
	then.AssertThat(t, strings.Contains(jsonString, "LKD001"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedLeistungskurvendefinition, is.Not(is.NilArray[byte]()))
	var deserializedLeistungskurvendefinition bo.Leistungskurvendefinition
	err = json.Unmarshal(serializedLeistungskurvendefinition, &deserializedLeistungskurvendefinition)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedLeistungskurvendefinition, is.EqualTo(validLeistungskurvendefinition))
}

// Test_Leistungskurvendefinition_Fields tests that all fields are correctly serialized
func Test_Leistungskurvendefinition_Fields(t *testing.T) {
	serializedLeistungskurvendefinition, _ := json.Marshal(validLeistungskurvendefinition)
	jsonString := string(serializedLeistungskurvendefinition)
	then.AssertThat(t, strings.Contains(jsonString, "ausgerollt"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "aenderungszeitpunkt"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "leistungskurvendefinitionscode"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "haeufigkeit"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "uebermittelbarkeit"), is.True())
	then.AssertThat(t, strings.Contains(jsonString, "obererSchwellwert"), is.True())
}
