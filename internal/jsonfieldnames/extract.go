// Package jsonfieldnames contains a function to extract JSON field names from a value.
//
// In addition to their fixed, hardcoded members, business objects for energy marshal additional values into Geschaeftsobjekt.ExtensionData.
// As Go's standard library JSON package has no option for a kitchen sink field accepting all remaining JSON fields,
// there are two passes of unmarshaling: Once into the struct, once into a generic map. To avoid duplicates in the struct and the map,
// fields that can be found in the struct are removed from the map. This package provides a mechanism to extract the names of
// these fields in a generic way.
package jsonfieldnames

import (
	"fmt"
	"reflect"
	"strings"
)

// Extract extracts JSON field names from a struct value or a pointer to a struct value. If the value passed to Extract is either nil
// or is not a struct value or a pointer to a struct value, an error is returned instead.
func Extract(value interface{}) ([]string, error) {
	typ := reflect.TypeOf(value)
	if typ == nil {
		return nil, fmt.Errorf("cannot extract field names: value is nil")
	}

	// When value is a pointer, dereference the type, but keep the original type to be used in the error message
	// in case the dereferenced type is also invalid.
	originalTyp := typ

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("cannot extract field names: value must either be struct nor pointer to a struct, but is %s", originalTyp.String())
	}

	return extractFromStruct(typ), nil
}

// extractFromStruct extracts the JSON field names from a struct type. Panics if typ is not a struct type.
func extractFromStruct(typ reflect.Type) []string {
	fieldNames := make([]string, 0)
	fieldCount := typ.NumField()

	// Cycle through all the fields and merge the JSON field names extracted from every field.
	for i := 0; i < fieldCount; i++ {
		field := typ.Field(i)
		fieldNames = append(fieldNames, extractFromField(field)...)
	}

	return fieldNames
}

// extractFromField extracts the JSON field names from a struct field. A single field can provide from zero to many JSON field names,
// because embedded structs are also fields and may contain an arbitrary count of fields themselves.
func extractFromField(field reflect.StructField) []string {
	jsonTag := field.Tag.Get("json")

	// If the json tag equals "-", it should not be included at all in the payload.
	if jsonTag == "-" {
		return fieldNames()
	}

	jsonFieldName := jsonFieldNameFromTag(jsonTag)

	switch {

	// field is explicitly given a JSON field name via struct tag.
	case jsonFieldName != "":
		return fieldNames(jsonFieldName)

	// field is an embedded struct without explicit JSON field name. That means its fields are merged into the embedding struct.
	case jsonFieldName == "" && field.Anonymous && field.Type.Kind() == reflect.Struct:
		return extractFromStruct(field.Type)

	// field is an embedded struct pointer without explicit JSON field name. That means its fields are merged into the embedding struct.
	case jsonFieldName == "" && field.Anonymous && field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct:
		return extractFromStruct(field.Type.Elem())

	// field has no explicit JSON field name, so use the field name instead. Both non-embedded fields and embedded fields that are neither
	// structs nor pointers to structs are handled this way.
	default:
		return fieldNames(field.Name)
	}
}

// jsonFieldNameFromTag extracts the desired json field name from a "json" struct tag. The form of these struct tags is "[<name>][,<options>]".
// This function returns the name part if it is present, else an empty string.
func jsonFieldNameFromTag(tag string) string {
	commaPosition := strings.Index(tag, ",")
	if commaPosition == -1 {
		return tag
	}
	return tag[:commaPosition]
}

func fieldNames(names ...string) []string {
	return names
}
