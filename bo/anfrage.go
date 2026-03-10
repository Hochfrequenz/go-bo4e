package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/anfragekategorie"
	"github.com/hochfrequenz/go-bo4e/enum/anfragetyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
)

// BO Anfrage repräsentiert eine Anfrage an einen Marktteilnehmer (ORDERS)
type Anfrage struct {
	Geschaeftsobjekt
	LokationsId             string                            `json:"lokationsId" validate:"required"`
	Lokationstyp            lokationstyp.Lokationstyp         `json:"lokationsTyp" validate:"required"`
	OBISKennzahl            *string                           `json:"obiskennzahl,omitempty"`
	ZeitraumMesswertanfrage *com.Zeitraum                     `json:"ZeitraumMesswertanfrage,omitempty"`
	Anfragekategorie        anfragekategorie.Anfragekategorie `json:"anfragekategorie,omitempty" validate:"required"`
	Anfragetyp              *anfragetyp.Anfragetyp            `json:"anfragetyp,omitempty"`
	// Positionsnummer is the position number from the ORDERS message
	// https://github.com/Hochfrequenz/BO4E-dotnet/blob/0.54.0/BO4E/BO/Anfrage.cs#L85
	Positionsnummer string `json:"positionsnummer,omitempty" validate:"omitempty"`
}
