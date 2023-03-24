package bo_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
	"reflect"
	"time"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/fehlercode"
	"github.com/hochfrequenz/go-bo4e/enum/fehlertyp"

	"github.com/hochfrequenz/go-bo4e/enum/berichtstatus"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

// Test_Statusbericht_Deserialization deserializes an Statusbericht json
func (s *Suite) Test_Statusbericht_Deserialization() {
	gegenstand := "Nachricht123"
	var bericht = bo.Statusbericht{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.STATUSBERICHT,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
			ExtensionData:     unmappeddatamarshaller.ExtensionData{},
		},
		Status:          berichtstatus.ERFOLGREICH,
		Pruefgegenstand: &gegenstand,
		DatumPruefung:   time.Date(2022, 03, 21, 21, 00, 00, 00, time.UTC),
		Fehler: &com.Fehler{
			Typ: fehlertyp.VERARBEITUNG,
			FehlerDetails: []com.FehlerDetail{
				{
					Code:         fehlercode.ERFORDERLICHE_ANGABE_FEHLT,
					Beschreibung: stringAsPointer("fehlt einfach"),
					Ursache: &com.FehlerUrsache{
						Dokument:     &gegenstand,
						Nachricht:    &gegenstand,
						Transaktion:  stringAsPointer("Transaktion"),
						Gruppe:       stringAsPointer("Gruppe"),
						Segment:      stringAsPointer("Segment"),
						Beschreibung: stringAsPointer("Beschreibung"),
					},
				},
			},
		},
	}

	serializedBericht, err := json.Marshal(bericht)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedBericht, is.Not(is.Nil()))
	var deserializedBericht bo.Statusbericht
	err = json.Unmarshal(serializedBericht, &deserializedBericht)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedBericht, is.EqualTo(bericht))
}

// Test_Successful_Vertrag_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Statusbericht_Validation() {
	object := bo.NewBusinessObject(botyp.STATUSBERICHT).(*bo.Statusbericht)
	object.Status = berichtstatus.ERFOLGREICH
	validate := validator.New()
	validVertrag := []bo.BusinessObject{
		*object,
	}
	VerfiySuccessfulValidations(s, validate, validVertrag)
}

func (s *Suite) Test_Empty_Statusbericht_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.STATUSBERICHT)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Statusbericht{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.STATUSBERICHT))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Statusbericht_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.STATUSBERICHT))
}

func stringAsPointer(s string) *string {
	return &s
}
