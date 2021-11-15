package bo

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungstyp"
	"github.com/shopspring/decimal"
	"time"
)

// Rechnung ist ein Modell für die Abbildung von Rechnungen im Kontext der Energiewirtschaft. Ausgehend von diesem Basismodell werden weitere spezifische Formen abgeleitet.
type Rechnung struct {
	Geschaeftsobjekt
	Rechnungstitel          string                          `json:"rechnungstitel"`                                             // Rechnungstitel ist die Bezeichnung für die vorliegende Rechnung.
	Rechnungsstatus         rechnungsstatus.Rechnungsstatus `json:"rechnungsstatus,omitempty"`                                  // Rechnungsstatus ist der Status der Rechnung zur Kennzeichnung des Bearbeitungsstandes
	Storno                  bool                            `json:"storno" validate:"required"`                                 // Storno ist eine Kennzeichnung, ob es sich um eine Stornorechnung handelt. Im Falle "true" findet sich im Attribut "originalrechnungsnummer" die Nummer der Originalrechnung
	Rechnungsnummer         string                          `json:"rechnungsnummer" validate:"required"`                        // Rechnungsnummer ist eine im Verwendungskontext eindeutige Nummer für die Rechnung
	Rechnungsdatum          time.Time                       `json:"rechnungsdatum" validate:"required"`                         // Rechnungsdatum ist das Ausstellungsdatum der Rechnung
	Faelligkeitsdatum       time.Time                       `json:"faelligkeitsdatum" validate:"required"`                      // Faelligkeitsdatum ist das Datum, zu dem die Zahlung fällig ist
	Rechnungstyp            rechnungstyp.Rechnungstyp       `json:"rechnungstyp,omitempty" validate:"required"`                 // Rechnungstyp ist ein kontextbezogener Rechnungstyp
	OriginalRechnungsnummer string                          `json:"originalRechnungsnummer" validate:"required_if=Storno true"` // OriginalRechnungsnummer: Im Falle einer Stornorechnung (Storno = true) steht hier die Rechnungsnummer der stornierten Rechnung
	Rechnungsperiode        com.Zeitraum                    `json:"rechnungsperiode" validate:"required"`                       // Rechnungsperiode ist der Zeitraum der zugrunde liegenden Lieferung zur Rechnung
	Rechnungsersteller      Geschaeftspartner               `json:"rechnungsersteller" validate:"required"`                     // Rechnungsersteller ist der Aussteller der Rechnung
	Rechnungsempfaenger     Geschaeftspartner               `json:"rechnungsempfaenger" validate:"required"`                    // Rechnungsempfaenger ist der Empfänger der Rechnung
	GesamtNetto             com.Betrag                      `json:"gesamtnetto" validate:"required"`                            // GesamtNetto ist die Summe der Nettobeträge der Rechnungsteile
	GesamtSteuer            com.Betrag                      `json:"gesamtsteuer" validate:"required"`                           // GesamtSteuer ist die Summe der Steuerbeträge der Rechnungsteile
	GesamtBrutto            com.Betrag                      `json:"gesamtbrutto" validate:"required"`                           // GesamtBrutto ist die Summe aus Netto- und Steuerbeträge
	Vorausgezahlt           *com.Betrag                     `json:"vorausgezahlt"`                                              // Vorausgezahlt ist die Summe eventuell vorausbezahlter Beträge, z.B. Abschläge. Angabe als Bruttowert.
	RabattBrutto            *com.Betrag                     `json:"rabattBrutto"`                                               // RabattBrutto ist der Gesamtrabatt auf den Bruttobetrag
	Zuzahlen                com.Betrag                      `json:"zuzahlen" validate:"required"`                               // Zuzahlen ist der zu zahlende Betrag, der sich aus (gesamtbrutto - vorausbezahlt - rabattBrutto) ergibt
	Steuerbetraege          []com.Steuerbetrag              `json:"steuerbetraege"`                                             // Steuerbetraege ist eine Liste mit Steuerbeträgen pro Steuerkennzeichen/Steuersatz. Die Summe dieser Beträge ergibt den Wert für GesamtSteuer
	Rechnungspositionen     []com.Rechnungsposition         `json:"rechnungspositionen" validate:"required,min=1"`              // Rechnungspositionen sind die einzelnen Rechnungsposition en.
}

// RechnungStructLevelValidation combines all the single validators
func RechnungStructLevelValidation(sl validator.StructLevel) {
	RechnungStructLevelValidationGesamtNetto(sl)
	RechnungStructLevelValidationGesamtSteuer(sl)
	RechnungStructLevelValidationGesamtBrutto(sl)
	RechnungStructLevelValidationZuZahlen(sl)
}

// RechnungStructLevelValidationGesamtNetto verifies that the sum of all Rechnungsposition.Netto equals the Rechnung.GesamtNetto
func RechnungStructLevelValidationGesamtNetto(sl validator.StructLevel) {
	rechnung := sl.Current().Interface().(Rechnung)
	if rechnung.Rechnungspositionen != nil {
		expectedGesamtNetto := com.Betrag{
			Wert:     decimal.Zero,
			Waehrung: 0,
		}
		for index, rp := range rechnung.Rechnungspositionen {
			if index == 0 {
				// by default use the waehrung of the first entry
				expectedGesamtNetto.Waehrung = rp.TeilsummeNetto.Waehrung
			} else if expectedGesamtNetto.Waehrung != rp.TeilsummeNetto.Waehrung {
				// the waehrung has to be consistent over all Rechnungspositionen; Otherwise adding stuff doesn't make sense
				sl.ReportError(rp.TeilsummeNetto.Waehrung, "Waehrung", "Rechnungspositionen", "Rechnungspositionen[0].Waehrung==Rechnungspositionen[j].Waehrung", "")
				return
			}
			expectedGesamtNetto.Wert = expectedGesamtNetto.Wert.Add(rp.TeilsummeNetto.Wert)
		}
		if expectedGesamtNetto.Waehrung != rechnung.GesamtNetto.Waehrung || !expectedGesamtNetto.Wert.Equals(rechnung.GesamtNetto.Wert) {
			sl.ReportError(rechnung.GesamtNetto, "Wert", "GesamtNetto", "GesamtNetto==sum(TeilsummeNetto)", "")
		}
	}
}

// RechnungStructLevelValidationZuZahlen verifies that Rechnung.Zuzahlen = Rechnung.GesamtBrutto - Rechnung.Vorausgezahlt - Rechnung.RabattBrutto
func RechnungStructLevelValidationZuZahlen(sl validator.StructLevel) {
	rechnung := sl.Current().Interface().(Rechnung)
	expectedZuZahlen := com.Betrag{
		Wert:     rechnung.GesamtBrutto.Wert,
		Waehrung: rechnung.GesamtBrutto.Waehrung,
	}
	if rechnung.Vorausgezahlt != nil {
		if rechnung.Vorausgezahlt.Waehrung != expectedZuZahlen.Waehrung {
			sl.ReportError(rechnung.Vorausgezahlt, "Waehrung", "Vorausgezahlt", "Vorausgezahlt.Waehrung==GesamtBrutto.Waehrung", "")
			return
		}
		expectedZuZahlen.Wert = expectedZuZahlen.Wert.Sub(rechnung.Vorausgezahlt.Wert)
	}
	if rechnung.RabattBrutto != nil {
		if rechnung.RabattBrutto.Waehrung != expectedZuZahlen.Waehrung {
			sl.ReportError(rechnung.Vorausgezahlt, "Waehrung", "RabattBrutto", "RabattBrutto.Waehrung==GesamtBrutto.Waehrung", "")
			return
		}
		expectedZuZahlen.Wert = expectedZuZahlen.Wert.Sub(rechnung.RabattBrutto.Wert)
	}
	if expectedZuZahlen.Waehrung != rechnung.Zuzahlen.Waehrung || !expectedZuZahlen.Wert.Equals(rechnung.Zuzahlen.Wert) {
		sl.ReportError(rechnung.Zuzahlen, "Wert", "Zuzahlen", "Zuzahlen==GesamtBrutto-Rechnung.Vorausgezahlt-Rechnung.RabattBrutto", "")
	}
}

// RechnungStructLevelValidationGesamtBrutto verifies that the sum of all Rechnungsposition. equals the Rechnung.GesamtBrutto
func RechnungStructLevelValidationGesamtBrutto(sl validator.StructLevel) {
	rechnung := sl.Current().Interface().(Rechnung)
	if rechnung.Rechnungspositionen != nil {
		expectedGesamtBrutto := com.Betrag{
			Wert:     decimal.Zero,
			Waehrung: 0,
		}
		for index, rp := range rechnung.Rechnungspositionen {
			if index == 0 {
				// by default use the waehrung of the first entry
				expectedGesamtBrutto.Waehrung = rp.TeilsummeSteuer.Waehrung
			} else if expectedGesamtBrutto.Waehrung != rp.TeilsummeSteuer.Waehrung {
				// the waehrung has to be consistent over all Rechnungspositionen; Otherwise adding stuff doesn't make sense
				sl.ReportError(rp.TeilsummeSteuer.Waehrung, "Waehrung", "Rechnungspositionen", "Rechnungspositionen[0].Waehrung==Rechnungspositionen[j].Waehrung", "")
				return
			}
			expectedGesamtBrutto.Wert = expectedGesamtBrutto.Wert.Add(rp.TeilsummeSteuer.Basiswert).Add(rp.TeilsummeSteuer.Steuerwert)
		}
		if expectedGesamtBrutto.Waehrung != rechnung.GesamtBrutto.Waehrung || !expectedGesamtBrutto.Wert.Equals(rechnung.GesamtBrutto.Wert) {
			sl.ReportError(rechnung.GesamtBrutto, "Wert", "GesamtBrutto", "GesamtBrutto==sum(TeilsummeSteuer)", "")
		}
	}
}

// RechnungStructLevelValidationGesamtSteuer verifies that the sum of all Rechnungsposition.Netto equals the Rechnung.GesamtSteuer
func RechnungStructLevelValidationGesamtSteuer(sl validator.StructLevel) {
	rechnung := sl.Current().Interface().(Rechnung)
	expectedGesamtSteuer := com.Betrag{
		Wert:     decimal.Zero,
		Waehrung: 0,
	}
	if rechnung.Steuerbetraege != nil {
		for index, sb := range rechnung.Steuerbetraege {
			if index == 0 {
				// by default use the waehrung of the first entry
				expectedGesamtSteuer.Waehrung = sb.Waehrung

			} else if expectedGesamtSteuer.Waehrung != sb.Waehrung {
				// the waehrung has to be consistent over all Steuerbetraege; Otherwise adding stuff doesn't make sense
				sl.ReportError(sb.Waehrung, "Waehrung", "Steuerbetraege", "Steuerbetraege[0].Waehrung==Steuerbetraege[j].Waehrung", "")
				return
			}
			expectedGesamtSteuer.Wert = expectedGesamtSteuer.Wert.Add(sb.Steuerwert)
		}
		if expectedGesamtSteuer.Waehrung != rechnung.GesamtSteuer.Waehrung || !expectedGesamtSteuer.Wert.Equals(rechnung.GesamtSteuer.Wert) {
			sl.ReportError(rechnung.GesamtSteuer, "Wert", "GesamtSteuer", "GesamtSteuer==sum(Steuerbetraege)", "")
		}
	}
}
