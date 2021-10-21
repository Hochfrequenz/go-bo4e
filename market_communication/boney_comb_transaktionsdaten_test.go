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
			"Nachrichtendatum": "adasdasd",
		},
	}
	nachrichtendatum, err := boneyCombWithMalformedDate.GetNachrichtendatum()
	then.AssertThat(s.T(), nachrichtendatum, is.Nil())
	then.AssertThat(s.T(), err, is.Not(is.Nil()))
}

func (s *Suite) Test_GetNachrichtendatum_Returns_Correct_Value() {
	var boneyCombWithMalformedDate = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"Nachrichtendatum": "2021-10-14T15:35:00Z",
		},
	}
	nachrichtendatum, err := boneyCombWithMalformedDate.GetNachrichtendatum()
	then.AssertThat(s.T(), *nachrichtendatum, is.EqualTo(time.Date(2021, 10, 14, 15, 35, 0, 0, time.UTC)))
	then.AssertThat(s.T(), err, is.Nil())
}

func (s *Suite) Test_GetDokumentennummer_Returns_Correct_Value() {
	var boneyCombWithDokumentennummer = market_communication.BOneyComb{
		Transaktionsdaten: map[string]string{
			"Dokumentennummer": "asdasdasd",
		},
	}
	then.AssertThat(s.T(), *boneyCombWithDokumentennummer.GetDokumentennummer(), is.EqualTo("asdasdasd"))
}
