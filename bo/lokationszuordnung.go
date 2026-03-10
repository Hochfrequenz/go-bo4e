package bo

import (
	"encoding/json"

	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
)

// The Lokationszuordnung ist ein Modell für die Abbildung der Referenz auf die Lokationsbündelstruktur. Diese gibt an welche Marktlokationen, Messlokationen, Netzlokationen, technische/steuerbaren Ressourcen an einer Lokation vorhanden sind.
type Lokationszuordnung struct {
	Geschaeftsobjekt
	// Marktlokationen ist eine Liste mit IDs der referenzierten Marktlokationen
	Marktlokationen []Marktlokation `json:"marktlokationen,omitempty"`

	// Messlokationen ist eine Liste mit IDs der referenzierten Messlokationen
	Messlokationen []Messlokation `json:"messlokationen,omitempty"`

	// Netzlokationen ist eine Liste mit IDs der referenzierten Netzlokationen
	Netzlokationen []Netzlokation `json:"netzlokationen,omitempty"`

	// TechnischeRessourcen ist eine Liste mit IDs der referenzierten technischen Ressourcen
	TechnischeRessourcen []TechnischeRessource `json:"technischeRessourcen,omitempty"`

	// SteuerbareRessourcen ist eine Liste mit IDs der referenzierten steuerbaren Ressourcen
	SteuerbareRessourcen []SteuerbareRessource `json:"steuerbareRessourcen,omitempty"`

	// Gueltigkeit ist die Zeitspanne in der die Zuordnung gültig ist
	Gueltigkeit []com.Zeitraum `json:"gueltigkeit,omitempty"`

	// Zuordnungstyp ist die Verknüpfungsrichtung z.B. Malo-Melo
	Zuordnungstyp string `json:"zuordnungstyp,omitempty"`

	// LokationsbuendelCode ist der Code, der angibt wie die Lokationsbündelstruktur zusammengesetzt ist (zu finden unter "Codeliste der Lokationsbündelstrukturen" auf https://www.edi-energy.de/index.php?id=38)
	LokationsbuendelCode string `json:"lokationsbuendelcode,omitempty"`
}

func (lokZu *Lokationszuordnung) UnmarshalJSON(bytes []byte) (err error) {
	return unmappeddatamarshaller.UnmarshallWithUnmappedData(lokZu, &lokZu.ExtensionData, bytes)
}

// lokationsZuordnungForUnmarshal is derived from Lokationszuordnung but uses a different unmarshaler/deserializer so that we don't run into an endless recursion when overriding the UnmarshalJSON func to include our hacky workaround
type lokationsZuordnungForUnmarshal Lokationszuordnung

func (lokZu Lokationszuordnung) MarshalJSON() (bytes []byte, err error) {
	s := lokationsZuordnungForUnmarshal(lokZu)
	byteArr, err := json.Marshal(s)
	if err != nil {
		return
	}
	return unmappeddatamarshaller.HandleUnmappedDataPropertyMarshalling(byteArr)
}
