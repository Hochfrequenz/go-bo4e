package com

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"time"
)

// TestZaehlerstaneLen tests the Zaehlerstaende.Len function
func (s *Suite) TestZaehlerstandeLen() {
	zaehlerstaende3 := Zaehlerstaende{
		Zaehlerstand{},
		Zaehlerstand{},
		Zaehlerstand{},
	}
	zaehlerstaende0 := Zaehlerstaende{}
	then.AssertThat(s.T(), zaehlerstaende3.Len(), is.EqualTo(3))
	then.AssertThat(s.T(), zaehlerstaende0.Len(), is.EqualTo(0))
	zaehlerstaende0 = nil
	then.AssertThat(s.T(), zaehlerstaende0.Len(), is.EqualTo(0))
}

//  Test_Zaehlerstande_Swap tests the Zaehlerstaende.Swap function
func (s *Suite) Test_Zaehlerstande_Swap() {
	zaehlerstaende := Zaehlerstaende{
		Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     17.0,
			Einheit:                  mengeneinheit.KW,
		},
		Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     42.1,
			Einheit:                  mengeneinheit.KW,
		},
		Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     123.2,
			Einheit:                  mengeneinheit.KW,
		},
	}
	then.AssertThat(s.T(), zaehlerstaende[1].Wert, is.EqualTo(float32(42.1)))
	then.AssertThat(s.T(), zaehlerstaende[2].Wert, is.EqualTo(float32(123.2)))
	zaehlerstaende.Swap(1, 2)
	then.AssertThat(s.T(), zaehlerstaende[1].Wert, is.EqualTo(float32(123.2)))
	then.AssertThat(s.T(), zaehlerstaende[2].Wert, is.EqualTo(float32(42.1)))
}

//  Test_Zaehlerstaende_Less tests the Zaehlerstaende.Less function
func (s *Suite) Test_Zaehlerstaende_Less() {
	zaehlerstaende := Zaehlerstaende{
		Zaehlerstand{
			Ablesedatum:              time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     17,
			Einheit:                  mengeneinheit.KW,
		},
		Zaehlerstand{
			Ablesedatum:              time.Date(2045, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     123,
			Einheit:                  mengeneinheit.KW,
		},
		Zaehlerstand{
			Ablesedatum:              time.Date(2038, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     42,
			Einheit:                  mengeneinheit.KW,
		},
		Zaehlerstand{
			Ablesedatum:              time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     18,
			Einheit:                  mengeneinheit.KW,
		},
		Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     18,
			Einheit:                  mengeneinheit.KW,
		},
	}
	then.AssertThat(s.T(), zaehlerstaende.Less(1, 2), is.True())
	then.AssertThat(s.T(), zaehlerstaende.Less(2, 1), is.False())
	then.AssertThat(s.T(), zaehlerstaende.Less(3, 0), is.False())
	then.AssertThat(s.T(), zaehlerstaende.Less(0, 3), is.False())
	then.AssertThat(s.T(), zaehlerstaende.Less(0, 4), is.True())
	then.AssertThat(s.T(), zaehlerstaende.Less(4, 0), is.False())
}
