package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// Test_Marktrechnung_Deserialization deserializes an Rechnung json
func (s *Suite) Test_Marktrechnung_Deserialization() {
	var rechnung = Marktrechnung{
		Rechnungsersteller: Marktteilnehmer{
			Marktrolle:       marktrolle.NB,
			Makoadresse:      "edifact@my-other-marketpartner.de",
			Rollencodetyp:    rollencodetyp.BDEW,
			Rollencodenummer: "9903100000007",
			Geschaeftspartner: Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Marktteilnehmer,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Anrede:               anrede.Divers,
				Name1:                "Müller",
				Name2:                "Lieschen",
				Name3:                "",
				Gewerbekennzeichnung: false,
				HrNummer:             "handelsregister foo",
				Amtsgericht:          "amtsgericht bar",
				Kontaktweg:           0,
				UmsatzsteuerId:       "umsatzsteuer foo",
				GlaeubigerId:         "glauebiger bar",
				EMailAdresse:         "email@lieschen-mueller.de",
				Website:              "https://lieschen-mueller.de",
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.Kunde,
					geschaeftspartnerrolle.Marktpartner,
				},
				Partneradresse: com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördlicher Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   landescode.DE,
				},
			},
		},
		Rechnungsempfaenger: Marktteilnehmer{
			Marktrolle:       marktrolle.LF,
			Makoadresse:      "edifact@my-favourite-marketpartner.de",
			Rollencodetyp:    rollencodetyp.DVGW,
			Rollencodenummer: "9903100000006",
			Geschaeftspartner: Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Marktteilnehmer,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Anrede:               anrede.Divers,
				Name1:                "Müller",
				Name2:                "Lieschen",
				Name3:                "",
				Gewerbekennzeichnung: false,
				HrNummer:             "handelsregister foo",
				Amtsgericht:          "amtsgericht bar",
				Kontaktweg:           0,
				UmsatzsteuerId:       "umsatzsteuer foo",
				GlaeubigerId:         "glauebiger bar",
				EMailAdresse:         "email@lieschen-mueller.de",
				Website:              "https://lieschen-mueller.de",
				Geschaeftspartnerrollen: []geschaeftspartnerrolle.Geschaeftspartnerrolle{
					geschaeftspartnerrolle.Kunde,
					geschaeftspartnerrolle.Marktpartner,
				},
				Partneradresse: com.Adresse{
					Postleitzahl: "82031",
					Ort:          "Grünwald",
					Strasse:      "Nördlicher Münchner Straße",
					Hausnummer:   "27A",
					Landescode:   landescode.DE,
				},
			},
		},
		Rechnung: Rechnung{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Marktrechnung,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			Rechnungstitel:          "De",
			Rechnungsstatus:         rechnungsstatus.Bezahlt,
			Storno:                  false,
			Rechnungsnummer:         "123",
			Rechnungsdatum:          time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
			Faelligkeitsdatum:       time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Rechnungstyp:            rechnungstyp.Endkundenrechnung,
			OriginalRechnungsnummer: "hallo",
			Rechnungsperiode: com.Zeitraum{
				Einheit:        zeiteinheit.Minute,
				Dauer:          decimal.NullDecimal{},
				Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
				Endzeitpunkt:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
			},
			GesamtNetto: com.Betrag{
				Wert:     decimal.NewFromFloat(18.36),
				Waehrung: waehrungscode.EUR,
			},
			GesamtSteuer: com.Betrag{
				Wert:     decimal.NewFromFloat(18.36),
				Waehrung: waehrungscode.EUR,
			},
			GesamtBrutto: com.Betrag{
				Wert:     decimal.NewFromFloat(18.36),
				Waehrung: waehrungscode.EUR,
			},
			Vorausgezahlt: nil,
			RabattBrutto:  nil,
			Zuzahlen: com.Betrag{
				Wert:     decimal.NewFromFloat(18.36),
				Waehrung: waehrungscode.EUR,
			},
			Steuerbetraege:      nil,
			Rechnungspositionen: nil,
		},
	}
	serializedMarktrechnung, err := json.Marshal(rechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMarktrechnung, is.Not(is.Nil()))
	jsonString := string(serializedMarktrechnung)
	then.AssertThat(s.T(), strings.Count(jsonString, "rechnungsempfaenger"), is.EqualTo(1))
	var deserializedMarktrechnung Marktrechnung
	err = json.Unmarshal(serializedMarktrechnung, &deserializedMarktrechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedMarktrechnung, is.EqualTo(rechnung))
}

func (s *Suite) Test_Marktrechnung_Serialization_Contains_No_Geschaeftspartner_As_Rechnungsempfaenger() {
	var marktrechnung = Marktrechnung{
		Rechnungsempfaenger: Marktteilnehmer{
			Rollencodenummer: "9876543210987",
		},
		Rechnung: Rechnung{
			Rechnungsempfaenger: Geschaeftspartner{
				Name1: "not in json",
			},
		},
	}
	serializedMarktrechnung, err := json.Marshal(marktrechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedMarktrechnung, is.Not(is.Nil()))
	jsonString := string(serializedMarktrechnung)
	then.AssertThat(s.T(), jsonString, is.Not(is.Nil()))
	then.AssertThat(s.T(), strings.Contains(jsonString, marktrechnung.Rechnungsempfaenger.Rollencodenummer), is.True())
	then.AssertThat(s.T(), strings.Contains(jsonString, marktrechnung.Rechnung.Rechnungsempfaenger.Name1), is.False())
}
