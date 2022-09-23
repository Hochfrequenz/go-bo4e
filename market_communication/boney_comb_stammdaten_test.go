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
	then.AssertThat(s.T(), boneyComb.ContainsAny(botyp.MESSLOKATION), is.True())
	then.AssertThat(s.T(), boneyComb.GetMasterDataCount(botyp.PREISBLATT), is.EqualTo(expectedPreisblattCount))
	then.AssertThat(s.T(), boneyComb.ContainsAny(botyp.PREISBLATT), is.False())
	then.AssertThat(s.T(), boneyComb.ContainsAny(botyp.RECHNUNG), is.True())
}

func (s *Suite) Test_GetAll() {
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
	then.AssertThat(s.T(), len(boneyComb.GetAll(botyp.PREISBLATT)), is.EqualTo(0))
	then.AssertThat(s.T(), len(boneyComb.GetAll(botyp.MARKTTEILNEHMER)), is.EqualTo(2))
	then.AssertThat(s.T(), len(boneyComb.GetAll(botyp.LASTGANG)), is.EqualTo(1))
}

func (s *Suite) Test_Generic_GetAll() {
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
	then.AssertThat(s.T(), len(market_communication.GetAll[bo.Preisblatt](&boneyComb)), is.EqualTo(0))
	then.AssertThat(s.T(), len(market_communication.GetAll[bo.Marktteilnehmer](&boneyComb)), is.EqualTo(2))
	then.AssertThat(s.T(), len(market_communication.GetAll[bo.Lastgang](&boneyComb)), is.EqualTo(1))
}

func (s *Suite) Test_GetSingle() {
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
	marktteilnehmer, marktteilnehmerErr := boneyComb.GetSingle(botyp.MARKTTEILNEHMER) // there are 2 marktteilnehmers
	then.AssertThat(s.T(), marktteilnehmer, is.Nil())
	then.AssertThat(s.T(), marktteilnehmerErr, is.Not(is.Nil()))

	preisblatt, preisblattErr := boneyComb.GetSingle(botyp.PREISBLATT) // there are 0 preisblatts
	then.AssertThat(s.T(), preisblatt, is.Nil())
	then.AssertThat(s.T(), preisblattErr, is.Not(is.Nil()))

	lastgang, lastgangErr := boneyComb.GetSingle(botyp.LASTGANG) // there is 1 lastgang
	then.AssertThat(s.T(), lastgang, is.Not(is.Nil()))
	then.AssertThat(s.T(), lastgangErr, is.Nil())
}

func (s *Suite) Test_Generic_GetSingle() {
	var untypedLastgang = bo.NewBusinessObject(botyp.LASTGANG)
	existingLastgang := untypedLastgang.(*bo.Lastgang)
	existingLastgang.LokationsId = "54321012345"
	var boneyComb = market_communication.BOneyComb{
		Stammdaten: []bo.BusinessObject{
			bo.NewBusinessObject(botyp.MARKTTEILNEHMER),
			bo.NewBusinessObject(botyp.MARKTTEILNEHMER),
			bo.NewBusinessObject(botyp.MESSLOKATION),
			bo.NewBusinessObject(botyp.ZAEHLER),
			bo.NewBusinessObject(botyp.RECHNUNG),
			bo.NewBusinessObject(botyp.MESSLOKATION),
			existingLastgang,
		},
	}
	marktteilnehmer, marktteilnehmerErr := market_communication.GetSingle[bo.Marktteilnehmer](&boneyComb) // there are 2 marktteilnehmers
	then.AssertThat(s.T(), marktteilnehmerErr, is.Not(is.Nil()))
	then.AssertThat(s.T(), marktteilnehmer, is.Nil())

	preisblatt, preisblattErr := market_communication.GetSingle[bo.Preisblatt](&boneyComb)
	then.AssertThat(s.T(), preisblattErr, is.Not(is.Nil()))
	then.AssertThat(s.T(), preisblatt, is.Nil())

	resultLastgang, lastgangErr := market_communication.GetSingle[bo.Lastgang](&boneyComb) // there is 1 lastgang
	then.AssertThat(s.T(), lastgangErr, is.Nil())
	then.AssertThat(s.T(), resultLastgang.LokationsId, is.EqualTo("54321012345"))
}
