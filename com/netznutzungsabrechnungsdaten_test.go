package com_test

import (
	"encoding/json"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/artikelidtyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/shopspring/decimal"
)

func Test_Netznutzungsabrechnungsdaten_Deserialization(t *testing.T) {
	artikelId := "foo"
	artikelIdTyp := artikelidtyp.GRUPPENARTIKELID
	anzahl := 17
	singulaereBetriebsmittel := com.Menge{
		Wert:    decimal.NewFromFloat(12.34),
		Einheit: internal.Ptr(mengeneinheit.KWH),
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

	serializedNnad, err := json.Marshal(netznutzungsabrechnungsdaten)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, serializedNnad, is.Not(is.NilArray[byte]()))

	var deserializedNnad com.Netznutzungsabrechnungsdaten
	err = json.Unmarshal(serializedNnad, &deserializedNnad)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedNnad, is.EqualTo(netznutzungsabrechnungsdaten))
}

// Test_Successful_Validation asserts that the validation does not fail for a valid Netznutzungsabrechnungsdaten
func Test_Successful_Netznutzungsabrechnungsdaten_Validation(t *testing.T) {
	validate := validator.New()
	validNetznutzungsabrechnungsdaten := []interface{}{
		com.Netznutzungsabrechnungsdaten{},
		// add more when there's something like a serious validation. as of now, all fields are optional
	}
	VerifySuccessfulValidations(t, validate, validNetznutzungsabrechnungsdaten)
}

func Test_Serialized_Empty_Netznutzungsabrechnungsdaten_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Netznutzungsabrechnungsdaten{})
}
