package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/emobilitaetsart"
	"github.com/hochfrequenz/go-bo4e/enum/erzeugungsart"
	"github.com/hochfrequenz/go-bo4e/enum/speicherart"
	"github.com/hochfrequenz/go-bo4e/enum/technischeressourcenutzung"
	"github.com/hochfrequenz/go-bo4e/enum/technischeressourceverbrauchsart"
	"github.com/hochfrequenz/go-bo4e/enum/waermenutzung"
	"regexp"
	"strconv"
	"unicode"
)

// TR-ID is short for Technische Ressource-ID
// trIdRegex is a regex that all Technische Ressourcen-IDs must match: A "D" followed by 9 upper case letters or digits and a trailing checksum
var trIdRegex = regexp.MustCompile(`^D[A-Z\d]{9}\d{1}$`)

// trIdRegexWithoutChecksum is a regex that all Technische Ressourcen-IDs[0:10] must match: A "D" followed by 9 upper case letters or digits BUT WITHOUT A TRAILING CHECKSUM
var trIdRegexWithoutChecksum = regexp.MustCompile(`^D[A-Z\d]{9}$`)

// GeTRIdCheckSum returns the checksum (11th character of the TR ID) that matches the first ten characters long provided in trIdWithoutCheckSum. This is going to crash if the length of the trIdWithoutCheckSum is <10. Use trIdWithoutCheckSum + strconv.Itoa(returnValue) to generate a TR ID
func GetTRIdCheckSum(trIdWithoutCheckSum string) int {
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
	inputMatchesRegex := trIdRegexWithoutChecksum.MatchString(trIdWithoutCheckSum)
	if !inputMatchesRegex {
		panic("You must provide a string that matches ^D[A-Z\\d]{9}$")
	}
	evenSum := 0
	oddSum := 0
	for index, digitRune := range trIdWithoutCheckSum[0:10] {
		var digit int
		if !unicode.IsDigit(digitRune) {
			// if the digitRune is a letter, then we du the usual ASCII conversion
			digit = int(digitRune) // digit is 65 for digitRune='A'
		} else {
			//, but if it's a "digit" character, then we use the digits value.
			digit = int(digitRune - '0') // digit is 0 for digitRune='0'
		}
		if index%2 == 0 {
			// this is "odd", because BDEW starts counting at 1, so the first index is odd
			oddSum = oddSum + digit
		} else {
			evenSum = evenSum + digit
		}
	}
	stepD := oddSum + (evenSum * 2)
	result := (((stepD/10)+1)*10 - stepD) % 10
	resultMatchesRegex := trIdRegex.MatchString(trIdWithoutCheckSum + strconv.Itoa(result))
	if !resultMatchesRegex {
		panic("This function is broken; And this should never happen")
	}
	return result
}

type TechnischeRessource struct {
	Geschaeftsobjekt
	TechnischeRessourceId            *string                                                            `json:"technischeRessourceId" validate:"required"`                                                                             //Identifikationsnummer einer TechnischeRessource
	VorgelagerteMesslokationsId      *string                                                            `json:"vorgelagerteMesslokationsId,omitempty" example:"DE00713739359S0000000000001222221" validate:"alphanum,required,len=33"` // Vorgelagerte Messlokation ID
	ZugeordneteMarktlokationsId      *string                                                            `json:"zugeordneteMarktlokationsId,omitempty" example:"20072281644"`                                                           // ZugeordneteMarktlokationsId Messlokation ID
	ZugeordneteSteuerbareRessourceId *string                                                            `json:"zugeordneteSteuerbareRessourceId" example:"20072281644"`                                                                // Referenz auf die der Technischen Ressource zugeordneten Steuerbaren Ressource
	NennleistungAufnahme             *com.Menge                                                         `json:"nennleistungAufnahme" example:"QTY+Z43:100:KWT"`                                                                        // Nennleistung (Aufnahme)
	NennleistungAbgabe               *com.Menge                                                         `json:"nennleistungAbgabe" example:"QTY+Z44:100:KWT"`                                                                          //Nennleistung (Abgabe)
	Speicherkapazitaet               *com.Menge                                                         `json:"speicherkapazitaet" example:"QTY+Z42:100:KWH"`                                                                          //Speicherkapazität
	TechnischeRessourceNutzung       *technischeressourcenutzung.TechnischeRessourceNutzung             `json:"technischeRessourceNutzung" example:"CCI+Z17"`                                                                          //Art und Nutzung der Technischen Ressource
	Verbrauchsart                    *technischeressourceverbrauchsart.TechnischeRessourceVerbrauchsart `json:"verbrauchsart" example:"CAV+Z64" `                                                                                      //Verbrauchsart der Technischen Ressource
	Waermenutzung                    *waermenutzung.Waermenutzung                                       `json:"waermenutzung" example:"CAV+Z56"`                                                                                       //Wärmenutzung
	EMobilitaetsart                  *emobilitaetsart.EMobilitaetsart                                   `json:"emobilitaetsart" example:"CAV+Z87"`                                                                                     //Art der E-Mobilität  Das Segment dient dazu, im Falle der E-Mobilität eine genauere Angabe über die Art der E-Mobilität zu definieren
	Erzeugungsart                    *erzeugungsart.Erzeugungsart                                       `json:"erzeugungsart" example:"CAV+ZF5"`                                                                                       //Art der Erzeugung der Energie
	Speicherart                      *speicherart.Speicherart                                           `json:"speicherart" example:"CAV+ZF7"`                                                                                         //Art der speicher. Details <see cref="ENUM.Speicherart" />
}

func (_ TechnischeRessource) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
