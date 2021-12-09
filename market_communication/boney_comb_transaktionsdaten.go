package market_communication

import (
	"sort"
	"time"
)

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

// dokumentnummerKey is the key under which the Dokumentennummer (BGM) is stored in the BOneyComb.Transaktionsdaten
const dokumentnummerKey = "Dokumentennummer"

// nachrichtenReferenzNummerKey is the key under which the NachrichtenReferenznummer (UNH) is stored in the BOneyComb.Transaktionsdaten
const nachrichtenReferenzNummerKey = "NachrichtenReferenznummer"

// nachrichtendatumKey is the key under which the Nachrichtendatum is stored in the BOneyComb.Transaktionsdaten
const nachrichtendatumKey = "Nachrichtendatum"

// nachrichtendatumKey is the key under which the Prüfidentifikator (RFF+Z13+.....) is stored in the BOneyComb.Transaktionsdaten
const pruefidentifikatorKey = "Pruefidentifikator"

// GetPruefidentifikator returns the Pruefidentifikator from BOneyComb.Transaktionsdaten if it's present and nil otherwise
func (boneyComb *BOneyComb) GetPruefidentifikator() *string {
	return boneyComb.GetTransactionData(pruefidentifikatorKey)
}

// SetPruefidentifikator sets the Pruefidentifikator in the BOneyComb.Transaktionsdaten
func (boneyComb *BOneyComb) SetPruefidentifikator(pruefi string) {
	boneyComb.initializeTransaktionsdaten()
	boneyComb.Transaktionsdaten[pruefidentifikatorKey] = pruefi
}

// GetDokumentennummer returns the Dokumentennummer (BMG) from BOneyComb.Transaktionsdaten if it's present and nil otherwise
func (boneyComb *BOneyComb) GetDokumentennummer() *string {
	return boneyComb.GetTransactionData(dokumentnummerKey)
}

// SetDokumentennummer sets the Dokumentennummer (BGM) in the BOneyComb.Transaktionsdaten
func (boneyComb *BOneyComb) SetDokumentennummer(dokumentnummer string) {
	boneyComb.initializeTransaktionsdaten()
	boneyComb.Transaktionsdaten[dokumentnummerKey] = dokumentnummer
}

// GetNachrichtenReferenznummer returns the Dokumentennummer (UNH) from BOneyComb.Transaktionsdaten if it's present and nil otherwise
func (boneyComb *BOneyComb) GetNachrichtenReferenznummer() *string {
	return boneyComb.GetTransactionData(nachrichtenReferenzNummerKey)
}

// SetNachrichtenReferenznummer sets the Reference Number (UNH) in the BOneyComb.Transaktionsdaten
func (boneyComb *BOneyComb) SetNachrichtenReferenznummer(referenzNummer string) {
	boneyComb.initializeTransaktionsdaten()
	boneyComb.Transaktionsdaten[nachrichtenReferenzNummerKey] = referenzNummer
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

// GetTransaktionsdatenKeys returns a sorted array of keys used in the Transaktionsdaten
// This function is useful for unit testing purposes
func (boneyComb *BOneyComb) GetTransaktionsdatenKeys() []string {
	var result []string
	if boneyComb.Transaktionsdaten != nil {
		for transaktionsdatenKey := range boneyComb.Transaktionsdaten {
			result = append(result, transaktionsdatenKey)
		}
		sort.Strings(result)
	} else {
		result = []string{}
	}
	return result
}
