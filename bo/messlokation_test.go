package bo

import (
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// TestFailedMesslokationValidation verifies that the validators of Messlokation work
func (s *Suite) TestFailedMesslokationValidation() {

}

//  TestSuccessfulMesslokationValidation verifies that a valid BO is validated without errors
func (s *Suite) TestSuccessfulMesslokationValidation() {
	validate := validator.New()
	validMelos := []interface{}{
		Messlokation{
			BusinessObject: BusinessObject{
				BoTyp:             botyp.Messlokation,
				VersionStruktur:   "1",
				ExterneReferenzen: nil,
			},
			MesslokationsId:              "DE000001111122222333334444455555666667",
			Sparte:                       sparte.Strom,
			NetzebeneMessung:             netzebene.MD,
			MessgebietNr:                 "",
			Gerate:                       nil,
			Messdienstleistung:           com.Dienstleistung{},
			GrundzustaendigerMsbCodeNr:   "",
			GrundzustaendigerMsbImCodeNr: "",
			Messadresse:                  com.Adresse{},
			Geoadresse:                   com.Geokoordinaten{},
			Katasterinformation:          com.Katasteradresse{},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMelos)
}
