package bo_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsart"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsstatus"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
)

// Test_Vertragspartner_Deserialization deserializes an Vertrag json
func Test_Vertrag_Deserialization(t *testing.T) {
	var contract = bo.Vertrag{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.VERTRAG,
			VersionStruktur:   "2",
			ExterneReferenzen: nil,
		},
		Vertragsnummer: internal.Ptr(""),
		Beschreibung:   internal.Ptr(""),
		Vertragsstatus: internal.Ptr(vertragsstatus.ABGELEHNT),
		Vertragsart:    internal.Ptr(vertragsart.BUENDELVERTRAG),
		Sparte:         sparte.STROM,
		Vertragsbeginn: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Vertragsende:   internal.Ptr(time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)),
		Vertragspartner1: internal.Ptr(bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.GESCHAEFTSPARTNER,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Name1:                "Mustermann",
			Anrede:               internal.Ptr(anrede.HERR),
			Gewerbekennzeichnung: false,
			Partneradresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "28A",
				Landescode:   internal.Ptr(landescode.DE),
			},
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.DIENSTLEISTER,
			},
		}),
		Vertragspartner2: internal.Ptr(bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.GESCHAEFTSPARTNER,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Anrede:               internal.Ptr(anrede.FRAU),
			Name1:                "Musterfrau",
			Gewerbekennzeichnung: false,
			Partneradresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördlicher Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   internal.Ptr(landescode.DE),
			},
			Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
				geschaeftspartnerrolle.DIENSTLEISTER,
			},
		}),
		UnterzeichnerVp1:    nil,
		UnterzeichnerVp2:    nil,
		Vertragskonditionen: nil,
		Vertragsteile: []com.Vertragsteil{
			{
				Vertragsteilbeginn: internal.Ptr(time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC)),
				Vertragsteilende:   internal.Ptr(time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)),
				Lokation:           internal.Ptr("DE0123456789012345678901234567890"),
			},
		},
		Gemeinderabatt: decimal.NewNullDecimal(decimal.NewFromInt(50)),
	}
	serializedContract, err := json.Marshal(contract)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedContract, is.Not(is.NilArray[byte]()))
	var deserializedVertrag bo.Vertrag
	err = json.Unmarshal(serializedContract, &deserializedVertrag)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedVertrag, is.EqualTo(contract))
}

// TestFailedGVertragValidation verifies that the validators of a Vertrag work
func Test_Failed_VertragValidation(t *testing.T) {
	validate := validator.New()
	invalidVertrags := map[string][]interface{}{
		"required": {
			bo.Vertrag{},
		},
		"gtfield": {
			bo.Vertrag{
				Vertragsbeginn: time.Now(),
				Vertragsende:   internal.Ptr(time.Now().Add(time.Hour * -12)),
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidVertrags)
}

// Test_Successful_Vertrag_Validation verifies that a valid BO is validated without errors
func Test_Successful_Vertrag_Validation(t *testing.T) {

	validate := validator.New()
	validVertrag := []bo.BusinessObject{
		bo.Vertrag{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.VERTRAG,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Vertragsnummer: internal.Ptr("asd"),
			Beschreibung:   internal.Ptr("asd"),
			Vertragsstatus: internal.Ptr(vertragsstatus.ANGENOMMEN),
			Vertragsart:    internal.Ptr(vertragsart.BUENDELVERTRAG),
			Sparte:         sparte.STROM,
			Vertragsbeginn: time.Now(),
			Vertragsende:   internal.Ptr(time.Now().Add(time.Hour * 24)),
			Vertragspartner1: internal.Ptr(bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.GESCHAEFTSPARTNER,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Name1:                "Musterfrau",
				Gewerbekennzeichnung: false,
				Partneradresse: &com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördlicher Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   internal.Ptr(landescode.DE),
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.DIENSTLEISTER,
				},
			}),
			Vertragspartner2: internal.Ptr(bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.GESCHAEFTSPARTNER,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Name1:                "Musterfrau",
				Gewerbekennzeichnung: false,
				Partneradresse: &com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördlicher Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   internal.Ptr(landescode.DE),
				},
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.DIENSTLEISTER,
				},
			}),
			UnterzeichnerVp1:    nil,
			UnterzeichnerVp2:    nil,
			Vertragskonditionen: nil,
			Vertragsteile: []com.Vertragsteil{
				{
					Vertragsteilbeginn: internal.Ptr(time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC)),
					Vertragsteilende:   internal.Ptr(time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC)),
					Lokation:           internal.Ptr("DE0123456789012345678901234567890"),
				},
			},
			Gemeinderabatt: decimal.NewNullDecimal(decimal.NewFromInt(50)),
		},
	}
	VerifySuccessfulValidations(t, validate, validVertrag)
}

func Test_Empty_Vertrag_Is_Creatable_Using_BoTyp(t *testing.T) {
	object := bo.NewBusinessObject(botyp.VERTRAG)
	then.AssertThat(t, object, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Vertrag{})))
	then.AssertThat(t, object.GetBoTyp(), is.EqualTo(botyp.VERTRAG))
	then.AssertThat(t, object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func Test_Serialized_Empty_Vertrag_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, bo.NewBusinessObject(botyp.VERTRAG))
}
