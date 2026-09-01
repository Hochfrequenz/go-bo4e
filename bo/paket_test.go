package bo_test

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
)

func Test_Get_PaketId_Checksum(t *testing.T) {
	actual, err := bo.GetPaketIdCheckSum("P9ABC12345")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, actual, is.EqualTo(4))
}

func Test_Get_PaketId_Checksum_Zero(t *testing.T) {
	actual, err := bo.GetPaketIdCheckSum("P95IK1KIK9")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, actual, is.EqualTo(0))
}

func Test_Get_PaketId_Checksum_Rejects_Invalid_Input(t *testing.T) {
	for _, invalidInput := range []string{
		"P8ABC12345", // the second character has to be a '9'
		"PAABC12345",
		"D9ABC12345", // wrong Codetyp, that's a TR-ID
		"P9abc12345", // lower case characters are not allowed
		"P9ABC1234",  // too short
		"P9ABC123456",
		"",
	} {
		_, err := bo.GetPaketIdCheckSum(invalidInput)
		then.AssertThat(t, err, is.Not(is.Nil()))
	}
}
