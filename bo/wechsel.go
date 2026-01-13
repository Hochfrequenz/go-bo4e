package bo

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Wechsel modelliert den Wechsel eines Gerätes oder Zählers.
type Wechsel struct {
	Geschaeftsobjekt
	// Sparte gibt an, ob es sich um Strom oder Gas handelt.
	Sparte sparte.Sparte `json:"sparte,omitempty"`

	// Geraete ist die Liste der Geräte, die gewechselt werden.
	Geraete []com.Geraet `json:"geraete,omitempty"`

	// Wechseldatum gibt an, wann der Wechsel (voraussichtlich) stattfinden wird.
	Wechseldatum *time.Time `json:"wechseldatum,omitempty"`

	// Vollstaendig gibt an, ob es sich um einen vollständigen Wechsel handelt.
	Vollstaendig *bool `json:"vollstaendig,omitempty"`
}
