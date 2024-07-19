package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/definitionennotwendigkeit"
	"time"
)

// / <summary>
// / Der NB bzw. LF nutzt Zählzeitdefinitionen für die Tarifierung von Werten.
// / <remarks>Zaehlzeitdefinitionen werden in der Marktkommunikation mit Prüfidentifikator 25001 (UTILTS) übermittelt</remarks>
// / Eine Zählzeitdefinition umfasst dabei eine Liste von möglichen Zählzeiten,
// / den dazugehörigen Registern und der tatsächlich ausgerollten Zählzeit (wenn diese elektronisch übermittelt wird)
// / </summary>
type Zaehlzeitdefinition struct {
	Geschaeftsobjekt
	Beginndatum            *time.Time                                           `json:"beginndatum,omitempty"`
	Endedatum              *time.Time                                           `json:"endedatum,omitempty"`
	Version                time.Time                                            `json:"version,omitempty"`
	Notwendigkeit          *definitionennotwendigkeit.DefinitionenNotwendigkeit `json:"notwendigkeit,omitempty"`
	Zaehlzeiten            []com.Zaehlzeit                                      `json:"zaehlzeiten,omitempty"`
	Zaehlzeitregister      []com.Zaehlzeitregister                              `json:"zaehlzeitregister,omitempty"`
	AusgerollteZaehlzeiten []com.AusgerollteZaehlzeit                           `json:"ausgerolltezaehlzeiten,omitempty"`
}

func (_ Zaehlzeitdefinition) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
