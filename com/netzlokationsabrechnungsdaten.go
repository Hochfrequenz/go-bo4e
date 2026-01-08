package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/artikelidtyp"
	"github.com/hochfrequenz/go-bo4e/enum/blindarbeitszahler"
)

// Netzlokationsabrechnungsdaten enthält Abrechnungsdaten der Netzlokation.
type Netzlokationsabrechnungsdaten struct {
	// ArtikelId ist die Artikel- oder Gruppen-ArtikelId.
	ArtikelId *string `json:"artikelId,omitempty"`

	// ArtikelIdTyp ist der Typ der ArtikelId (Einzel oder Gruppe).
	ArtikelIdTyp *artikelidtyp.ArtikelIdTyp `json:"artikelIdTyp,omitempty"`

	// Blindarbeitszahler gibt den Zahler der Blindarbeit an.
	Blindarbeitszahler *blindarbeitszahler.Blindarbeitszahler `json:"blindarbeitszahler,omitempty"`

	// FindetBlindarbeitsAbrechnungStatt gibt an, ob eine Abrechnung der Blindarbeit stattfindet.
	FindetBlindarbeitsAbrechnungStatt *bool `json:"findetBlindarbeitsAbrechnungStatt,omitempty"`

	// LieferantBereitZurZahlungBlindarbeit gibt an, ob der Lieferant bereit ist, die Blindarbeit zu zahlen.
	LieferantBereitZurZahlungBlindarbeit *bool `json:"lieferantBereitZurZahlungBlindarbeit,omitempty"`
}
