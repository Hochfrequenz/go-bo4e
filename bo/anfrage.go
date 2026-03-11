package bo

import (
	"encoding/json"
	"fmt"

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
	// Positionsnummer: Aus der ORDERS gemappte Positionsnummer der Anfrage
	Positionsnummer string `json:"positionsnummer,omitempty" validate:"omitempty"`
}

// UnmarshalJSON supports deserializing both string and integer Positionsnummer for backwards compatibility.
func (a *Anfrage) UnmarshalJSON(data []byte) error {
	type Alias Anfrage
	aux := &struct {
		*Alias
		RawPositionsnummer json.RawMessage `json:"positionsnummer,omitempty"`
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.RawPositionsnummer) > 0 {
		var s string
		if err := json.Unmarshal(aux.RawPositionsnummer, &s); err == nil {
			a.Positionsnummer = s
		} else {
			var n json.Number
			if err2 := json.Unmarshal(aux.RawPositionsnummer, &n); err2 == nil {
				a.Positionsnummer = n.String()
			} else {
				return fmt.Errorf("positionsnummer must be a string or number, got %s", string(aux.RawPositionsnummer))
			}
		}
	}
	return nil
}
