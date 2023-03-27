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

// The generic Marshalling is currently not possible due to a malfunction of the "json.Marshal" method.
// The method resolving the '*T' pointer of the 'Shadow' type instead of treading it as its own type.
// This is causing us to always end up in an infinite loop when calling 'json.Marshal(s)' where 's' is the 'Shadow' type.
//func MarshalWithUnmappedData[T any](t T) (bytes []byte, err error) {
//	type Shadow *T
//	s := Shadow(&t)
//
//	byteArr, err := json.Marshal(s)
//	if err != nil {
//		return
//	}
//
//	unmappedDataFieldNames, err := jsonfieldnames.Extract(UnmappedData{})
//	if err != nil {
//		return
//	} else if len(unmappedDataFieldNames) != 1 {
//		err = errors.New("expected exactly one field ('ExtensionData' in struct 'UnmappedData')")
//		return
//	}
//	unmappedDataFieldName := unmappedDataFieldNames[0]
//
//	var structFields map[string]any
//	err = json.Unmarshal(byteArr, &structFields)
//
//	for fieldName, value := range structFields {
//		if fieldName == unmappedDataFieldName {
//			unmappedDataMap := value.(map[string]any)
//			for k, v := range unmappedDataMap {
//				structFields[k] = v
//			}
//			delete(structFields, fieldName)
//		}
//	}
//
//	return json.Marshal(structFields)
//}

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

func UnmarshallWithUnmappedData[T any](targetStruct *T, unmappedDataInTargetStruct *ExtensionData, bytes []byte) (err error) {
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
