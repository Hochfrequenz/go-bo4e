package bo

import (
	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
)

// Marktteilnehmer models participants of the German energy market
type Marktteilnehmer struct {
	Geschaeftspartner
	Marktrolle       marktrolle.Marktrolle       `json:"marktrolle" validate:"required"`                             // Gibt im Klartext die Bezeichnung der Marktrolle an. Details siehe ENUM Marktrolle
	Rollencodenummer string                      `json:"rollencodenummer" validate:"required,numeric,min=13,max=13"` // Gibt die Codenummer der Marktrolle an (13 digits)
	Rollencodetyp    rollencodetyp.Rollencodetyp `json:"rollencodetyp" validate:"required"`                          // Gibt den Typ des Codes an/Vergabestelle. Details siehe ENUM Rollencodetyp
	Makoadresse      string                      `json:"makoadresse,omitempty" validate:"omitempty"`                 // Die 1:1-Kommunikationsadresse des Marktteilnehmers. Diese wird in der Marktkommunikation verwendet.
}
