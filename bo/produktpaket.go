package bo

import "github.com/hochfrequenz/go-bo4e/com"

// The Produktpaket ist eine Sammlung von Produktkonfigurationen im Rahmen der Marktkommunikation.
type Produktpaket struct {
	Geschaeftsobjekt
	// PaketId ist zur Paket-Identifikation (Durchnummerierung).
	PaketId int `json:"paketId,omitempty"`

	// Konfigurationen ist eine Liste an Produktkonfigurationen
	Konfigurationen []com.Produktkonfiguration `json:"konfigurationen,omitempty"`

	// Prioritaet gibt die Prioritaet des Pakets an von 1-5 (1 ist die hoechste Prioritaet)
	Prioritaet int `json:"prioritaet,omitempty"`

	// MussVollstaendigSein gibt an, ob das Paket vollstaendig umgesetzt sein werden muss
	MussVollstaendigSein bool `json:"mussVollstaendigSein,omitempty"`
}
