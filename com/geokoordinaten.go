package com

import "github.com/shopspring/decimal"

// Geokoordinaten are GPS coordinates
type Geokoordinaten struct {
	Breitengrad decimal.Decimal `json:"breitengrad" validate:"required"` // Breitengrad is the latitude
	Laengengrad decimal.Decimal `json:"laengengrad" validate:"required"` // Laengengrad is the longitude
	// todo: change to decimal.Decimal -> validator latitude/longitude lte=-90,gte=90/lte=-180,gte=180 won't work
}
