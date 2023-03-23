package unmappeddatamarshaller

import (
	"encoding/json"
	"errors"
	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
)

type UnmappedData struct {
	Data map[string]any
}

func (umd *UnmappedData) storeUnmappedData(key string, value any) {
	umd.Data[key] = value
}

func MarshalWithUnmappedData[T any, TShadow interface{ *T }](s T) (bytes []byte, err error) {
	byteArr, err := json.Marshal(s)
	if err != nil {
		return
	}

	unmappedDataFieldNames, err := jsonfieldnames.Extract(UnmappedData{})
	if err != nil {
		return
	} else if len(unmappedDataFieldNames) != 1 {
		err = errors.New("expected exactly one field ('Data' in struct 'UnmappedData')")
		return
	}
	unmappedDataFieldName := unmappedDataFieldNames[0]

	var structFields map[string]any
	err = json.Unmarshal(byteArr, &structFields)

	for fieldName, value := range structFields {
		if fieldName == unmappedDataFieldName {
			unmappedDataMap := value.(map[string]any)
			for k, v := range unmappedDataMap {
				structFields[k] = v
			}
			delete(structFields, fieldName)
		}
	}

	return json.Marshal(structFields)
}

func UnmarshallWithUnmappedData[T any, TPtr interface{ *T }](targetStruct *T, shadow TPtr, unmappedDataInTargetStruct *UnmappedData, bytes []byte) (err error) {
	var unmarshalledFields map[string]any
	err = json.Unmarshal(bytes, &unmarshalledFields)
	if err != nil {
		return
	}

	targetFieldNames, err := jsonfieldnames.Extract(targetStruct)
	if err != nil {
		return
	}

	unmarshalledFieldNames := getKeysFromMap(unmarshalledFields)
	for _, unmarshalledFieldName := range unmarshalledFieldNames {
		isMappedField := false

		for _, targetFieldName := range targetFieldNames {
			if targetFieldName == unmarshalledFieldName {
				isMappedField = true
				break
			}
		}

		if !isMappedField {
			unmappedDataInTargetStruct.storeUnmappedData(unmarshalledFieldName, unmarshalledFields[unmarshalledFieldName])
			delete(unmarshalledFields, unmarshalledFieldName)
		}
	}

	type Shadow *T
	s := Shadow(shadow)

	byteArr, err := json.Marshal(unmarshalledFields)
	err = json.Unmarshal(byteArr, s)
	if err != nil {
		return
	}

	targetStruct = s

	return
}

func getKeysFromMap[T comparable](m map[T]any) (keys []T) {
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
