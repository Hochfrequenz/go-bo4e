package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/schwachlastfaehigkeit"
	"github.com/shopspring/decimal"
)

// A Zaehlwerk is the counting part of a meter. A meter consists of one or more Zaehlwerke
type Zaehlwerk struct {
	ZaehlwerkId    string                          `json:"zaehlwerkId,omitempty" validate:"required" example:"47110815_1"`          // ZaehlwerkId ist die Identifikation des Zählwerks (Registers) innerhalb des Zählers. Oftmals eine laufende Nummer hinter der Zählernummer.
	Bezeichnung    string                          `json:"bezeichnung,omitempty" validate:"required" example:"Zählwerk_Wirkarbeit"` // Bezeichnung ist eine zusätzliche Bezeichnung
	Richtung       energierichtung.Energierichtung `json:"richtung,omitempty" validate:"required"`                                  // Richtung beschreibt die Energierichtung: Einspeisung oder Ausspeisung.
	ObisKennzahl   string                          `json:"obisKennzahl,omitempty" validate:"required" example:"1-0:1.8.1"`          // Die ObisKennzahl für das Zählwerk, die festlegt, welche auf die gemessene Größe mit dem Stand gemeldet wird. Nur Zählwerkstände mit dieser OBIS-Kennzahl werden an diesem Zählwerk registriert.
	Wandlerfaktor  decimal.Decimal                 `json:"wandlerfaktor" validate:"required"`                                       // Der Wandlerfaktor ist der Faktor, mit dem die Zählerstandsdifferenz multipliziert wird, um zum eigentlichen Verbrauch im Zeitraum zu kommen.
	Einheit        mengeneinheit.Mengeneinheit     `json:"einheit,omitempty" validate:"required"`                                   // Die Einheit der gemessenen Größe
	Zaehlerstaende Zaehlerstaende                  `json:"zaehlerstaende,omitempty"`                                                // A list of Zaehlerstand (Non BO4E Standard)

	// The following fields are non bo4e standard (yet)

	// note that both "kommastellen" (plural) serialize as singular "...stelle" for compatability reasons (with bo4e dotnet)
	Vorkommastellen   *int                                        `json:"vorkommastelle,omitempty" validate:"omitempty,min=0"`  // Vorkommastellen ist the number of pre-decimal places
	Nachkommastellen  *int                                        `json:"nachkommastelle,omitempty" validate:"omitempty,min=0"` // Nachkommastellen is the number of post decimal places
	Schwachlastfaehig schwachlastfaehigkeit.Schwachlastfaehigkeit `json:"schwachlastfaehig,omitempty"`
}
