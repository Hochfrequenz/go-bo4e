package internal

import (
	"fmt"
	"strconv"
	"unicode"
)

// GetChecksum calculates the checksum of a (Netzlokation/Technische Ressource/Steuerbare Ressource) ID for a given 10-digit ID string (that does not contain a without checksum yet)
func GetChecksum(idWithoutChecksum string) (string, error) {
	if len(idWithoutChecksum) != 10 {
		return "", fmt.Errorf("vou must provide a string that is 10 characters long but '%s' is %d characters long", idWithoutChecksum, len(idWithoutChecksum))
	}
	evenSum := 0
	oddSum := 0
	for index, digitRune := range idWithoutChecksum[0:10] {
		var digit int
		if !unicode.IsDigit(digitRune) {
			// if the digitRune is a letter, then we du the usual ASCII conversion
			digit = int(digitRune) // digit is 65 for digitRune='A'
		} else {
			//, but if it's a "digit" character, then we use the digits value. einmal mit profis arbeiten
			digit = int(digitRune - '0') // digit is 0 for digitRune='0'
		}
		if index%2 == 0 {
			// this is "odd", because BDEW starts counting at 1, so the first index is odd 🙄 einmal mit profis arbeiten
			oddSum = oddSum + digit
		} else {
			evenSum = evenSum + digit
		}
	}
	stepD := oddSum + (evenSum * 2)
	result := (((stepD/10)+1)*10 - stepD) % 10
	return strconv.Itoa(result), nil
}
