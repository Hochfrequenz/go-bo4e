package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/status"
	"github.com/hochfrequenz/go-bo4e/enum/statusart"
)

// StatusZusatzInformation enthält die Auflistung der STS Segmente
// Plausibilisierungshinweis, Ersatzwertbildungsverfahren, Korrekturgrund,
// Gasqualität, Tarif, Grundlage der Energiemenge.
type StatusZusatzInformation struct {
	// StatusArt enthält die Zusatzinformation Art des angegebenen Wertes.
	Art *statusart.StatusArt `json:"art,omitempty"`

	// Status enthält die Zusatzinformation Status des angegebenen Wertes.
	Status status.Status `json:"status,omitempty" validate:"required"`
}
