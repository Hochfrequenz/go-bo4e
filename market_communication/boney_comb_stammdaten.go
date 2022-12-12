package market_communication

import (
	"fmt"

	"github.com/hochfrequenz/go-bo4e/bo"
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

// ContainsAny returns true if any business object of the given type is present in BOneyComb.Stammdaten
func (boneyComb *BOneyComb) ContainsAny(typ botyp.BOTyp) bool {
	return boneyComb.GetMasterDataCount(typ) > 0
}

// GetAll returns all business objects from BOneyComb.Stammdaten that have the given type
func (boneyComb *BOneyComb) GetAll(typ botyp.BOTyp) bo.BusinessObjectSlice {
	result := bo.BusinessObjectSlice{}
	if boneyComb.Stammdaten != nil {
		for _, businessObject := range boneyComb.Stammdaten {
			if businessObject.GetBoTyp() == typ {
				result = append(result, businessObject)
			}
		}
	}
	return result
}

// GetSingle returns the single present business object of the given type if it is present in BOneyComb.Stammdaten. It returns an error if there is none or more than one
func (boneyComb *BOneyComb) GetSingle(typ botyp.BOTyp) (bo.BusinessObject, error) {
	allBOsOfType := boneyComb.GetAll(typ)
	if len(allBOsOfType) == 0 {
		return nil, fmt.Errorf("There is no business object with type {%v}", typ)
	}
	if len(allBOsOfType) > 1 {
		return nil, fmt.Errorf("There is more than one (%d) business object with type {%v}", len(allBOsOfType), typ)
	}
	return allBOsOfType[0], nil
}

// GetAll filters a BOneyComb's Stammdaten for business objects of a specific type, given as type parameter T. T must be a non-pointer type,
// even if bo.BusinessObject is implemented only via pointer receivers (*T).
func GetAll[T any, Ptr interface {
	bo.BusinessObject
	*T
}](boneyComb *BOneyComb) []*T {
	result := make([]*T, 0)

	for _, businessObject := range boneyComb.Stammdaten {
		if convertedBusinessObject, ok := businessObject.(T); ok {
			result = append(result, &convertedBusinessObject)
		}

		if convertedBusinessObject, ok := businessObject.(Ptr); ok {
			result = append(result, (*T)(convertedBusinessObject))
		}
	}

	return result
}

// GetSingle searches BOneyComb's Stammdaten for the single business object of the specified type, given as type parameter T.
// T must be non-pointer type, even if bo.BusinessObject is implemented only via pointer receivers (*T). Returns an error
// if there is either no business object of that type or there is more than one.
func GetSingle[T any, Ptr interface {
	bo.BusinessObject
	*T
}](boneyComb *BOneyComb) (*T, error) {
	all := GetAll[T, Ptr](boneyComb)

	if len(all) == 0 {
		return nil, fmt.Errorf("there is no business object with the given type")
	}

	if len(all) > 1 {
		return nil, fmt.Errorf("there is more than one (%d) business object with the given type", len(all))
	}

	return all[0], nil
}
