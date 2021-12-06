package market_communication_test

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

func (s *Suite) Test_GetMasterDataCount_For_Empty_BoneyComb() {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(s.T(), emptyBoneyComb.GetMasterDataCounts(), is.EqualTo(map[botyp.BOTyp]uint{}))
}

func (s *Suite) Test_GetMasterDataCount_For_Nil_Stammdaten() {
	var emptyBoneyComb = market_communication.BOneyComb{}
	emptyBoneyComb.Stammdaten = nil
	then.AssertThat(s.T(), emptyBoneyComb.GetMasterDataCounts(), is.EqualTo(map[botyp.BOTyp]uint{}))
}

func (s *Suite) Test_GetMasterDataCount() {
	var boneyComb = market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{
			bo.NewBusinessObject(botyp.MARKTTEILNEHMER),
			bo.NewBusinessObject(botyp.MARKTTEILNEHMER),
			bo.NewBusinessObject(botyp.MESSLOKATION),
			bo.NewBusinessObject(botyp.ZAEHLER),
			bo.NewBusinessObject(botyp.RECHNUNG),
			bo.NewBusinessObject(botyp.MESSLOKATION),
			bo.NewBusinessObject(botyp.LASTGANG),
		},
	}
	expectedResult := map[botyp.BOTyp]uint{
		botyp.MARKTTEILNEHMER: 2,
		botyp.MESSLOKATION:    2,
		botyp.ZAEHLER:         1,
		botyp.RECHNUNG:        1,
		botyp.LASTGANG:        1,
	}
	then.AssertThat(s.T(), boneyComb.GetMasterDataCounts(), is.EqualTo(expectedResult))
	var expectedMarktteilnehmerCount uint = 2
	var expectedPreisblattCount uint = 0
	then.AssertThat(s.T(), boneyComb.GetMasterDataCount(botyp.MARKTTEILNEHMER), is.EqualTo(expectedMarktteilnehmerCount))
	then.AssertThat(s.T(), boneyComb.GetMasterDataCount(botyp.PREISBLATT), is.EqualTo(expectedPreisblattCount))
}
