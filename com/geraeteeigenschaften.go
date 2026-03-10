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

	// Wandlerfaktor is the conversion factor for current/voltage transformers
	// https://github.com/Hochfrequenz/BO4E-dotnet/blob/0.54.0/BO4E/COM/Geraeteeigenschaften.cs#L52
	Wandlerfaktor *decimal.Decimal `json:"wandlerfaktor,omitempty"`
}
