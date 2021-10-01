package com

import "github.com/shopspring/decimal"

// Geokoordinaten are GPS coordinates
type Geokoordinaten struct {
	Breitengrad decimal.Decimal `json:"breitengrad" validate:"latitude"`  // latitude
	Laengengrad decimal.Decimal `json:"laengengrad" validate:"longitude"` // longitude
}
