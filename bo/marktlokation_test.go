package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/arithmetischeoperation"
	"github.com/hochfrequenz/go-bo4e/enum/bilanzierungsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/gebiettyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/verbrauchsart"
	"reflect"
	"strings"
)

// Test_Marktlokation_Deserialization tests serialization and deserialization of Marktlokation
func (s *Suite) Test_Marktlokation_Deserialization() {
	f := false
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
		Unterbrechbar:        &f,
		Netzebene:            netzebene.MSP,
		Netzbetreibercodenr:  "0815",
		Gebiettyp:            gebiettyp.GRUNDVERSORGUNGSGEBIET,
		Netzgebietnr:         "1234",
		Bilanzierungsgebiet:  "Foo",
		Grundversorgercodenr: "5678",
		Endkunde: &bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.GESCHAEFTSPARTNER,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Name1:                "Kruemel",
			Gewerbekennzeichnung: false,
			Partneradresse: com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.DIENSTLEISTER,
			},
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
	stringSerializedMalo := string(serializedMalo)
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "STROM"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "AUSSP"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "RLM"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "KL"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "MSP"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "GRUNDVERSORGUNGSGEBIET"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "DE"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "ADDITION"), is.True())
	then.AssertThat(s.T(), strings.Contains(stringSerializedMalo, "\"unterbrechbar\":false"), is.True())
	var deserializedMalo bo.Marktlokation
	err = json.Unmarshal(serializedMalo, &deserializedMalo)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMalo, is.EqualTo(malo))
}

// TestFailedMarktlokationValidation verifies that the validators of Marktlokation work
func (s *Suite) Test_Failed_MarktlokationValidation() {
	validate := validator.New()
	registerError := validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(s.T(), registerError, is.Nil())
	validate.RegisterStructValidation(bo.XorStructLevelValidation, bo.Marktlokation{})

	invalidMarktlokationMap := map[string][]interface{}{
		"required": {
			bo.Marktlokation{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				MarktlokationsId:     "",
				Sparte:               0,
				Energierichtung:      0,
				Bilanzierungsmethode: 0,
				Netzebene:            0,
			},
		},
		"onlyOneAddressType": {
			bo.Marktlokation{
				Lokationsadresse: &com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördliche Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   landescode.DE,
				},
				Katasterinformation: &com.Katasteradresse{
					GemarkungFlur: "Foo",
					Flurstueck:    "Bar",
				},
			},
		},
		"maloid": {
			bo.Marktlokation{
				MarktlokationsId: "12345678918",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMarktlokationMap)
}

//  Test_Successful_Marktlokation_Validation verifies that a valid Marktlokation is validated without errors
func (s *Suite) Test_Successful_Marktlokation_Validation() {
	validate := validator.New()
	registerError := validate.RegisterValidation("maloid", bo.MaloIdFieldLevelValidation)
	then.AssertThat(s.T(), registerError, is.Nil())
	validate.RegisterStructValidation(bo.XorStructLevelValidation, bo.Marktlokation{})
	validMalos := []bo.BusinessObject{
		// Minimal attributes
		bo.Marktlokation{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.MARKTLOKATION,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			MarktlokationsId:     "12345678913",
			Sparte:               sparte.STROM,
			Energierichtung:      energierichtung.EINSP,
			Bilanzierungsmethode: bilanzierungsmethode.SLP,
			Netzebene:            netzebene.NSP,
			Lokationsadresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördliche Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMalos)
}

func (s *Suite) Test_Empty_Marktlokation_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.MARKTLOKATION)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Marktlokation{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.MARKTLOKATION))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}
