package bo_test

import (
	"encoding/json"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/reklamationsgrund"
	"github.com/hochfrequenz/go-bo4e/internal"

	"reflect"
	"time"
)

// Test_Reklamation_Deserialization tests serialization and deserialization of Reklamation
func Test_Reklamation_Deserialization(t *testing.T) {
	obis := func(s string) *string { return &s }
	var reklamation = bo.Reklamation{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.REKLAMATION,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		LokationsId:  "12345678910",
		LokationsTyp: lokationstyp.MALO,
		ObisKennzahl: obis("1-0:1.8.0"),
		ZeitraumMesswertanfrage: com.Zeitraum{
			Startzeitpunkt: internal.Ptr(time.Date(2021, 12, 31, 22, 0, 0, 0, time.UTC)),
		}.AsPointer(),
		Reklamationsgrund: reklamationsgrund.WERTE_FEHLEN,
	}
	serializedReklamation, err := json.Marshal(reklamation)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedReklamation, is.Not(is.NilArray[byte]()))
	var deserializedReklamation bo.Reklamation
	err = json.Unmarshal(serializedReklamation, &deserializedReklamation)
	then.AssertThat(t, err, is.Nil())

	then.AssertThat(t, deserializedReklamation, is.EqualTo(reklamation))
}

// Test_Failed_ReklamationValidation verifies that the validators of Reklamation work
func Test_Failed_ReklamationValidation(t *testing.T) {
	validate := validator.New()
	invalidReklamationMap := map[string][]interface{}{
		"required": {
			bo.Reklamation{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				LokationsId:             "",
				LokationsTyp:            0,
				Reklamationsgrund:       0,
				ObisKennzahl:            nil,
				ZeitraumMesswertanfrage: nil,
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidReklamationMap)
}

// Test_Successful_Reklamation_Validation verifies that a valid BO is validated without errors
func Test_Successful_Reklamation_Validation(t *testing.T) {
	obis := func(s string) *string { return &s }
	validate := validator.New()
	validPricat := []bo.BusinessObject{
		bo.Reklamation{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.REKLAMATION,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			LokationsId:  "12345678910",
			LokationsTyp: lokationstyp.MALO,
			ObisKennzahl: obis("1-0:1.8.0"),
			ZeitraumMesswertanfrage: com.Zeitraum{
				Startzeitpunkt: internal.Ptr(time.Date(2021, 11, 30, 22, 0, 0, 0, time.UTC)),
				Startdatum:     internal.Ptr(time.Date(2021, 11, 30, 22, 0, 0, 0, time.UTC)),
				Endzeitpunkt:   internal.Ptr(time.Date(2021, 12, 31, 22, 0, 0, 0, time.UTC)),
				Enddatum:       internal.Ptr(time.Date(2021, 12, 31, 22, 0, 0, 0, time.UTC)),
			}.AsPointer(),
			Reklamationsgrund: reklamationsgrund.WERTE_FEHLEN,
		},
	}
	VerifySuccessfulValidations(t, validate, validPricat)
}

func Test_Empty_Reklamation_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.REKLAMATION)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Reklamation{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.REKLAMATION))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Reklamation_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.REKLAMATION))
}
