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

type SomeStructWithMoreFields struct {
	ExtensionData `json:"-"`
	A             string
	B             int
	X             string
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

	expectedUnmappedData := map[string]any{
		"X": "very nice",
	}
	expectedStrongTypedFields := SomeStruct{A: someStructWithUnmappedData.A, B: someStructWithUnmappedData.B, ExtensionData: expectedUnmappedData}

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

	unmappedData := map[string]any{
		"X": "very nice",
	}
	structWithUnmappedData := SomeStruct{A: expectedStructWithUnmappedData.A, B: expectedStructWithUnmappedData.B, ExtensionData: unmappedData}

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
