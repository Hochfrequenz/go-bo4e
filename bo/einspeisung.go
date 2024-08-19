package bo

import (
	eegvermarktungsform "github.com/hochfrequenz/go-bo4e/enum/eegvermarktungsform"
	"github.com/hochfrequenz/go-bo4e/enum/fernsteuerbarkeitstatus"
	"github.com/hochfrequenz/go-bo4e/enum/geschaeftspartnerrolle"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
)

// Einspeisung
type Einspeisung struct {
	Geschaeftsobjekt
	MarktlokationsId        string                                          `json:"marktlokationsId,omitempty"`        // Für welche Marktlokation gelten diese Einspeisedaten
	TrancheId               string                                          `json:"trancheId,omitempty"`               // Für welche Tranche gelten diese Einspeisedaten
	Verguetungsempfaenger   geschaeftspartnerrolle.Geschaeftspartnerrolle   `json:"verguetungsempfaenger,omitempty"`   // Empfänger der Vergütung zur Einspeisung
	EEGVermarktungsform     eegvermarktungsform.EEGVermarktungsform         `json:"eEGVermarktungsform,omitempty"`     // Vermarktungsformen gemäß dem Erneuerbare-Energien-Gesetz (EEG).
	Landescode              *landescode.Landescode                          `json:"landescode,omitempty"`              // Land der Förderung
	FernsteuerbarkeitStatus fernsteuerbarkeitstatus.FernsteuerbarkeitStatus `json:"fernsteuerbarkeitStatus,omitempty"` // Status der Fernsteuerbarkeit einer Marktlokation
}

func (_ Einspeisung) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
