package bo

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hochfrequenz/go-bo4e/internal"
)

// LoBü-ID is short for Lokationsbündel-Identifikationsnummer
// loBueIdRegex is a regex that all Lokationsbündel-IDs must match: A "G" followed by 9 upper case letters or digits and a trailing checksum
var loBueIdRegex = regexp.MustCompile(`^G[A-Z\d]{9}\d{1}$`)

// loBueIdRegexWithoutChecksum is a regex that all Lokationsbündel-IDs[0:10] must match: A "G" followed by 9 upper case letters or digits BUT WITHOUT A TRAILING CHECKSUM
var loBueIdRegexWithoutChecksum = regexp.MustCompile(`^G[A-Z\d]{9}$`)

// GetLoBueIdCheckSum returns the checksum (11th character of the LoBü-ID) that matches the first ten characters provided in loBueIdWithoutCheckSum. It returns an error if loBueIdWithoutCheckSum does not match ^G[A-Z\d]{9}$. Use loBueIdWithoutCheckSum + strconv.Itoa(returnValue) to generate a LoBü-ID
func GetLoBueIdCheckSum(loBueIdWithoutCheckSum string) (int, error) {
	// Quote from https://www.bdew.de/media/documents/Anwendungshilfe_Identifikatoren_in_der_Marktkommunikation_V1.3.pdf chapter 9.2
	// > Das ASCII-Verfahren zur Berechnung der Prüfziffer findet bei der Ressourcen-ID, NeLo-ID, NeBe-ID, LoBü-ID und Paket-ID Anwendung.
	// Verfahren:
	// a) Umwandlung der Buchstaben mittels ASCII-Tabelle in Zahlenwerte
	// b) Addition aller Zahlen in ungerader Position
	// c) Addition aller Zahlen auf gerader Position multipliziert mit 2
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
	// The Bildungsvorschrift of the LoBü-ID is described in chapter 7.2 of the same document.
	inputMatchesRegex := loBueIdRegexWithoutChecksum.MatchString(loBueIdWithoutCheckSum)
	if !inputMatchesRegex {
		return 0, fmt.Errorf("you must provide a string that matches ^G[A-Z\\d]{9}, but '%s' does not", loBueIdWithoutCheckSum)
	}
	checksum, checksumErr := internal.GetChecksum(loBueIdWithoutCheckSum)
	if checksumErr != nil {
		return 0, checksumErr
	}
	result := loBueIdWithoutCheckSum + checksum
	resultMatchesRegex := loBueIdRegex.MatchString(result)
	if !resultMatchesRegex {
		return 0, fmt.Errorf("this function is broken; And this should never happen")
	}
	return strconv.Atoi(checksum)
}
