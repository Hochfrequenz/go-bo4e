package com

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
)

// Adresse is a general purpose address struct
type Adresse struct {
	Postleitzahl string                 `json:"postleitzahl,omitempty" example:"82031" validate:"numeric,required"` // Postleitzahl is the zip code, mandatory
	Ort          string                 `json:"ort,omitempty" example:"Grünwald" validate:"alphaunicode,required"`  // Ort is the city name, mandatory
	Strasse      string                 `json:"strasse,omitempty" example:"Nördliche Münchner Straße 27A"`          // Strasse is the street name
	Hausnummer   string                 `json:"hausnummer,omitempty" example:"27A"`                                 // Hausnummer is the house number including suffixes like "a" or "-1"
	Postfach     string                 `json:"postfach,omitempty"`                                                 // Postfach is a post box
	Adresszusatz string                 `json:"adresszusatz,omitempty" example:"3. Stock linke Wohnung"`            // Adresszusatz are additional data, like f.e. floor or "Hinterhaus"
	CoErgaenzung string                 `json:"co_ergaenzung,omitempty" example:"c/o Veronica Hauptmieterin"`       //
	Landescode   *landescode.Landescode `json:"landescode,omitempty"`                                               // Landescode is a ISO 3166 country code
}

// AdresseStructLevelValidation does a cross check on a Adresse object
func AdresseStructLevelValidation(sl validator.StructLevel) {
	// ToDo: use required_without/required_if instead of own validator
	// see https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Required_Without
	address := sl.Current().Interface().(Adresse)
	if (len(address.Strasse) == 0 && len(address.Postfach) == 0) || (len(address.Strasse) > 0 && len(address.Postfach) > 0) {
		sl.ReportError(address.Strasse, "Strasse", "Postfach", "StrasseXORPostfach", "")
	}
	if (len(address.Strasse) > 0 && len(address.Hausnummer) == 0) || (len(address.Strasse) == 0 && len(address.Hausnummer) > 0) {
		sl.ReportError(address.Strasse, "Strasse", "Hausnummer", "StrasseRequiresHausnummer", "")
	}
}
