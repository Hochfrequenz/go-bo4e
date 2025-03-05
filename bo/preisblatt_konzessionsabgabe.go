package bo

import "github.com/hochfrequenz/go-bo4e/enum/kundengruppeKA"

// PreisblattKonzessionsabgabe ist eine Preisblatt-Variante mit zusätzlichen Merkmalen für die Kundengruppe
type PreisblattKonzessionsabgabe struct {
	Preisblatt
	KundengruppeKA kundengruppeKA.KundengruppeKA
}
