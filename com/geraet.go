package com

import "github.com/hochfrequenz/go-bo4e/enum/geraeteart"

// Geraet modelliert alle Geräte, die keine Zähler sind.
type Geraet struct {
	// Geraetenummer ist die auf dem Gerät aufgedruckte Nummer, die vom MSB vergeben wird.
	Geraetenummer *string `json:"geraetenummer,omitempty"`

	// Geraeteeigenschaften beschreibt die Eigenschaften des Gerätes, z.B. Wandler MS/NS.
	Geraeteeigenschaften *Geraeteeigenschaften `json:"geraeteeigenschaften,omitempty"`

	// Geraeteart gibt die Art des Gerätes an, z.B. ZAEHLEINRICHTUNG.
	Geraeteart *geraeteart.Geraeteart `json:"geraeteart,omitempty"`
}
