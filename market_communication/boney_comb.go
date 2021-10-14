package market_communication

import "github.com/hochfrequenz/go-bo4e/bo"

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
