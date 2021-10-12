package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

// BusinessObject is the interface that all Business Objects implement.
type BusinessObject interface {
	GetBoTyp() botyp.BOTyp
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
