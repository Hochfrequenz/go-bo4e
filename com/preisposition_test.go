package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/artikelid"
	"github.com/shopspring/decimal"
	"strings"
)

// Test_Preisposition_Deserialization deserializes an Preisposition json
func (s *Suite) Test_Preisposition_Deserialization() {
	artikelId := func(i artikelid.ArtikelId) *artikelid.ArtikelId{ return &i }
	var preisposition = com.Preisposition{
		Berechnungsmethode:       0,
		Leistungstyp:             0,
		Preiseinheit:             0,
		Bezugsgroesse:            0,
		Zeitbasis:                nil,
		Tarifzeit:                nil,
		BDEWArtikelnummer:        nil,
		ArtikelId:                artikelId(artikelid.WIEDERHERSTELLUNG_REGULAER),
		Zonungsgroesse:           nil,
		FreimengeBlindarbeit:     decimal.NullDecimal{},
		FreimengeLeistungsfaktor: decimal.NullDecimal{},
		Preisstaffeln:          []com.Preisstaffel{
	{
		Einheitspreis: decimal.NewFromFloat(1.0),
	},
	},
	}
	then.AssertThat(s.T(), preisposition, is.Not(is.Nil()))

	serializedPreisposition, err := json.Marshal(preisposition)
	then.AssertThat(s.T(), serializedPreisposition, is.Not(is.Nil()))

	jsonString := string(serializedPreisposition)
	then.AssertThat(s.T(), strings.Contains(jsonString, "WIEDERHERSTELLUNG_REGULAER"), is.True()) // stringified enum
	then.AssertThat(s.T(), err, is.Nil())

	var deserializedPreisposition com.Preisposition
	err = json.Unmarshal(serializedPreisposition, &deserializedPreisposition)
	then.AssertThat(s.T(), err, is.Nil())

}