package bo

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/gasqualitaet"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
)

// Messlokation contains information about a metering location aka "MeLo"
type Messlokation struct {
	Geschaeftsobjekt
	MesslokationsId                 string                     `json:"messlokationsId,omitempty" example:"DE0123456789012345678901234567890" validate:"alphanum,required,len=33"` // MesslokationsId is the ID of the metering location
	Sparte                          sparte.Sparte              `json:"sparte,omitempty" validate:"required"`                                                                      // Sparte is the division
	NetzebeneMessung                *netzebene.Netzebene       `json:"netzebeneMessung,omitempty"`                                                                                // NetzebeneMessung is the grid level of measurement
	MessgebietNr                    string                     `json:"messgebietNr,omitempty"`                                                                                    // MessgebietNr is the number of the measurement area in ene't database
	Geraete                         []com.Hardware             `json:"geraete,omitempty"`                                                                                         // Geraete is a list of devices
	Messdienstleistung              *com.Dienstleistung        `json:"messdienstleistung,omitempty"`                                                                              // Messdienstleistung is a metering services
	GrundzustaendigerMsbCodeNr      string                     `json:"grundzustaendigerMSBCodeNr,omitempty" validate:"omitempty,numeric,len=13"`                                  // GrundzustaendigerMsbCodeNr is the code number of the "grundzuständige Messstellenbetreiber", responsitble for this MeLo
	GrundzustaendigerMsbImCodeNr    string                     `json:"GrundzustaendigerMsbImCodeNr,omitempty" validate:"omitempty,numeric,len=13"`                                // GrundzustaendigerMsbImCodeNr si the code number of the "grundzuständige Messsstellenbetreiber", responsible for intelligent meters at this MeLo
	Messlokationszaehler            []Zaehler                  `json:"messlokationszaehler,omitempty"`                                                                            // Messlokationszaehler meters associated to this Messlokation
	Gasqualitaet                    *gasqualitaet.Gasqualitaet `json:"gasqualitaet,omitempty"`                                                                                    // gasqualitaet für EDIFACT mapping
	Abrechnungmessstellenbetriebnna *bool                      `json:"abrechnungmessstellenbetriebnna,omitempty"`                                                                 //Dieser Wert ist true, falls die Abrechnungs des Messstellenbetriebs die Netznutzungsabrechnung enthält. false andernfalls
	// only one of the following three optional address attributes can be set
	Messadresse         *com.Adresse         `json:"messadresse,omitempty" validate:"required_without_all=Geoadresse Katasterinformation"` // Messadresse is a street address of the Messlokation
	Geoadresse          *com.Geokoordinaten  `json:"geoadresse,omitempty" validate:"required_without_all=Messadresse Katasterinformation"` // Geoadresse are gps coordinates of the Messlokation
	Katasterinformation *com.Katasteradresse `json:"katasterinformation,omitempty" validate:"required_without_all=Messadresse Geoadresse"` // Katasterinformation is a cadastre address of the Messlokation
}

// XorStructLevelMesslokationValidation ensures that only one of the possible address types is given
func XorStructLevelMesslokationValidation(sl validator.StructLevel) {
	melo := sl.Current().Interface().(Messlokation)
	numberOfAdresses := 0
	addressesExists := []bool{melo.Messadresse != nil, melo.Geoadresse != nil, melo.Katasterinformation != nil}
	for _, adressExists := range addressesExists {
		if adressExists {
			numberOfAdresses++
		}
	}
	if numberOfAdresses > 1 {
		sl.ReportError(Messlokation{}.Messadresse, "", "", "onlyOneAddressType", "")
	}
}

func (melo Messlokation) GetDefaultJsonTags() []string {
	// We know we pass a struct here so ignore the error.
	fields, _ := jsonfieldnames.Extract(melo)
	return fields
}

// the below code is just copy pasted from the malo
// so sorry

func (melo *Messlokation) UnmarshalJSON(bytes []byte) (err error) {
	if melo.UnmappedData.Data == nil {
		melo.UnmappedData.Data = map[string]any{}
	}
	return unmappeddatamarshaller.UnmarshallWithUnmappedData(melo, &melo.UnmappedData, bytes)
}

func (melo Messlokation) MarshalJSON() ([]byte, error) {
	return unmappeddatamarshaller.MarshalWithUnmappedData(melo)
}
