package bo

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/berechnungsformelnotwendigkeit"
	"github.com/hochfrequenz/go-bo4e/enum/verwendungszweck"
)

// Berechnungsformel bildet ein BO zur Abbildung von Berechnungsformeln ab.
type Berechnungsformel struct {
	Geschaeftsobjekt
	// Beginndatum ist das Beginndatum der Gültigkeit der Formel.
	Beginndatum time.Time `json:"beginndatum,omitempty"`

	// Notwendigkeit gibt an, ob eine Berechnungsformel erforderlich ist.
	Notwendigkeit berechnungsformelnotwendigkeit.BerechnungsformelNotwendigkeit `json:"notwendigkeit,omitempty"`

	// RechenschrittId ist die ID des Rechenschritts.
	RechenschrittId *int `json:"rechenschrittId,omitempty"`

	// Verwendungszweck gibt den Verwendungszweck der Berechnungsformel an.
	Verwendungszweck *verwendungszweck.Verwendungszweck `json:"verwendungszweck,omitempty"`

	// Rechenschritt enthält den zugehörigen Rechenschritt.
	Rechenschritt *com.Rechenschritt `json:"rechenschritt,omitempty"`
}
