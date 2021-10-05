package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

// BusinessObject is the common base struct of all Business Objects
type BusinessObject struct {
	BoTyp             botyp.BOTyp           `json:"boTyp" validate:"required"`                     // type of business object, may be used as discriminator
	VersionStruktur   string                `json:"versionStruktur,omitempty" validate:"required"` // version of BO4E used
	ExterneReferenzen []com.ExterneReferenz `json:"externeReferenzen,omitempty"`                   // external references of this object in various systems
}
