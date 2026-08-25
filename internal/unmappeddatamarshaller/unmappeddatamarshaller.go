package unmappeddatamarshaller

import (
	"encoding/json"
	"reflect"

	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
)

type ExtensionData map[string]any

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

	byteArr, err := json.Marshal(unmarshalledFields)
	if err != nil {
		return
	}

	// original is the reflection value we use to set the fields of targetStruct.
	original := reflect.Indirect(reflect.ValueOf(targetStruct))

	// create reflection struct fields and field name mappers from the type we want to write to.
	shadowFields, targetValueSetters := createFieldsForShadowType(original.Type())

	shadow := reflect.New(reflect.StructOf(shadowFields))
	if err = json.Unmarshal(byteArr, shadow.Interface()); err != nil {
		return
	}

	for field, value := range reflect.Indirect(shadow).Fields() {
		targetValueSetters[field.Name](value, original)
	}

	return
}

// createFieldsForShadowType takes a reflection type and returns a list of fields that represent the targetType's fields,
// but flattened. When a target type embeds a type, both the target type's and the embedded type's fields
// can be found in the list. This also works with nested embeddings, e.g. A embeds B, B embeds C.
// The second return value targetValueSetters returns a map from field names to setters. When constructing a struct type
// from the fields list via reflection and unmarshalling JSON into a value of that type, you can set the field values
// of the target type by iterating over the shadow's fields and calling the target value setter by fieldname with
// the shadow field value as the first argument and the target as the second.
func createFieldsForShadowType(targetType reflect.Type) (fields []reflect.StructField, targetValueSetters map[string]func(value reflect.Value, target reflect.Value)) {
	targetValueSetters = make(map[string]func(value reflect.Value, target reflect.Value))

	if targetType.Kind() != reflect.Struct {
		return
	}

	for field := range targetType.Fields() {
		// Unexported fields cannot be used with reflect.StructOf. The json package can't unmarshal into them anyway.
		// So we take only the exported fields.
		if field.PkgPath != "" {
			continue
		}

		// Named fields are simple. Append them to the list of fields. Setting the value is just finding the field with the same name
		// and setting its value.
		if !field.Anonymous {
			fields = append(fields, field)
			targetValueSetters[field.Name] = func(value, target reflect.Value) {
				target.FieldByName(field.Name).Set(value)
			}

			continue
		}

		// Anonymous fields sadly can't just be set directly as reflect.StructOf does not support embedded types with methods
		// (which we have). So we flatten them into one big struct.
		subFields, subMappings := createFieldsForShadowType(field.Type)
		fields = append(fields, subFields...)
		for name, subMapping := range subMappings {
			targetValueSetters[name] = func(value, target reflect.Value) {
				target = target.FieldByName(field.Name)
				subMapping(value, target)
			}
		}
	}

	return
}
