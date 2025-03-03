package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/geraetemerkmal"
	"github.com/hochfrequenz/go-bo4e/enum/geraetetyp"
)

// Geraeteeigenschaften ist eine Komponente mit der die Eigenschaften eines Gerätes in Bezug auf den Typ und weitere Merkmale modelliert werden.
type Geraeteeigenschaften struct {
	Geraetetyp     geraetetyp.Geraetetyp          `json:"geraetetyp"`
	Geraetemerkmal *geraetemerkmal.Geraetemerkmal `json:"geraetemerkmal,omitempty"`
}
