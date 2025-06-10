package bo_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
)

func TestMarshalBusinessObjectTypesExtensionData(t *testing.T) {
	boTypes := []botyp.BOTyp{
		botyp.BILANZIERUNG,
		botyp.LOKATIONSZUORDNUNG,
		botyp.MARKTLOKATION,
		botyp.MARKTTEILNEHMER,
		botyp.MESSLOKATION,
	}

	for _, boType := range boTypes {
		t.Run(
			boType.String(),
			createMarshalBusinessObjectExtensionDataTest(boType),
		)
	}
}

func createMarshalBusinessObjectExtensionDataTest(typ botyp.BOTyp) func(t *testing.T) {
	return func(t *testing.T) {
		businessObject := reflect.ValueOf(bo.NewBusinessObject(typ))
		if businessObject.Kind() != reflect.Pointer {
			t.Fatalf("expected bo.NewBusinessObject() to create pointer")
		}
		if businessObject.IsNil() {
			t.Fatalf("could not create business object for type %s", typ)
		}

		concreteValue := businessObject.Elem()
		if concreteValue.Kind() != reflect.Struct {
			t.Fatalf("expected kind, expected struct, got %s when creating BO type %s", concreteValue.Kind(), typ)
		}

		field := concreteValue.FieldByName("ExtensionData")
		if isZero(field) {
			t.Fatalf("missing embedded ExtensionData")
		}

		_, ok := field.Interface().(unmappeddatamarshaller.ExtensionData)
		if !ok {
			t.Fatalf("ExtensionData is type %T instead of unmappedmarshaller.ExtensionData", field.Interface())
		}
		if !field.CanSet() {
			t.Fatalf("ExtensionData cannot be written to")
		}
		field.Set(
			reflect.ValueOf(
				map[string]any{
					"extensionField": "value of extension field",
				},
			),
		)

		data, err := json.Marshal(businessObject.Interface())
		if err != nil {
			t.Fatalf("could not marshal business object: %+v", err)
		}

		var dest map[string]any
		if err := json.Unmarshal(data, &dest); err != nil {
			t.Fatalf("could not unmarshal back: %+v", err)
		}

		if _, ok := dest["ExtensionData"]; ok {
			t.Errorf("did not expect extension data to be marshalled as ExtensionData")
		}

		then.AssertThat(t, dest["extensionField"], is.EqualTo(any("value of extension field")))
	}
}

// isZero returns true if the given value is the zero value.
func isZero[T comparable](value T) bool {
	var zero T
	return zero == value
}
