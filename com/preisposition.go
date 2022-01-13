package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/artikelid"
	"github.com/hochfrequenz/go-bo4e/enum/bdewartikelnummer"
	"github.com/hochfrequenz/go-bo4e/enum/bemessungsgroesse"
	"github.com/hochfrequenz/go-bo4e/enum/kalkulationsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/leistungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/tarifzeit"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
	"github.com/shopspring/decimal"
)

// Preisposition is part of a Preisblatt and includes a price for a product or service
type Preisposition struct {
	Berechnungsmethode kalkulationsmethode.Kalkulationsmethode `json:"berechnungsmethode,omitempty" validate:"required"` //	Das Modell, das der Preisbildung zugrunde liegt. Details siehe ENUM Kalkulationsmethode
	Leistungstyp       leistungstyp.Leistungstyp               `json:"leistungstyp,omitempty" validate:"required"`       // Standardisierte Bezeichnung für die abgerechnete Leistungserbringung. Details siehe ENUM Leistungstyp
	//Leistungsbezeichung string	// Bezeichnung für die in der Position abgebildete Leistungserbringung
	Preiseinheit             waehrungseinheit.Waehrungseinheit    `json:"preiseinheit,omitempty" validate:"required"`                       // Festlegung, mit welcher Preiseinheit abgerechnet wird, z.B. Ct. oder €
	Bezugsgroesse            mengeneinheit.Mengeneinheit          `json:"bezugsgroesse,omitempty" validate:"required"`                      // Hier wird festgelegt, auf welche Bezugsgrösse sich der Preis bezieht, z.B. kWh oder Stück
	Zeitbasis                *zeiteinheit.Zeiteinheit             `json:"zeitbasis,omitempty" validate:"required"`                // Die Zeit(dauer) auf die sich der Preis bezieht. Z.B. ein Jahr für einen Leistungspreis der in €/kW/Jahr ausgegeben wird
	Tarifzeit                *tarifzeit.Tarifzeit                 `json:"tarifzeit,omitempty" validate:"required"`                // Festlegung, für welche Tarifzeit der Preis hier festgelegt ist
	BDEWArtikelnummer        *bdewartikelnummer.BDEWArtikelnummer `json:"bdewartikelnummer,omitempty" validate:"required"`        // Eine vom BDEW standardisierte Bezeichnug für die abgerechnete Leistungserbringung. Diese Artikelnummer wird auch im Rechnungsteil der INVOIC verwendet
	ArtikelId                *artikelid.ArtikelId                 `json:"artikelid,omitempty" validate:"required"`                // Standardisierte vom BDEW herausgegebene Liste, welche im Strommarkt die BDEW-Artikelnummer ablöst
	Zonungsgroesse           *bemessungsgroesse.Bemessungsgroesse `json:"zonungsgroesse,omitempty" validate:"required"`           // Mit der Menge der hier angegebenen Größe wird die Staffelung/Zonung durchgeführt. Z.B. Vollbenutzungsstunden
	FreimengeBlindarbeit     decimal.NullDecimal                  `json:"freimengeBlindarbeit,omitempty" validate:"required"`     // Der Anteil der Menge der Blindarbeit in Prozent von der Wirkarbeit, für die keine Abrechnung erfolgt
	FreimengeLeistungsfaktor decimal.NullDecimal                  `json:"freimengeLeistungsfaktor,omitempty" validate:"required"` // Der cos phi (Verhältnis Wirkleistung/Scheinleistung) aus dem die Freimenge für die Blindarbeit berechnet wird als tan phi (Verhältnis Blindleistung/Wirkleistung)
	Preisstaffeln            []Preisstaffel                       `json:"preisstaffel" validate:"required"`                      // Preisstaffeln, die zu dieser Preisposition gehören
}
