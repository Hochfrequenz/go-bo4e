package market_communication

import (
	"math/rand"
	"time"
)

// unhAllowedCharacters are characters that are allowed to be used in an UNH segment
var unhAllowedCharacters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// maxCharactersUNH is the max length of an UNH (both international and German)
const maxCharactersUNH = 14

// bgmAllowedCharacters are characters that are allowed to be used in a BGM segment
var bgmAllowedCharacters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// maxCharactersBGM is the max length of a BGM key (in Germany)
const maxCharactersBGM = 35

// generateRandomString returns a random combination of the allowed characters with given length
func generateRandomString(allowedCharacters []rune, length uint) string {
	// source: https://stackoverflow.com/a/22892986/10009545
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = allowedCharacters[rand.Intn(len(unhAllowedCharacters))]
	}
	time.Sleep(1 * time.Nanosecond)
	return string(b)
}

// GenerateRandomNachrichtenReferenznummer returns a random UNH key
func GenerateRandomNachrichtenReferenznummer() string {
	return generateRandomString(unhAllowedCharacters, maxCharactersUNH)
}

// GenerateRandomDokumentennummer returns a random BGM key
func GenerateRandomDokumentennummer() string {
	return generateRandomString(bgmAllowedCharacters, maxCharactersBGM)
}
