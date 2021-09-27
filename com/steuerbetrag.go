package com

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
)

// Steuerbetrag ist ein besteuerter Betrag
type Steuerbetrag struct {
	Steuerkennzeichen steuerkennzeichen.Steuerkennzeichen `json:"steuerkennzeichen" validate:"required" example:"Ust19"` // Kennzeichnung des Steuersatzes, bzw. Verfahrens
	// Basiswert and Steuerwert are _not_ marked as required because the steuerwert 0 is actually valid
	Basiswert  float32                     `json:"basiswert" example:"100"`                    // Nettobetrag für den die Steuer berechnet wurde
	Steuerwert float32                     `json:"steuerwert" example:"19"`                    // Aus dem Basiswert berechnete Steuer
	Waehrung   waehrungscode.Waehrungscode `json:"waehrung" example:"EUR" validate:"required"` // currency
}

// SteuerbetragStructLevelValidation does a cross check on a Steuerbetrag object and checks if Steuerbetrag.Steuerkennzeichen, Steuerbetrag.Basiswert and Steuerbetrag.Steuerbetrag are consistent
func SteuerbetragStructLevelValidation(sl validator.StructLevel) {
	steuerbetrag := sl.Current().Interface().(Steuerbetrag)
	if steuerbetrag.Steuerkennzeichen == steuerkennzeichen.RCV {
		return
	}
	expectedSteuerwert := float32(0)
	switch sk := steuerbetrag.Steuerkennzeichen; sk {
	case 0:
		return // already the field level validation should fail on this, we don't need struct level validation
	case steuerkennzeichen.Vst7, steuerkennzeichen.Ust7:
		expectedSteuerwert = steuerbetrag.Basiswert * 0.07
	case steuerkennzeichen.Vst19, steuerkennzeichen.Ust19:
		expectedSteuerwert = steuerbetrag.Basiswert * 0.19
	case steuerkennzeichen.Vst0, steuerkennzeichen.Ust0:
		expectedSteuerwert = steuerbetrag.Basiswert * 0.0
	default:
		err := errors.New(fmt.Sprintf("Validation of Steuerkennzeichen %v", steuerbetrag.Steuerkennzeichen))
		panic(err)
	}
	if expectedSteuerwert != steuerbetrag.Steuerwert {
		sl.ReportError(steuerbetrag.Steuerwert, "Steuerwert", "Steuerwert", "Steuerwert=Basiswert*Steuerkennzeichen", "")
	}
}
