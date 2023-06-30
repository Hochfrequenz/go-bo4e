package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/reklamationsgrund"
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
	Positionsnummer            int                                 `json:"positionsnummer,omitempty" validate:"omitempty"`                                                               // Positionsnummer ist eine fortlaufende Nummer für die Position
}

func (_ Reklamation) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
