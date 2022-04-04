package sparte

import (
	"testing"

	"github.com/corbym/gocrest/then"

	"github.com/corbym/gocrest/is"
)

// TestStringifySparteForDB checks if the Value method converts the Sparte to its String representaion and if the Scan reads it
func TestStringifySparteForDB(t *testing.T) {
	s := STROM
	v, err := s.Value()

	sb := v.(string)

	var rs Sparte
	err = rs.Scan(&sb)
	then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
	then.AssertThat(t, err, is.Nil().Reason("No error should occur in this test case"))
	then.AssertThat(t, sb, is.EqualTo("STROM").Reason("Value() should return the string representation"))
	then.AssertThat(t, rs, is.EqualTo(STROM).Reason("Scan() result is original Sparte Enum"))
}
