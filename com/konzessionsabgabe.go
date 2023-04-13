package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/abgabeart"
	"github.com/shopspring/decimal"
)

type Konzessionsabgabe struct {
	Satz      abgabeart.AbgabeArt
	Kosten    *decimal.Decimal
	Kategorie *string
}
