package com_test

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/enum/ermittlungleistungsmaximum"
	"github.com/hochfrequenz/go-bo4e/enum/haeufigkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/uebermittelbarkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlzeitdefinitiontyp"
	"github.com/hochfrequenz/go-bo4e/internal"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
)

func Test_Zaehlzeit_Deserialization(t *testing.T) {
	zaehlzeit := com.Zaehlzeit{
		Code:                       internal.Ptr("HL1"),
		Haeufigkeit:                internal.Ptr(haeufigkeitzaehlzeit.JAEHRLICH),
		Uebermittelbarkeit:         internal.Ptr(uebermittelbarkeitzaehlzeit.ELEKTRONISCH),
		ErmittlungLeistungsmaximum: internal.Ptr(ermittlungleistungsmaximum.VERWENDUNG_HOCHLASTFENSTER),
		IstBestellbar:              internal.Ptr(true),
		Typ:                        internal.Ptr(zaehlzeitdefinitiontyp.HOCHLASTZEITFENSTER),
		BeschreibungTyp:            internal.Ptr("EineBeschreibung"),
	}
	serializedZaehlzeit, err := json.Marshal(zaehlzeit)
	then.AssertThat(t, serializedZaehlzeit, is.Not(is.NilArray[byte]()))
	then.AssertThat(t, err, is.Nil())
	deserializedZaehlzeit := com.Zaehlzeit{}
	err = json.Unmarshal(serializedZaehlzeit, &deserializedZaehlzeit)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, deserializedZaehlzeit, is.EqualTo(zaehlzeit))
}

func Test_Successful_Zaehlzeit_Validation(t *testing.T) {
	validate := validator.New()

	validAbweichung := []interface{}{
		com.Zaehlzeit{Code: internal.Ptr("HL1"),
			Haeufigkeit:                internal.Ptr(haeufigkeitzaehlzeit.JAEHRLICH),
			Uebermittelbarkeit:         internal.Ptr(uebermittelbarkeitzaehlzeit.ELEKTRONISCH),
			ErmittlungLeistungsmaximum: internal.Ptr(ermittlungleistungsmaximum.VERWENDUNG_HOCHLASTFENSTER),
			IstBestellbar:              internal.Ptr(true),
			Typ:                        internal.Ptr(zaehlzeitdefinitiontyp.HOCHLASTZEITFENSTER),
			BeschreibungTyp:            internal.Ptr("EineBeschreibung"),
		},
		com.Zaehlzeit{},
	}
	VerifySuccessfulValidations(t, validate, validAbweichung)
}

func Test_Serialized_Empty_Zaehlzeit_Contains_No_Enum_Defaults(t *testing.T) {
	assertDoesNotSerializeDefaultEnums(t, com.Zaehlzeit{})
}
