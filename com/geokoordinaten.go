package com

import "github.com/shopspring/decimal"

// Geokoordinaten are GPS coordinates
type Geokoordinaten struct {
	Breitengrad decimal.Decimal `json:"breitengrad" validate:"required"` // Breitengrad is the latitude
	Laengengrad decimal.Decimal `json:"laengengrad" validate:"required"` // Laengengrad is the longitude
	// todo: change to decimal.Decimal -> validator latitude/longitude lte=-90,gte=90/lte=-180,gte=180 won't work

	// OestlicheLaenge gibt die östliche Länge in UTM-Koordinaten eines entsprechenden Ortes an.
	OestlicheLaenge *decimal.Decimal `json:"oestlicheLaenge,omitempty"`

	// NoerdlicheBreite gibt die nördliche Breite in UTM-Koordinaten eines entsprechenden Ortes an.
	NoerdlicheBreite *decimal.Decimal `json:"noerdlicheBreite,omitempty"`

	// Zone gibt die UTM-Zone des Ortes an.
	Zone *string `json:"zone,omitempty"`

	// NordWert gibt den Nordwert des Ortes in UTM-Koordinaten an.
	NordWert *string `json:"nordWert,omitempty"`

	// OstWert gibt den Ostwert des Ortes in UTM-Koordinaten an.
	OstWert *string `json:"ostWert,omitempty"`
}
