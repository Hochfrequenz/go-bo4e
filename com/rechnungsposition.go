package com

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/enum/bdewartikelnummer"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungscode"
	"github.com/hochfrequenz/go-bo4e/enum/waehrungseinheit"
	"github.com/hochfrequenz/go-bo4e/enum/zeiteinheit"
)

// Rechnungsposition en sind Teil von Rechnung en. In einem Rechnungsteil wird jeweils eine in sich geschlossene LEISTUNG abgerechnet.
type Rechnungsposition struct {
	Positionsnummer   int                                 `json:"positionsnummer,omitempty" validate:"required"`                        // Positionsnummer ist eine fortlaufende Nummer für die Rechnungsposition
	LieferungVon      time.Time                           `json:"lieferungVon,omitempty"`                                               // LieferungVon ist ein _inklusiver_ Start der Lieferung für die abgerechnete LEISTUNG
	LieferungBis      time.Time                           `json:"lieferungBis,omitempty" validate:"omitempty,gtfield=LieferungVon"`     // LieferungBis ist ein _exklusives_ Ende der Lieferung für die abgerechnete LEISTUNG
	Positionstext     string                              `json:"positionstext,omitempty" validate:"required"`                          // Positionstext ist eine Bezeichung für die abgerechnete Position.
	Zeiteinheit       zeiteinheit.Zeiteinheit             `json:"zeiteinheit,omitempty"`                                                // Zeiteinheit wird angegeben, falls sich der Preis auf eine Zeit bezieht, steht hier die Einheit, z.B. JAHR
	Artikelnummer     bdewartikelnummer.BDEWArtikelnummer `json:"artikelnummer,omitempty"`                                              // Artikelnummer ist eine Kennzeichnung der Rechnungsposition mit der Standard-Artikelnummer des BDEW
	LokationsId       string                              `json:"lokationsId,omitempty" validate:"omitempty,min=11,max=11,numeric"`     // LokationsId ist die MarktlokationsId zu der diese Position gehört
	PositionsMenge    Menge                               `json:"positionsMenge,omitempty" validate:"required"`                         // PositionsMenge ist die abgerechnete Menge mit Einheit. Z.B. 4372 kWh
	ZeitbezogeneMenge *Menge                              `json:"zeitbezogeneMenge,omitempty"`                                          // ZeitbezogeneMenge ist eine optionale, auf die Zeiteinheit bezogene Untermenge. Z.B. bei einem Jahrespreis, 3 Monate oder 146 Tage. Basierend darauf wird der Preis aufgeteilt
	Korrekturfaktor   *float32                            `json:"korrekturfaktor,omitempty" validate:"omitempty,min=-1,max=-1,numeric"` // Korrekturfaktor ist ein Faktor -1 der ggf. in Mehrmengen-Rechnungen vorkommt
	Einzelpreis       Preis                               `json:"einzelpreis,omitempty" validate:"required"`                            // Einzelpreis ist der Preis für eine Einheit der energetischen Menge
	TeilsummeNetto    Betrag                              `json:"teilsummeNetto,omitempty" validate:"required"`                         // TeilsummeNetto ist das Ergebnis der Multiplikation aus einzelpreis * positionsMenge * (Faktor aus zeitbezogeneMenge). Z.B. 12,60€ * 120 kW * 3/12 (für 3 Monate).
	TeilsummeSteuer   Steuerbetrag                        `json:"teilsummeSteuer,omitempty" validate:"required"`                        // TeilsummeSteuer ist der auf die Position entfallende Steuer, bestehend aus Steuersatz und Betrag
	TeilrabattNetto   *Betrag                             `json:"teilrabattNetto,omitempty"`                                            // TeilrabattNetto ist der Rabatt für diese Position
	ArtikelId         string                              `json:"artikelId,omitempty"`                                                  // Die ArtikelId ist Artikel-ID (zu verwenden seit 2022-10-01)
	Ausfuehrungsdatum time.Time                           `json:"ausfuehrungsdatum,omitempty"`                                          // Um Qualifier 203 Ausführungsdatum/-zeit abzubilden
}

// RechnungspositionStructLevelValidation does a cross check on a Rechnungsposition object
func RechnungspositionStructLevelValidation(sl validator.StructLevel) {
	rp := sl.Current().Interface().(Rechnungsposition)
	if rp.Einzelpreis.Bezugswert != rp.PositionsMenge.Einheit {
		sl.ReportError(rp.PositionsMenge.Einheit, "Einheit", "PositionsMenge", "PositionsMenge.Einheit==Einzelpreis.Bezugswert", "")
	}
	expectedTeilsummeNetto := Menge{
		Wert:    rp.PositionsMenge.Wert.Mul(rp.Einzelpreis.Wert), // anzahl*preis
		Einheit: rp.PositionsMenge.Einheit,
	}
	if rp.ZeitbezogeneMenge != nil {
		if rp.ZeitbezogeneMenge.Einheit != rp.Einzelpreis.Bezugswert {
			sl.ReportError(rp.ZeitbezogeneMenge.Einheit, "Einheit", "ZeitbezogeneMenge", "ZeitbezogeneMenge.Einheit==Einzelpreis.Bezugswert", "")
			// further checks are not meaningful at this point because there is no implicit conversion like: "1 year = 12 months"
			return
		}
		expectedTeilsummeNetto.Wert = expectedTeilsummeNetto.Wert.Mul(rp.ZeitbezogeneMenge.Wert)
	}
	if rp.TeilsummeNetto.Waehrung != waehrungscode.EUR || rp.Einzelpreis.Einheit != waehrungseinheit.EUR {
		// the bo4e standard is broken at this point.
		// To make a meaningful comparison both the teilsumme which currently uses waehrungsCODE and the Einzelpreis which currently uses waehrungsEINHEIT should refer to the same enum.
		// I emailed this to Peter Martin Schroer. Until he comes up with an official fix, we'll skip the validation.
		// todo: write a bo4e modification proposal: https://github.com/Hochfrequenz/bo4e-modification-proposals/issues/7
		return
	}
	if !expectedTeilsummeNetto.Wert.Equals(rp.TeilsummeNetto.Wert) {
		sl.ReportError(rp.TeilsummeNetto.Wert, "Wert", "TeilsummeNetto", "TeilsummeNetto.Wert==Einzelpreis*Positionsmenge", "")
	}
}
