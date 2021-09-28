package com

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/bdewartikelnummer"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// Test_Rechnungsposition_Deserialization deserializes an Rechnungsposition json
func (s *Suite) Test_Rechnungsposition_Deserialization() {
	var rechnungsposition = Rechnungsposition{
		Positionsnummer: 17,
		LieferungVon:    time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		LieferungBis:    time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Positionstext:   "foo",
		Zeiteinheit:     zeiteinheit.Jahr,
		Artikelnummer:   bdewartikelnummer.Abgabekwkg,
		LokationsId:     "54321012345",
		PositionsMenge: Menge{
			Wert:    newDecimalFromString("20"),
			Einheit: mengeneinheit.KWH,
		},
		ZeitbezogeneMenge: &Menge{
			Wert:    newDecimalFromString("23"),
			Einheit: mengeneinheit.KUBIKMETER,
		},
		Einzelpreis: Preis{
			Wert:       newDecimalFromString("1"),
			Einheit:    waehrungseinheit.EUR,
			Bezugswert: mengeneinheit.Jahr,
			Status:     preisstatus.Endgueltig,
		},
		TeilsummeNetto: Betrag{
			Wert:     newDecimalFromString("42"),
			Waehrung: waehrungscode.AMD,
		},
		TeilsummeSteuer: Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.Ust7,
			Basiswert:         newDecimalFromString("100"),
			Steuerwert:        newDecimalFromString("7"),
			Waehrung:          waehrungscode.EUR,
		},
		TeilrabattNetto: nil,
	}
	serializedRechnungsposition, err := json.Marshal(rechnungsposition)
	jsonString := string(serializedRechnungsposition)
	then.AssertThat(s.T(), strings.Contains(jsonString, "Abgabekwkg"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedRechnungsposition, is.Not(is.Nil()))
	var deserializedRechnungsposition Rechnungsposition
	err = json.Unmarshal(serializedRechnungsposition, &deserializedRechnungsposition)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedRechnungsposition, is.EqualTo(rechnungsposition))
}

// TestFailedRechnungspositionValidation asserts that the validation fails for invalid Rechnungsposition
func (s *Suite) Test_Failed_RechnungspositionValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(RechnungspositionStructLevelValidation, Rechnungsposition{})
	invalidRechnungsposition := map[string][]interface{}{
		"required": {
			Rechnungsposition{},
		},
		"min": {
			Rechnungsposition{
				LokationsId: "123",
			},
		},
		"max": {
			Rechnungsposition{
				LokationsId: "111111111111111",
			},
		},
		"gtfield": {
			Rechnungsposition{
				LieferungVon: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
				LieferungBis: time.Date(1900, 8, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		"PositionsMenge.Einheit==Einzelpreis.Bezugswert": {
			Rechnungsposition{
				PositionsMenge: Menge{
					Einheit: mengeneinheit.KW,
				},
				Einzelpreis: Preis{
					Bezugswert: mengeneinheit.KUBIKMETER,
				},
			},
		},
		"ZeitbezogeneMenge.Einheit==Einzelpreis.Bezugswert": {
			Rechnungsposition{
				PositionsMenge: Menge{
					Einheit: mengeneinheit.KW,
				},
				Einzelpreis: Preis{
					Bezugswert: mengeneinheit.KW,
				},
				ZeitbezogeneMenge: &Menge{
					Einheit: mengeneinheit.KUBIKMETER,
				},
			},
		},
		"TeilsummeNetto.Wert==Einzelpreis*Positionsmenge": {
			Rechnungsposition{
				PositionsMenge: Menge{
					Wert:    decimal.NewFromFloat(10),
					Einheit: mengeneinheit.KWH,
				},
				Einzelpreis: Preis{
					Wert:       decimal.NewFromFloat(1.5),
					Bezugswert: mengeneinheit.KWH,
					Einheit:    waehrungseinheit.EUR,
				},
				ZeitbezogeneMenge: &Menge{
					Wert:    decimal.NewFromFloat(3),
					Einheit: mengeneinheit.KWH,
				},
				TeilsummeNetto: Betrag{
					Wert:     decimal.NewFromFloat(44), // expected 45 = 3*1.5*10 => validation error
					Waehrung: waehrungscode.EUR,
				},
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidRechnungsposition)
}

// TestSuccessfulRechnungspositionValidation asserts that the validation does not fail for a valid Rechnungsposition
func (s *Suite) Test_Successful_RechnungspositionValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(RechnungspositionStructLevelValidation, Rechnungsposition{})
	validRechnungsposition := []interface{}{
		Rechnungsposition{
			Positionsnummer: 17,
			LieferungVon:    time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			LieferungBis:    time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Positionstext:   "foo",
			Zeiteinheit:     zeiteinheit.Jahr,
			Artikelnummer:   bdewartikelnummer.Abgabekwkg,
			LokationsId:     "54321012345",
			PositionsMenge: Menge{
				Wert:    decimal.NewFromFloat(20),
				Einheit: mengeneinheit.KWH,
			},
			Einzelpreis: Preis{
				Wert:       decimal.NewFromFloat(12),
				Einheit:    waehrungseinheit.EUR,
				Bezugswert: mengeneinheit.KWH,
				Status:     preisstatus.Endgueltig,
			},
			TeilsummeNetto: Betrag{
				Wert:     decimal.NewFromFloat(240),
				Waehrung: waehrungscode.EUR,
			},
			TeilsummeSteuer: Steuerbetrag{
				Steuerkennzeichen: steuerkennzeichen.Ust19,
				Basiswert:         decimal.NewFromFloat(240),
				Steuerwert:        decimal.NewFromFloat(0),
				Waehrung:          waehrungscode.EUR,
			},
			TeilrabattNetto: nil,
		},
	}
	VerfiySuccessfulValidations(s, validate, validRechnungsposition)
}
