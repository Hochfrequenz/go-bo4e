package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/geraetemerkmal"
	"github.com/hochfrequenz/go-bo4e/enum/geraetetyp"
	"github.com/shopspring/decimal"
)

// Geraeteeigenschaften ist eine Komponente mit der die Eigenschaften eines Gerätes in Bezug auf den Typ und weitere Merkmale modelliert werden.
type Geraeteeigenschaften struct {
	Geraetetyp     geraetetyp.Geraetetyp          `json:"geraetetyp"`
	Geraetemerkmal *geraetemerkmal.Geraetemerkmal `json:"geraetemerkmal,omitempty"`

	// Wandlerfaktor für Strom-/Spannungswandler.
	Wandlerfaktor *decimal.Decimal `json:"wandlerfaktor,omitempty"`

	// WeitereGeraetenummern: Weitere Gerätenummern (z.B. von Wandlern).
	WeitereGeraetenummern []string `json:"weitereGeraetenummern,omitempty"`
}
