package bo_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anrede"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/tarifart"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlerauspraegung"
	"github.com/hochfrequenz/go-bo4e/enum/zaehlertyp"
	"github.com/hochfrequenz/go-bo4e/internal"
	"github.com/hochfrequenz/go-bo4e/market_communication"
	"github.com/shopspring/decimal"
)

var SliceWithThreeValidBos = bo.BusinessObjectSlice{
	&bo.Geschaeftspartner{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.GESCHAEFTSPARTNER,
			VersionStruktur: "1",
		},
		Anrede: internal.Ptr(anrede.FRAU),
		Name1:  "Musterfrau",
		Name2:  "Erika",
	},
	&bo.Zaehler{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.ZAEHLER,
			VersionStruktur: "1",
		},
		Sparte:             sparte.STROM,
		Zaehlernummer:      "1ASD23",
		Zaehlerauspraegung: internal.Ptr(zaehlerauspraegung.EINRICHTUNGSZAEHLER),
		Zaehlertyp:         internal.Ptr(zaehlertyp.DREHKOLBENZAEHLER),
		Tarifart:           internal.Ptr(tarifart.EINTARIF),
	},
	&bo.Energiemenge{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.ENERGIEMENGE,
			VersionStruktur: "1",
		},
		LokationsId:  "54321012345",
		LokationsTyp: internal.Ptr(lokationstyp.MALO),
		Verbrauch: []com.Verbrauch{
			{
				Startdatum:               time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
				Enddatum:                 time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 15),
				Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
				Obiskennzahl:             "1-0:1.8.1",
				Wert:                     decimal.NewFromFloat(17),
				Einheit:                  mengeneinheit.KWH,
			},
		},
	},
}

func Test_Successful_Slice_Deserialization(t *testing.T) {
	boList := SliceWithThreeValidBos
	jsonBytes, err := json.Marshal(boList)
	then.AssertThat(t, err, is.Nil())
	var unmarshalledSlice bo.BusinessObjectSlice
	err = json.Unmarshal(jsonBytes, &unmarshalledSlice)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, len(unmarshalledSlice), is.EqualTo(3))
	then.AssertThat(t, unmarshalledSlice, is.EqualTo(boList))
}

func Test_Slice_Deserialization_Fails_For_Invalid_BoTyps(t *testing.T) {
	jsonWithInvalidBoTyp := `{"stammdaten": [{"boTyp":"foo"}]}` // foo is not a valid boTyp => deserialization should fail
	var deserializedBoneyComb market_communication.BOneyComb
	err := json.Unmarshal([]byte(jsonWithInvalidBoTyp), &deserializedBoneyComb)
	then.AssertThat(t, err, is.Not(is.Nil()))
}

func Test_Slice_Deserialization_Fails_For_Unimplemented_BoTyps(t *testing.T) {
	jsonWithUnimplementedBoTyp := `{"stammdaten": [{"boTyp":"PREISBLATTUMLAGEN"}], "transaktionsdaten": {"pruefidentifikator": "12345"}}` // PREISBLATTUMLAGEN is not (yet) an implemented boTyp => deserialization should fail
	var deserializedBoneyComb market_communication.BOneyComb
	err := json.Unmarshal([]byte(jsonWithUnimplementedBoTyp), &deserializedBoneyComb)
	then.AssertThat(t, err, is.Not(is.Nil()))
	then.AssertThat(t, errors.Is(err, bo.ErrorUnimplementedBusinessObject), is.True())
	then.AssertThat(t, err.Error(), is.StringContaining("PREISBLATTUMLAGEN"))
	then.AssertThat(t, len(deserializedBoneyComb.Transaktionsdaten), is.EqualTo(1))
	then.AssertThat(t, deserializedBoneyComb.GetPruefidentifikator(), is.Not(is.NilPtr[string]()))
	then.AssertThat(t, *deserializedBoneyComb.GetPruefidentifikator(), is.EqualTo("12345"))
}

func TestBusinessObjectSliceUnmarshalErrors(t *testing.T) {
	rawJSON := `
	[
		{
			"boTyp": "foo"
		},
		{
			"boTyp": "PREISBLATT",
			"Herausgeber": true,
			"Gueltigkeit": "always"
		},
		{
			"boTyp": 5000
		}
	]
	`

	var boSlice bo.BusinessObjectSlice

	err := json.Unmarshal([]byte(rawJSON), &boSlice)

	unwrapper, ok := err.(interface{ Unwrap() []error })
	if !ok {
		t.Errorf("expected error to provide wrapped errors, got %+v (%T)", err, err)
	}

	errs := unwrapper.Unwrap()
	if len(errs) != 3 {
		t.Errorf("expected 3 errors, got %+v (%d errors)", errs, len(errs))
	}

	// We do not check the types of errors as there's no reliable error API yet.
}
