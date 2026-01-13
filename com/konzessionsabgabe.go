package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/abgabeart"
	"github.com/shopspring/decimal"
)

type Konzessionsabgabe struct {
	Satz      abgabeart.AbgabeArt `json:"satz,omitempty"`
	Kosten    *decimal.Decimal    `json:"kosten,omitempty"`
	Kategorie *string             `json:"kategorie,omitempty"`
}
