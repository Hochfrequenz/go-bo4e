package unmappeddatamarshaller

import (
	"encoding/json"
	"fmt"
	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
	"reflect"
)

type ExtensionData map[string]any

func (ed ExtensionData) CompareTo(otherEd ExtensionData) bool {
	mapAContent := fmt.Sprint(ed)
	mapBContent := fmt.Sprint(otherEd)
	return mapAContent == mapBContent
}

func (ed ExtensionData) storeUnmappedData(key string, value any) {
	ed[key] = value
}

// HandleUnmappedDataPropertyMarshalling expects the bytes of a marshalled struct. If the marshalled struct contains
// 'unmapped' fields meaning ones that had no corresponding, strong-typed field when initially unmarshalled, those fields
// will be extracted again. The extracted field from the maps containing the unmapped data will be placed on the top level
// of the marshalled struct.
func HandleUnmappedDataPropertyMarshalling(b []byte) (bytes []byte, err error) {
	unmappedDataFieldName := reflect.TypeOf(ExtensionData{}).Name()

	var structFields map[string]any
	err = json.Unmarshal(b, &structFields)
	if err != nil {
		return
	}

	for fieldName, value := range structFields {
		if fieldName == unmappedDataFieldName {
			if value != nil {
				unmappedDataMap := value.(map[string]any)
				for k, v := range unmappedDataMap {
					structFields[k] = v
				}
				delete(structFields, fieldName)
			}
		}
	}

	return json.Marshal(structFields)
}

// UnmarshallWithUnmappedData will unmarshal a given type by mapping all strong-typed fields to the 'targetStruct'. All
// other fields will be preserved in the 'unmappedDataInTargetStruct' dictionary.
func UnmarshallWithUnmappedData[T any](targetStruct *T, unmappedDataInTargetStruct *ExtensionData, bytes []byte) (err error) {
	if *unmappedDataInTargetStruct == nil {
		*unmappedDataInTargetStruct = ExtensionData{}
	}

	var unmarshalledFields map[string]any
	err = json.Unmarshal(bytes, &unmarshalledFields)
	if err != nil {
		return
	}

	targetFieldNames, err := jsonfieldnames.Extract(targetStruct)
	if err != nil {
		return
	}

	for fieldName, value := range unmarshalledFields {
		isMappedField := false

		for _, targetFieldName := range targetFieldNames {
			if targetFieldName == fieldName {
				isMappedField = true
				break
			}
		}

		if !isMappedField {
			unmappedDataInTargetStruct.storeUnmappedData(fieldName, value)
			delete(unmarshalledFields, fieldName)
		}
	}

	s := T(any(*targetStruct))

	byteArr, err := json.Marshal(unmarshalledFields)
	if err != nil {
		return
	}
	err = json.Unmarshal(byteArr, s)
	if err != nil {
		return
	}

	return
}
