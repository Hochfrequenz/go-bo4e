package bo

import (
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
)

// Marktteilnehmer models participants of the German energy market
type Marktteilnehmer struct {
	Geschaeftspartner
	Marktrolle       *marktrolle.Marktrolle      `json:"marktrolle,omitempty" validate:"required"`                             // Marktrolle gibt im Klartext die Bezeichnung der Marktrolle an. Details siehe ENUM Marktrolle
	Rollencodenummer string                      `json:"rollencodenummer,omitempty" validate:"required,numeric,min=13,max=13"` // Rollencodenummer gibt die Codenummer der Marktrolle an (13 digits)
	Rollencodetyp    rollencodetyp.Rollencodetyp `json:"rollencodetyp,omitempty" validate:"required"`                          // Rollencodetyp gibt den Typ des Codes an/Vergabestelle. Details siehe ENUM Rollencodetyp
	Makoadresse      string                      `json:"makoadresse,omitempty" validate:"omitempty"`                           // Makoadresse ist die 1:1-Kommunikationsadresse des Marktteilnehmers. Diese wird in der Marktkommunikation verwendet.
	Ansprechpartner  *Ansprechpartner            `json:"ansprechpartner,omitempty" validate:"omitempty"`                       // Ansprechpartner ist ein Kontakt zur bilateralen Klärung
}

func (_ Marktteilnehmer) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
