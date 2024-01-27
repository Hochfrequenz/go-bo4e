package landescode_test

import (
	"fmt"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

// Test_Landescode_ToString deserializes an address json
func (s *Suite) Test_Landescode_ToString() {
	then.AssertThat(s.T(), fmt.Sprintf("%v", landescode.DE), is.EqualTo("DE"))
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
