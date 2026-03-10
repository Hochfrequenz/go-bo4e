package bo

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/enum/haeufigkeitzaehlzeit"
	"github.com/hochfrequenz/go-bo4e/enum/uebermittelbarkeitzaehlzeit"
)

// Leistungskurvendefinition beschreibt die Eigenschaften einer (nicht) ausgerollten Leistungskurvendefinition.
type Leistungskurvendefinition struct {
	Geschaeftsobjekt
	// Ausgerollt gibt an, ob die Leistungskurvendefinition ausgerollt ist oder nicht.
	Ausgerollt *bool `json:"ausgerollt,omitempty"`

	// Aenderungszeitpunkt ist der Leistungskurvenänderungszeitpunkt.
	Aenderungszeitpunkt time.Time `json:"aenderungszeitpunkt,omitempty"`

	// LeistungskurvendefinitionsCode ist der Code der Leistungskurvendefinition.
	LeistungskurvendefinitionsCode *string `json:"leistungskurvendefinitionscode,omitempty"`

	// Haeufigkeit beschreibt die Häufigkeit der Übermittlung.
	Haeufigkeit *haeufigkeitzaehlzeit.HaeufigkeitZaehlzeit `json:"haeufigkeit,omitempty"`

	// Uebermittelbarkeit beschreibt die Übermittelbarkeit der ausgerollten Leistungskurvendefinition.
	Uebermittelbarkeit *uebermittelbarkeitzaehlzeit.UebermittelbarkeitZaehlzeit `json:"uebermittelbarkeit,omitempty"`

	// ObererSchwellwert ist der obere Schwellwert.
	ObererSchwellwert *string `json:"obererSchwellwert,omitempty"`

	// AenderungszeitpunktDateFormat is the EDIFACT date format of the Aenderungszeitpunkt (e.g. "401")
	// https://github.com/Hochfrequenz/BO4E-dotnet/blob/0.54.0/BO4E/BO/Leistungskurvendefinition.cs#L95
	AenderungszeitpunktDateFormat *string `json:"aenderungszeitpunktDateFormat,omitempty"`
}
