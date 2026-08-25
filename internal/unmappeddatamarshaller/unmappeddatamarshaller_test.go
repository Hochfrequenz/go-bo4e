package unmappeddatamarshaller

import (
	"encoding/json"
	"reflect"
	"testing"
)

type SomeStruct struct {
	ExtensionData
	A string
	B int
}

type someStructShadow SomeStruct

func (s SomeStruct) MarshalJSON() (b []byte, e error) {
	shadow := someStructShadow(s)
	b, e = json.Marshal(shadow)
	if e != nil {
		return
	}

	return HandleUnmappedDataPropertyMarshalling(b)
}

func (s *SomeStruct) UnmarshalJSON(bytes []byte) (err error) {
	return UnmarshallWithUnmappedData(s, &s.ExtensionData, bytes)
}

func Test_Unmarshalling_WithUnmappedData_PreservesUnmappedDataInStruct(t *testing.T) {
	expectedUnmappedData := map[string]any{
		"X": "very nice",
	}
	expectedStrongTypedFields := SomeStruct{
		A:             "nice",
		B:             911,
		ExtensionData: expectedUnmappedData,
	}

	someStructWithUnmappedData := map[string]any{
		"A": expectedStrongTypedFields.A,
		"B": expectedStrongTypedFields.B,
		"X": expectedUnmappedData["X"],
	}

	bytes, _ := json.Marshal(someStructWithUnmappedData)
	actualStrongTypedFields := SomeStruct{}
	err := json.Unmarshal(bytes, &actualStrongTypedFields)
	if err != nil {
		t.Errorf("Error occured while unmarshalling: %v", err)
	}

	if !reflect.DeepEqual(actualStrongTypedFields, expectedStrongTypedFields) {
		t.Errorf("Unmarshalling struct with unmapped data failed:\nexpected: %v,\nactual: %v", expectedStrongTypedFields, actualStrongTypedFields)
	}
}

func Test_Marshalling_WitUnmappedData_PreservesUnmappedDataInJson(t *testing.T) {
	unmappedData := map[string]any{
		"X": "very nice",
	}
	structWithUnmappedData := SomeStruct{A: "nice", B: 911, ExtensionData: unmappedData}

	actual, err := json.Marshal(structWithUnmappedData)
	if err != nil {
		t.Errorf("Error occured while marshalling: %v", err)
	}

	expectedJson := `{"A":"nice","B":911,"X":"very nice"}`
	actualJson := string(actual)
	if expectedJson != actualJson {
		t.Errorf("Marshalling struct with unmapped data failed:\nexpected: %s,\nactual: %s", expectedJson, actualJson)
	}
}

func Test_Marshalling_WithNilUnmappedData_OmitsExtensionDataFromJson(t *testing.T) {
	structWithoutUnmappedData := SomeStruct{A: "nice", B: 911}

	actual, err := json.Marshal(structWithoutUnmappedData)
	if err != nil {
		t.Errorf("Error occurred while marshalling: %v", err)
	}

	expectedJson := `{"A":"nice","B":911}`
	actualJson := string(actual)
	if expectedJson != actualJson {
		t.Errorf("Marshalling struct with nil unmapped data failed:\nexpected: %s,\nactual: %s", expectedJson, actualJson)
	}
}

func TestUnmarshalWithEmbeddedType(t *testing.T) {
	type Embedded struct {
		ExtensionData
		A string `json:"not_a"`
	}

	type Embedder struct {
		Embedded
		B string `json:"not_b"`
	}

	raw := `{ "not_a": "foo", "not_b": "bar", "extra": "baz" }`

	embedder := Embedder{}

	if err := UnmarshallWithUnmappedData(&embedder, &embedder.ExtensionData, []byte(raw)); err != nil {
		t.Fatalf("could not unmarshal: %v", err)
	}

	if embedder.A != "foo" {
		t.Errorf("expected A to be '%s', got '%s'", "foo", embedder.A)
	}

	if embedder.B != "bar" {
		t.Errorf("expected B to be '%s', got '%s'", "bar", embedder.B)
	}

	if len(embedder.ExtensionData) != 1 {
		t.Fatalf("expected %d extension data entry, got %v", 1, embedder.ExtensionData)
	}

	extra, ok := embedder.ExtensionData["extra"]
	if !ok {
		t.Error("expected extra")
	}

	s, ok := extra.(string)
	if !ok || s != "baz" {
		t.Errorf("expected extra to be string '%[1]s', got %[2]v (%[2]T)", "baz", extra)
	}
}
