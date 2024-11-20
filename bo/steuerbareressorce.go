package bo

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hochfrequenz/go-bo4e/enum/steuerkanalsleistungsbeschreibung"
	"github.com/hochfrequenz/go-bo4e/internal"
)

// SR-ID is short for Steuerbare Ressource-ID
// srIdRegex is a regex that all Steuerbare Ressourcen-IDs must match: A "C" followed by 9 upper case letters or digits and a trailing checksum
var srIdRegex = regexp.MustCompile(`^C[A-Z\d]{9}\d{1}$`)

// srIdRegexWithoutChecksum is a regex that all Steuerbare Ressourcen-IDs[0:10] must match: A "C" followed by 9 upper case letters or digits BUT WITHOUT A TRAILING CHECKSUM
var srIdRegexWithoutChecksum = regexp.MustCompile(`^C[A-Z\d]{9}$`)

// GetSRIdCheckSum returns the checksum (11th character of the SR ID) that matches the first ten characters long provided in srIdWithoutCheckSum. This is going to crash if the length of the srIdWithoutCheckSum is <10. Use srIdWithoutCheckSum + strconv.Itoa(returnValue) to generate a SR ID
func GetSRIdCheckSum(srIdWithoutCheckSum string) (int, error) {
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
	inputMatchesRegex := srIdRegexWithoutChecksum.MatchString(srIdWithoutCheckSum)
	if !inputMatchesRegex {
		return 0, fmt.Errorf("you must provide a string that matches ^C[A-Z\\d]{9}, but '%s' does not", srIdWithoutCheckSum)
	}
	checksum, checksumErr := internal.GetChecksum(srIdWithoutCheckSum)
	if checksumErr != nil {
		return 0, checksumErr
	}
	result := srIdWithoutCheckSum + checksum
	resultMatchesRegex := srIdRegex.MatchString(result)
	if !resultMatchesRegex {
		return 0, fmt.Errorf("this function is broken; And this should never happen")
	}
	return strconv.Atoi(checksum)
}

type SteuerbareRessource struct {
	Geschaeftsobjekt
	SteuerbareRessourceId             string                                                               `json:"steuerbareRessourceId" validate:"required"` // Identifikationsnummer einer SteuerbareRessource
	SteuerkanalsLeistungsbeschreibung *steuerkanalsleistungsbeschreibung.Steuerkanalsleistungsbeschreibung `json:"steuerkanalsLeistungsbeschreibung"`         // Leistungsbeschreibung des Steuerkanals
	ZugeordnetMSBCodeNr               *string                                                              `json:"zugeordnetMSBCodeNr"`                       // Angabe des Messstellenbetreibers, der der Steuerbaren Ressource zugeordnet ist
}
