package bdewartikelnummer

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/corbym/gocrest/then"

	"github.com/corbym/gocrest/is"
)

// TestStringifyArtikelnummerForDB checks if the Value method converts the artikelnummer to its string representation and if the Scan reads it
func TestStringifyArtikelnummerForDB(t *testing.T) {
	for key, artikelnummer := range _BDEWArtikelnummerNameToValue {
		v, err := artikelnummer.Value()
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		sb := v.(string)

		var rs BDEWArtikelnummer
		err = rs.Scan(&sb)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))

		then.AssertThat(t, sb, is.EqualTo(key).Reason("Value() should return the string representation"))
		then.AssertThat(t, rs, is.EqualTo(artikelnummer).Reason("Scan() result is original artikelnummer Enum"))
	}
}

// TestArtikelnummerMarshalling checks if the Value method converts the artikelnummer to its string representation and if the Scan reads it
func TestArtikelnummerMarshalling(t *testing.T) {
	for key, grund := range _BDEWArtikelnummerNameToValue {

		j, err := json.Marshal(grund)
		expectedString := fmt.Sprintf("\"%s\"", key)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, string(j), is.EqualTo(expectedString).Reason("Marshal() should return the string representation"))

		var rs BDEWArtikelnummer
		err = json.Unmarshal(j, &rs)
		then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
		then.AssertThat(t, rs, is.EqualTo(grund).Reason("Unmarshal() result is original artikelnummer Enum"))
	}
}

func TestScanArtikelnummer_With_Malformed_Value(t *testing.T) {
	var rs BDEWArtikelnummer
	var invalidGrundString = "foobar"
	err := rs.Scan(&invalidGrundString)
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestScanArtikelnummer_With_Nil_Value_Should_Be_Ok(t *testing.T) {
	rs := WIRKARBEIT
	var nullString *string = nil
	var expectedArtikel BDEWArtikelnummer = 0

	err := rs.Scan(nullString)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedArtikel))

	err = rs.Scan(nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedArtikel))
}
func TestValueArtikelnummer_With_Nonexistent_Value(t *testing.T) {
	var rs BDEWArtikelnummer = 99

	_, err := rs.Value()
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}
