package com

import "time"

// Aufgabe ist eine Aufgabe als Teil einer Benachrichtigung.
// Aufgaben entsprechen den Klärfall-Lösungsmethoden im SAP.
type Aufgabe struct {
	// AufgabenId ist die eindeutige Kennzeichnung der Aufgabe.
	AufgabenId string `json:"aufgabenId,omitempty" validate:"required"`

	// Beschreibung ist eine optionale Beschreibung der Aufgabe.
	Beschreibung *string `json:"beschreibung,omitempty"`

	// Deadline ist die optionale Deadline bis zu der die Aufgabe ausgeführt werden kann
	// oder ihre Ausführung sinnvoll ist.
	Deadline *time.Time `json:"deadline,omitempty"`

	// Ausgefuehrt gibt an, ob diese Aufgabe schon ausgeführt wurde (true)
	// oder noch zur Bearbeitung ansteht (false).
	Ausgefuehrt bool `json:"ausgefuehrt,omitempty"`

	// Ausfuehrungszeitpunkt ist der Zeitpunkt, zu dem die Aufgabe ausgeführt wurde.
	// Nur sinnvoll, wenn Ausgefuehrt==true.
	Ausfuehrungszeitpunkt *time.Time `json:"ausfuehrungszeitpunkt,omitempty"`

	// Ausfuehrender ist die eindeutige Kennung des Benutzers, der diese Aufgabe ausgeführt hat.
	// Nur sinnvoll, wenn Ausgefuehrt==true.
	Ausfuehrender *string `json:"ausfuehrender,omitempty"`
}
