package com

// Die Komponente wird dazu verwendet die Ursache bzw. Herkunft von Fehlern genauer zu spezifizieren
type FehlerUrsache struct {
	Dokument     *string `json:"dokument,omitempty"`     //
	Nachricht    *string `json:"nachricht,omitempty"`    //
	Transaktion  *string `json:"transaktion,omitempty"`  //
	Gruppe       *string `json:"gruppe,omitempty"`       //
	Segment      *string `json:"segment,omitempty"`      //
	Beschreibung *string `json:"beschreibung,omitempty"` //
}
