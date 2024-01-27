package market_communication_test

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

func Test_GetMasterDataCount_For_Empty_BoneyComb(t *testing.T) {
	var emptyBoneyComb = market_communication.BOneyComb{}
	then.AssertThat(t, emptyBoneyComb.GetMasterDataCounts(), is.EqualTo(map[botyp.BOTyp]uint{}))
}

func Test_GetMasterDataCount_For_Nil_Stammdaten(t *testing.T) {
	var emptyBoneyComb = market_communication.BOneyComb{}
	emptyBoneyComb.Stammdaten = nil
	then.AssertThat(t, emptyBoneyComb.GetMasterDataCounts(), is.EqualTo(map[botyp.BOTyp]uint{}))
}

func Test_GetMasterDataCount(t *testing.T) {
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
	then.AssertThat(t, boneyComb.GetMasterDataCounts(), is.EqualTo(expectedResult))
	var expectedMarktteilnehmerCount uint = 2
	var expectedPreisblattCount uint = 0
	then.AssertThat(t, boneyComb.GetMasterDataCount(botyp.MARKTTEILNEHMER), is.EqualTo(expectedMarktteilnehmerCount))
	then.AssertThat(t, boneyComb.ContainsAny(botyp.MESSLOKATION), is.True())
	then.AssertThat(t, boneyComb.GetMasterDataCount(botyp.PREISBLATT), is.EqualTo(expectedPreisblattCount))
	then.AssertThat(t, boneyComb.ContainsAny(botyp.PREISBLATT), is.False())
	then.AssertThat(t, boneyComb.ContainsAny(botyp.RECHNUNG), is.True())
}

func Test_GetAll(t *testing.T) {
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
	then.AssertThat(t, len(boneyComb.GetAll(botyp.PREISBLATT)), is.EqualTo(0))
	then.AssertThat(t, len(boneyComb.GetAll(botyp.MARKTTEILNEHMER)), is.EqualTo(2))
	then.AssertThat(t, len(boneyComb.GetAll(botyp.LASTGANG)), is.EqualTo(1))
}

func Test_GetSingle(t *testing.T) {
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
	then.AssertThat(t, marktteilnehmer, is.EqualTo[bo.BusinessObject](nil))
	then.AssertThat(t, marktteilnehmerErr, is.Not(is.Nil()))

	preisblatt, preisblattErr := boneyComb.GetSingle(botyp.PREISBLATT) // there are 0 preisblatts
	then.AssertThat(t, preisblatt, is.EqualTo[bo.BusinessObject](nil))
	then.AssertThat(t, preisblattErr, is.Not(is.Nil()))

	lastgang, lastgangErr := boneyComb.GetSingle(botyp.LASTGANG) // there is 1 lastgang
	then.AssertThat(t, lastgang, is.Not(is.EqualTo[bo.BusinessObject](nil)))
	then.AssertThat(t, lastgangErr, is.Nil())
}
