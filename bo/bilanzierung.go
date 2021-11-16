package bo

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/zeitreihentyp"
	"github.com/shopspring/decimal"
	"regexp"
	"time"
)

// Bilanzierung is a business object used for balancing.
// Note that this is not official BO4E standard (yet)!
type Bilanzierung struct {
	Geschaeftsobjekt
	Lastprofile              []com.Lastprofil `json:"lastprofile"`
	Bilanzierungsbeginn      time.Time        `json:"bilanzierungsbeginn,omitempty" validate:"omitempty,ltefield=Bilanzierungsende"`
	Bilanzierungsende        time.Time        `json:"bilanzierungsende,omitempty" validate:"omitempty,gtefield=Bilanzierungsbeginn"`
	Bilanzkreis              string           `json:"bilanzkreis" validate:"omitempty,len=16,eic"`
	Jahresverbrauchsprognose *com.Menge       `json:"jahresverbrauchsprognose,omitempty"`
	Kundenwert               *com.Menge       `json:"kundenwert,omitempty"`
	Verbrauchsaufteilung     decimal.NullDecimal `json:"verbrauchsaufteilung,omitempty"`
	Zeitreihentyp zeitreihentyp.Zeitreihentyp `json:"zeitreihentyp"`
}

var eicRegex = regexp.MustCompile(`(?P<vergabestelle>\\d{2})(?P<typ>A|T|V|W|X|Y|Z)([-A-Z\\d]{12})(?P<pruefziffer>[A-Z0-9])`)

// EICFieldLevelValidation ensures that the string stored in the annotated/tagged field is a valid EIC code
// Therefore it has to obey the "Bildungsvorschrift" according to https://bdew-codes.de/Content/Files/EIC/Awh_20171218_EIC-Vergabe_V1-0.pdf section 2.2.1
func EICFieldLevelValidation(fl validator.FieldLevel) bool {
	eicCandidate := fl.Field().String()
	matched := eicRegex.MatchString(eicCandidate)
	if matched {
		actualCheckCharacter := eicCandidate[len(eicCandidate)-1:]
		expectedCheckCharacter := getEICCheckCharacter(eicCandidate[0 : len(eicCandidate)-1])
		return actualCheckCharacter == expectedCheckCharacter
	}
	return false
}

// getEICCheckCharacter calculates the check character according to the EIC rules (ENTSO-E)
// See https://eepublicdownloads.entsoe.eu/clean-documents/EDI/Library/cim_based/02%20EIC%20Code%20implementation%20guide_final.pdf line 461 (section 7.1)
func getEICCheckCharacter(eicWithoutChecksum string) string {
	// a priori assumption: eicWithoutChecksum is exactly 15 characters long
	numericValues := []rune(eicWithoutChecksum) // step 2
	sum := 0
	for i, numericValue := range numericValues {
		// step 3&4
		positionWeight := rune(16 - i)
		numericValues[i] = numericValue * positionWeight
		sum = sum + int(numericValues[i]) // step 5
	}
	var checkCharacter string
	checkNumber := rune(36 - (sum-1)%37)
	if checkNumber < 10 {
		checkCharacter = string(checkNumber)
	} else {
		checkCharacter, err := numberToEicCharacter(checkNumber)
		if err != nil {
			panic(err)
		}
		if checkCharacter == "-" {
			errorMessage := fmt.Sprintf("The check character must not be '-'. Please choose other 15 characters than '{eicWithoutChecksum}' to start with", eicWithoutChecksum)
			// https://eepublicdownloads.entsoe.eu/clean-documents/EDI/Library/cim_based/02%20EIC%20Code%20implementation%20guide_final.pdf line 509
			panic(errorMessage)
		}
	}
	return checkCharacter
}

// numberToEicCharacter converts the given number/rune to a character according to the ENTSO-E/EIC algorithm
// returns an error if the rune is out of range
func numberToEicCharacter(r rune) (string, error) {
	if r < 10 || r > 36 {
		return "", fmt.Errorf("only the range [10,36] is supported")
	}
	if r == rune(36) {
		return "-", nil
	}
	return string(r + 55), nil
}
