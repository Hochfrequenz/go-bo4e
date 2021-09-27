package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungsstatus"
	"github.com/hochfrequenz/go-bo4e/enum/rechnungstyp"
	"time"
)

// Rechnung ist ein Modell für die Abbildung von Rechnungen im Kontext der Energiewirtschaft. Ausgehend von diesem Basismodell werden weitere spezifische Formen abgeleitet.
type Rechnung struct {
	BusinessObject
	Rechnungstitel          string                          `json:"rechnungstitel"`                                             // Bezeichnung für die vorliegende Rechnung.
	Rechnungsstatus         rechnungsstatus.Rechnungsstatus `json:"rechnungsstatus"`                                            // Status der Rechnung zur Kennzeichnung des Bearbeitungsstandes
	Storno                  bool                            `json:"storno" validate:"required"`                                 // Kennzeichnung, ob es sich um eine Stornorechnung handelt. Im Falle "true" findet sich im Attribut "originalrechnungsnummer" die Nummer der Originalrechnung
	Rechnungsnummer         string                          `json:"rechnungsnummer" validate:"required"`                        // Eine im Verwendungskontext eindeutige Nummer für die Rechnung
	Rechnungsdatum          time.Time                       `json:"rechnungsdatum" validate:"required"`                         // Ausstellungsdatum der Rechnung
	Faelligkeitsdatum       time.Time                       `json:"faelligkeitsdatum" validate:"required"`                      // Zu diesem Datum ist die Zahlung fällig
	Rechnungstyp            rechnungstyp.Rechnungstyp       `json:"rechnungstyp" validate:"required"`                           // Ein kontextbezogender Rechnungstyp
	OriginalRechnungsnummer *string                         `json:"originalRechnungsnummer" validate:"required_if=Storno true"` // Im Falle einer Stornorechnung (Storno = true) steht hier die Rechnungsnummer der stornierten Rechnung
	Rechnungsperiode        com.Zeitraum                    `json:"rechnungsperiode" validate:"required"`                       // Der Zeitraum der zugrunde liegenden Lieferung zur Rechnung
	Rechnungsersteller      Geschaeftspartner               `json:"rechnungsersteller" validate:"required"`                     // Der Aussteller der Rechnung
	Rechnungsempfaenger     Geschaeftspartner               `json:"rechnungsempfaenger" validate:"required"`                    // Der Empfänger der Rechnung
	GesamtNetto             com.Betrag                      `json:"gesamtnetto" validate:"required"`                            // Die Summe der Nettobeträge der Rechnungsteile // todo: structlevel validation
	GesamtSteuer            com.Betrag                      `json:"gesamtsteuer" validate:"required"`                           // Die Summe der Steuerbeträge der Rechnungsteile // todo: structlevel validation
	GesamtBrutto            com.Betrag                      `json:"gesamtbrutto" validate:"required"`                           // Die Summe aus Netto- und Steuerbeträge // todo: structlevel validation
	Vorausgezahlt           *com.Betrag                     `json:"vorausgezahlt"`                                              // Die Summe evtl. vorausgezahlter Beträge, z.B. Abschläge. Angabe als Bruttowert.
	RabattBrutto            *com.Betrag                     `json:"rabattBrutto"`                                               // Gesamtrabatt auf den Bruttobetrag
	Zuzahlen                com.Betrag                      `json:"zuzahlen" validate:"required"`                               // Der zu zahlende Betrag, der sich aus (gesamtbrutto - vorausbezahlt - rabattBrutto) ergibt // todo: struct level validation
	Steuerbetraege          []com.Steuerbetrag              `json:"steuerbetraege"`                                             // Eine Liste mit Steuerbeträgen pro Steuerkennzeichen/Steuersatz. Die Summe dieser Beträge ergibt den Wert für GesamtSteuer. // todo: structlevel validation
	Rechnungspositionen     []com.Rechnungsposition         `json:"rechnungspositionen" validate:"required,min=1"`              // Die Rechnungspositionen
}
