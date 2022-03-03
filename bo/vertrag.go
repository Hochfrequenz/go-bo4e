package bo

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsart"
	"github.com/hochfrequenz/go-bo4e/enum/vertragsstatus"
	"time"
)

// Vertrag ist ein Modell für die Abbildung von Vertragsbeziehungen. Das Objekt dient dazu, alle Arten von Verträgen, die in der Energiewirtschaft Verwendung finden, abzubilden.
type Vertrag struct {
	Geschaeftsobjekt
	Vertragsnummer string                        `json:"vertragsnummer,omitempty" validate:"alphanum,required"` // Vertragsnummer ist eine im Verwendungskontext eindeutige Nummer für den Vertrag
	Beschreibung   string                        `json:"beschreibung,omitempty"`                                // Beschreibung zum Vertrag
	Vertragsstatus vertragsstatus.Vertragsstatus `json:"vertragstatus,omitempty" validate:"required"`           // Vertragsstatus ist der Status des Vertrags
	// vertragsstatus serializes as "vertragstatus" with single "s" for compatability.
	Vertragsart         vertragsart.Vertragsart  `json:"vertragsart,omitempty" validate:"required"`                         // Vertragsart legt fest, um welche Art von Vertrag es sich handelt. Z.B. Netznutzungvertrag.
	Sparte              sparte.Sparte            `json:"sparte,omitempty" validate:"required"`                              // Sparte sind Unterscheidungsmöglichkeiten für die Sparte
	Vertragsbeginn      time.Time                `json:"vertragsbeginn,omitempty" validate:"required"`                      // Vertragsbeginn is the inclusive start
	Vertragsende        time.Time                `json:"vertragsende,omitempty" validate:"required,gtfield=Vertragsbeginn"` // Vertragsende is the exclusive end
	Vertragspartner1    Geschaeftspartner        `json:"vertragspartner1,omitempty" validate:"required"`                    // Vertragspartner1 ist der "erstgenannte" Vertragspartner. In der Regel der Aussteller des Vertrags. Beispiel: "Vertrag zwischen Vertragspartner 1 ..."
	Vertragspartner2    Geschaeftspartner        `json:"vertragspartner2,omitempty" validate:"required"`                    // Vertragspartner2 ist der "zweitgenannte" Vertragspartner. In der Regel der Empfänger des Vertrags. Beispiel "Vertrag zwischen Vertragspartner 1 und Vertragspartner 2"
	UnterzeichnerVp1    []com.Unterschrift       `json:"unterzeichnervp1,omitempty"`                                        // UnterzeichnerVp1 ist der Unterzeichner des Vertragspartner1
	UnterzeichnerVp2    []com.Unterschrift       `json:"unterzeichnervp2,omitempty"`                                        // UnterzeichnerVp2 ist der Unterzeichner des Vertragspartner2
	Vertragskonditionen *com.Vertragskonditionen `json:"vertragskonditionen,omitempty"`                                     // Vertragskonditionen ist eine Festlegungen zu Laufzeiten und Kündigungsfristen
	Vertragsteile       []com.Vertragsteil       `json:"vertragsteile,omitempty" validate:"required,min=1"`                 // Vertragsteile sind die Vertragsteile, die dazu verwendet werden, eine vertragliche Leistung in Bezug zu einer Lokation (Markt- oder Messlokation) festzulegen.
}

func (_ Vertrag) GetDefaultJsonTags() []string {
	return vertragJsonKeys // this is needed for (un)marshaling of non-default/unknown json fields
}

// vertragJsonKeys is a list of all keys in the standard bo4e json Vertrag.
// It is used to distinguish fields that can be mapped to the Vertrag struct and those that are moved to Geschaeftsobjekt.ExtensionData
var vertragJsonKeys = []string{
	// https://c.tenor.com/71HGq_GX1pMAAAAC/kill-me-simpsons.gif
	// there has to be a better way than this.
	"boTyp",
	"versionStruktur",
	"vertragsnummer",
	"beschreibung",
	"vertragstatus", // watch out: not "ss"
	"vertragsart",
	"sparte",
	"vertragsbeginn",
	"vertragsende",
	"vertragspartner1",
	"vertragspartner2",
	"unterzeichnervp1",
	"unterzeichnervp2",
	"vertragskonditionen",
	"vertragsteile",
}

// the vertragForUnmarshal type is derived from Vertrag but uses a different unmarshaler/deserializer so that we don't run into an endless recursion when overriding the UnmarshalJSON func to include our hacky workaround
type vertragForUnmarshal Vertrag

func (vertrag *Vertrag) UnmarshalJSON(bytes []byte) (err error) {
	// the approach we use here is often referred to as "unmarshal twice"
	// it's a workaround for the not-so-mature/feature rich encoding/json framework in go
	_vertrag := vertragForUnmarshal{}
	// first we deserialize into the helper/intermediate type. this is to _not_ run into this Unmarshal func in an endless recursion
	if err = json.Unmarshal(bytes, &_vertrag); err == nil {
		*vertrag = Vertrag(_vertrag)
		// the vertrag contains only those fields that are defined in the Vertrag struct by now
	} else {
		return err
	}

	// now we unmarshal the same bytes into the extension data
	if err = json.Unmarshal(bytes, &vertrag.ExtensionData); err != nil {
		return nil
	}
	// But now the extension data also contain those fields that in fact have a representation in the Bilanzierung struct
	vertrag.RemoveStronglyTypedFieldsFromExtensionData(vertrag.GetDefaultJsonTags()) // remove those fields from the extension data that have a representation in the Bilanzierung struct
	return nil
}

// vertragForMarshal is a struct similar to the original Vertrag but uses a different Marshaller so that we don't run into an endless recursion
type vertragForMarshal Vertrag

//nolint:dupl // This can only be simplified if we use generics. anything else seems overly complicated but maybe it's just me
func (vertrag Vertrag) MarshalJSON() ([]byte, error) {
	if vertrag.ExtensionData == nil || len(vertrag.ExtensionData) == 0 {
		// no special handling needed
		return json.Marshal(vertragForMarshal(vertrag)) // just marshal but use a helper type to not run into an endless recursion
	}
	// there is probably a better way than this, like f.e. creating an adhoc-struct that has an embedded bila-like type and f.e. a map or
	// we first convert the Bilanzierung into a map[string]interface{}
	// we serialize the bilanzierung via our helper type
	bilaDictBytes, bilaMarshalErr := json.Marshal(vertragForMarshal(vertrag)) // we must use a different type here to not run into an endless recursion
	if bilaMarshalErr != nil {
		return []byte{}, bilaMarshalErr
	}
	// now we deserialize the malo again but we use a generic dictionary as target type
	bilaAsMap := map[string]interface{}{}
	extensionUnmarshalErr := json.Unmarshal(bilaDictBytes, &bilaAsMap)
	if extensionUnmarshalErr != nil {
		return []byte{}, extensionUnmarshalErr
	}
	// now we join/merge the original malo and its extension data (which is already a map[string]interface{} into a single result
	result := map[string]interface{}{}
	for key, value := range bilaAsMap {
		result[key] = value
	}
	for key, value := range vertrag.ExtensionData {
		result[key] = value
	}
	return json.Marshal(result)
}
