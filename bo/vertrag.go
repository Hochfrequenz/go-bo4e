package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsart"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsstatus"
	"time"
)

// Vertrag ist ein Modell für die Abbildung von Vertragsbeziehungen. Das Objekt dient dazu, alle Arten von Verträgen, die in der Energiewirtschaft Verwendung finden, abzubilden.
type Vertrag struct {
	BusinessObject
	Vertragsnummer      string                        `json:"vertragsnummer" validate:"alphanum,required"` // Eine im Verwendungskontext eindeutige Nummer für den Vertrag
	Beschreibung        string                        `json:"beschreibung,omitempty"`                      // Beschreibung zum Vertrag
	Vertragsstatus      vertragsstatus.Vertragsstatus `json:"vertragsstatus" validate:"required"`          //Gibt den Status des Vertrags an
	Vertragsart         vertragsart.Vertragsart       `json:"vertragsart" validate:"required"`             // Hier ist festgelegt, um welche Art von Vertrag es sich handelt. Z.B. Netznutzungvertrag.
	Sparte              sparte.Sparte                 `json:"sparte" validate:"required"`                  // Unterscheidungsmöglichkeiten für die Sparte
	Vertragsbeginn      time.Time                     `json:"startdatum" validate:"required"`              // inclusive start
	Vertragsende        time.Time                     `json:"enddatum" validate:"required"`                // exclusive end
	Vertragspartner1    string                        //	Der "erstgenannte" Vertragspartner. In der Regel der Aussteller des Vertrags. Beispiel: "Vertrag zwischen Vertragspartner 1 ..." Siehe BO Geschaeftspartner
	Vertragspartner2    string                        // Der "zweitgenannte" Vertragspartner. In der Regel der Empfänger des Vertrags. Beispiel "Vertrag zwischen Vertragspartner 1 und Vertragspartner 2". Siehe BO Geschaeftspartner
	UnterzeichnerVp1    []com.Unterschrift            `json:"unterzeichnervp1,omitempty"`              //	Unterzeichner des Vertragspartner1
	UnterzeichnerVp2    []com.Unterschrift            `json:"unterzeichnervp2,omitempty"`              //	Unterzeichner des Vertragspartner2
	Vertragskonditionen *com.Vertragskonditionen      `json:"vertragskonditionen,omitempty"`           // Festlegungen zu Laufzeiten und Kündigungsfristen
	Vertragsteile       []com.Vertragsteil            `json:"vertragsteile" validate:"required,min=1"` // Der Vertragsteil wird dazu verwendet, eine vertragliche Leistung in Bezug zu einer Lokation (Markt- oder Messlokation) festzulegen.
}
