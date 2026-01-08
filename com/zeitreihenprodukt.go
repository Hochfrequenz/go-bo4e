package com

import "github.com/shopspring/decimal"

// Zeitreihenprodukt modelliert Produkte von Summenzeitreihen.
type Zeitreihenprodukt struct {
	// Identifikation ist die OBIS-Kennzahl für das Produkt.
	Identifikation *string `json:"identifikation,omitempty"`

	// Korrekturfaktor ist der Korrekturfaktor der Zeitreihe.
	Korrekturfaktor *decimal.Decimal `json:"korrekturfaktor,omitempty"`

	// Verbrauch ist der Verbrauch der Zeitreihe.
	Verbrauch *Verbrauch `json:"verbrauch,omitempty"`
}
