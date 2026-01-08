package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/erzeugungsart"
	"github.com/shopspring/decimal"
)

// Energieherkunft beschreibt die Herkunft der Energie nach Erzeugungsart.
type Energieherkunft struct {
	// Erzeugungsart ist die Art der Erzeugung der Energie.
	Erzeugungsart *erzeugungsart.Erzeugungsart `json:"erzeugungsart,omitempty"`

	// AnteilProzent ist der prozentuale Anteil der jeweiligen Erzeugungsart.
	AnteilProzent *decimal.Decimal `json:"anteilProzent,omitempty"`
}
