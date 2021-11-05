package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
	"github.com/hochfrequenz/go-bo4e/enum/bilanzierungsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/gebiettyp"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/verbrauchsart"
	"reflect"
	"strings"
)

// Test_Marktlokation_Deserialization tests serialization and deserialization of Marktlokation
func (s *Suite) Test_Marktlokation_Deserialization() {
	var malo = bo.Marktlokation{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.MARKTLOKATION,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		MarktlokationsId:     "51238696781",
		Sparte:               sparte.STROM,
		Energierichtung:      energierichtung.AUSSP,
		Bilanzierungsmethode: bilanzierungsmethode.RLM,
		Verbrauchsart:        verbrauchsart.KL,
		Unterbrechbar:        true,
		Netzebene:            netzebene.MSP,
		Netzbetreibercodenr:  "0815",
		Gebiettyp:            gebiettyp.GRUNDVERSORGUNGSGEBIET,
		Netzgebietnr:         "1234",
		Bilanzierungsgebiet:  "Foo",
		Grundversorgercodenr: "5678",
		Endkunde: bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.GESCHAEFTSPARTNER,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Name1:                "Kruemel",
			Gewerbekennzeichnung: false,
		},
		Lokationsadresse: &com.Adresse{
			Postleitzahl: "82031",
			Ort:          "Grünwald",
			Strasse:      "Nördliche Münchner Straße",
			Hausnummer:   "27A",
			Landescode:   landescode.DE,
		},
		ZugehoerigeMesslokationen: []com.Messlokationszuordnung{
			{
				MesslokationsId: "DE0123456789012345678901234567890",
				Arithmetik:      arithmetischeoperation.ADDITION,
			},
		},
	}
	serializedMalo, err := json.Marshal(malo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMalo, is.Not(is.Nil()))
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "STROM"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "AUSSP"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "RLM"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "KL"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "MSP"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "GRUNDVERSORGUNGSGEBIET"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "DE"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedMalo), "ADDITION"), is.True())
	var deserializedMalo bo.Marktlokation
	err = json.Unmarshal(serializedMalo, &deserializedMalo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMalo, is.EqualTo(malo))
}

func (s *Suite) Test_Empty_Marktlokation_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.MARKTLOKATION)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Marktlokation{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.MARKTLOKATION))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}