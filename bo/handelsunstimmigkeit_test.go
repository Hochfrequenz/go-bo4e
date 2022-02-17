package bo_test

import (
	"encoding/json"
	"reflect"
	"strings"

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

func (s *Suite) Test_Handelsunstimmigkeit_Deserialization() {
	h := bo.Handelsunstimmigkeit{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.HANDELSUNSTIMMIGKEIT,
			VersionStruktur: "1",
		},
		Nummer: "123",
		Typ:    handelsunstimmigkeitstyp.HANDELSRECHNUNG,
		Begruendung: com.Handelsunstimmigkeitsbegruendung{
			Referenzen: []string{"ref1", "ref2"},
			Grund:      handelsunstimmigkeitsgrund.NN_MSCONS_UEBERSENDET,
		},
		Betrag: &com.Betrag{
			Wert:     decimal.NewFromInt(456),
			Waehrung: waehrungscode.EUR,
		},
	}

	serialized, err := json.Marshal(h)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serialized, is.Not(is.Nil()))
	then.AssertThat(s.T(), strings.Contains(string(serialized), "NN_MSCONS_UEBERSENDET"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serialized), "HANDELSRECHNUNG"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serialized), "456"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serialized), "123"), is.True())
	var deserialized bo.Handelsunstimmigkeit
	err = json.Unmarshal(serialized, &deserialized)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserialized, is.EqualTo(h))
}

func (s *Suite) Test_Failed_Handelsunstimmigkeit_Validation() {
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
	VerfiyFailedValidations(s, validate, invalidMap)
}

func (s *Suite) Test_Successful_Handelsunstimmigkeit_Validation() {
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
				Grund:      handelsunstimmigkeitsgrund.NN_MSCONS_UEBERSENDET,
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validBO)
}

func (s *Suite) Test_Empty_Handelsunstimmigkeit_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.HANDELSUNSTIMMIGKEIT)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Handelsunstimmigkeit{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.HANDELSUNSTIMMIGKEIT))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Handelsunstimmigkeit_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.HANDELSUNSTIMMIGKEIT))
}
