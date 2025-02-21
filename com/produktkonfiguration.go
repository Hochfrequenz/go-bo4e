package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/produktcode"
)

// Produktkonfiguration als einzelnes Bestandteil eines Produktpakets.
type Produktkonfiguration struct {
	Code              produktcode.Produktcode `json:"code,omitempty"`              // Eindeutiger Code der Konfiguration
	Eigenschaft       produktcode.Produktcode `json:"eigenschaft,omitempty"`       // Eigenschaftswert zur Konfiguration (als Code)
	Zusatzeigenschaft string                  `json:"zusatzeigenschaft,omitempty"` // Zusätzlicher Eigenschaftswert, z.B. Angabe der Jahresverbrauchsprognose (4000). Im Allgemeinen zur Angabe von Werten, die nicht als Produktcode definiert sind
}
