package bo

import (
	"time"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/definitionennotwendigkeit"
)

// Zaehlzeitdefinition en nutzt der NB bzw. LF  für die Tarifierung von Werten.
// Zaehlzeitdefinitionen werden in der Marktkommunikation mit Prüfidentifikator 25001 (UTILTS) übermittelt.
// Eine Zählzeitdefinition umfasst dabei eine Liste von möglichen Zählzeiten, den dazugehörigen Registern und der tatsächlich ausgerollten Zählzeit (wenn diese elektronisch übermittelt wird).
type Zaehlzeitdefinition struct {
	Geschaeftsobjekt
	Beginndatum            *time.Time                                           `json:"beginndatum,omitempty"`
	Endedatum              *time.Time                                           `json:"endedatum,omitempty"`
	Version                time.Time                                            `json:"version,omitempty"`
	Notwendigkeit          *definitionennotwendigkeit.DefinitionenNotwendigkeit `json:"notwendigkeit,omitempty"`
	Zaehlzeiten            []com.Zaehlzeit                                      `json:"zaehlzeiten,omitempty"`
	Zaehlzeitregister      []com.Zaehlzeitregister                              `json:"zaehlzeitregister,omitempty"`
	AusgerollteZaehlzeiten []com.AusgerollteZaehlzeit                           `json:"ausgerollteZaehlzeiten,omitempty"`
}
