package com

import "time"

// Vorauszahlung wird dazu verwendet Vorauszahlungen einer Rechnung abzubilden (inkl. Referenz auf die bezahlte Rechnung)
type Vorauszahlung struct {
	Betrag        Betrag     `json:"betrag"`                  // Betrag gibt den Betrag des Preises an
	Referenz      *string    `json:"referenz,omitempty"`      // Referenz auf die Rechnungsnummer, die durch diesen Betrag bezahlt wurde
	ReferenzDatum *time.Time `json:"referenzDatum,omitempty"` // ReferenzDatum ist die Referenz auf das Datum der Rechnung, die durch diesen Betrag bezahlt wurde
}
