package com

import (
	"time"
)

type AvisPosition struct {
	RechnungsNummer      string                  `json:"rechnungsNummer,omitempty" validate:"required"` // Die Rechnungsnummer der Rechnung, auf die sich das Avis bezieht
	RechnungsDatum       time.Time               `json:"rechnungsDatum,omitempty" validate:"required"`  // Das Rechnungsdatum der Rechnung, auf die sich das Avis bezieht
	Storno               bool                    `json:"istStorno"`                                     // Kennzeichnung, ob es sich bei der Rechnung auf die sich das Avis bezieht, um eine Stornorechnung handelt
	IstSelbstausgestellt *bool                   `json:"istSelbstausgestellt,omitempty"`                // Kennzeichnung, ob es sich bei der Rechnung auf die sich das Avis bezieht, um eine selbst aus gestellte Rechnung handelt
	GesamtBrutto         Betrag                  `json:"gesamtBrutto,omitempty" validate:"required"`    // Überweisungsbetrag
	ZuZahlen             Betrag                  `json:"zuZahlen,omitempty" validate:"required"`        // Geforderter Rechnungsbetrag
	Abweichungen         []Abweichung            `json:"abweichungen,omitempty"`                        // Abweichungen ist eine Liste der Gründe der Ablehnung einer INVOIC oder COMDIS
	Positionen           []Rueckmeldungsposition `json:"positionen,omitempty"`                          // Rückmeldungspositionen
}
