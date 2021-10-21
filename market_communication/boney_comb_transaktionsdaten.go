package market_communication

import "time"

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
