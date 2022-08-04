package bo_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/bdewartikelnummer"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"time"
)

// Test_Preisblatt_Deserialization tests serialization and deserialization of Preisblatt
func (s *Suite) Test_Preisblatt_Deserialization() {
	artikel := func(i bdewartikelnummer.BDEWArtikelnummer) *bdewartikelnummer.BDEWArtikelnummer { return &i }
	var pricat = bo.Preisblatt{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:             botyp.PREISBLATT,
			VersionStruktur:   "1",
			ExterneReferenzen: nil,
		},
		Bezeichnung: "Preisblatt",
		Sparte:      sparte.STROM,
		Preisstatus: preisstatus.ENDGUELTIG,
		Herausgeber: bo.Marktteilnehmer{
			Geschaeftspartner: bo.Geschaeftspartner{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             botyp.MARKTTEILNEHMER,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				}},
			Marktrolle:       marktrolle.NB,
			Rollencodenummer: "0815",
			Rollencodetyp:    rollencodetyp.BDEW,
			Makoadresse:      "test@test.com",
		},
		Gueltigkeit: com.Zeitraum{
			Startzeitpunkt: time.Date(2021, 12, 31, 22, 0, 0, 0, time.UTC),
			Startdatum:     time.Date(2021, 12, 31, 22, 0, 0, 0, time.UTC),
		},
		Preispositionen: []com.Preisposition{
			{
				Berechnungsmethode:  0,
				Leistungstyp:        0,
				Leistungsbezeichung: "Preisblatt MSB",
				Preiseinheit:        waehrungseinheit.EUR,
				Bezugsgroesse:       0,
				Zeitbasis:           nil,
				Tarifzeit:           nil,
				BdewArtikelnummer:   artikel(bdewartikelnummer.MSB_INKL_MESSUNG),
				ArtikelId:           "ArtikelID",
				Zonungsgroesse:      nil,

				FreimengeBlindarbeit:     decimal.NullDecimal{Valid: false},
				FreimengeLeistungsfaktor: decimal.NullDecimal{Valid: false},
				Preisstaffeln: []com.Preisstaffel{
					{
						Einheitspreis: decimal.NewFromFloat(1.0),
					},
				},
			},
		},
	}
	serializedPricat, err := json.Marshal(pricat)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedPricat, is.Not(is.Nil()))
	then.AssertThat(s.T(), strings.Contains(string(serializedPricat), "STROM"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedPricat), "MSB_INKL_MESSUNG"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedPricat), "EUR"), is.True())
	then.AssertThat(s.T(), strings.Contains(string(serializedPricat), "BDEW"), is.True())
	var deserializedPricat bo.Preisblatt
	err = json.Unmarshal(serializedPricat, &deserializedPricat)
	then.AssertThat(s.T(), err, is.Nil())

	// The following lines prepare the deserialized pricat for the equality; The capacity of the slices is different, so the slice has to be created manually
	// Also somehow the Preisstaffel raises the assertion, therefore it is replaced with the original
	deserializedPricat.Preispositionen = []com.Preisposition{deserializedPricat.Preispositionen[0]}
	//deserializedPricat.Preispositionen[0].Preisstaffeln = []com.Preisstaffel{deserializedPricat.Preispositionen[0].Preisstaffeln[0]}
	deserializedPricat.Preispositionen[0].Preisstaffeln = []com.Preisstaffel{pricat.Preispositionen[0].Preisstaffeln[0]}
	then.AssertThat(s.T(), deserializedPricat, is.EqualTo(pricat))
}

// Test_Failed_PreisblattValidation verifies that the validators of Preisblatt work
func (s *Suite) Test_Failed_PreisblattValidation() {
	validate := validator.New()
	invalidPreisblattMap := map[string][]interface{}{
		"required": {
			bo.Preisblatt{
				Geschaeftsobjekt: bo.Geschaeftsobjekt{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				Bezeichnung:     "",
				Sparte:          0,
				Preisstatus:     0,
				Herausgeber:     bo.Marktteilnehmer{},
				Gueltigkeit:     com.Zeitraum{},
				Preispositionen: nil,
			},
		},
		"required_without_all": {
			bo.Preisblatt{},
		},
	}
	VerfiyFailedValidations(s, validate, invalidPreisblattMap)
}

//  Test_Successful_Preisblatt_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Preisblatt_Validation() {
	ad := func(adresse com.Adresse) *com.Adresse { return &adresse }
	validate := validator.New()
	validPricat := []bo.BusinessObject{
		bo.Preisblatt{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:             botyp.PREISBLATT,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Bezeichnung: "Preisblatt",
			Sparte:      sparte.STROM,
			Preisstatus: preisstatus.ENDGUELTIG,
			Herausgeber: bo.Marktteilnehmer{
				Geschaeftspartner: bo.Geschaeftspartner{
					Geschaeftsobjekt: bo.Geschaeftsobjekt{
						BoTyp:             botyp.MARKTTEILNEHMER,
						VersionStruktur:   "1",
						ExterneReferenzen: nil,
					},
					Name1:                   "Testor von Test",
					Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{geschaeftspartnerrolle.MARKTPARTNER},
					Partneradresse: ad(com.Adresse{
						Postleitzahl: "12345",
						Ort:          "Testhausen",
						Strasse:      "Testallee",
						Hausnummer:   "12",
						Landescode:   landescode.DE,
					}),
				},
				Marktrolle:       marktrolle.NB,
				Rollencodenummer: "0815081508151",
				Rollencodetyp:    rollencodetyp.BDEW,
				Makoadresse:      "test@test.com",
			},
			Gueltigkeit: com.Zeitraum{
				Startzeitpunkt: time.Date(2021, 12, 31, 22, 00, 00, 00, time.UTC),
				Startdatum:     time.Date(2021, 12, 31, 22, 00, 00, 00, time.UTC),
				Enddatum:       time.Date(9999, 12, 31, 23, 59, 59, 00, time.UTC),
				Endzeitpunkt:   time.Date(9999, 12, 31, 23, 59, 59, 00, time.UTC),
			},
			Preispositionen: []com.Preisposition{},
		},
	}
	VerfiySuccessfulValidations(s, validate, validPricat)
}

func (s *Suite) Test_Empty_Preisblatt_Is_Creatable_Using_BoTyp() {
	object := bo.NewBusinessObject(botyp.PREISBLATT)
	then.AssertThat(s.T(), object, is.Not(is.Nil()))
	then.AssertThat(s.T(), reflect.TypeOf(object), is.EqualTo(reflect.TypeOf(&bo.Preisblatt{})))
	then.AssertThat(s.T(), object.GetBoTyp(), is.EqualTo(botyp.PREISBLATT))
	then.AssertThat(s.T(), object.GetVersionStruktur(), is.EqualTo("1.1"))
}

func (s *Suite) Test_Serialized_Empty_Preisblatt_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(bo.NewBusinessObject(botyp.PREISBLATT))
}
