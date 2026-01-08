package bo

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Auftrag beschreibt einen Vorgang, der von einem anderen Marktpartner auszuführen ist.
// In C# ist dies eine abstrakte Klasse - in Go kann diese Struktur eingebettet werden.
type Auftrag struct {
	Geschaeftsobjekt
	// Ausfuehrungsdatum beschreibt, zu welchem Zeitpunkt ein Auftrag ausgeführt werden soll.
	Ausfuehrungsdatum *time.Time `json:"ausfuehrungsdatum,omitempty"`

	// Fertigstellungsdatum beschreibt, zu welchem Zeitpunkt ein Auftrag ausgeführt wurde/wird.
	Fertigstellungsdatum *time.Time `json:"fertigstellungsdatum,omitempty"`

	// Sparte gibt die Sparte an, in der der Auftrag relevant ist.
	Sparte *sparte.Sparte `json:"sparte,omitempty"`

	// Lieferanschrift ist die Adresse, die sich in Belieferung befindet.
	Lieferanschrift *com.Adresse `json:"lieferanschrift,omitempty"`

	// MarktlokationsId ist die ID der Marktlokation, der der betroffene Zähler zugeordnet ist.
	MarktlokationsId *string `json:"marktlokationsId,omitempty"`

	// Bemerkungen enthält zusätzliche Freitexte.
	Bemerkungen []string `json:"bemerkungen,omitempty"`

	// Mindestpreis ist der Mindestpreis eines Auftrags (z.B. für eine Sperrung).
	Mindestpreis *com.Preis `json:"mindestpreis,omitempty"`

	// Hoechstpreis ist der Höchstpreis eines Auftrags (z.B. für eine Sperrung).
	Hoechstpreis *com.Preis `json:"hoechstpreis,omitempty"`
}
