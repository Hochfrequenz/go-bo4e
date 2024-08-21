package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"time"
)

// THIS IS STILL WIP - missing Angebotsvariante, Angebotsteil, Angebotsstatus etc.

// Mit diesem BO kann ein Versorgungsangebot zur Strom- oder Gasversorgung oder die Teilnahme an einer Ausschreibung übertragen werden.
// Es können verschiedene Varianten enthalten sein (z.B.ein- und mehrjährige Laufzeit). Innerhalb jeder Variante können Teile enthalten sein,
// die jeweils für eine oder mehrere Marktlokationen erstellt werden.
type Angebot struct {
	Geschaeftsobjekt
	Angebotsnummer              string                 `json:"angebotsnummer,omitempty" validate:"required"` // Eindeutige Nummer des Angebotes.
	Anfragereferenz             *string                `json:"anfragereferenz,omitempty"`                    // Referenz auf eine Anfrage oder Ausschreibung. Kann dem Empfänger des Angebotes bei Zuordnung des Angebotes zur Anfrage bzw.Ausschreibung helfen.
	Angebotsdatum               time.Time              `json:"angebotsdatum,omitempty"`                      // Erstellungsdatum des Angebots
	Sparte                      sparte.Sparte          `json:"sparte,omitempty" validate:"required"`         // Sparte, für die das Angebot abgegeben wird (Strom/Gas).
	Bindefrist                  *time.Time             `json:"bindefrist,omitempty"`                         // Bis zu diesem Zeitpunkt(Tag/Uhrzeit) inklusive gilt das Angebot.
	Angebotgeber                *Geschaeftspartner     `json:"angebotgeber,omitempty" validate:"required"`   // Link auf den Ersteller des Angebots.
	Angebotnehmer               *Geschaeftspartner     `json:"angebotnehmer,omitempty" validate:"required"`  // Link auf den Empfänger des Angebots.
	UnterzeichnerAngebotsnehmer *Ansprechpartner       `json:"unterzeichnerAngebotsnehmer,omitempty"`        // Link auf die Person, die als Angebotsnehmer das Angebot angenommen hat.
	UnterzeichnerAngebotsgeber  *Ansprechpartner       `json:"unterzeichnerAngebotsgeber,omitempty"`         // Link auf die Person, die als Angebotsgeber das Angebot ausgestellt hat.
	Varianten                   []com.Angebotsvariante `json:"varianten,omitempty"`                          // Eine oder mehrere Varianten des Angebots mit den Angebotsteilen. Ein Angebot besteht mindestens aus einer Variante.
}

func (_ Angebot) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
