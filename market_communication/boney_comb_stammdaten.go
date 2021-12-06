package market_communication

import (
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

// GetMasterDataCounts counts the different object types in BOneyComb.Stammdaten
func (boneyComb *BOneyComb) GetMasterDataCounts() map[botyp.BOTyp]uint {
	result := map[botyp.BOTyp]uint{}
	if boneyComb.Stammdaten != nil {
		for _, businessObject := range boneyComb.Stammdaten {
			typ := businessObject.GetBoTyp()
			result[typ] = result[typ] + 1
		}
	}
	return result
}

// GetMasterDataCount counts the objects of the given type in BOneyComb.Stammdaten
func (boneyComb *BOneyComb) GetMasterDataCount(typ botyp.BOTyp) uint {
	result := boneyComb.GetMasterDataCounts()
	return result[typ]
}
