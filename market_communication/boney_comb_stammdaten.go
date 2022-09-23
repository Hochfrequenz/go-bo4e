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

// GetAll is a generic version of BOneyComb.GetAll. It returns all business objects of a given type in the boneycomb
func GetAll[T bo.BusinessObject](boneyComb *BOneyComb) []*T {
	var result []*T
	var emptyInstance = new(T)
	expectedType := (*emptyInstance).GetBoTyp()
	if boneyComb.Stammdaten != nil {
		for _, businessObject := range boneyComb.Stammdaten {
			if HasType(businessObject, expectedType) {
				// typ
				var x *T
				x = businessObject
				result = append(result, x)
			}
		}
	}
	return result
}

func HasType(businessObject bo.BusinessObject, typ botyp.BOTyp) bool {
	actualType := businessObject.GetBoTyp()
	return actualType == typ
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

// GetSingle is a generic function that writes the single present business object of the specified type into target if it is present in the given boneyComb.
// It returns an error if there is none or more than one.
// This is a generic version of BOneyComb.GetSingle but methods are not allowed to have type parameters.
func GetSingle[T bo.BusinessObject](boneyComb *BOneyComb) (*T, error) {
	allBOsOfType := GetAll[T](boneyComb)
	if len(allBOsOfType) == 0 {
		return nil, fmt.Errorf("There is no business object with the given type")
	}
	if len(allBOsOfType) > 1 {
		return nil, fmt.Errorf("There is more than one (%d) business object with the given type", len(allBOsOfType))
	}
	singleResult := allBOsOfType[0]
	return singleResult, nil
}
