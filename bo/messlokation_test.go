package bo

import (
	"github.com/go-playground/validator"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/landescode"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// TestFailedMesslokationValidation verifies that the validators of Messlokation work
func (s *Suite) TestFailedMesslokationValidation() {
	validate := validator.New()
	invalidMesslokationMap := map[string][]interface{}{
		"required": {
			Messlokation{
				BusinessObject: BusinessObject{
					BoTyp:             0,
					VersionStruktur:   "",
					ExterneReferenzen: nil,
				},
				MesslokationsId:              "",
				Sparte:                       0,
				NetzebeneMessung:             0,
				MessgebietNr:                 "",
				Gerate:                       nil,
				Messdienstleistung:           nil,
				GrundzustaendigerMsbCodeNr:   "",
				GrundzustaendigerMsbImCodeNr: "",
				Messadresse:                  nil,
			},
		},
		"len": {
			Messlokation{
				MesslokationsId: "not33",
			},
		},
	}
	VerfiyFailedValidations(s, validate, invalidMesslokationMap)
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
			MesslokationsId:              "DE0000011111222223333344444555556",
			Sparte:                       sparte.Strom,
			NetzebeneMessung:             netzebene.MD,
			MessgebietNr:                 "",
			Gerate:                       nil,
			Messdienstleistung:           nil,
			GrundzustaendigerMsbCodeNr:   "",
			GrundzustaendigerMsbImCodeNr: "",
			Messadresse: &com.Adresse{
				Postleitzahl: "82031",
				Ort:          "Grünwald",
				Strasse:      "Nördliche Münchner Straße",
				Hausnummer:   "27A",
				Landescode:   landescode.DE,
			},
		},
	}
	VerfiySuccessfulValidations(s, validate, validMelos)
}
