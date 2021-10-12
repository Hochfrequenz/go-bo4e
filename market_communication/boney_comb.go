package market_communication

import "github.com/hochfrequenz/go-bo4e/bo"

// BOneyComb is a structure that is used when dealing with business objects that are embedded into market communication messages.
// The BOneyComb combines an array of business objects named "Stammdaten" with a key value dict of process data named "Transaktionsdaten"
type BOneyComb struct {
	Stammdaten        []bo.BusinessObject    `json:"stammdaten" validate:"required"`        // Stammdaten is an array of business objects
	Transaktionsdaten map[string]interface{} `json:"transaktionsdaten" validate:"required"` // Transaktionsdaten are data relevant only in the context of this market communication message
}

type Transaktionsdaten interface {
	GetPruefidentifikator() string // GetPruefidentifikator returns the Pruefidentifikator
}
