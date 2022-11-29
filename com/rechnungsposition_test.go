package com_test

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/hochfrequenz/go-bo4e/internal"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/bdewartikelnummer"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/preisstatus"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
)

var korrekturfaktor float32 = -1

// Test_Rechnungsposition_Deserialization deserializes an Rechnungsposition json
func (s *Suite) Test_Rechnungsposition_Deserialization() {
	var rechnungsposition = com.Rechnungsposition{
		Positionsnummer: 17,
		LieferungVon:    time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		LieferungBis:    time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		Positionstext:   "foo",
		Zeiteinheit:     zeiteinheit.JAHR,
		Artikelnummer:   internal.Ptr(bdewartikelnummer.ABGABE_KWKG),
		LokationsId:     "54321012345",
		PositionsMenge: com.Menge{
			Wert:    newDecimalFromString("20"),
			Einheit: mengeneinheit.KWH,
		},
		ZeitbezogeneMenge: &com.Menge{
			Wert:    newDecimalFromString("23"),
			Einheit: mengeneinheit.KUBIKMETER,
		},
		Korrekturfaktor: &korrekturfaktor,
		Einzelpreis: com.Preis{
			Wert:       newDecimalFromString("1"),
			Einheit:    waehrungseinheit.EUR,
			Bezugswert: mengeneinheit.JAHR,
			Status:     preisstatus.ENDGUELTIG,
		},
		TeilsummeNetto: com.Betrag{
			Wert:     newDecimalFromString("42"),
			Waehrung: waehrungscode.AMD,
		},
		TeilsummeSteuer: com.Steuerbetrag{
			Steuerkennzeichen: steuerkennzeichen.UST_7,
			Basiswert:         newDecimalFromString("100"),
			Steuerwert:        newDecimalFromString("7"),
			Waehrung:          waehrungscode.EUR,
		},
		TeilrabattNetto: nil,
	}
	serializedRechnungsposition, err := json.Marshal(rechnungsposition)
	jsonString := string(serializedRechnungsposition)
	then.AssertThat(s.T(), strings.Contains(jsonString, "ABGABE_KWKG"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedRechnungsposition, is.Not(is.Nil()))
	var deserializedRechnungsposition com.Rechnungsposition
	err = json.Unmarshal(serializedRechnungsposition, &deserializedRechnungsposition)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedRechnungsposition, is.EqualTo(rechnungsposition))
}

// TestFailedRechnungspositionValidation asserts that the validation fails for invalid Rechnungsposition
func (s *Suite) Test_Failed_RechnungspositionValidation() {
	validate := validator.New()
	validate.RegisterStructValidation(com.RechnungspositionStructLevelValidation, com.Rechnungsposition{})
	invalidRechnungsposition := map[string][]interface{}{
		"required": {
			com.Rechnungsposition{},
		},
		"min": {
			com.Rechnungsposition{
				LokationsId: "123",
			},
		},
		"max": {
			com.Rechnungsposition{
				LokationsId: "111111111111111",
			},
		},
		"gtfield": {
			com.Rechnungsposition{
				LieferungVon: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
				LieferungBis: time.Date(1900, 8, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		"PositionsMenge.Einheit==Einzelpreis.Bezugswert": {
			com.Rechnungsposition{
				PositionsMenge: com.Menge{
					Einheit: mengeneinheit.KW,
				},
				Einzelpreis: com.Preis{
					Bezugswert: mengeneinheit.KUBIKMETER,
				},
			},
		},
		"ZeitbezogeneMenge.Einheit==Einzelpreis.Bezugswert": {
			com.Rechnungsposition{
				PositionsMenge: com.Menge{
					Einheit: mengeneinheit.KW,
				},
				Einzelpreis: com.Preis{
					Bezugswert: mengeneinheit.KW,
				},
				ZeitbezogeneMenge: &com.Menge{
					Einheit: mengeneinheit.KUBIKMETER,
				},
			},
		},
		"TeilsummeNetto.Wert==Einzelpreis*Positionsmenge": {
			com.Rechnungsposition{
				PositionsMenge: com.Menge{
					Wert:    decimal.NewFromFloat(10),
					Einheit: mengeneinheit.KWH,
				},
				Einzelpreis: com.Preis{
					Wert:       decimal.NewFromFloat(1.5),
					Bezugswert: mengeneinheit.KWH,
					Einheit:    waehrungseinheit.EUR,
				},
				ZeitbezogeneMenge: &com.Menge{
					Wert:    decimal.NewFromFloat(3),
					Einheit: mengeneinheit.KWH,
				},
				TeilsummeNetto: com.Betrag{
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
	validate.RegisterStructValidation(com.RechnungspositionStructLevelValidation, com.Rechnungsposition{})
	validRechnungsposition := []interface{}{
		com.Rechnungsposition{
			Positionsnummer: 17,
			LieferungVon:    time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
			LieferungBis:    time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			Positionstext:   "foo",
			Zeiteinheit:     zeiteinheit.JAHR,
			Artikelnummer:   internal.Ptr(bdewartikelnummer.ABGABE_KWKG),
			LokationsId:     "54321012345",
			PositionsMenge: com.Menge{
				Wert:    decimal.NewFromFloat(20),
				Einheit: mengeneinheit.KWH,
			},
			Einzelpreis: com.Preis{
				Wert:       decimal.NewFromFloat(12),
				Einheit:    waehrungseinheit.EUR,
				Bezugswert: mengeneinheit.KWH,
				Status:     preisstatus.ENDGUELTIG,
			},
			TeilsummeNetto: com.Betrag{
				Wert:     decimal.NewFromFloat(240),
				Waehrung: waehrungscode.EUR,
			},
			TeilsummeSteuer: com.Steuerbetrag{
				Steuerkennzeichen: steuerkennzeichen.UST_19,
				Basiswert:         decimal.NewFromFloat(240),
				Steuerwert:        decimal.NewFromFloat(0),
				Waehrung:          waehrungscode.EUR,
			},
			TeilrabattNetto: nil,
		},
	}
	VerfiySuccessfulValidations(s, validate, validRechnungsposition)
}

func (s *Suite) Test_Serialized_Empty_Rechnungspositionen_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Rechnungsposition{})
}
