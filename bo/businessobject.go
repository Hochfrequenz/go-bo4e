package bo

import (
	"encoding/json"
	"errors"
	"fmt"
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

// BusinessObjectSlice is a slice that contains 0-n BusinessObject s
type BusinessObjectSlice []BusinessObject

// GetNewBusinessObject creates an empty BusinessObject based on the provided type; Returns nil if the type is not implemented.
func GetNewBusinessObject(typ botyp.BOTyp) BusinessObject {
	switch typ {
	case botyp.Energiemenge:
		return new(Energiemenge)
	case botyp.Geschaeftspartner:
		return new(Geschaeftspartner)
	case botyp.Lastgang:
		return new(Lastgang)
	case botyp.Messlokation:
		return new(Messlokation)
	case botyp.Netznutzungsrechnung:
		return new(Netznutzungsrechnung)
	case botyp.Rechnung:
		return new(Rechnung)
	case botyp.Vertrag:
		return new(Vertrag)
	case botyp.Zaehler:
		return new(Zaehler)
	default:
		return nil
	}
}

func (boSlice *BusinessObjectSlice) UnmarshalJSON(data []byte) error {
	// https://stackoverflow.com/a/69557652/10009545
	var array []json.RawMessage
	if err := json.Unmarshal(data, &array); err != nil {
		return err
	}

	*boSlice = make(BusinessObjectSlice, len(array))
	for i := range array {
		base := Geschaeftsobjekt{}
		data := []byte(array[i])
		if err := json.Unmarshal(data, &base); err != nil {
			return err
		}
		var elem BusinessObject
		elem = GetNewBusinessObject(base.GetBoTyp())
		if elem == nil {
			errorMessage := fmt.Sprintf("The BusinessObject with type %v is not implemented (or not mapped yet)", base.GetBoTyp())
			return errors.New(errorMessage)
		}
		if err := json.Unmarshal(data, elem); err != nil {
			return err
		}
		(*boSlice)[i] = elem
	}
	return nil
}
