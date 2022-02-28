package com

import (
	"time"
)

type AvisPosition struct {
	RechnungsNummer string     `json:"rechnungsNummer,omitempty" validate:"required"` // Die Rechnungsnummer der Rechnung, auf die sich das Avis bezieht.
	RechnungsDatum  time.Time  `json:"rechnungsDatum,omitempty" validate:"required"`  // Das Rechnungsdatum der Rechnung, auf die sich das Avis bezieht.
	Storno          bool       `json:"istStorno"`                                     // Kennzeichnung, ob es sich bei der Rechnung auf die sich das Avis bezieht, um eine Stornorechnung handelt.
	GesamtBrutto    Betrag     `json:"gesamtBrutto,omitempty" validate:"required"`    // Überweisungsbetrag
	ZuZahlen        Betrag     `json:"zuZahlen,omitempty" validate:"required"`        // Geforderter Rechnungsbetrag
	Abweichung      Abweichung `json:"abweichung,omitempty"`                          // Abweichung bei Ablehnung einer COMDIS
}
