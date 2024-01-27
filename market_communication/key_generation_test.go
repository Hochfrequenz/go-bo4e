package market_communication_test

import (
	"regexp"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/hochfrequenz/go-bo4e/market_communication"
)

var expectedBgmPattern = regexp.MustCompile(`[A-Z\d]{35}`)
var expectedUnhPattern = regexp.MustCompile(`[A-Z\d]{14}`)

func Test_GenerationRandomDokumentennummer(t *testing.T) {
	actual := market_communication.GenerateRandomDokumentennummer()
	then.AssertThat(t, expectedBgmPattern.MatchString(actual), is.True())
}

func Test_GenerationRandomNachrichtenReferenznummer(t *testing.T) {
	actual := market_communication.GenerateRandomNachrichtenReferenznummer()
	then.AssertThat(t, expectedUnhPattern.MatchString(actual), is.True())
}

const numberOfTries = 1000 // <-- the higher the more "random" the functions return values are

// funcGeneratesDuplicates returns true iff the function returned duplicates within numberOfTries invocations
func funcGeneratesDuplicates(stringGenerator func() string) bool {
	set := map[string]struct{}{} // in native go, a set is a map with empty values - oh wow!
	for i := 0; i < numberOfTries; i++ {
		set[stringGenerator()] = struct{}{}
	}
	return len(set) != numberOfTries
}

func Test_RandomDokumentennummer_Is_Random_Enough(t *testing.T) {
	then.AssertThat(t, funcGeneratesDuplicates(market_communication.GenerateRandomDokumentennummer), is.False())
}

func Test_RandomNachrichtenReferenznummer_Is_Random_Enough(t *testing.T) {
	then.AssertThat(t, funcGeneratesDuplicates(market_communication.GenerateRandomNachrichtenReferenznummer), is.False())
}
