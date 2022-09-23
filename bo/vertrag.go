package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsart"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsstatus"
	"time"
)

// Vertrag ist ein Modell für die Abbildung von Vertragsbeziehungen. Das Objekt dient dazu, alle Arten von Verträgen, die in der Energiewirtschaft Verwendung finden, abzubilden.
type Vertrag struct {
	Geschaeftsobjekt
	Vertragsnummer string                        `json:"vertragsnummer,omitempty" validate:"alphanum,required"` // Vertragsnummer ist eine im Verwendungskontext eindeutige Nummer für den Vertrag
	Beschreibung   string                        `json:"beschreibung,omitempty"`                                // Beschreibung zum Vertrag
	Vertragsstatus vertragsstatus.Vertragsstatus `json:"vertragstatus,omitempty" validate:"required"`           // Vertragsstatus ist der Status des Vertrags
	// vertragsstatus serializes as "vertragstatus" with single "s" for compatability.
	Vertragsart         vertragsart.Vertragsart  `json:"vertragsart,omitempty" validate:"required"`                         // Vertragsart legt fest, um welche Art von Vertrag es sich handelt. Z.B. Netznutzungvertrag.
	Sparte              sparte.Sparte            `json:"sparte,omitempty" validate:"required"`                              // Sparte sind Unterscheidungsmöglichkeiten für die Sparte
	Vertragsbeginn      time.Time                `json:"vertragsbeginn,omitempty" validate:"required"`                      // Vertragsbeginn is the inclusive start
	Vertragsende        time.Time                `json:"vertragsende,omitempty" validate:"required,gtfield=Vertragsbeginn"` // Vertragsende is the exclusive end
	Vertragspartner1    Geschaeftspartner        `json:"vertragspartner1,omitempty" validate:"required"`                    // Vertragspartner1 ist der "erstgenannte" Vertragspartner. In der Regel der Aussteller des Vertrags. Beispiel: "Vertrag zwischen Vertragspartner 1 ..."
	Vertragspartner2    Geschaeftspartner        `json:"vertragspartner2,omitempty" validate:"required"`                    // Vertragspartner2 ist der "zweitgenannte" Vertragspartner. In der Regel der Empfänger des Vertrags. Beispiel "Vertrag zwischen Vertragspartner 1 und Vertragspartner 2"
	UnterzeichnerVp1    []com.Unterschrift       `json:"unterzeichnervp1,omitempty"`                                        // UnterzeichnerVp1 ist der Unterzeichner des Vertragspartner1
	UnterzeichnerVp2    []com.Unterschrift       `json:"unterzeichnervp2,omitempty"`                                        // UnterzeichnerVp2 ist der Unterzeichner des Vertragspartner2
	Vertragskonditionen *com.Vertragskonditionen `json:"vertragskonditionen,omitempty"`                                     // Vertragskonditionen ist eine Festlegungen zu Laufzeiten und Kündigungsfristen
	Vertragsteile       []com.Vertragsteil       `json:"vertragsteile,omitempty" validate:"required,min=1"`                 // Vertragsteile sind die Vertragsteile, die dazu verwendet werden, eine vertragliche Leistung in Bezug zu einer Lokation (Markt- oder Messlokation) festzulegen.
}

func (_ Vertrag) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
func (_ Vertrag) GetBoTyp() botyp.BOTyp {
	// this is useful for generics to work
	return botyp.VERTRAG
}
