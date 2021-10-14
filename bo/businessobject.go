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
}

// Geschaeftsobjekt is the common base struct of all Business Objects
type Geschaeftsobjekt struct {
	BoTyp             botyp.BOTyp           `json:"boTyp" validate:"required"`           // type of business object, may be used as discriminator
	VersionStruktur   string                `json:"versionStruktur" validate:"required"` // version of BO4E used
	ExterneReferenzen []com.ExterneReferenz `json:"externeReferenzen"`                   // external references of this object in various systems
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
const defaultVersionStruktur = "1.1" // because "externeReferenzen" not "externeReferenz"

// NewBusinessObject creates an empty BusinessObject based on the provided type; Returns nil if the type is not implemented.
func NewBusinessObject(typ botyp.BOTyp) BusinessObject {
	var bo BusinessObject
	switch typ {
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
	case botyp.RECHNUNG:
		bo = new(Rechnung)
		bo.(*Rechnung).BoTyp = typ
		bo.(*Rechnung).VersionStruktur = defaultVersionStruktur
	case botyp.VERTRAG:
		bo = new(Vertrag)
		bo.(*Vertrag).BoTyp = typ
		bo.(*Vertrag).VersionStruktur = defaultVersionStruktur
	case botyp.ZAEHLER:
		bo = new(Zaehler)
		bo.(*Zaehler).BoTyp = typ
		bo.(*Zaehler).VersionStruktur = defaultVersionStruktur
	default:
		return nil
	}
	return bo
}
