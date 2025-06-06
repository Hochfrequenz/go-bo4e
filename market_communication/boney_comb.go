package market_communication

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
)

// BOneyComb is a structure that is used when dealing with business objects that are embedded into market communication messages.
// The BOneyComb combines an array of business objects named "Stammdaten" with a key value dict of process data named "Transaktionsdaten"
type BOneyComb struct {
	Stammdaten        bo.BusinessObjectSlice `json:"stammdaten" validate:"required"`        // Stammdaten is an array of business objects
	Transaktionsdaten map[string]string      `json:"transaktionsdaten" validate:"required"` // Transaktionsdaten are data relevant only in the context of this market communication message
	Links             map[string][]string    `json:"links"`                                 // Links describes relations between different BusinessObjects in Stammdaten
	Anfrage           *BOneyComb             `json:"anfrage,omitempty"`                     // Anfrage can contain the related market communication message e.g. ORDERS => ORDRSP or MSCONS => Storno
}

var pruefiPattern = regexp.MustCompile(`^[1-9]\d{4}$`)

// PruefidentifikatorInTransaktionsdatenValidation returns true iff a valid Pruefidentifikator (see GetPruefidentifikator) is present
func PruefidentifikatorInTransaktionsdatenValidation(sl validator.StructLevel) {
	boneyComb := sl.Current().Interface().(BOneyComb)
	pruefi := boneyComb.GetPruefidentifikator()
	if pruefi == nil {
		sl.ReportError(boneyComb.Stammdaten, "Pruefidentifikator", "Pruefidentifikator", "Pruefi must not be null", "")
	} else if !pruefiPattern.MatchString(*pruefi) {
		sl.ReportError(boneyComb.Stammdaten, "Pruefidentifikator", "Pruefidentifikator", "Pruefi must be set and 5 digits long", "")
	}
}

// UnmarshalJSON is a custom method which enables the caller to ignore unimplemented business objects.
// This is handy for cases where the consumer of the BOneyComb does only care about specific BOs inside the BoneyComb.
func (boc *BOneyComb) UnmarshalJSON(data []byte) error {
	var boneyComb struct {
		Stammdaten        json.RawMessage     `json:"stammdaten" validate:"required"`
		Transaktionsdaten map[string]string   `json:"transaktionsdaten" validate:"required"`
		Links             map[string][]string `json:"links"`
		Anfrage           json.RawMessage     `json:"anfrage"`
	}

	errs := errors.Join(json.Unmarshal(data, &boneyComb))

	stammdaten := bo.BusinessObjectSlice{}
	if err := json.Unmarshal(boneyComb.Stammdaten, &stammdaten); err != nil {
		errs = errors.Join(errs, err)
	}

	*boc = BOneyComb{
		Stammdaten:        stammdaten,
		Transaktionsdaten: boneyComb.Transaktionsdaten,
		Links:             boneyComb.Links,
	}

	if boneyComb.Anfrage != nil {
		var anfrage BOneyComb
		if err := json.Unmarshal(boneyComb.Anfrage, &anfrage); err != nil {
			errs = errors.Join(errs, fmt.Errorf("anfrage: %w", err))
		} else {
			boc.Anfrage = &anfrage
		}
	}

	return errs
}
