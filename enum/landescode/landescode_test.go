package landescode

import (
	"fmt"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

// SetupSuite sets up the tests
func (s *Suite) SetupSuite() {
}

// Test_Landescode_ToString deserializes an address json
func (s *Suite) Test_Landescode_ToString() {
	then.AssertThat(s.T(), fmt.Sprintf("%v", DE), is.EqualTo("DE"))
}

func (s *Suite) AfterTest(_, _ string) {
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}
