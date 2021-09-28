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
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
	"time"
)

// Test_Rechnung_Deserialization deserializes an Rechnung json
func (s *Suite) Test_Rechnung_Deserialization() {
	var rechnung = Rechnung{
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
	}
	serializedRechnung, err := json.Marshal(rechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedRechnung, is.Not(is.Nil()))
	var deserializedRechnung Rechnung
	err = json.Unmarshal(serializedRechnung, &deserializedRechnung)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedRechnung, is.EqualTo(rechnung))
}

// TestFailedRechnungValidation verifies that the validators of a Rechnung work
func (s *Suite) TestFailedRechnungValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(RechnungStructLevelValidation, Rechnung{})
	invalidRechnungs := map[string][]interface{}{
		"required": {
			Rechnung{},
		},
		"required_if": {
			Rechnung{
				Storno: true,
				// missing OriginalRechnungsnummer
			},
		},
		"min": {
			Rechnung{
				Rechnungspositionen: []com.Rechnungsposition{}, // min 1 entry required
			},
		},
		"Rechnungspositionen[0].Waehrung==Rechnungspositionen[j].Waehrung": {
			Rechnung{
				Rechnungspositionen: []com.Rechnungsposition{
					{
						TeilsummeNetto: com.Betrag{
							Waehrung: waehrungscode.AMD,
						},
					},
					{
						TeilsummeNetto: com.Betrag{
							Waehrung: waehrungscode.ZAR, // != AMD
						},
					},
				},
			},
		},
		"Steuerbetraege[0].Waehrung==Steuerbetraege[j].Waehrung": {
			Rechnung{
				Steuerbetraege: []com.Steuerbetrag{
					{
						Waehrung: waehrungscode.ALL,
					},
					{
						Waehrung: waehrungscode.ZWL, // != ALL
					},
				},
			},
		},
		"GesamtNetto==sum(TeilsummeNetto)": {
			Rechnung{
				GesamtNetto: com.Betrag{
					Wert:     decimal.NewFromFloat(8), // expected 7
					Waehrung: waehrungscode.EUR,
				},
				Rechnungspositionen: []com.Rechnungsposition{
					{
						TeilsummeNetto: com.Betrag{
							Wert:     decimal.NewFromFloat(6),
							Waehrung: waehrungscode.EUR,
						},
					},
					{
						TeilsummeNetto: com.Betrag{
							Wert:     decimal.NewFromFloat(1),
							Waehrung: waehrungscode.EUR, // != AMD
						},
					},
				},
			},
		},
		"GesamtBrutto==sum(TeilsummeSteuer)": {
			Rechnung{
				GesamtBrutto: com.Betrag{
					Wert:     decimal.NewFromFloat(9), // expected 1+2+3+4 = 10
					Waehrung: waehrungscode.EUR,
				},
				Rechnungspositionen: []com.Rechnungsposition{
					{
						TeilsummeSteuer: com.Steuerbetrag{
							Basiswert:  decimal.NewFromFloat(1),
							Steuerwert: decimal.NewFromFloat(2),
							Waehrung:   waehrungscode.EUR,
						},
					},
					{
						TeilsummeSteuer: com.Steuerbetrag{
							Basiswert:  decimal.NewFromFloat(3),
							Steuerwert: decimal.NewFromFloat(4),
							Waehrung:   waehrungscode.EUR,
						},
					},
				},
			},
		},
		"Zuzahlen==GesamtBrutto-Rechnung.Vorausgezahlt-Rechnung.RabattBrutto": {
			Rechnung{
				Zuzahlen: com.Betrag{
					Wert:     decimal.NewFromFloat(4), // expected 10-2-3 = 5
					Waehrung: waehrungscode.EUR,
				},
				GesamtBrutto: com.Betrag{
					Wert:     decimal.NewFromFloat(10),
					Waehrung: waehrungscode.EUR,
				},
				RabattBrutto: &com.Betrag{
					Wert:     decimal.NewFromFloat(2),
					Waehrung: waehrungscode.EUR,
				},
				Vorausgezahlt: &com.Betrag{
					Wert:     decimal.NewFromFloat(3),
					Waehrung: waehrungscode.EUR,
				},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidRechnungs)
}

//  Test_Successful_Rechnung_Validation verifies that a valid BO is validated without errors
func (s *Suite) Test_Successful_Rechnung_Validation() {
	validate := validator.New()
	validate.RegisterStructValidation(RechnungStructLevelValidation, Rechnung{})
	validRechnung := []interface{}{
		Rechnung{
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
	}
	VerfiySuccessfulValidations(s, validate, validRechnung)
}
