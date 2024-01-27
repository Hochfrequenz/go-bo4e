package bo_test

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/bo"
)

func Test_Get_NeloId_Checksum(t *testing.T) {
	actual := bo.GetNeLoIdCheckSum("E113735592")
	then.AssertThat(t, actual, is.EqualTo(1))
}

func Test_Get_NeloId_Doesnt_Panic(t *testing.T) {
	actual := bo.GetNeLoIdCheckSum("E5345G7F6F")
	then.AssertThat(t, actual, is.EqualTo(0))
}
