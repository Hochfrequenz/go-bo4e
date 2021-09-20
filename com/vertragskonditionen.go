package com

// Vertragskonditionen ist eine Komponente zur Abbildung von Vertragskonditionen. Die Komponente wird sowohl im Vertrag als auch im Tarif verwendet.
type Vertragskonditionen struct {
	Beschreibung          string   `json:"beschreibung,omitempty"`          // Freitext zur Beschreibung der Konditionen, z.B. "Standardkonditionen Gas"
	AnzahlAbschlaege      int      `json:"anzahlAbschlaege,omitempty"`      // Anzahl der vereinbarten Abschläge pro Jahr, z.B. 12
	Vertragslaufzeit      Zeitraum `json:"vertragslaufzeit,omitempty"`      // 	Über diesen Zeitraum läuft der Vertrag.
	Kuendigungsfrist      Zeitraum `json:"kuendigungsfrist,omitempty"`      // 	Innerhalb dieser Frist kann der Vertrag gekündigt werden.
	Vertragsverlaengerung Zeitraum `json:"vertragsverlaengerung,omitempty"` //  Falls der Vertrag nicht gekündigt wird, verlängert er sich automatisch um die hier angegebene Zeit.
	Abschlagszyklus       Zeitraum `json:"abschlagszyklus,omitempty"`       //	In diesen Zyklen werden Abschläge gestellt. Alternativ kann auch die Anzahl in den Konditionen angeben werden.
}
