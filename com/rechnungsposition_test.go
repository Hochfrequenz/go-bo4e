package com_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/hochfrequenz/go-bo4e/enum/rechnungspositionszuschlag"

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
func Test_Rechnungsposition_Deserialization(t *testing.T) {
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
			Einheit: internal.Ptr(mengeneinheit.KWH),
		},
		ZeitbezogeneMenge: &com.Menge{
			Wert:    newDecimalFromString("23"),
			Einheit: internal.Ptr(mengeneinheit.KUBIKMETER),
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
		TeilrabattNetto:         nil,
		Zuschlag:                internal.Ptr(rechnungspositionszuschlag.PAUSCHALE_NETZENTGELTREDUZIERUNG_ENWG_14A),
		Gesamtzuabschlagsbetrag: internal.Ptr(newDecimalFromString("13.5")),
	}
	serializedRechnungsposition, err := json.Marshal(rechnungsposition)
	jsonString := string(serializedRechnungsposition)
	then.AssertThat(t, strings.Contains(jsonString, "ABGABE_KWKG"), is.True()) // stringified enum
	then.AssertThat(t, strings.Contains(jsonString, "zuschlag"), is.True())
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedRechnungsposition, is.Not(is.NilArray[byte]()))
	var deserializedRechnungsposition com.Rechnungsposition
	err = json.Unmarshal(serializedRechnungsposition, &deserializedRechnungsposition)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedRechnungsposition, is.EqualTo(rechnungsposition))
}

// TestFailedRechnungspositionValidation asserts that the validation fails for invalid Rechnungsposition
func Test_Failed_RechnungspositionValidation(t *testing.T) {
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
					Einheit: internal.Ptr(mengeneinheit.KW),
				},
				Einzelpreis: com.Preis{
					Bezugswert: mengeneinheit.KUBIKMETER,
				},
			},
		},
		"ZeitbezogeneMenge.Einheit==Einzelpreis.Bezugswert": {
			com.Rechnungsposition{
				PositionsMenge: com.Menge{
					Einheit: internal.Ptr(mengeneinheit.KW),
				},
				Einzelpreis: com.Preis{
					Bezugswert: mengeneinheit.KW,
				},
				ZeitbezogeneMenge: &com.Menge{
					Einheit: internal.Ptr(mengeneinheit.KUBIKMETER),
				},
			},
		},
		"TeilsummeNetto.Wert==Einzelpreis*Positionsmenge": {
			com.Rechnungsposition{
				PositionsMenge: com.Menge{
					Wert:    decimal.NewFromFloat(10),
					Einheit: internal.Ptr(mengeneinheit.KWH),
				},
				Einzelpreis: com.Preis{
					Wert:       decimal.NewFromFloat(1.5),
					Bezugswert: mengeneinheit.KWH,
					Einheit:    waehrungseinheit.EUR,
				},
				ZeitbezogeneMenge: &com.Menge{
					Wert:    decimal.NewFromFloat(3),
					Einheit: internal.Ptr(mengeneinheit.KWH),
				},
				TeilsummeNetto: com.Betrag{
					Wert:     decimal.NewFromFloat(44), // expected 45 = 3*1.5*10 => validation error
					Waehrung: waehrungscode.EUR,
				},
			},
		},
	}
	VerifyFailedValidations(t, validate, invalidRechnungsposition)
}

// TestSuccessfulRechnungspositionValidation asserts that the validation does not fail for a valid Rechnungsposition
func Test_Successful_RechnungspositionValidation(t *testing.T) {
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
				Einheit: internal.Ptr(mengeneinheit.KWH),
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
	VerifySuccessfulValidations(t, validate, validRechnungsposition)
}

func Test_Serialized_Empty_Rechnungspositionen_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Rechnungsposition{})
}
