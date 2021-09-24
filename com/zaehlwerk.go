package com

import (
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
)

// A Zaehlwerk is the counting part of a meter. A meter consists of one or more Zaehlwerke
type Zaehlwerk struct {
	ZaehlwerkId string `json:"zaehlwerkId" validate:"required"`// Identifikation des Zählwerks (Registers) innerhalb des Zählers. Oftmals eine laufende Nummer hinter der Zählernummer. Z.B. 47110815_1
	Bezeichnung string `json:"bezeichnung"`// Zusätzliche Bezeichnung, z.B. Zählwerk_Wirkarbeit.
	Richtung energierichtung.Energierichtung `json:"richtung" validate:"required"` // Die Energierichtung, Einspeisung oder Ausspeisung. Details siehe ENUM Energierichtung
	ObisKennzahl string `json:"obisKennzahl" validate:"required"` //Die OBIS-Kennzahl für das Zählwerk, die festlegt, welche auf die gemessene Größe mit dem Stand gemeldet wird. Nur Zählwerkstände mit dieser OBIS-Kennzahl werden an diesem Zählwerk registriert. Beispiel: 1-0:1.8.1 für elektrische Wirkarbeit.
	Wandlerfaktor float32 `json:"wandlerfaktor"` // Mit diesem Faktor wird eine Zählerstandsdifferenz multipliziert, um zum eigentlichen Verbrauch im Zeitraum zu kommen.
	Einheit mengeneinheit.Mengeneinheit `json:"einheit" validate:"required"`	// Die Einheit der gemessenen Größe, z.B. kWh.
	Zaehlerstaende Zaehlerstaende `json:"zaehlerstaende,omitempty"` // Non BO4E Standard
}
