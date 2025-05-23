package com

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/steuerkennzeichen"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/shopspring/decimal"
)

// Steuerbetrag ist ein besteuerter Betrag
type Steuerbetrag struct {
	Steuerkennzeichen steuerkennzeichen.Steuerkennzeichen `json:"steuerkennzeichen,omitempty" validate:"required" example:"UST19"` // Kennzeichnung des Steuersatzes, bzw. Verfahrens
	// Basiswert and Steuerwert are _not_ marked as required because the steuerwert 0 is actually valid
	Sondersteuersatz        *decimal.Decimal            `json:"sondersteuersatz,omitempty"`                           // Wert eines besonderen Steuersatzes, wichig wenn Steuerkennzeichen den Wert UST_SONDER hat
	Basiswert               decimal.Decimal             `json:"basiswert" example:"100"`                              // Basiswert ist der Nettobetrag für den die Steuer berechnet wurde
	Steuerwert              decimal.Decimal             `json:"steuerwert" example:"19"`                              // Steuerwert ist die aus dem Basiswert berechnete Steuer
	Waehrung                waehrungscode.Waehrungscode `json:"waehrung,omitempty" example:"EUR" validate:"required"` // Waehrung is the currency
	BasiswertVorausgezahlt  *decimal.Decimal            `json:"basiswertVorausgezahlt,omitempty"`                     // BasiswertVorausgezahlt ist die Brutto-Summe der vorausbezahlten Beträge
	SteuerwertVorausgezahlt *decimal.Decimal            `json:"steuerwertVorausgezahlt,omitempty"`                    // SteuerwertVorausgezahlt ist die Steuer-Summe der vorausbezahlten Beträge
}

// SteuerbetragStructLevelValidation does a cross check on a Steuerbetrag object and checks if Steuerbetrag.Steuerkennzeichen, Steuerbetrag.Basiswert and Steuerbetrag.Steuerbetrag are consistent
func SteuerbetragStructLevelValidation(sl validator.StructLevel) {
	steuerbetrag := sl.Current().Interface().(Steuerbetrag)
	if steuerbetrag.Steuerkennzeichen == steuerkennzeichen.RCV {
		return
	}
	var expectedSteuerwert decimal.Decimal
	switch sk := steuerbetrag.Steuerkennzeichen; sk {
	case 0:
		return // already the field level validation should fail on this, we don't need struct level validation
	case steuerkennzeichen.VST_7, steuerkennzeichen.UST_7:
		expectedSteuerwert = steuerbetrag.Basiswert.Mul(decimal.NewFromFloat(0.07))
	case steuerkennzeichen.VST_19, steuerkennzeichen.UST_19:
		expectedSteuerwert = steuerbetrag.Basiswert.Mul(decimal.NewFromFloat(0.19))
	case steuerkennzeichen.VST_0, steuerkennzeichen.UST_0:
		expectedSteuerwert = steuerbetrag.Basiswert.Mul(decimal.Zero)
	default:
		err := fmt.Errorf("validation of Steuerkennzeichen %v is not implemented", steuerbetrag.Steuerkennzeichen)
		panic(err)
	}
	if !expectedSteuerwert.Equal(steuerbetrag.Steuerwert) {
		sl.ReportError(steuerbetrag.Steuerwert, "Steuerwert", "Steuerwert", "Steuerwert=Basiswert*Steuerkennzeichen", "")
	}
}
