package abweichungsgrund

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/corbym/gocrest/then"

	"github.com/corbym/gocrest/is"
)

// TestStringifyAbweichungsgrundForDB checks if the Value method converts the Sparte to its string representation and if the Scan reads it
func TestStringifyAbweichungsgrundForDB(t *testing.T) {
	for key, grund := range _AbweichungsGrundNameToValue {

		v, err := grund.Value()
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		sb := v.(string)

		var rs AbweichungsGrund
		err = rs.Scan(&sb)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))

		then.AssertThat(t, sb, is.EqualTo(key).Reason("Value() should return the string representation"))
		then.AssertThat(t, rs, is.EqualTo(grund).Reason("Scan() result is original Abweichungsgrund Enum"))
	}
}

// TestStringifyAbweichungsgrundForDB checks if the Value method converts the Sparte to its string representation and if the Scan reads it
func TestAbweichungsgrundMarshalling(t *testing.T) {
	for key, grund := range _AbweichungsGrundNameToValue {

		j, err := json.Marshal(grund)
		expectedString := fmt.Sprintf("\"%s\"", key)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, string(j), is.EqualTo(expectedString).Reason("Marshal() should return the string representation"))

		var rs AbweichungsGrund
		err = json.Unmarshal(j, &rs)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, rs, is.EqualTo(grund).Reason("Unmarshal() result is original Abweichungsgrund Enum"))
	}
}

// TestScanEdiCodeAbweichungsgrundForDB checks if the Value method converts the Sparte to its string representation and if the Scan reads it
func TestScanEdiCodeAbweichungsgrundForDB(t *testing.T) {
	for key, grund := range _AbweichungsGrundEdiToValue {
		var rs AbweichungsGrund
		err := rs.Scan(&key)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, rs, is.EqualTo(grund).Reason("Scan() result is expected Abweichungsgrund Enum"))
	}
}

func TestScanAbweichungsgrund_With_Malformed_Value(t *testing.T) {
	var rs AbweichungsGrund
	var invalidGrundString = "foobar"
	err := rs.Scan(&invalidGrundString)
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestScanAbweichungsgrund_With_Nil_Value_Should_Be_Ok(t *testing.T) {
	rs := SONSTIGER_ABWEICHUNGSGRUND
	var nullString *string = nil
	var expectedGrund AbweichungsGrund = 0

	err := rs.Scan(nullString)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedGrund))

	err = rs.Scan(nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedGrund))
}
func TestValueAbweichungsgrund_With_Nonexistent_Value(t *testing.T) {
	var rs AbweichungsGrund = 99

	_, err := rs.Value()
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestEnumToEdifactCodeMapping(t *testing.T) {
	for expectedCode, grund := range _AbweichungsGrundEdiToValue {
		code, err := grund.EdiValue()
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, code, is.EqualTo(expectedCode).Reason("EDIFACT codes are misaligned between the two mapping tables"))
	}
}

func TestEdifactCodeToEnumMapping(t *testing.T) {
	for expectedGrund, code := range _AbweichungsGrundValueToEdi {
		var grund AbweichungsGrund
		err := grund.FromEdi(code)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, grund, is.EqualTo(expectedGrund).Reason("EDIFACT codes are misaligned between the two mapping tables"))
	}
}

func TestNonexistentCode(t *testing.T) {
	var grund AbweichungsGrund
	err := grund.FromEdi("LOL")
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestNonexistentGrund(t *testing.T) {
	var grund AbweichungsGrund = 99
	_, err := grund.EdiValue()
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}
