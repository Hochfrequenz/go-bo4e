package bo_test

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
)

func Test_Get_LoBueId_Checksum(t *testing.T) {
	actual, err := bo.GetLoBueIdCheckSum("G816417ST7")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, actual, is.EqualTo(3))
}

func Test_Get_LoBueId_Checksum_Zero(t *testing.T) {
	actual, err := bo.GetLoBueIdCheckSum("GJIN5LOWFB")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, actual, is.EqualTo(0))
}

func Test_Get_LoBueId_Checksum_Rejects_Invalid_Input(t *testing.T) {
	for _, invalidInput := range []string{
		"E816417ST7", // wrong Codetyp, that's a NeLo-ID
		"Gabc417ST7", // lower case characters are not allowed
		"G816417ST",  // too short
		"G816417ST73",
		"",
	} {
		_, err := bo.GetLoBueIdCheckSum(invalidInput)
		then.AssertThat(t, err, is.Not(is.Nil()))
	}
}
