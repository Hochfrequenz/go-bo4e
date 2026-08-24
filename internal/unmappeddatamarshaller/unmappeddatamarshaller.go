package unmappeddatamarshaller

import (
	"encoding/json"
	"fmt"

	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
)

type ExtensionData map[string]any

func (ed ExtensionData) CompareTo(otherEd ExtensionData) bool {
	mapAContent := fmt.Sprint(ed)
	mapBContent := fmt.Sprint(otherEd)
	return mapAContent == mapBContent
}

// HandleUnmappedDataPropertyMarshalling expects the bytes of a marshalled struct. If the marshalled struct contains
// 'unmapped' fields meaning ones that had no corresponding, strong-typed field when initially unmarshalled, those fields
// will be extracted again. The extracted field from the maps containing the unmapped data will be placed on the top level
// of the marshalled struct.
func HandleUnmappedDataPropertyMarshalling(b []byte) (bytes []byte, err error) {
	var structFields map[string]any
	err = json.Unmarshal(b, &structFields)
	if err != nil {
		return
	}

	if unmappedDataMap, ok := structFields[unmappedDataFieldName].(map[string]any); ok {
		for k, v := range unmappedDataMap {
			structFields[k] = v
		}
	}

	delete(structFields, unmappedDataFieldName)

	return json.Marshal(structFields)
}

const unmappedDataFieldName = "ExtensionData"

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
			(*unmappedDataInTargetStruct)[fieldName] = value
			delete(unmarshalledFields, fieldName)
		}
	}

	type Shadow *T
	s := Shadow(targetStruct)

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
