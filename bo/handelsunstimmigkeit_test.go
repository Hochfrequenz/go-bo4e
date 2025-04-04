package bo_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/internal"
	"reflect"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/corbym/gocrest/is"

	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/handelsunstimmigkeitsgrund"
	"github.com/hochfrequenz/go-bo4e/enum/handelsunstimmigkeitstyp"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

func Test_Handelsunstimmigkeit_Deserialization(t *testing.T) {
	h := bo.Handelsunstimmigkeit{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.HANDELSUNSTIMMIGKEIT,
			VersionStruktur: "1",
		},
		Nummer: "123",
		Typ:    handelsunstimmigkeitstyp.HANDELSRECHNUNG,
		Begruendung: com.Handelsunstimmigkeitsbegruendung{
			Referenzen: []string{"ref1", "ref2"},
			Grund:      internal.Ptr(handelsunstimmigkeitsgrund.NN_MSCONS_UEBERSENDET),
		},
		Betrag: &com.Betrag{
			Wert:     decimal.NewFromInt(456),
			Waehrung: waehrungscode.EUR,
		},
	}

	serialized, err := json.Marshal(h)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serialized, is.Not(is.NilArray[byte]()))
	then.AssertThat(t, strings.Contains(string(serialized), "NN_MSCONS_UEBERSENDET"), is.True())
	then.AssertThat(t, strings.Contains(string(serialized), "HANDELSRECHNUNG"), is.True())
	then.AssertThat(t, strings.Contains(string(serialized), "456"), is.True())
	then.AssertThat(t, strings.Contains(string(serialized), "123"), is.True())
	var deserialized bo.Handelsunstimmigkeit
	err = json.Unmarshal(serialized, &deserialized)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserialized, is.EqualTo(h))
}

func Test_Failed_Handelsunstimmigkeit_Validation(t *testing.T) {
	validate := validator.New()
	invalidMap := map[string][]interface{}{
		"required": {
			bo.Handelsunstimmigkeit{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Nummer:      "",
				Typ:         0,
				Begruendung: com.Handelsunstimmigkeitsbegruendung{},
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidMap)
}

func Test_Successful_Handelsunstimmigkeit_Validation(t *testing.T) {
	validate := validator.New()
	validBO := []bo.BusinessObject{
		bo.Handelsunstimmigkeit{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.HANDELSUNSTIMMIGKEIT,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Nummer: "123",
			Typ:    handelsunstimmigkeitstyp.LIEFERSCHEIN_HANDELSUNSTIMMIGKEITSTYP,
			Begruendung: com.Handelsunstimmigkeitsbegruendung{
				Referenzen: nil,
				Grund:      internal.Ptr(handelsunstimmigkeitsgrund.NN_MSCONS_UEBERSENDET),
			},
		},
	}
	VerifySuccessfulValidations(t, validate, validBO)
}

func Test_Empty_Handelsunstimmigkeit_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.HANDELSUNSTIMMIGKEIT)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Handelsunstimmigkeit{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.HANDELSUNSTIMMIGKEIT))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Handelsunstimmigkeit_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.HANDELSUNSTIMMIGKEIT))
}
