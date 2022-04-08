package sparte

import (
	"testing"

	"github.com/corbym/gocrest/then"

	"github.com/corbym/gocrest/is"
)

// TestStringifySparteForDB checks if the Value method converts the Sparte to its string representation and if the Scan reads it
func TestStringifySparteForDB(t *testing.T) {
	s := STROM
	v, err := s.Value()
	then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
	sb := v.(string)

	var rs Sparte
	err = rs.Scan(&sb)
	then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))

	then.AssertThat(t, sb, is.EqualTo("STROM").Reason("Value() should return the string representation"))
	then.AssertThat(t, rs, is.EqualTo(STROM).Reason("Scan() result is original Sparte Enum"))
}

func TestScanSparte_With_Malformed_Value(t *testing.T) {
	var rs Sparte
	var invalidSpartenString = "foobar"
	err := rs.Scan(&invalidSpartenString)
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}

func TestScanSparte_With_Nil_Value_Should_Be_Ok(t *testing.T) {
	rs := STROM
	var nullString *string = nil
	var expectedSparte Sparte = 0

	err := rs.Scan(nullString)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedSparte))

	err = rs.Scan(nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, rs, is.EqualTo(expectedSparte))
}
func TestValueSparte_With_Nonexistent_Value(t *testing.T) {
	var rs Sparte = 99

	_, err := rs.Value()
	then.AssertThat(t, err, is.Not(is.Nil()).Reason("An error should occur in this test case"))
}
