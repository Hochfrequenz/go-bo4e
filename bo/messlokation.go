package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Messlokation contains information about a metering location aka "MeLo"
type Messlokation struct {
	BusinessObject
	MesslokationsId              string              `json:"messlokationsId" example:"DE0123456789012345678901234567890" validate:"alphanum,required,len=33"` // ID of the metering location
	Sparte                       sparte.Sparte       `json:"sparte" validate:"required"`                                                                             // division
	NetzebeneMessung             netzebene.Netzebene `json:"netzebeneMessung"`                                                                                       // grid level of measurement
	MessgebietNr                 string              `json:"messgebietNr,omitempty"`                                                                                 // number of the measurement area in ene't database
	Gerate                       []com.Hardware      `json:"geraete,omitempty"`                                                                                      // list of devices
	Messdienstleistung           com.Dienstleistung  `json:"messdienstleistung,omitempty"`                                                                           // list of metering services
	GrundzustaendigerMsbCodeNr   string              `json:"grundzustaendigerMSBCodeNr,omitempty" validate:"numeric"`                                                // Code number of the "grundzuständige Messstellenbetreiber", responsitble for this MeLo
	GrundzustaendigerMsbImCodeNr string              `json:"GrundzustaendigerMsbImCodeNr,omitempty" validate:"numeric"`                                              // Code number of the "grundzuständige Messsstellenbetreiber", responsible for intelligent meters at this MeLo
	// ToDo implement BO zaehler
	//	messlokationszaehler: Optional[List[Zaehler]] = attr.ib(default=None)
	//
	//
	// only one of the following three optional address attributes can be set
	Messadresse         com.Adresse         `json:"messadresse" validate:"required_without_all=Geoadresse KatasterInformation"` // address of the melo
	Geoadresse          com.Geokoordinaten  `json:"geoadresse" validate:"required_without_all=Messadresse KatasterInformation"` // gps coordinates
	Katasterinformation com.Katasteradresse `json:"katasterinformation" validate:"required_without_all=Messadresse Geoadresse"` // cadastre address
}
