package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/handelsunstimmigkeitstyp"
)

// Handelsunstimmigkeit contains information about discrepancies in market communication.
type Handelsunstimmigkeit struct {
	Geschaeftsobjekt

	// Nummer contains the "Handelsunstimmigkeitsnummer".
	Nummer string `json:"nummer,omitempty" validate:"required"`

	// Typ specifies the type of discrepancy.
	Typ handelsunstimmigkeitstyp.Handelsunstimmigkeitstyp `json:"typ,omitempty" validate:"required"`

	// Begruendung contains the reason of discrepancy.
	Begruendung com.Handelsunstimmigkeitsbegruendung `json:"begruendung" validate:"required"`

	// Betrag is the requested sum amount (optional).
	Betrag *com.Betrag `json:"zuZahlen,omitempty"`
}

func (_ Handelsunstimmigkeit) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
