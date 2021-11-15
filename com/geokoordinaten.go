package com

import "github.com/shopspring/decimal"

// Geokoordinaten are GPS coordinates
type Geokoordinaten struct {
	Breitengrad decimal.Decimal `json:"breitengrad" validate:"latitude"`  // Breitengrad is the latitude
	Laengengrad decimal.Decimal `json:"laengengrad" validate:"longitude"` // Laengengrad is the longitude
}
