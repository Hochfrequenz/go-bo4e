package rechnungstyp

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/corbym/gocrest/then"

	"github.com/corbym/gocrest/is"
)

// TestStringifyAbweichungsgrundForDB checks if the Value method converts the Sparte to its string representation and if the Scan reads it
func TestStringifyAbweichungsgrundForDB(t *testing.T) {
	for key, rechnungstyp := range _RechnungstypNameToValue {

		v, err := rechnungstyp.Value()
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		sb := v.(string)

		var rs Rechnungstyp
		err = rs.Scan(&sb)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))

		then.AssertThat(t, sb, is.EqualTo(key).Reason("Value() should return the string representation"))
		then.AssertThat(t, rs, is.EqualTo(rechnungstyp).Reason("Scan() result is original Abweichungsgrund Enum"))
	}
}

// TestStringifyAbweichungsgrundForDB checks if the Value method converts the Sparte to its string representation and if the Scan reads it
func TestAbweichungsgrundMarshalling(t *testing.T) {
	for key, grund := range _RechnungstypNameToValue {

		j, err := json.Marshal(grund)
		expectedString := fmt.Sprintf("\"%s\"", key)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, string(j), is.EqualTo(expectedString).Reason("Marshal() should return the string representation"))

		var rs Rechnungstyp
		err = json.Unmarshal(j, &rs)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, rs, is.EqualTo(grund).Reason("Unmarshal() result is original Abweichungsgrund Enum"))
	}
}

// TestScanEdiCodeAbweichungsgrundForDB checks if the Value method converts the Sparte to its string representation and if the Scan reads it
func TestScanEdiCodeAbweichungsgrundForDB(t *testing.T) {
	for key, grund := range _RechnungstypEdiToValue {
		var rs Rechnungstyp
		err := rs.Scan(&key)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, rs, is.EqualTo(grund).Reason("Scan() result is expected Abweichungsgrund Enum"))
	}
}

func TestScanAbweichungsgrund_With_Malformed_Value(t *testing.T) {
	var rs Rechnungstyp
	var invalidGrundString = "foobar"
	err := rs.Scan(&invalidGrundString)
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestScanAbweichungsgrund_With_Nil_Value_Should_Be_Ok(t *testing.T) {
	rs := ABSCHLUSSRECHNUNG
	var nullString *string = nil
	var expectedGrund Rechnungstyp = 0

	err := rs.Scan(nullString)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedGrund))

	err = rs.Scan(nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedGrund))
}
func TestValueAbweichungsgrund_With_Nonexistent_Value(t *testing.T) {
	var rs Rechnungstyp = 99

	_, err := rs.Value()
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestEnumToEdifactCodeMapping(t *testing.T) {
	for expectedCode, grund := range _RechnungstypEdiToValue {
		code, err := grund.EdiValue()
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, code, is.EqualTo(expectedCode).Reason("EDIFACT codes are misaligned between the two mapping tables"))
	}
}

func TestEdifactCodeToEnumMapping(t *testing.T) {
	for expectedRechnungstyp, code := range _RechnungstypValueToEdi {
		var rechnungstyp Rechnungstyp
		err := rechnungstyp.FromEdi(code)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, rechnungstyp, is.EqualTo(expectedRechnungstyp).Reason("EDIFACT codes are misaligned between the two mapping tables"))
	}
}

func TestNonexistentCode(t *testing.T) {
	var rechnungstyp Rechnungstyp
	err := rechnungstyp.FromEdi("LOL")
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestNonexistentGrund(t *testing.T) {
	var rechnungstyp Rechnungstyp = 99
	_, err := rechnungstyp.EdiValue()
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}
