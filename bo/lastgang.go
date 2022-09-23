package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
	"github.com/hochfrequenz/go-bo4e/enum/mengeneinheit"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Lastgang ist ein Modell zur Abbildung eines Lastganges. In diesem Modell werden die Messwerte mit einem vollständigen Zeitintervall angegeben und es bietet daher eine hohe Flexibilität in der Übertragung jeglicher zeitlich veränderlicher Messgrössen.
type Lastgang struct {
	Geschaeftsobjekt
	Sparte       sparte.Sparte               `json:"sparte,omitempty" example:"STROM" validate:"required"`                                                         // Sparte ist eine Angabe, ob es sich um einen GAS- oder STROM-Lastgang handelt
	Version      string                      `json:"version,omitempty" validate:"alphanum"`                                                                        // Version ist die Versionsnummer des Lastgangs
	LokationsId  string                      `json:"lokationsId,omitempty" example:"DE0123456789012345678901234567890" validate:"alphanum,required,min=11,max=33"` // LokationsId is the ID of the location (either a LokationsTyp MALO ID (11 digits) or a LokationsTyp MELO ID (33 alphanum))
	LokationsTyp lokationstyp.Lokationstyp   `json:"lokationsTyp,omitempty" example:"MELO" validate:"required"`                                                    // LokationsTyp is the type of the location in LokationsId
	Messgroesse  mengeneinheit.Mengeneinheit `json:"einheit,omitempty" validate:"required" example:"KWH"`                                                          // Messgroesse ist die Definition der gemessenen Größe anhand ihrer Einheit.
	Obiskennzahl string                      `json:"obiskennzahl,omitempty" example:"1-0:1.8.1"`                                                                   // Obiskennzahl ist die genormte OBIS-Kennzahl zur Kennzeichnung der Messgröße
	Werte        []com.Zeitreihenwert        `json:"energieverbrauch,omitempty" validate:"required,min=1"`                                                         // Werte sind die im Lastgang enthaltenen Messwerte.
}

func (_ Lastgang) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
func (_ Lastgang) GetBoTyp() botyp.BOTyp {
	// this is useful for generics to work
	return botyp.LASTGANG
}
