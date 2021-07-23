package com

import "../enum"

// Adresse is a general purpose address struct
type Adresse struct {
	Postleitzahl string          `json:"postleitzahl" example:"82031"`                       // zip code, mandatory
	Ort          string          `json:"ort" example:"Grünwald"`                             // city name, mandatory
	Strasse      string          `json:"strasse" example:"Nördliche Münchner Straße 27A"`    // street name
	Hausnummer   string          `json:"hausnummer" example:"27A"`                           // house number including suffixes like "a" or "-1"
	Postfach     string          `json:"postfach"`                                           // post box
	Adresszusatz string          `json:"adresszusatz" example:"3. Stock linke Wohnung"`      // additional data, like f.e. floor or "Hinterhaus"
	CoErgaenzung string          `json:"co_ergaenzung" example:"c/o Veronica Hauptmieterin"` //
	Landescode   enum.Landescode `json:"landescode"`                                         // ISO 3166 country code
}
