package market_communication_test

import (
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
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

func (s *Suite) Test_GetAbsenderCode_Returns_Correct_Value_Without_Uri() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"absender": "9876543210987",
		},
	}
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo("9876543210987"))
	then.AssertThat(s.T(), boneyCombWithDokumentennummer.GetAbsender(), is.Nil())
}

func (s *Suite) Test_GetAbsenderCode_Returns_Correct_Value_With_Uri() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"absender": "bo4e://Marktteilnehmer/9876543210987",
		},
	}
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo("9876543210987"))
}

func (s *Suite) Test_SetAbsenderCode() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetAbsenderCode("9876543210987", false)
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo("9876543210987"))
}

func (s *Suite) Test_SetAbsenderCode_With_Uri() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetAbsenderCode("9876543210987", true)
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetAbsenderCode(), is.EqualTo("9876543210987"))
}

func (s *Suite) Test_GetEmpfaengerCode_Returns_Correct_Value_With_Uri() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"empfaenger": "bo4e://Marktteilnehmer/9876543210987",
		},
	}
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetEmpfaengerCode(), is.EqualTo("9876543210987"))
	then.AssertThat(s.T(), boneyCombWithDokumentennummer.GetEmpfaenger(), is.Nil())
}

func (s *Suite) Test_SetEmpfaengerCode() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetEmpfaengerCode("9876543210987", false)
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetEmpfaengerCode(), is.EqualTo("9876543210987"))
}

func (s *Suite) Test_SetEmpfaengerCode_With_Uri() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{}
	boneyCombWithDokumentennummer.SetEmpfaengerCode("9876543210987", true)
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetEmpfaengerCode(), is.EqualTo("9876543210987"))
}

func (s *Suite) Test_GetEmpfaenger_Returns_Correct_Value_If_Present() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{
			&bo.Marktteilnehmer{
				Marktrolle:       marktrolle.LF,
				Makoadresse:      "edifact@my-favourite-marketpartner.de",
				Rollencodetyp:    rollencodetyp.DVGW,
				Rollencodenummer: "9903100000006",
				Geschaeftspartner: bo.Geschaeftspartner{
					Geschaeftsobjekt: bo.Geschaeftsobjekt{
						BoTyp:             botyp.MARKTTEILNEHMER,
						VersionStruktur:   "1",
						ExterneReferenzen: nil,
					},
					Anrede:               anrede.DIVERS,
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
						Landescode:   landescode.DE,
					},
				},
			},
			&bo.Marktteilnehmer{
				Marktrolle:       marktrolle.LF,
				Makoadresse:      "edifact@my-favourite-marketpartner.de",
				Rollencodetyp:    rollencodetyp.DVGW,
				Rollencodenummer: "9903100000007",
				Geschaeftspartner: bo.Geschaeftspartner{
					Geschaeftsobjekt: bo.Geschaeftsobjekt{
						BoTyp:             botyp.MARKTTEILNEHMER,
						VersionStruktur:   "1",
						ExterneReferenzen: nil,
					},
					Anrede:               anrede.DIVERS,
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
						Landescode:   landescode.DE,
					},
				},
			},
		},

		Transaktionsdaten: map[string]string{
			"empfaenger": "bo4e://Marktteilnehmer/9903100000006",
			"absender":   "9903100000007",
		},
	}
	then.AssertThat(s.T(), boneyCombWithDokumentennummer.GetEmpfaenger(), is.EqualTo(boneyCombWithDokumentennummer.Stammdaten[0]))
	then.AssertThat(s.T(), boneyCombWithDokumentennummer.GetAbsender(), is.EqualTo(boneyCombWithDokumentennummer.Stammdaten[1]))
}

func (s *Suite) Test_GetAbsenderCode_Returns_Nil_For_Malformed_Data() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"Absender": "asdasd",
		},
	}
	then.AssertThat(s.T(), boneyCombWithDokumentennummer.GetAbsenderCode(), is.Nil())
	then.AssertThat(s.T(), boneyCombWithDokumentennummer.GetAbsender(), is.Nil())
}

func (s *Suite) Test_SetAbsender() {
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
	then.AssertThat(s.T(), *bc.GetAbsenderCode(), is.EqualTo("9903100000006"))
	then.AssertThat(s.T(), bc.GetAbsender(), is.EqualTo(&mt))
}

func (s *Suite) Test_SetEmpfaenger() {
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
	then.AssertThat(s.T(), *bc.GetEmpfaengerCode(), is.EqualTo("9903100000006"))
	then.AssertThat(s.T(), bc.GetEmpfaenger(), is.EqualTo(&mt))
}
