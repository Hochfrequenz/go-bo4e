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
	Bezeichnung          string                            `json:"bezeichnung" validate:"required"`          // Bezeichnung des Auf-/Abschlags
	Beschreibung         string                            `json:"beschreibung" validate:"required"`         // Beschreibung zum Auf-/Abschlags
	AufAbschlagstyp      aufabschlagstyp.AufAbschlagstyp   `json:"aufAbschlagstyp" validate:"required"`      // Typ des AufAbschlages
	AufAbschlagswert     decimal.NullDecimal               `json:"aufAbschlagswert" validate:"required"`     // Höhe des AufAbschlages
	AufAbschlagswaehrung waehrungseinheit.Waehrungseinheit `json:"aufAbschlagswaehrung" validate:"required"` //  Einheit, in der der Auf-/Abschlag angegeben ist (z.B. ct/kWh)
}
