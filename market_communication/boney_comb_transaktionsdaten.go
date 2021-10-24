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

// initializeTransaktionsdaten sets BOneyComb.Transaktionsdaten to an empty map iff it is nil
func (boneyComb *BOneyComb) initializeTransaktionsdaten() {
	if boneyComb.Transaktionsdaten == nil {
		boneyComb.Transaktionsdaten = map[string]string{}
	}
}

// dokumentnummerKey is the key under which the Dokumentennummer is stored in the BOneyComb.Transaktionsdaten
const dokumentnummerKey = "Dokumentennummer"

// nachrichtendatumKey is the key under which the Nachrichtendatum is stored in the BOneyComb.Transaktionsdaten
const nachrichtendatumKey = "Nachrichtendatum"

// GetDokumentennummer returns the Dokumentennummer from BOneyComb.Transaktionsdaten if it's present and nil otherwise
func (boneyComb *BOneyComb) GetDokumentennummer() *string {
	return boneyComb.GetTransactionData(dokumentnummerKey)
}

// SetDokumentennummer sets the Dokumentennummer in the BOneyComb.Transaktionsdaten
func (boneyComb *BOneyComb) SetDokumentennummer(dokumentnummer string) {
	boneyComb.initializeTransaktionsdaten()
	boneyComb.Transaktionsdaten[dokumentnummerKey] = dokumentnummer
}

// GetNachrichtendatum checks if the message date is present in BOneyComb.Transaktionsdaten, returns its value if it is present and nil otherwise. Returns an error iff the parsing fails
func (boneyComb *BOneyComb) GetNachrichtendatum() (*time.Time, error) {
	stringlyTypedNachrichtenDatum := boneyComb.GetTransactionData(nachrichtendatumKey)
	if stringlyTypedNachrichtenDatum == nil || *stringlyTypedNachrichtenDatum == "" {
		return nil, nil
	}
	nachrichtendatum, err := time.Parse(time.RFC3339, *stringlyTypedNachrichtenDatum)
	if err != nil {
		return nil, err
	}
	return &nachrichtendatum, nil
}

// SetNachrichtendatum sets the nachrichtendatum in BOneyComb.Transaktionsdaten
func (boneyComb *BOneyComb) SetNachrichtendatum(nachrichtendatum time.Time) {
	boneyComb.initializeTransaktionsdaten()
	boneyComb.Transaktionsdaten[nachrichtendatumKey] = nachrichtendatum.Format(time.RFC3339)
}