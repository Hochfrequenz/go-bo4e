package com

import "github.com/hochfrequenz/go-bo4e/enum/preisgarantietyp"

// Preisgarantie definiert eine Preisgarantie mit der Möglichkeit verschiedener Ausprägungen.
type Preisgarantie struct {
	// Beschreibung ist ein Freitext zur Beschreibung der Preisgarantie.
	Beschreibung *string `json:"beschreibung,omitempty"`

	// Preisgarantietyp legt fest, auf welche Preisbestandteile die Garantie gewährt wird.
	Preisgarantietyp preisgarantietyp.Preisgarantietyp `json:"preisgarantietyp,omitempty" validate:"required"`

	// ZeitlicheGueltigkeit ist der Zeitraum, bis zu dem die Preisgarantie gilt.
	ZeitlicheGueltigkeit Zeitraum `json:"zeitlicheGueltigkeit,omitempty" validate:"required"`
}
