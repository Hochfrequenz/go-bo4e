package bo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/reklamationsgrund"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
)

// Reklamation wird benutzt um eine fehlende, oder unplausible Energiemenge anzuzeigen. Für Reklamationen von Rechnungen, wird ein Avis vom Typ Nicht-Zahlungsavis benutzt
type Reklamation struct {
	Geschaeftsobjekt
	LokationsId                string                              `json:"lokationsId,omitempty" example:"DE0123456789012345678901234567890" validate:"alphanum,required,min=11,max=33"` // LokationsId gibt die ID der Lokation an (entweder LokationsTyp MALO ID (11 Stellen) oder LokationsTyp MELO ID (33 alphanum)).
	LokationsTyp               lokationstyp.Lokationstyp           `json:"lokationsTyp,omitempty" example:"MELO" validate:"required"`                                                    // LokationsTyp gibt an, ob es sich um eine Markt- oder Messlokation handelt.
	Reklamationsgrund          reklamationsgrund.Reklamationsgrund `json:"reklamationsgrund,omitempty" validate:"required"`                                                              // Reklamationsgrund gibt den Grund der Reklamation an
	ObisKennzahl               *string                             `json:"obisKennzahl,omitempty" validate:"required" example:"1-0:1.8.1"`                                               // ObisKennzahl für das Zählwerk, die festlegt, welche auf die gemessene Größe mit dem Stand gemeldet wird. Nur Zählwerkstände mit dieser OBIS-Kennzahl werden an diesem Zählwerk registriert.
	ZeitraumMesswertanfrage    *com.Zeitraum                       `json:"zeitraumMesswertanfrage,omitempty" validate:"required"`                                                        // ZeitraumMesswertanfrage gibt den com.Zeitraum an, auf die sich die Reklamation bezieht.
	ReklamationsgrundBemerkung string                              `json:"reklamationsgrundBemerkung,omitempty" validate:"omitempty"`                                                    // ReklamationsgrundBemerkung ist ein Freitext für eine weitere Beschreibung des Reklamationsgrunds
	// Positionsnummer: Aus der ORDERS gemappte Positionsnummer der Anfrage
	Positionsnummer          string     `json:"positionsnummer,omitempty" validate:"omitempty"`
	ZeitpunktFuerWertanfrage *time.Time `json:"zeitpunktFuerWertanfrage,omitempty" validate:"omitempty"` // ZeitpunktFuerWertanfrage gibt den com.Zeitpunkt an, zu dem die Wertanfrage erfolgt ist.
}

// UnmarshalJSON supports deserializing both string and integer Positionsnummer for backwards compatibility,
// while preserving unmapped data handling via unmappeddatamarshaller.
func (r *Reklamation) UnmarshalJSON(bytes []byte) (err error) {
	// Pre-process: if positionsnummer is a JSON number, convert it to a string for backwards compatibility
	var raw map[string]json.RawMessage
	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}
	if posRaw, ok := raw["positionsnummer"]; ok && len(posRaw) > 0 {
		var s string
		if json.Unmarshal(posRaw, &s) != nil {
			// Not a string, try as number
			var n json.Number
			if err2 := json.Unmarshal(posRaw, &n); err2 == nil {
				quoted, _ := json.Marshal(n.String())
				raw["positionsnummer"] = quoted
				bytes, err = json.Marshal(raw)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("positionsnummer must be a string or number, got %s", string(posRaw))
			}
		}
	}
	return unmappeddatamarshaller.UnmarshallWithUnmappedData(r, &r.ExtensionData, bytes)
}

// reklamationForMarshal is a struct similar to the original Reklamation but uses a different Marshaller so that we don't run into an endless recursion
type reklamationForMarshal Reklamation

//nolint:dupl // This can only be simplified if we use generics. anything else seems overly complicated but maybe it's just me
func (r Reklamation) MarshalJSON() (bytes []byte, err error) {
	s := reklamationForMarshal(r)
	byteArr, err := json.Marshal(s)
	if err != nil {
		return
	}
	return unmappeddatamarshaller.HandleUnmappedDataPropertyMarshalling(byteArr)
}
