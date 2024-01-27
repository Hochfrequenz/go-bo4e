package landescode_test

import (
	"fmt"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
)

// Test_Landescode_ToString deserializes an address json
func Test_Landescode_ToString(t *testing.T) {
	then.AssertThat(t, fmt.Sprintf("%v", landescode.DE), is.EqualTo("DE"))
}
