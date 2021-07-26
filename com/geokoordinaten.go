package com

// Geokoordinaten are GPS coordinates
type Geokoordinaten struct {
	Breitengrad float32 `json:"breitengrad" validate:"latitude"`  // latitude
	Laengengrad float32 `json:"laengengrad" validate:"longitude"` // longitude
}
