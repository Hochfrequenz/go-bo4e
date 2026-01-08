package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/aufabschlagstyp"
	"github.com/hochfrequenz/go-bo4e/enum/aufabschlagsziel"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
)

// AufAbschlag modelliert preiserhöhende (Aufschlag) bzw. preisvermindernde (Abschlag)
// Zusatzvereinbarungen, die individuell zu einem neuen oder bestehenden Liefervertrag abgeschlossen wurden.
type AufAbschlag struct {
	// Bezeichnung ist die Bezeichnung des Auf-/Abschlags.
	Bezeichnung string `json:"bezeichnung,omitempty" validate:"required"`

	// Beschreibung ist eine Beschreibung zum Auf-/Abschlag.
	Beschreibung *string `json:"beschreibung,omitempty"`

	// AufAbschlagstyp ist der Typ des Aufabschlages (z.B. absolut oder prozentual).
	AufAbschlagstyp *aufabschlagstyp.AufAbschlagstyp `json:"aufAbschlagstyp,omitempty"`

	// AufAbschlagsziel gibt an, welchem Preis oder welchen Kosten der Auf/Abschlag zugeordnet ist.
	AufAbschlagsziel *aufabschlagsziel.AufAbschlagsziel `json:"aufAbschlagsziel,omitempty"`

	// Einheit gibt an, in welcher Währungseinheit der Auf/Abschlag berechnet wird (nur bei absoluten Aufschlagstypen).
	Einheit *waehrungseinheit.Waehrungseinheit `json:"einheit,omitempty"`

	// Website ist die Internetseite, auf der die Informationen zum Auf-/Abschlag veröffentlicht sind.
	Website *string `json:"website,omitempty"`

	// Gueltigkeitszeitraum ist der Zeitraum, in dem der Abschlag zur Anwendung kommen kann.
	Gueltigkeitszeitraum *Zeitraum `json:"gueltigkeitszeitraum,omitempty"`

	// Staffeln enthält die Werte für die gestaffelten Auf/Abschläge.
	Staffeln []Preisstaffel `json:"staffeln,omitempty"`
}
