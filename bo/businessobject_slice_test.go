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
	"github.com/shopspring/decimal"
	"time"
)

var SliceWithThreeValidBos = bo.BusinessObjectSlice{
	&bo.Geschaeftspartner{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.Geschaeftspartner,
			VersionStruktur: "1",
		},
		Anrede: anrede.Frau,
		Name1:  "Musterfrau",
		Name2:  "Erika",
	},
	&bo.Zaehler{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.Zaehler,
			VersionStruktur: "1",
		},
		Sparte:             sparte.Strom,
		Zaehlernummer:      "1ASD23",
		Zaehlerauspraegung: zaehlerauspraegung.Einrichtungszaehler,
		Zaehlertyp:         zaehlertyp.Drehkolbenzaehler,
		Tarifart:           tarifart.Eintarif,
	},
	&bo.Energiemenge{
		Geschaeftsobjekt: bo.Geschaeftsobjekt{
			BoTyp:           botyp.Energiemenge,
			VersionStruktur: "1",
		},
		LokationsId:  "54321012345",
		LokationsTyp: lokationstyp.MaLo,
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

func (s *Suite) Test_Deserialization_Of_Slice() {
	boList := SliceWithThreeValidBos
	jsonBytes, err := json.Marshal(boList)
	then.AssertThat(s.T(), err, is.Nil())
	var unmarshaledSlice bo.BusinessObjectSlice
	err = json.Unmarshal(jsonBytes, &unmarshaledSlice)
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), len(unmarshaledSlice), is.EqualTo(3))
	then.AssertThat(s.T(), unmarshaledSlice, is.EqualTo(boList))
}
