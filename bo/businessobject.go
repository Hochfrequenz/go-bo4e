package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

// BusinessObject is the interface that all Business Objects implement.
type BusinessObject interface {
	// GetBoTyp returns the Geschaeftsobjekt.BoTyp
	GetBoTyp() botyp.BOTyp
	// GetVersionStruktur returns the Geschaeftsobjekt.VersionStruktur
	GetVersionStruktur() string
	// GetDefaultJsonTags returns an array of string that represent the json tags that are (by default) used by this business object
	GetDefaultJsonTags() []string
}

// Geschaeftsobjekt is the common base struct of all Business Objects
type Geschaeftsobjekt struct {
	BoTyp             botyp.BOTyp            `json:"boTyp" validate:"required"`           // BoTyp is the type of business object, may be used as discriminator
	VersionStruktur   string                 `json:"versionStruktur" validate:"required"` // VersionStruktur is the version of BO4E used
	ExterneReferenzen []com.ExterneReferenz  `json:"externeReferenzen,omitempty"`         // ExterneReferenzen are external references of this object in various systems
	ExtensionData     map[string]interface{} `json:"-"`                                   // ExtensionData contain those values that are included in a Geschaeftsobjekt but not part of the BO4E standard/this implementation
	// ExtensionData is derived from the .NET naming (JsonExtensionData). It may also be called "MissingMemberHandling"
}

// RemoveStronglyTypedFieldsFromExtensionData removes the keys of the given key set from the Geschaeftsobjekt.ExtensionData
func (gob *Geschaeftsobjekt) RemoveStronglyTypedFieldsFromExtensionData(keys []string) {
	// https://stackoverflow.com/a/33437853/10009545
	// Maybe I'm not a perfect go expert; But if this frickeling and hardcoding is the way to go, then the typical go programmers pain threshold is admirable high.
	for _, key := range keys {
		gob.DeleteExtensionData(key)
	}
}

// DeleteExtensionData removes the given jsonKey from the Geschaeftsobjekt.ExtensionData
func (gob *Geschaeftsobjekt) DeleteExtensionData(jsonKey string) {
	if gob.ExtensionData != nil {
		delete(gob.ExtensionData, jsonKey)
	}
}

func (gob Geschaeftsobjekt) GetBoTyp() botyp.BOTyp {
	return gob.BoTyp
}

func (gob Geschaeftsobjekt) GetVersionStruktur() string {
	if gob.VersionStruktur == "" {
		return defaultVersionStruktur
	}
	return gob.VersionStruktur
}

// defaultVersionStruktur is the Geschaeftsobjekt.VersionStruktur that a NewBusinessObject has by default
const defaultVersionStruktur = "1.1" // because Geschaeftsobjekt.ExterneReferenzen (plural) was introduced in v1.1. In v1.0 the name was "externeReferenz" (singular)

// NewBusinessObject creates an empty BusinessObject based on the provided type; Returns nil if the type is not implemented.
func NewBusinessObject(typ botyp.BOTyp) BusinessObject {
	var bo BusinessObject
	switch typ {
	case botyp.ANSPRECHPARTNER:
		bo = new(Ansprechpartner)
		bo.(*Ansprechpartner).BoTyp = typ
		bo.(*Ansprechpartner).VersionStruktur = defaultVersionStruktur
	case botyp.BILANZIERUNG:
		bo = new(Bilanzierung)
		bo.(*Bilanzierung).BoTyp = typ
		bo.(*Bilanzierung).VersionStruktur = defaultVersionStruktur
	case botyp.ENERGIEMENGE:
		bo = new(Energiemenge)
		bo.(*Energiemenge).BoTyp = typ
		bo.(*Energiemenge).VersionStruktur = defaultVersionStruktur
	case botyp.GESCHAEFTSPARTNER:
		bo = new(Geschaeftspartner)
		bo.(*Geschaeftspartner).BoTyp = typ
		bo.(*Geschaeftspartner).VersionStruktur = defaultVersionStruktur
	case botyp.LASTGANG:
		bo = new(Lastgang)
		bo.(*Lastgang).BoTyp = typ
		bo.(*Lastgang).VersionStruktur = defaultVersionStruktur
	case botyp.MARKTLOKATION:
		bo = new(Marktlokation)
		bo.(*Marktlokation).BoTyp = typ
		bo.(*Marktlokation).VersionStruktur = defaultVersionStruktur
	case botyp.MARKTTEILNEHMER:
		bo = new(Marktteilnehmer)
		bo.(*Marktteilnehmer).BoTyp = typ
		bo.(*Marktteilnehmer).VersionStruktur = defaultVersionStruktur
	case botyp.MESSLOKATION:
		bo = new(Messlokation)
		bo.(*Messlokation).BoTyp = typ
		bo.(*Messlokation).VersionStruktur = defaultVersionStruktur
	case botyp.NETZNUTZUNGSRECHNUNG:
		bo = new(Netznutzungsrechnung)
		bo.(*Netznutzungsrechnung).BoTyp = typ
		bo.(*Netznutzungsrechnung).VersionStruktur = defaultVersionStruktur
		bo.(*Netznutzungsrechnung).Rechnungsempfaenger.BoTyp = botyp.GESCHAEFTSPARTNER
		bo.(*Netznutzungsrechnung).Rechnungsersteller.BoTyp = botyp.GESCHAEFTSPARTNER
	case botyp.PREISBLATT:
		bo = new(Preisblatt)
		bo.(*Preisblatt).BoTyp = typ
		bo.(*Preisblatt).VersionStruktur = defaultVersionStruktur
		bo.(*Preisblatt).Herausgeber.BoTyp = botyp.MARKTTEILNEHMER
	case botyp.RECHNUNG:
		bo = new(Rechnung)
		bo.(*Rechnung).BoTyp = typ
		bo.(*Rechnung).VersionStruktur = defaultVersionStruktur
		bo.(*Rechnung).Rechnungsempfaenger.BoTyp = botyp.GESCHAEFTSPARTNER
		bo.(*Rechnung).Rechnungsersteller.BoTyp = botyp.GESCHAEFTSPARTNER
	case botyp.REKLAMATION:
		bo = new(Reklamation)
		bo.(*Reklamation).BoTyp = typ
		bo.(*Reklamation).VersionStruktur = defaultVersionStruktur
	case botyp.VERTRAG:
		bo = new(Vertrag)
		bo.(*Vertrag).BoTyp = typ
		bo.(*Vertrag).VersionStruktur = defaultVersionStruktur
		bo.(*Vertrag).Vertragspartner1.BoTyp = botyp.GESCHAEFTSPARTNER
		bo.(*Vertrag).Vertragspartner2.BoTyp = botyp.GESCHAEFTSPARTNER
	case botyp.ZAEHLER:
		bo = new(Zaehler)
		bo.(*Zaehler).BoTyp = typ
		bo.(*Zaehler).VersionStruktur = defaultVersionStruktur
	case botyp.HANDELSUNSTIMMIGKEIT:
		bo = new(Handelsunstimmigkeit)
		bo.(*Handelsunstimmigkeit).BoTyp = typ
		bo.(*Handelsunstimmigkeit).VersionStruktur = defaultVersionStruktur
	case botyp.AVIS:
		bo = new(Avis)
		bo.(*Avis).BoTyp = typ
		bo.(*Avis).VersionStruktur = defaultVersionStruktur
	case botyp.STATUSBERICHT:
		bo = new(Statusbericht)
		bo.(*Statusbericht).BoTyp = typ
		bo.(*Statusbericht).VersionStruktur = defaultVersionStruktur
	default:
		return nil
	}
	return bo
}
