package bo

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hochfrequenz/go-bo4e/internal"
)

// Paket-ID is short for Paket-Identifikationsnummer; it identifies the Lokationen that are affected by a Netzbetreiberwechsel

// paketIdRegex is a regex that all Paket-IDs must match: A "P", a "9", then 8 upper case letters or digits and a trailing checksum
var paketIdRegex = regexp.MustCompile(`^P9[A-Z\d]{8}\d{1}$`)

// paketIdRegexWithoutChecksum is a regex that all Paket-IDs[0:10] must match: A "P", a "9" and 8 upper case letters or digits BUT WITHOUT A TRAILING CHECKSUM
var paketIdRegexWithoutChecksum = regexp.MustCompile(`^P9[A-Z\d]{8}$`)

// GetPaketIdCheckSum returns the checksum (11th character of the Paket-ID) that matches the first ten characters provided in paketIdWithoutCheckSum. It returns an error if paketIdWithoutCheckSum does not match ^P9[A-Z\d]{8}$. Use paketIdWithoutCheckSum + strconv.Itoa(returnValue) to generate a Paket-ID
func GetPaketIdCheckSum(paketIdWithoutCheckSum string) (int, error) {
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
	// The Bildungsvorschrift of the Paket-ID is described in chapter 8.2 of the same document: the second character is always a "9".
	inputMatchesRegex := paketIdRegexWithoutChecksum.MatchString(paketIdWithoutCheckSum)
	if !inputMatchesRegex {
		return 0, fmt.Errorf("you must provide a string that matches ^P9[A-Z\\d]{8}, but '%s' does not", paketIdWithoutCheckSum)
	}
	checksum, checksumErr := internal.GetChecksum(paketIdWithoutCheckSum)
	if checksumErr != nil {
		return 0, checksumErr
	}
	result := paketIdWithoutCheckSum + checksum
	resultMatchesRegex := paketIdRegex.MatchString(result)
	if !resultMatchesRegex {
		return 0, fmt.Errorf("this function is broken; And this should never happen")
	}
	return strconv.Atoi(checksum)
}
