package bo

import (
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungsart"
	"github.com/hochfrequenz/go-bo4e/enum/nnrechnungstyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Netznutzungsrechnung models grid usage invoices
type Netznutzungsrechnung struct {
	Rechnung
	Sparte               sparte.Sparte                 `json:"sparte" validate:"required"`                                        // Sparte (STROM, GAS ...) für die die Rechnung ausgestellt ist
	Absendercodenummer   string                        `json:"absendercodenummer" validate:"required,numeric,min=13,max=13"`      // Absendercodenummer ist die Rollencodenummer des Absenders (Rechnung.Rechnungsersteller). Über die Nummer können weitere Informationen zum Marktteilnehmer ermittelt werden.
	Empfaengercodenummer string                        `json:"empfaengercodenummer" validate:"required,numeric,min=13,max=13"`    // Absendercodenummer ist die Rollencodenummer des Empfängers (Rechnung.Rechnungsempfaenger). Über die Nummer können weitere Informationen zum Marktteilnehmer ermittelt werden.
	Nnrechnungsart       nnrechnungsart.NNRechnungsart `json:"nnrechnungsart" validate:"required"`                                // Aus der INVOIC entnommen
	Nnrechnungstyp       nnrechnungstyp.NNRechnungstyp `json:"nnrechnungstyp" validate:"required"`                                // Aus der INVOIC entnommen
	Original             bool                          `json:"original"`                                                          // Original ist ein Kennzeichen, ob es sich um ein Original (true) oder eine Kopie handelt (false).
	Simuliert            bool                          `json:"simuliert"`                                                         // Simuliert ist ein Kennzeichen, ob es sich um eine simulierte Rechnung, z.B. zur Rechnungsprüfung handelt.
	LokationsId          string                        `json:"lokationsId,omitempty" validate:"omitempty,alphanum,max=33,min=11"` // LokationsId ist die Markt- oder Messlokations-Identifikation (als Malo/Melo-Id) der Lokation, auf die sich die Rechnung bezieht.
}
