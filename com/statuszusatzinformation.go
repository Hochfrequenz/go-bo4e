package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/statusart"
	"github.com/hochfrequenz/go-bo4e/enum/statuszusatz"
)

// Statuszusatzinformation Enthält die Auflistung der STS Segmente Plausibilisierungshinweis, Ersatzwertbildungsverfahren,
// Korrekturgrund, Gasqualität, Tarif, Grundlage der Energiemenge.
// Statuszusatzinformation ist inoffiziell
type Statuszusatzinformation struct {
	Art    *statusart.Statusart       `json:"art,omitempty" validate:"omitempty"`    // Art enthält die Zusatzinformation Art des angegebenen Wertes
	Status *statuszusatz.Statuszusatz `json:"status,omitempty" validate:"omitempty"` // Status Enthält die Zusatzinformation Status des angegebenen Wertes
}
