package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/aufabschlagstyp"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/shopspring/decimal"
)

// PositionsAufAbschlag Differenzierung der zu betrachtenden Produkte anhand der preiserhöhenden (Aufschlag) bzw. preisvermindernden
// (Abschlag) Zusatzvereinbarungen, die individuell zu einem neuen oder bestehenden Liefervertrag abgeschlossen werden
// können. Es können mehrere Auf-/Abschläge gleichzeitig ausgewählt werden.
type PositionsAufAbschlag struct {
	Bezeichnung          string                            `json:"bezeichnung,omitempty" validate:"required"`          // Bezeichnung des Auf-/Abschlags
	Beschreibung         string                            `json:"beschreibung,omitempty" validate:"required"`         // Beschreibung zum Auf-/Abschlags
	AufAbschlagstyp      aufabschlagstyp.AufAbschlagstyp   `json:"aufabschlagstyp,omitempty" validate:"required"`      // Typ des AufAbschlages
	AufAbschlagswert     decimal.NullDecimal               `json:"aufabschlagswert,omitempty" validate:"required"`     // Höhe des AufAbschlages
	AufAbschlagswaehrung waehrungseinheit.Waehrungseinheit `json:"aufabschlagswaehrung,omitempty" validate:"required"` //  Einheit, in der der Auf-/Abschlag angegeben ist (z.B. ct/kWh)
}
