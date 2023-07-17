package bo

import (
	"regexp"
	"strconv"
	"unicode"
)

// NeLo is short for Netzlokation.
// There should be a nelo struct in the future. for now we just add the nelo-id validation

// neloIdRegex is a regex that all Netzlokation-IDs must match: An "E" followed by 9 upper case letters or digits and a trailing checksum
var neloIdRegex = regexp.MustCompile(`^E[A-Z\d]{9}\d{1}$`)

// neloIdRegexWithoutChecksum is a regex that all Netzlokation-IDs[0:10] must match: An "E" followed by 9 upper case letters or digits BUT WITHOUT A TRAILING CHECKSUM
var neloIdRegexWithoutChecksum = regexp.MustCompile(`^E[A-Z\d]{9}$`)

// GetNeLoIdCheckSum returns the checksum (11th character of the nelo ID) that matches the first ten characters long provided in neloIdWithoutCheckSum. This is going to crash if the length of the neloIdWithoutCheckSum is <10. Use neloIdWithoutCheckSum + strconv.Itoa(returnValue) to generate a NeLoId
func GetNeLoIdCheckSum(neloIdWithoutCheckSum string) int {
	// Quote from https://bdew-codes.de/Content/Files/Anwdh_2023-01-18-AWH-Identifikatoren-MaKo-Bildungsvorschrift_Version.1.0.pdf chapter 6.2
	// > Das ASCII-Verfahren zur Berechnung der Prüfziffer findet bei der Ressourcen-ID und der NeLo-ID Anwendung.
	// Verfahren:
	// a) Umwandlung der Buchstaben mittels ASCII-Tabelle in Zahlenwerte
	// b) Quersumme aller Ziffern in ungerader Position
	// c) Quersumme aller Ziffern auf gerader Position multipliziert mit 2
	// d) Summe von b) und c)
	// e) Differenz von d) zum nächsthöheren Vielfachen von 10 (ergibt sich hier 10, wird die
	// Prüfziffer 0 genommen)
	// Beispiel: Code: A 1 1 3 7 3 5 5 9 2 PZ
	// a) A = 65
	// b) 65 + 1 + 7 + 5 + 9 = 87
	// c) (1 + 3 + 3 + 5 + 2) * 2 = 28
	// d) 87 + 28 = 115
	// e) 120 - 115 = 5 => Prüfziffer 5
	// Identifikationsnummer: A 1 1 3 7 3 5 5 9 2 5
	// Find an online tool for the check here: https://bdew-codes.de/Codenumbers/NetLocationId (click "Prüfziffernrechner" on the right sidebar)
	inputMatchesRegex := neloIdRegexWithoutChecksum.MatchString(neloIdWithoutCheckSum)
	if !inputMatchesRegex {
		panic("You must provide a string that matches ^E[A-Z\\d]{9}$")
	}
	evenSum := 0
	oddSum := 0
	for index, digitRune := range neloIdWithoutCheckSum[0:10] {
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
	resultMatchesRegex := neloIdRegex.MatchString(neloIdWithoutCheckSum + strconv.Itoa(result))
	if !resultMatchesRegex {
		panic("This function is broken; And this should never happen")
	}
	return result
}
