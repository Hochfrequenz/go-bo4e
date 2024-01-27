package com_test

import (
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/wertermittlungsverfahren"
	"github.com/shopspring/decimal"
)

// TestZaehlerstaneLen tests the Zaehlerstaende.Len function
func Test_Zaehlerstande_Len(t *testing.T) {
	zaehlerstaende3 := com.Zaehlerstaende{
		com.Zaehlerstand{},
		com.Zaehlerstand{},
		com.Zaehlerstand{},
	}
	zaehlerstaende0 := com.Zaehlerstaende{}
	then.AssertThat(t, zaehlerstaende3.Len(), is.EqualTo(3))
	then.AssertThat(t, zaehlerstaende0.Len(), is.EqualTo(0))
	zaehlerstaende0 = nil
	then.AssertThat(t, zaehlerstaende0.Len(), is.EqualTo(0))
}

// Test_Zaehlerstande_Swap tests the Zaehlerstaende.Swap function
func Test_Zaehlerstande_Swap(t *testing.T) {
	zaehlerstaende := com.Zaehlerstaende{
		com.Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(17.0),
			Einheit:                  mengeneinheit.KW,
		},
		com.Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(42.1),
			Einheit:                  mengeneinheit.KW,
		},
		com.Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(123.2),
			Einheit:                  mengeneinheit.KW,
		},
	}
	then.AssertThat(t, zaehlerstaende[1].Wert, is.EqualTo(decimal.NewFromFloat(42.1)))
	then.AssertThat(t, zaehlerstaende[2].Wert, is.EqualTo(decimal.NewFromFloat(123.2)))
	zaehlerstaende.Swap(1, 2)
	then.AssertThat(t, zaehlerstaende[1].Wert, is.EqualTo(decimal.NewFromFloat(123.2)))
	then.AssertThat(t, zaehlerstaende[2].Wert, is.EqualTo(decimal.NewFromFloat(42.1)))
}

// Test_Zaehlerstaende_Less tests the Zaehlerstaende.Less function
func Test_Zaehlerstaende_Less(t *testing.T) {
	zaehlerstaende := com.Zaehlerstaende{
		com.Zaehlerstand{
			Ablesedatum:              time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(17),
			Einheit:                  mengeneinheit.KW,
		},
		com.Zaehlerstand{
			Ablesedatum:              time.Date(2045, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(123),
			Einheit:                  mengeneinheit.KW,
		},
		com.Zaehlerstand{
			Ablesedatum:              time.Date(2038, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(42),
			Einheit:                  mengeneinheit.KW,
		},
		com.Zaehlerstand{
			Ablesedatum:              time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC),
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(18),
			Einheit:                  mengeneinheit.KW,
		},
		com.Zaehlerstand{
			Ablesedatum:              time.Time{},
			Wertermittlungsverfahren: wertermittlungsverfahren.MESSUNG,
			Wert:                     decimal.NewFromFloat(18),
			Einheit:                  mengeneinheit.KW,
		},
	}
	then.AssertThat(t, zaehlerstaende.Less(1, 2), is.True())
	then.AssertThat(t, zaehlerstaende.Less(2, 1), is.False())
	then.AssertThat(t, zaehlerstaende.Less(3, 0), is.False())
	then.AssertThat(t, zaehlerstaende.Less(0, 3), is.False())
	then.AssertThat(t, zaehlerstaende.Less(0, 4), is.True())
	then.AssertThat(t, zaehlerstaende.Less(4, 0), is.False())
}
