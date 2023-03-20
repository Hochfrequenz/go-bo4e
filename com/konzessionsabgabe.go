package com

import "github.com/hochfrequenz/go-bo4e/enum/abgabeart"

type Konzessionsabgabe struct {
	Satz      abgabeart.AbgabeArt `json:"satz"`      // Art der Abgabe
	Kosten    *float64            `json:"kosten"`    // Konzessionsabgabe in E/kWh
	Kategorie *string             `json:"kategorie"` // Gebührenkategorie der Konzessionsabgabe
}
