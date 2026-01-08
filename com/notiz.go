package com

import "time"

// Notiz enthält beliebige, unstrukturierte Zusatzinformationen
type Notiz struct {
	// Autor ist die Person oder das System, das die Notiz angelegt hat.
	Autor string `json:"autor,omitempty"`

	// Zeitpunkt ist der Zeitpunkt, zu dem die Notiz angelegt wurde.
	Zeitpunkt time.Time `json:"zeitpunkt,omitempty"`

	// Inhalt ist der Inhalt der Notiz (Freitext).
	Inhalt string `json:"inhalt,omitempty"`
}
