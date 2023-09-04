package com

// Geokoordinaten are GPS coordinates
type Geokoordinaten struct {
	Breitengrad string `json:"breitengrad" validate:"latitude"`  // Breitengrad is the latitude
	Laengengrad string `json:"laengengrad" validate:"longitude"` // Laengengrad is the longitude
	// todo: change to decimal.Decimal -> validator latidude
}
