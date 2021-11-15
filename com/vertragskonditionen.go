package com

// Vertragskonditionen ist eine Komponente zur Abbildung von Vertragskonditionen. Die Komponente wird sowohl im Vertrag als auch im Tarif verwendet.
type Vertragskonditionen struct {
	Beschreibung          string   `json:"beschreibung,omitempty"`          // Beschreibung ist ein Freitext zur Beschreibung der Konditionen, z.B. "Standardkonditionen Gas"
	AnzahlAbschlaege      int      `json:"anzahlAbschlaege,omitempty"`      // AnzahlAbschlaege enthält die Anzahl der vereinbarten Abschläge pro Jahr, z.B. 12
	Vertragslaufzeit      Zeitraum `json:"vertragslaufzeit,omitempty"`      // Vertragslaufzeit enthält den Zeitraum, über den dieser Vertrag läuft
	Kuendigungsfrist      Zeitraum `json:"kuendigungsfrist,omitempty"`      // Kuendigungsfrist ist die Frist, innerhalb derer der Vertrag gekündigt werden kann
	Vertragsverlaengerung Zeitraum `json:"vertragsverlaengerung,omitempty"` // Vertragsverlaengerung beschreibt, dass, falls der Vertrag nicht gekündigt wird, er sich automatisch um die angegebene Zeit verlängert
	Abschlagszyklus       Zeitraum `json:"abschlagszyklus,omitempty"`       // Abschlagszyklus sind die Zyklen, in denen Abschläge erstellt werden. Alternativ kann auch die Anzahl in den Konditionen angeben werden.
}
