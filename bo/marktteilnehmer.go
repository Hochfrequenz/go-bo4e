package bo

import (
	"encoding/json"

	"github.com/hochfrequenz/go-bo4e/enum/marktrolle"
	"github.com/hochfrequenz/go-bo4e/enum/rollencodetyp"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
)

// Marktteilnehmer models participants of the German energy market
type Marktteilnehmer struct {
	Geschaeftspartner
	Marktrolle       *marktrolle.Marktrolle       `json:"marktrolle,omitempty"`                                                 // Marktrolle gibt im Klartext die Bezeichnung der Marktrolle an. Details siehe ENUM Marktrolle
	Rollencodenummer *string                      `json:"rollencodenummer,omitempty" validate:"required,numeric,min=13,max=13"` // Rollencodenummer gibt die Codenummer der Marktrolle an (13 digits)
	Rollencodetyp    *rollencodetyp.Rollencodetyp `json:"rollencodetyp,omitempty" validate:"required"`                          // Rollencodetyp gibt den Typ des Codes an/Vergabestelle. Details siehe ENUM Rollencodetyp
	Makoadresse      string                       `json:"makoadresse,omitempty" validate:"omitempty"`                           // Makoadresse ist die 1:1-Kommunikationsadresse des Marktteilnehmers. Diese wird in der Marktkommunikation verwendet.
	Ansprechpartner  *Ansprechpartner             `json:"ansprechpartner,omitempty" validate:"omitempty"`                       // Ansprechpartner ist ein Kontakt zur bilateralen Klärung
}

func (marktteilnehmer *Marktteilnehmer) UnmarshalJSON(bytes []byte) (err error) {
	return unmappeddatamarshaller.UnmarshallWithUnmappedData(marktteilnehmer, &marktteilnehmer.ExtensionData, bytes)
}

// martktteilnehmerForMarshal is a struct similar to the original Marktteilnehmer but uses a different Marshaller so that we don't run into an endless recursion
type marktteilnehmerForMarshal Marktteilnehmer

//nolint:dupl // This can only be simplified if we use generics. anything else seems overly complicated but maybe it's just me
func (marktteilnehmer Marktteilnehmer) MarshalJSON() (bytes []byte, err error) {
	s := marktteilnehmerForMarshal(marktteilnehmer)
	byteArr, err := json.Marshal(s)
	if err != nil {
		return
	}
	return unmappeddatamarshaller.HandleUnmappedDataPropertyMarshalling(byteArr)
}
