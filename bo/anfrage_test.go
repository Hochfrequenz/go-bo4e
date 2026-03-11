package bo_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anfragekategorie"
	"github.com/hochfrequenz/go-bo4e/enum/anfragetyp"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
)

var validAnfrage = bo.Anfrage{
	Geschaeftsobjekt: bo.Geschaeftsobjekt{
		BoTyp:           botyp.ANFRAGE,
		VersionStruktur: "1.1",
	},
	LokationsId:  "0815",
	Lokationstyp: lokationstyp.MALO,
	OBISKennzahl: internal.Ptr("1-0:1.8.1"),
	ZeitraumMesswertanfrage: &com.Zeitraum{
		Einheit:        0,
		Dauer:          decimal.NullDecimal{},
		Startdatum:     nil,
		Enddatum:       nil,
		Startzeitpunkt: nil,
		Endzeitpunkt:   nil,
	},
	Anfragekategorie: anfragekategorie.ENTSPERRUNG,
	Anfragetyp:       internal.Ptr(anfragetyp.ZAEHLERSTAENDE),
	Positionsnummer:  "13",
}

// Test_Anfrage_Deserialization deserializes an Anfrage json
func Test_Anfrage_Deserialization(t *testing.T) {
	serializedAnfrage, err := json.Marshal(validAnfrage)
	jsonString := string(serializedAnfrage)
	then.AssertThat(t, strings.Contains(jsonString, "MALO"), is.True()) // stringified enum
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedAnfrage, is.Not(is.NilArray[byte]()))
	var deserializedAnfrage bo.Anfrage
	err = json.Unmarshal(serializedAnfrage, &deserializedAnfrage)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedAnfrage, is.EqualTo(validAnfrage))
	then.AssertThat(t, deserializedAnfrage.Positionsnummer, is.EqualTo(validAnfrage.Positionsnummer))
}

// Test_Anfrage_IntegerPositionsnummer_BackwardsCompat ensures integer positionsnummer values deserialize correctly
func Test_Anfrage_IntegerPositionsnummer_BackwardsCompat(t *testing.T) {
	jsonWithIntPositionsnummer := `{"boTyp":"ANFRAGE","versionStruktur":"1.1","lokationsId":"0815","lokationsTyp":"MALO","anfragekategorie":"ENTSPERRUNG","positionsnummer":13}`
	var anfrage bo.Anfrage
	err := json.Unmarshal([]byte(jsonWithIntPositionsnummer), &anfrage)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, anfrage.Positionsnummer, is.EqualTo("13"))
}
