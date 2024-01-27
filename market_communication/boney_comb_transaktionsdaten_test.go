package market_communication_test

import (
	"testing"
	"time"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

func Test_GetTransactionData_Returns_Nil_For_Nil_Transaktionsdaten(t *testing.T) {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(t, emptyBoneyComb.Transaktionsdaten, is.NilMap[string, string]())
	then.AssertThat(t, emptyBoneyComb.GetTransactionData("foo"), is.NilPtr[string]())
}

func Test_GetTransactionData_Returns_Nil_For_Not_Found_Key_And_Value_Otherwise(t *testing.T) {
	var emptyBoneyComb = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"foo": "bar",
		},
	}
	then.AssertThat(t, emptyBoneyComb.GetTransactionData("asd"), is.NilPtr[string]())
	then.AssertThat(t, emptyBoneyComb.GetTransactionData(""), is.NilPtr[string]())
	then.AssertThat(t, emptyBoneyComb.GetTransactionData("foo"), is.Not(is.NilPtr[string]()))
	then.AssertThat(t, *emptyBoneyComb.GetTransactionData("foo"), is.EqualTo("bar"))
}

func Test_GetNachrichtendatum_Returns_Nil_For_Nil_Transaktionsdaten(t *testing.T) {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(t, emptyBoneyComb.Transaktionsdaten, is.NilMap[string, string]())
	nachrichtendatum, err := emptyBoneyComb.GetNachrichtendatum()
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, nachrichtendatum, is.NilPtr[time.Time]())
}

func Test_GetNachrichtendatum_Returns_Error_For_Unparsable_Date(t *testing.T) {
	var boneyCombWithMalformedDate = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"nachrichtendatum": "adasdasd",
		},
	}
	nachrichtendatum, err := boneyCombWithMalformedDate.GetNachrichtendatum()
	then.AssertThat(t, nachrichtendatum, is.NilPtr[time.Time]())
	then.AssertThat(t, err, is.Not(is.Nil()))
}

func Test_GetNachrichtendatum_Returns_Correct_Value(t *testing.T) {
	var boneyCombWithMalformedDate = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"nachrichtendatum": "2021-10-14T15:35:00Z",
		},
	}
	nachrichtendatum, err := boneyCombWithMalformedDate.GetNachrichtendatum()
	then.AssertThat(t, *nachrichtendatum, is.EqualTo(time.Date(2021, 10, 14, 15, 35, 0, 0, time.UTC)))
	then.AssertThat(t, err, is.Nil())
}

func Test_SetNachrichtendatum(t *testing.T) {
	boneyComb := market_communication.BOneyComb{}
	arbitraryDate := time.Date(2021, 10, 24, 16, 32, 0, 0, time.UTC)
	boneyComb.SetNachrichtendatum(arbitraryDate)
	setDate, err := boneyComb.GetNachrichtendatum()
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, *setDate, is.EqualTo(arbitraryDate))
}

func Test_GetDokumentennummer_Returns_Correct_Value(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"dokumentennummer": "asdasdasd",
		},
	}
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetDokumentennummer(), is.EqualTo("asdasdasd"))
}

func Test_SetDokumentennummer(t *testing.T) {
	boneyComb := market_communication.BOneyComb{}
	boneyComb.SetDokumentennummer("1234567ASDFGH")
	then.AssertThat(t, *boneyComb.GetDokumentennummer(), is.EqualTo("1234567ASDFGH"))
}

func Test_GtNachrichtenReferenznummer_Returns_Correct_Value(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"nachrichtenReferenznummer": "asdasdasd",
		},
	}
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetNachrichtenReferenznummer(), is.EqualTo("asdasdasd"))
}

func Test_SetNachrichtenReferenznummer(t *testing.T) {
	boneyComb := market_communication.BOneyComb{}
	boneyComb.SetNachrichtenReferenznummer("1234567ASDFGH")
	then.AssertThat(t, *boneyComb.GetNachrichtenReferenznummer(), is.EqualTo("1234567ASDFGH"))
}

func Test_GetPruefidentifikator_Returns_Correct_Value(t *testing.T) {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"pruefidentifikator": "13002",
		},
	}
	then.AssertThat(t, *boneyCombWithDokumentennummer.GetPruefidentifikator(), is.EqualTo("13002"))
}

func Test_GetDokumentennummer(t *testing.T) {
	boneyComb := market_communication.BOneyComb{}
	boneyComb.SetPruefidentifikator("11042")
	then.AssertThat(t, *boneyComb.GetPruefidentifikator(), is.EqualTo("11042"))
}

func Test_GetTransaktionsdatenKeys(t *testing.T) {
	var boneyComb = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"ZZZ":              "z",
			"nachrichtendatum": "2021-10-14T15:35:00Z",
			"Foo":              "Bar",
			"Asd":              "xyz",
		},
	}
	then.AssertThat(t, boneyComb.GetTransaktionsdatenKeys(), is.EqualTo([]string{"Asd", "Foo", "ZZZ", "nachrichtendatum"}))
}

func Test_GetTransaktionsdatenKeys_Works_For_Nil(t *testing.T) {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(t, emptyBoneyComb.GetTransaktionsdatenKeys(), is.EqualTo([]string{}))
}
