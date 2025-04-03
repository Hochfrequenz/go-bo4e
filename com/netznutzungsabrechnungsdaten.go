package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/artikelidtyp"
	"github.com/shopspring/decimal"
)

// Netznutzungsabrechnungsdaten sind Daten zur Abrechnung der Netznutzung (no shit, sherlock)
type Netznutzungsabrechnungsdaten struct {
	ArtikelId                     *string                    `json:"artikelId" validate:"omitempty"`    // ArtikelId enthält die Artikel oder Gruppen-ArtikelId
	ArtikelIdTyp                  *artikelidtyp.ArtikelIdTyp `json:"artikelIdTyp"`                      // ArtikelIdTyp ist der Typ der ArtikelId (Einzel oder Gruppe)
	Anzahl                        *int                       `json:"anzahl"`                            // Anzahl ist die Anzahl der Positionen für diese ArtikelId
	Gemeinderabatt                decimal.NullDecimal        `json:"gemeinderabatt"`                    // Gemeinderabatt
	Zuschlag                      decimal.NullDecimal        `json:"zuschlag"`                          // Zuschlag
	Abschlag                      decimal.NullDecimal        `json:"abschlag"`                          // Abschlag
	SingulaereBetriebsmittel      *Menge                     `json:"singulaereBetriebsmittel"`          // SingulaereBetriebsmittel sind Singuläre Betriebsmittel
	PreisSingulaereBetriebsmittel *Preis                     `json:"preisSingulaereBetriebsmittel"`     // PreisSingulaereBetriebsmittel ist der Preis für SingulaereBetriebsmittel
	Zaehlzeit                     *Zaehlzeitregister         `json:"zaehlzeit"`                         // Zaehlzeit
	IstDifferenz                  *bool                      `json:"istDifferenz"`                      // IstDifferenz gibt an, ob es sich um Original- oder Differenz-Netznutzungsabrechnungsdaten handelt
	Marktrollen                   *[]MarktpartnerDetails     `json:"marktrollen"  validate:"omitempty"` // Marktrollen gibt die zu den Netznutzungsabrechnungsdaten gehörende Marktrollen (z.B. Netzbetreiber) an

}
