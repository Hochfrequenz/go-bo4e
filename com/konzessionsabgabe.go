package com

import "github.com/hochfrequenz/go-bo4e/enum/abgabeart"

type Konzessionsabgabe struct {
	Satz      abgabeart.AbgabeArt
	Kosten    *float64
	Kategorie *string
}
