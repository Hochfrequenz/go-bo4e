package market_communication_test

import (
	"encoding/json"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/kontaktart"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

func Test_GetAbsenderCode_Returns_Correct_Value_Without_Uri(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"absender": "9876543210987",
		},
	}
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo("9876543210987"))
	then.AssertThat(t, boneyCombWithDokumentennummer.GetAbsender(), is.NilPtr[bo.Marktteilnehmer]())
}

func Test_GetAbsenderCode_Returns_Correct_Value_With_Uri(t *testing.T) {
	params := []struct{ transaktionsDatenValue, expectedId, testCaseInfo string }{
		{"bo4e://Marktteilnehmer/9876543210987", "9876543210987", "Strom/Gas Sparten Id"},
		{"bo4e://Marktteilnehmer/L34SWH", "L34SWH", "Wasser Sparten Id (kein offizielles Format definiert)"},
	}

	for _, testdata := range params {
		t.Run(testdata.testCaseInfo, func(t *testing.T) {
			var boneyCombWithDokumentennummer = market_communication.BOneyComb{
				Transaktionsdaten: map[string]string{
					"absender": testdata.transaktionsDatenValue,
				},
			}
			then.AssertThat(t, *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo(testdata.expectedId))
		})
	}
}

func Test_SetAbsenderCode(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetAbsenderCode("9876543210987", false)
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo("9876543210987"))
}

func Test_SetAbsenderCode_With_Uri(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetAbsenderCode("9876543210987", true)
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo("9876543210987"))
}

func Test_GetEmpfaengerCode_Returns_Correct_Value_With_Uri(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"empfaenger": "bo4e://Marktteilnehmer/9876543210987",
		},
	}
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetEmpfaengerCode(), is.EqualTo("9876543210987"))
	then.AssertThat(t, boneyCombWithDokumentennummer.GetEmpfaenger(), is.NilPtr[bo.Marktteilnehmer]())
}

func Test_SetEmpfaengerCode(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetEmpfaengerCode("9876543210987", false)
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetEmpfaengerCode(), is.EqualTo("9876543210987"))
}

func Test_SetEmpfaengerCode_With_Uri(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetEmpfaengerCode("9876543210987", true)
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetEmpfaengerCode(), is.EqualTo("9876543210987"))
}

func Test_GetEmpfaenger_Returns_Correct_Value_If_Present(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{
			&bo.Marktteilnehmer{
				Marktrolle:       internal.Ptr(marktrolle.LF),
				Makoadresse:      "edifact@my-favourite-marketpartner.de",
				Rollencodetyp:    rollencodetyp.DVGW,
				Rollencodenummer: "9903100000006",
				Geschaeftspartner: bo.Geschaeftspartner{
					Geschaeftsobjekt: bo.Geschaeftsobjekt{
						BoTyp:             botyp.MARKTTEILNEHMER,
						VersionStruktur:   "1",
						ExterneReferenzen: nil,
					},
					Anrede:               internal.Ptr(anrede.DIVERS),
					Name1:                "Müller",
					Name2:                "Lieschen",
					Name3:                "",
					Gewerbekennzeichnung: false,
					HrNummer:             "handelsregister foo",
					Amtsgericht:          "amtsgericht bar",
					Kontaktwege:          []kontaktart.Kontaktart{},
					UmsatzsteuerId:       "umsatzsteuer foo",
					GlaeubigerId:         "glauebiger bar",
					EMailAdresse:         "email@lieschen-mueller.de",
					Website:              "https://lieschen-mueller.de",
					Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
						geschaeftspartnerrolle.KUNDE,
						geschaeftspartnerrolle.MARKTPARTNER,
					},
					Partneradresse: &com.Adresse{
						Postleitzahl: "82031",
						Ort:          "Grünwald",
						Strasse:      "Nördlicher Münchner Straße",
						Hausnummer:   "27A",
						Landescode:   internal.Ptr(landescode.DE),
					},
				},
			},
			&bo.Marktteilnehmer{
				Marktrolle:       internal.Ptr(marktrolle.LF),
				Makoadresse:      "edifact@my-favourite-marketpartner.de",
				Rollencodetyp:    rollencodetyp.DVGW,
				Rollencodenummer: "9903100000007",
				Geschaeftspartner: bo.Geschaeftspartner{
					Geschaeftsobjekt: bo.Geschaeftsobjekt{
						BoTyp:             botyp.MARKTTEILNEHMER,
						VersionStruktur:   "1",
						ExterneReferenzen: nil,
					},
					Anrede:               internal.Ptr(anrede.DIVERS),
					Name1:                "Müller",
					Name2:                "Lieschen",
					Name3:                "",
					Gewerbekennzeichnung: false,
					HrNummer:             "handelsregister foo",
					Amtsgericht:          "amtsgericht bar",
					Kontaktwege:          []kontaktart.Kontaktart{},
					UmsatzsteuerId:       "umsatzsteuer foo",
					GlaeubigerId:         "glauebiger bar",
					EMailAdresse:         "email@lieschen-mueller.de",
					Website:              "https://lieschen-mueller.de",
					Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
						geschaeftspartnerrolle.KUNDE,
						geschaeftspartnerrolle.MARKTPARTNER,
					},
					Partneradresse: &com.Adresse{
						Postleitzahl: "82031",
						Ort:          "Grünwald",
						Strasse:      "Nördlicher Münchner Straße",
						Hausnummer:   "27A",
						Landescode:   internal.Ptr(landescode.DE),
					},
				},
			},
		},

		Transaktionsdaten: map[string]string{
			"empfaenger": "bo4e://Marktteilnehmer/9903100000006",
			"absender":   "9903100000007",
		},
	}
	then.AssertThat(t, boneyCombWithDokumentennummer.GetEmpfaenger().BoTyp, is.EqualTo[botyp.BOTyp](boneyCombWithDokumentennummer.Stammdaten[0].GetBoTyp()))
	then.AssertThat(t, boneyCombWithDokumentennummer.GetAbsender().BoTyp, is.EqualTo[botyp.BOTyp](boneyCombWithDokumentennummer.Stammdaten[1].GetBoTyp()))
}

func Test_GetAbsenderCode_Returns_Nil_For_Malformed_Data(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"Absender": "asdasd",
		},
	}
	then.AssertThat(t, boneyCombWithDokumentennummer.GetAbsenderCode(), is.NilPtr[string]())
	then.AssertThat(t, boneyCombWithDokumentennummer.GetAbsender(), is.NilPtr[bo.Marktteilnehmer]())
}

func Test_SetAbsender(t *testing.T) {
	mt := bo.Marktteilnehmer{
		Rollencodenummer: "9903100000006",
		Geschaeftspartner: bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:           botyp.MARKTTEILNEHMER,
				VersionStruktur: "1.1",
			},
		},
	}
	bc := market_communication.BOneyComb{}
	bc.SetAbsender(mt, false)
	then.AssertThat(t, *bc.GetAbsenderCode(), is.EqualTo("9903100000006"))
	then.AssertThat(t, bc.GetAbsender(), is.EqualTo(&mt))
}

func Test_SetEmpfaenger(t *testing.T) {
	mt := bo.Marktteilnehmer{
		Rollencodenummer: "9903100000006",
		Geschaeftspartner: bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:           botyp.MARKTTEILNEHMER,
				VersionStruktur: "1.1",
			},
		},
	}
	bc := market_communication.BOneyComb{}
	bc.SetEmpfaenger(mt, true)
	then.AssertThat(t, *bc.GetEmpfaengerCode(), is.EqualTo("9903100000006"))
	then.AssertThat(t, bc.GetEmpfaenger(), is.EqualTo(&mt))
}

func Test_ExtensionData(t *testing.T) {
	mt := bo.Marktteilnehmer{
		Rollencodenummer: "9903100000006",
		Geschaeftspartner: bo.Geschaeftspartner{
			Geschaeftsobjekt: bo.Geschaeftsobjekt{
				BoTyp:           botyp.MARKTTEILNEHMER,
				VersionStruktur: "1.1",
			},
		},
	}
	mt.ExtensionData = map[string]any{}
	mt.ExtensionData["sparte"] = "STROM"
	bc := market_communication.BOneyComb{}
	bc.SetEmpfaenger(mt, true)
	bcJson, _ := json.Marshal(bc)

	unmarshalBc := market_communication.BOneyComb{}
	err := json.Unmarshal(bcJson, &unmarshalBc)
	if err != nil {
		t.Errorf("Error occured while unmarshalling: %v", err)
	}
	empfaenger := unmarshalBc.GetEmpfaenger()
	then.AssertThat(t, empfaenger.ExtensionData["sparte"].(string), is.EqualTo("STROM"))
}
