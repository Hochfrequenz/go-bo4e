package com_test

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/hochfrequenz/go-bo4e/enum/ersatzwertbildungsverfahren"
	tarif2 "github.com/hochfrequenz/go-bo4e/enum/tarif"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/gasqualitaet"
	"github.com/hochfrequenz/go-bo4e/enum/grundlageenergiemenge"
	"github.com/hochfrequenz/go-bo4e/enum/korrekturgrund"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/plausibilisierungshinweis"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
)

// Test_Deserialization deserializes an Verbrauch json
func (s *Suite) Test_Verbrauch_Deserialization() {
	gasQualitaet := gasqualitaet.H_GAS
	plausibilisierung := plausibilisierungshinweis.PLAUSIBEL_WG_KUNDENHINWEIS
	grundlageEnergiemenge := grundlageenergiemenge.ZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT
	korrektur := korrekturgrund.AENDERUNG_DER_BERECHNUNG
	ersatzwertbildung := ersatzwertbildungsverfahren.INTERPOLATION
	tarif := tarif2.T1_TARIF1

	var verbrauch = com.Verbrauch{
		Startdatum:               time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		Enddatum:                 time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
		Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
		Obiskennzahl:             "1-0:1.8.1",
		Wert:                     decimal.NewFromFloat(17),
		Einheit:                  mengeneinheit.KWH,
		Statuszusatz: &com.Statuszusatz{
			Gasqualitaet:                &gasQualitaet,
			Plausibilisierungshinweis:   &plausibilisierung,
			GrundlageEnergiemenge:       &grundlageEnergiemenge,
			Korrekturgrund:              &korrektur,
			Ersatzwertbildungsverfahren: &ersatzwertbildung,
			Tarif:                       &tarif,
		},
	}
	serializedVerbrauch, err := json.Marshal(verbrauch)
	jsonString := string(serializedVerbrauch)
	then.AssertThat(s.T(), strings.Contains(jsonString, "KWH"), is.True())                                                                            // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "H_GAS"), is.True())                                                                          // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "PLAUSIBEL_WG_KUNDENHINWEIS"), is.True())                                                     // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "ZAEHLERSTAND_ZUM_BEGIN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT"), is.True()) // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "AENDERUNG_DER_BERECHNUNG"), is.True())                                                       // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "INTERPOLATION"), is.True())                                                                  // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "T1_TARIF1"), is.True())                                                                      // stringified enum
	then.AssertThat(s.T(), strings.Contains(jsonString, "\"2021-08-01T00:00:00Z\""), is.True())                                                       // ISO8601 ftw
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedVerbrauch, is.Not(is.Nil()))
	var deserializedVerbrauch com.Verbrauch
	err = json.Unmarshal(serializedVerbrauch, &deserializedVerbrauch)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedVerbrauch, is.EqualTo(verbrauch))
}

//  Test_Successful_Verbrauch_Validation asserts that the validation does not fail for a valid Verbrauch
func (s *Suite) Test_Successful_Verbrauch_Validation() {
	validate := validator.New()
	validVerbrauch := []interface{}{
		com.Verbrauch{
			Startdatum:               time.Now(),
			Enddatum:                 time.Now().Add(time.Minute * 15),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Obiskennzahl:             "1-0:1.8.1",
			Wert:                     decimal.NewFromFloat(17),
			Einheit:                  mengeneinheit.KWH,
		},
	}
	VerfiySuccessfulValidations(s, validate, validVerbrauch)
}

//  TestVerbrauchFailedValidation verifies that invalid verbrauch values are considered invalid
func (s *Suite) Test_Verbrauch_FailedValidation() {
	validate := validator.New()
	invalidVerbrauchMap := map[string][]interface{}{
		"required": {
			com.Verbrauch{
				Startdatum:               time.Now(),
				Enddatum:                 time.Now().Add(time.Minute * 15),
				Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
				Obiskennzahl:             "1-0:1.8.1",
				Wert:                     decimal.NewFromFloat(0),
				Einheit:                  0,
			},
		},
		"gtfield": {
			com.Verbrauch{
				Startdatum:               time.Now(),
				Enddatum:                 time.Now().Add(time.Minute * -15),
				Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
				Obiskennzahl:             "1-0:1.8.1",
				Wert:                     decimal.NewFromFloat(17),
				Einheit:                  mengeneinheit.KWH,
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidVerbrauchMap)
}

func (s *Suite) Test_Serialized_Empty_Verbrauch_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Verbrauch{})
}
