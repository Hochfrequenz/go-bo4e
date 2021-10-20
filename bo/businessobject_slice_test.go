package bo_test

import (
	"encoding/json"
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
	"github.com/hochfrequenz/go-bo4e/market_communication"
	"github.com/shopspring/decimal"
	"time"
)

var SliceWithThreeValidBos = bo.BusinessObjectSlice{
	&bo.Geschaeftspartner{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.GESCHAEFTSPARTNER,
			VersionStruktur: "1",
		},
		Anrede: anrede.FRAU,
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
		Zaehlerauspraegung: zaehlerauspraegung.EINRICHTUNGSZAEHLER,
		Zaehlertyp:         zaehlertyp.DREHKOLBENZAEHLER,
		Tarifart:           tarifart.EINTARIF,
	},
	&bo.Energiemenge{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.ENERGIEMENGE,
			VersionStruktur: "1",
		},
		LokationsId:  "54321012345",
		LokationsTyp: lokationstyp.MALO,
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

func (s *Suite) Test_Successful_Slice_Deserialization() {
	boList := SliceWithThreeValidBos
	jsonBytes, err := json.Marshal(boList)
	then.AssertThat(s.T(), err, is.Nil())
	var unmarshalledSlice bo.BusinessObjectSlice
	err = json.Unmarshal(jsonBytes, &unmarshalledSlice)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), len(unmarshalledSlice), is.EqualTo(3))
	then.AssertThat(s.T(), unmarshalledSlice, is.EqualTo(boList))
}

func (s *Suite) Test_Slice_Deserialization_Fails_For_Invalid_BoTyps() {
	jsonWithInvalidBoTyp := "[{\"boTyp\":\"foo\"}]" // foo is not a valid boTyp => deserialization should fail
	var deserializedBoneyComb market_communication.BOneyComb
	err := json.Unmarshal([]byte(jsonWithInvalidBoTyp), &deserializedBoneyComb)
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) Test_Slice_Deserialization_Fails_For_Unimplemented_BoTyps() {
	jsonWithUnimplementedBoTyp := "[{\"boTyp\":\"PREISBLATTUMLAGEN\"}]" // PREISBLATTUMLAGEN is not (yet) an implemented boTyp => deserialization should fail
	var deserializedBoneyComb market_communication.BOneyComb
	err := json.Unmarshal([]byte(jsonWithUnimplementedBoTyp), &deserializedBoneyComb)
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}
