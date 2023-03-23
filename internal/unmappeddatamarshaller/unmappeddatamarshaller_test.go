package unmappeddatamarshaller

import (
	"encoding/json"
	"reflect"
	"testing"
)

type SomeStruct struct {
	UnmappedData
	A string
	B int
}

type Shadow struct {
	*SomeStruct
}

func (s *SomeStruct) UnmarshalJSON(bytes []byte) (err error) {
	if s.UnmappedData.Data == nil {
		s.UnmappedData.Data = map[string]any{}
	}

	return UnmarshallWithUnmappedData(s, s, &s.UnmappedData, bytes)
}

func (s SomeStruct) MarshalJSON() ([]byte, error) {
	return MarshalWithUnmappedData(s)
}

type SomeStructWithMoreFields struct {
	UnmappedData `json:"-"`
	A            string
	B            int
	X            string
}

func Test_Unmarshalling_WithUnmappedData_PreservesUnmappedDataInStruct(t *testing.T) {
	someStructWithUnmappedData := SomeStructWithMoreFields{
		A: "nice",
		B: 911,
		X: "very nice",
	}

	bytes, _ := json.Marshal(someStructWithUnmappedData)
	actualStrongTypedFields := SomeStruct{}
	err := json.Unmarshal(bytes, &actualStrongTypedFields)
	if err != nil {
		t.Errorf("Error occured while unmarshalling: %v", err)
	}

	expectedUnmappedData := UnmappedData{Data: map[string]any{
		"X": "very nice",
	}}
	expectedStrongTypedFields := SomeStruct{A: someStructWithUnmappedData.A, B: someStructWithUnmappedData.B, UnmappedData: expectedUnmappedData}

	if !reflect.DeepEqual(actualStrongTypedFields, expectedStrongTypedFields) {
		t.Errorf("Unmarshalling struct with unmapped data failed:\nexpected: %v,\nactual: %v", expectedStrongTypedFields, actualStrongTypedFields)
	}
}

func Test_Marshalling_WitUnmappedData_PreservesUnmappedDataInJson(t *testing.T) {
	expectedStructWithUnmappedData := SomeStructWithMoreFields{
		A: "nice",
		B: 911,
		X: "very nice",
	}

	unmappedData := UnmappedData{Data: map[string]any{
		"X": "very nice",
	}}
	structWithUnmappedData := SomeStruct{A: expectedStructWithUnmappedData.A, B: expectedStructWithUnmappedData.B, UnmappedData: unmappedData}

	actual, err := json.Marshal(structWithUnmappedData)
	if err != nil {
		t.Errorf("Error occured while marshalling: %v", err)
	}

	expected, err := json.Marshal(expectedStructWithUnmappedData)
	if err != nil {
		t.Errorf("Error occured while marshalling: %v", err)
	}

	expectedJson := string(expected)
	actualJson := string(actual)
	if expectedJson != actualJson {
		t.Errorf("Marshalling struct with unmapped data failed:\nexpected: %s,\nactual: %s", expectedJson, actualJson)
	}
}
