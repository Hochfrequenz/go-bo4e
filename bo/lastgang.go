package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Lastgang ist ein Modell zur Abbildung eines Lastganges. In diesem Modell werden die Messwerte mit einem vollständigen Zeitintervall angegeben und es bietet daher eine hohe Flexibilität in der Übertragung jeglicher zeitlich veränderlicher Messgrössen.
type Lastgang struct {
	BusinessObject
	Sparte       sparte.Sparte               `json:"sparte,omitempty" example:"STROM" validate:"required"`                                               // Angabe, ob es sich um einen Gas- oder Stromlastgang handelt
	Version      string                      `json:"version,omitempty" validate:"alphanum"`                                                              // Versionsnummer des Lastgangs
	LokationsId  string                      `json:"lokationsId" example:"DE0123456789012345678901234567890" validate:"alphanum,required,min=11,max=33"` // ID of the location (either a LokationsTyp MaLo ID (11 digits) or a LokationsTyp MeLo ID (33 alphanum))
	LokationsTyp lokationstyp.Lokationstyp   `json:"lokationsTyp,omitempty" example:"MeLo" validate:"required"`                                          // type of the location in LokationsId
	Messgroesse  mengeneinheit.Mengeneinheit `json:"einheit,omitempty" validate:"required" example:"KWH"`                                                // Definition der gemessenen Größe anhand ihrer Einheit.
	Obiskennzahl string                      `json:"obiskennzahl,omitempty" example:"1-0:1.8.1"`                                                         // Genormte OBIS-Kennzahl zur Kennzeichnung der Messgröße
	Werte        []com.Zeitreihenwert        `json:"energieverbrauch,omitempty" validate:"required,min=1"`                                               // Die im Lastgang enthaltenen Messwerte.
}
