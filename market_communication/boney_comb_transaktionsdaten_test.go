package market_communication_test

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/market_communication"
	"time"
)

func (s *Suite) Test_GetTransactionData_Returns_Nil_For_Nil_Transaktionsdaten() {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(s.T(), emptyBoneyComb.Transaktionsdaten, is.Nil())
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData("foo"), is.Nil())
}

func (s *Suite) Test_GetTransactionData_Returns_Nil_For_Not_Found_Key_And_Value_Otherwise() {
	var emptyBoneyComb = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"foo": "bar",
		},
	}
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData("asd"), is.Nil())
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData(""), is.Nil())
	then.AssertThat(s.T(), emptyBoneyComb.GetTransactionData("foo"), is.Not(is.Nil()))
	then.AssertThat(s.T(), *emptyBoneyComb.GetTransactionData("foo"), is.EqualTo("bar"))
}

func (s *Suite) Test_GetNachrichtendatum_Returns_Nil_For_Nil_Transaktionsdaten() {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(s.T(), emptyBoneyComb.Transaktionsdaten, is.Nil())
	nachrichtendatum, err := emptyBoneyComb.GetNachrichtendatum()
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), nachrichtendatum, is.Nil())
}

func (s *Suite) Test_GetNachrichtendatum_Returns_Error_For_Unparsable_Date() {
	var boneyCombWithMalformedDate = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"nachrichtendatum": "adasdasd",
		},
	}
	nachrichtendatum, err := boneyCombWithMalformedDate.GetNachrichtendatum()
	then.AssertThat(s.T(), nachrichtendatum, is.Nil())
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) Test_GetNachrichtendatum_Returns_Correct_Value() {
	var boneyCombWithMalformedDate = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"nachrichtendatum": "2021-10-14T15:35:00Z",
		},
	}
	nachrichtendatum, err := boneyCombWithMalformedDate.GetNachrichtendatum()
	then.AssertThat(s.T(), *nachrichtendatum, is.EqualTo(time.Date(2021, 10, 14, 15, 35, 0, 0, time.UTC)))
	then.AssertThat(s.T(), err, is.Nil())
}

func (s *Suite) Test_SetNachrichtendatum() {
	boneyComb := market_communication.BOneyComb{}
	arbitraryDate := time.Date(2021, 10, 24, 16, 32, 0, 0, time.UTC)
	boneyComb.SetNachrichtendatum(arbitraryDate)
	setDate, err := boneyComb.GetNachrichtendatum()
	then.AssertThat(s.T(), err, is.Nil())
	then.AssertThat(s.T(), *setDate, is.EqualTo(arbitraryDate))
}

func (s *Suite) Test_GetDokumentennummer_Returns_Correct_Value() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"dokumentennummer": "asdasdasd",
		},
	}
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetDokumentennummer(), is.EqualTo("asdasdasd"))
}

func (s *Suite) Test_SetDokumentennummer() {
	boneyComb := market_communication.BOneyComb{}
	boneyComb.SetDokumentennummer("1234567ASDFGH")
	then.AssertThat(s.T(), *boneyComb.GetDokumentennummer(), is.EqualTo("1234567ASDFGH"))
}

func (s *Suite) Test_GtNachrichtenReferenznummer_Returns_Correct_Value() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"nachrichtenReferenznummer": "asdasdasd",
		},
	}
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetNachrichtenReferenznummer(), is.EqualTo("asdasdasd"))
}

func (s *Suite) Test_SetNachrichtenReferenznummer() {
	boneyComb := market_communication.BOneyComb{}
	boneyComb.SetNachrichtenReferenznummer("1234567ASDFGH")
	then.AssertThat(s.T(), *boneyComb.GetNachrichtenReferenznummer(), is.EqualTo("1234567ASDFGH"))
}

func (s *Suite) Test_GetPruefidentifikator_Returns_Correct_Value() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"pruefidentifikator": "13002",
		},
	}
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetPruefidentifikator(), is.EqualTo("13002"))
}

func (s *Suite) Test_GetDokumentennummer() {
	boneyComb := market_communication.BOneyComb{}
	boneyComb.SetPruefidentifikator("11042")
	then.AssertThat(s.T(), *boneyComb.GetPruefidentifikator(), is.EqualTo("11042"))
}

func (s *Suite) Test_GetTransaktionsdatenKeys() {
	var boneyComb = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"ZZZ":              "z",
			"nachrichtendatum": "2021-10-14T15:35:00Z",
			"Foo":              "Bar",
			"Asd":              "xyz",
		},
	}
	then.AssertThat(s.T(), boneyComb.GetTransaktionsdatenKeys(), is.EqualTo([]string{"Asd", "Foo", "ZZZ", "nachrichtendatum"}))
}

func (s *Suite) Test_GetTransaktionsdatenKeys_Works_For_Nil() {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(s.T(), emptyBoneyComb.GetTransaktionsdatenKeys(), is.EqualTo([]string{}))
}
