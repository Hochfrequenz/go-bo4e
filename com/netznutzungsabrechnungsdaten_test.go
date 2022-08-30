package com_test

import (
	"encoding/json"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/artikelidtyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/shopspring/decimal"
)

func (s *Suite) Test_Netznutzungsabrechnungsdaten_Deserialization() {
	artikelId := "foo"
	artikelIdTyp := artikelidtyp.GRUPPENARTIKELID
	anzahl := 17
	singulaereBetriebsmittel := com.Menge{
		Wert:    decimal.NewFromFloat(12.34),
		Einheit: mengeneinheit.KWH,
	}
	preisSingulaereBetriebsmittel := com.Preis{
		Wert:    decimal.NewFromFloat(12.34),
		Einheit: waehrungseinheit.EUR,
	}
	var netznutzungsabrechnungsdaten = com.Netznutzungsabrechnungsdaten{
		ArtikelId:                     &artikelId,
		ArtikelIdTyp:                  &artikelIdTyp,
		Anzahl:                        &anzahl,
		Gemeinderabatt:                decimal.NewNullDecimal(decimal.NewFromFloat(12.34)),
		Zuschlag:                      decimal.NewNullDecimal(decimal.NewFromFloat(56.78)),
		Abschlag:                      decimal.NewNullDecimal(decimal.NewFromFloat(9.10)),
		SingulaereBetriebsmittel:      &singulaereBetriebsmittel,
		PreisSingulaereBetriebsmittel: &preisSingulaereBetriebsmittel,
	}
	then.AssertThat(s.T(), netznutzungsabrechnungsdaten, is.Not(is.Nil()))

	serializedNnad, err := json.Marshal(netznutzungsabrechnungsdaten)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), serializedNnad, is.Not(is.Nil()))

	var deserializedNnad com.Netznutzungsabrechnungsdaten
	err = json.Unmarshal(serializedNnad, &deserializedNnad)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), deserializedNnad, is.EqualTo(netznutzungsabrechnungsdaten))
}

// Test_Successful_Validation asserts that the validation does not fail for a valid Netznutzungsabrechnungsdaten
func (s *Suite) Test_Successful_Netznutzungsabrechnungsdaten_Validation() {
	validate := validator.New()
	validNetznutzungsabrechnungsdaten := []interface{}{
		com.Netznutzungsabrechnungsdaten{},
		// add more when there's something like a serious validation. as of now, all fields are optional
	}
	VerfiySuccessfulValidations(s, validate, validNetznutzungsabrechnungsdaten)
}

func (s *Suite) Test_Serialized_Empty_Netznutzungsabrechnungsdaten_Contains_No_Enum_Defaults() {
	s.assert_Does_Not_Serialize_Default_Enums(com.Netznutzungsabrechnungsdaten{})
}
