package market_communication

import (
	"github.com/hochfrequenz/go-bo4e/bo"
	"regexp"
	"time"
)

// BOneyComb is a structure that is used when dealing with business objects that are embedded into market communication messages.
// The BOneyComb combines an array of business objects named "Stammdaten" with a key value dict of process data named "Transaktionsdaten"
type BOneyComb struct {
	Stammdaten        bo.BusinessObjectSlice `json:"stammdaten" validate:"required"`        // Stammdaten is an array of business objects
	Transaktionsdaten map[string]string      `json:"transaktionsdaten" validate:"required"` // Transaktionsdaten are data relevant only in the context of this market communication message
}

// GetTransactionData checks if the key is present in BOneyComb.Transaktionsdaten, returns its value if it is present and nil otherwise
func (boneyComb *BOneyComb) GetTransactionData(key string) *string {
	if key == "" {
		return nil
	}
	if boneyComb.Transaktionsdaten == nil {
		return nil
	}
	result, found := boneyComb.Transaktionsdaten[key]
	if !found {
		return nil
	}
	return &result
}

// the following 4 properties are present in any boneycomb: Dokumentennummer, Nachrichtendatum, Absender, Empfaenger,

// GetDokumentennummer returns the Dokumentennummer from BOneyComb.Transaktionsdaten if it's present and nil otherwise
func (boneyComb *BOneyComb) GetDokumentennummer() *string {
	return boneyComb.GetTransactionData("Dokumentennummer")
}

// GetNachrichtendatum checks if the message date is present in BOneyComb.Transaktionsdaten, returns its value if it is present and nil otherwise. Returns an error iff the parsing fails
func (boneyComb *BOneyComb) GetNachrichtendatum() (*time.Time, error) {
	stringlyTypedNachrichtenDatum := boneyComb.GetTransactionData("Nachrichtendatum")
	if stringlyTypedNachrichtenDatum == nil || *stringlyTypedNachrichtenDatum == "" {
		return nil, nil
	}
	nachrichtendatum, err := time.Parse(time.RFC3339, *stringlyTypedNachrichtenDatum)
	if err != nil {
		return nil, err
	}
	return &nachrichtendatum, nil
}

// GetAbsenderCode returns the 13 digit ID of the sending Marktteilnehmer if present in the Transaktionsdaten
func (boneyComb *BOneyComb) GetAbsenderCode() *string {
	return boneyComb.getMpCode("Absender")
}

// GetEmpfaengerCode returns the 13 digit ID of the receiving Marktteilnehmer if present in the Transaktionsdaten
func (boneyComb *BOneyComb) GetEmpfaengerCode() *string {
	return boneyComb.getMpCode("Empfaenger")
}

var bo4eUriRegex = regexp.MustCompile(`(bo4e://Marktteilnehmer/)?(?P<mpid>\d{13})`)

// getMpCode returns the 13 digit ID of a Marktteilnehmer if present in the Transaktionsdaten
func (boneyComb *BOneyComb) getMpCode(transactionDataKey string) *string {
	mpString := boneyComb.GetTransactionData(transactionDataKey)
	if mpString == nil {
		return nil
	}
	groupNames := bo4eUriRegex.SubexpNames()
	for _, match := range bo4eUriRegex.FindAllStringSubmatch(*mpString, -1) {
		for groupIdx, group := range match {
			name := groupNames[groupIdx]
			if name == "mpid" {
				return &group
			}
		}
	}
	return nil
}
