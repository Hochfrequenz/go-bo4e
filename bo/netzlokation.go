package bo

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal"
)

// NeLo is short for Netzlokation.

// neloIdRegex is a regex that all Netzlokation-IDs must match: An "E" followed by 9 upper case letters or digits and a trailing checksum
var neloIdRegex = regexp.MustCompile(`^E[A-Z\d]{9}\d{1}$`)

// neloIdRegexWithoutChecksum is a regex that all Netzlokation-IDs[0:10] must match: An "E" followed by 9 upper case letters or digits BUT WITHOUT A TRAILING CHECKSUM
var neloIdRegexWithoutChecksum = regexp.MustCompile(`^E[A-Z\d]{9}$`)

// GetNeLoIdCheckSum returns the checksum (11th character of the nelo ID) that matches the first ten characters long provided in neloIdWithoutCheckSum. This is going to crash if the length of the neloIdWithoutCheckSum is <10. Use neloIdWithoutCheckSum + strconv.Itoa(returnValue) to generate a NeLoId
func GetNeLoIdCheckSum(neloIdWithoutCheckSum string) (int, error) {
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
		return 0, fmt.Errorf("you must provide a string that matches E[A-Z\\d]{9}, but '%s' does not", neloIdWithoutCheckSum)
	}
	checksum, checksumErr := internal.GetChecksum(neloIdWithoutCheckSum)
	if checksumErr != nil {
		return 0, checksumErr
	}
	result := neloIdWithoutCheckSum + checksum
	resultMatchesRegex := neloIdRegex.MatchString(result)
	if !resultMatchesRegex {
		return 0, fmt.Errorf("this function is broken; And this should never happen")
	}
	return strconv.Atoi(checksum)
}

// Netzlokation is a minimalistic implementation of the BO Netzlokation. But this small implementation alone, allows use to unmarshall boneycombs that contain Netzlokation-BOs
type Netzlokation struct {
	Geschaeftsobjekt
	NetzlokationsId            *string               `json:"netzlokationsId,omitempty" example:"EOI05HSBJG0" validate:"alphanum,len=11"` // NetzlokationsId is the ID of the Netzlokation
	Sparte                     *sparte.Sparte        `json:"sparte,omitempty"`                                                           // Sparte describes the Division
	Netzanschlussleistung      *com.Menge            `json:"netzanschlussleistung,omitempty"`                                            // Netzanschlussleistungsmenge der Netzlokation
	GrundzustaendigerMSBCodeNr *string               `json:"grundzustaendigerMSBCodeNr,omitempty"`                                       // Codenummer des grundzuständigen Messstellenbetreibers, der für diese Netzlokation zuständig ist.
	Steuerkanal                *bool                 `json:"steuerkanal,omitempty"`                                                      // Ob ein Steuerkanal der Netzlokation zugeordnet ist und somit die Netzlokation gesteuert werden kann. ZF2: Kein Steuerkanal vorhanden. ZF3: Steuerkanal vorhanden.
	ObisKennzahl               *string               `json:"obisKennzahl,omitempty"`                                                     // Die OBIS-Kennzahl für die Netzlokation
	Verwendungszweck           *com.Verwendungszweck `json:"verwendungszweck,omitempty"`                                                 // Verwendungungszweck der Werte Netzlokation
	// Konfigurationsprodukte is missing the Marktteilnehmer because adding it causes a serious issue: "could not import github.com/hochfrequenz/go-bo4e/bo (-: import cycle not allowed in test) (typecheck)".
	Konfigurationsprodukte     []com.Konfigurationsprodukt `json:"konfigurationsprodukte,omitempty"`     // Produkt-Daten der Netzlokation
	EigenschaftMSBLokation     *marktrolle.Marktrolle      `json:"eigenschaftMSBLokation,omitempty"`     // Eigenschaft des Messstellenbetreibers an der Lokation
	LokationsbuendelObjektcode *string                     `json:"lokationsbuendelObjektcode,omitempty"` // Lokationsbuendel-Code, der die Funktion dieses BOs an der Lokationsbuendelstruktur beschreibt.
}
