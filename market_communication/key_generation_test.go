package market_communication_test

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/market_communication"
	"regexp"
	"time"
)

var expectedBgmPattern = regexp.MustCompile(`[A-Z\d]{35}`)
var expectedUnhPattern = regexp.MustCompile(`[A-Z\d]{14}`)

func (s *Suite) Test_GenerationRandomDokumentennummer() {
	actual := market_communication.GenerateRandomDokumentennummer()
	then.AssertThat(s.T(), expectedBgmPattern.MatchString(actual), is.True())
}

func (s *Suite) Test_GenerationRandomNachrichtenReferenznummer() {
	actual := market_communication.GenerateRandomNachrichtenReferenznummer()
	then.AssertThat(s.T(), expectedUnhPattern.MatchString(actual), is.True())
}

const numberOfTries = 1000 // <-- the higher the more "random" the functions return values are

func funcGeneratesDuplicates(stringGenerator func() string) bool {
	set := map[string]struct{}{} // in native go, a set is a map with empty values - oh wow!
	for i := 0; i < numberOfTries; i++ {
		time.Sleep(1)
		set[stringGenerator()] = struct{}{}
	}
	return len(set) == numberOfTries
}

func (s *Suite) Test_RandomDokumentennummer_Is_Random_Enough() {
	then.AssertThat(s.T(), funcGeneratesDuplicates(market_communication.GenerateRandomDokumentennummer), is.False())
}

func (s *Suite) Test_RandomNachrichtenReferenznummer_Is_Random_Enough() {
	then.AssertThat(s.T(), funcGeneratesDuplicates(market_communication.GenerateRandomNachrichtenReferenznummer), is.False())
}
