package bo

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/bdewartikelnummer"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungsart"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
	"time"
)

// Test_Netznutzungsrechnung_Deserialization deserializes an Netznutzungsrechnung json
func (s *Suite) Test_Netznutzungsrechnung_Deserialization() {
	var nnrechnung = Netznutzungsrechnung{
		Rechnung: Rechnung{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Rechnung,
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
			Rechnungsersteller: Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Geschaeftspartner,
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
			Rechnungsempfaenger: Geschaeftspartner{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Geschaeftspartner,
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
		Sparte:               sparte.Gas,
		Absendercodenummer:   "9876543210987",
		Empfaengercodenummer: "0123456789012",
		Nnrechnungsart:       nnrechnungsart.SELBSTAUSGESTELLT,
		Nnrechnungstyp:       nnrechnungstyp.ABSCHLUSSRECHNUNG,
		Original:             true,
		Simuliert:            false,
		LokationsId:          "DE0123456789012345678901234567890",
	}
	serializedNnrechnung, err := json.Marshal(nnrechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedNnrechnung, is.Not(is.Nil()))
	var deserializedRechnung Netznutzungsrechnung
	err = json.Unmarshal(serializedNnrechnung, &deserializedRechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedRechnung, is.EqualTo(nnrechnung))
}

// Test_Failed_Netznutzungsrechnung_Validation verifies that the validators of a Netznutzungsrechnung work
func (s *Suite) Test_Failed_Netznutzungsrechnung_Validation() {
	validate := validator.New()
	invalidNnrs := map[string][]interface{}{
		"required": {
			Netznutzungsrechnung{},
		},
		"max": {
			Netznutzungsrechnung{LokationsId: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			Netznutzungsrechnung{Absendercodenummer: "012345678901234"},
			Netznutzungsrechnung{Empfaengercodenummer: "012345678901234"},
		},
		"min": {
			Netznutzungsrechnung{LokationsId: "a"},
			Netznutzungsrechnung{Absendercodenummer: "012345678901"},
			Netznutzungsrechnung{Empfaengercodenummer: "012345678901"},
		},
		"alphanum": {
			Netznutzungsrechnung{LokationsId: "%!23"},
		},
		"numeric": {
			Netznutzungsrechnung{Absendercodenummer: "asd"},
			Netznutzungsrechnung{Empfaengercodenummer: "asd"},
		},
	}
	VerfiyFailedValidations(s, validate, invalidNnrs)
}

// Test_Successful_Netznutzungsrechnung_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Netznutzungsrechnung_Validation() {
	validate := validator.New()
	validNnrs := []interface{}{
		Netznutzungsrechnung{
			Rechnung: Rechnung{
				BusinessObject: BusinessObject{
					BoTyp:             botyp.Rechnung,
					VersionStruktur:   "1",
					ExterneReferenzen: nil,
				},
				Rechnungstitel:          "De",
				Rechnungsstatus:         rechnungsstatus.Bezahlt,
				Storno:                  true,
				Rechnungsnummer:         "123",
				Rechnungsdatum:          time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
				Faelligkeitsdatum:       time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
				Rechnungstyp:            rechnungstyp.Endkundenrechnung,
				OriginalRechnungsnummer: "hallo",
				Rechnungsperiode: com.Zeitraum{
					Einheit:        zeiteinheit.Minute,
					Dauer:          decimal.NullDecimal{Valid: true, Decimal: decimal.NewFromFloat(15)},
					Startzeitpunkt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
					Endzeitpunkt:   time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
				},
				Rechnungsersteller: Geschaeftspartner{
					BusinessObject: BusinessObject{
						BoTyp:             botyp.Geschaeftspartner,
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
				Rechnungsempfaenger: Geschaeftspartner{
					BusinessObject: BusinessObject{
						BoTyp:             botyp.Geschaeftspartner,
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
				GesamtNetto: com.Betrag{
					Wert:     decimal.NewFromFloat(240),
					Waehrung: waehrungscode.EUR,
				},
				GesamtSteuer: com.Betrag{
					Wert:     decimal.NewFromFloat(18.36),
					Waehrung: waehrungscode.EUR,
				},
				GesamtBrutto: com.Betrag{
					Wert:     decimal.NewFromFloat(285.6),
					Waehrung: waehrungscode.EUR,
				},
				Vorausgezahlt: &com.Betrag{
					Wert:     decimal.NewFromFloat(85.6),
					Waehrung: waehrungscode.EUR,
				},
				RabattBrutto: nil,
				Zuzahlen: com.Betrag{
					Wert:     decimal.NewFromFloat(200.),
					Waehrung: waehrungscode.EUR,
				},
				Steuerbetraege: nil,
				Rechnungspositionen: []com.Rechnungsposition{
					{
						Positionsnummer: 17,
						LieferungVon:    time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
						LieferungBis:    time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
						Positionstext:   "foo",
						Zeiteinheit:     zeiteinheit.Jahr,
						Artikelnummer:   bdewartikelnummer.Abgabekwkg,
						LokationsId:     "54321012345",
						PositionsMenge: com.Menge{
							Wert:    decimal.NewFromFloat(20),
							Einheit: mengeneinheit.KWH,
						},
						Einzelpreis: com.Preis{
							Wert:       decimal.NewFromFloat(12),
							Einheit:    waehrungseinheit.EUR,
							Bezugswert: mengeneinheit.KWH,
							Status:     preisstatus.Endgueltig,
						},
						TeilsummeNetto: com.Betrag{
							Wert:     decimal.NewFromFloat(240),
							Waehrung: waehrungscode.EUR,
						},
						TeilsummeSteuer: com.Steuerbetrag{
							Steuerkennzeichen: steuerkennzeichen.Ust19,
							Basiswert:         decimal.NewFromFloat(240),
							Steuerwert:        decimal.NewFromFloat(45.6),
							Waehrung:          waehrungscode.EUR,
						},
						TeilrabattNetto: nil,
					},
				},
			},
			Sparte:               sparte.Gas,
			Absendercodenummer:   "9876543210987",
			Empfaengercodenummer: "0123456789012",
			Nnrechnungsart:       nnrechnungsart.SELBSTAUSGESTELLT,
			Nnrechnungstyp:       nnrechnungstyp.ABSCHLUSSRECHNUNG,
			Original:             true,
			Simuliert:            false,
			LokationsId:          "DE0123456789012345678901234567890",
		},
	}
	VerfiySuccessfulValidations(s, validate, validNnrs)
}
