package market_communication

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/bo"
	"regexp"
)

// BOneyComb is a structure that is used when dealing with business objects that are embedded into market communication messages.
// The BOneyComb combines an array of business objects named "Stammdaten" with a key value dict of process data named "Transaktionsdaten"
type BOneyComb struct {
	Stammdaten        bo.BusinessObjectSlice `json:"stammdaten" validate:"required"`        // Stammdaten is an array of business objects
	Transaktionsdaten map[string]string      `json:"transaktionsdaten" validate:"required"` // Transaktionsdaten are data relevant only in the context of this market communication message
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
