package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
)

// Messlokation contains information about a metering location aka "MeLo"
type Messlokation struct {
	BusinessObject
	MesslokationsId  string              `json:"messlokationsId" example:"DE0123456789012345678901234567890" validate:"alphanumeric,required"` // ID of the metering location
	Sparte           sparte.Sparte       `json:"sparte" validate:"required"`                                                                   // division
	NetzebeneMessung netzebene.Netzebene `json:"netzebeneMessung"`                                                                             // grid level of measurement
	// ToDo
	//	messgebietnr: Optional[str] = attr.ib(default=None)
	//	geraete: Optional[List[Hardware]] = attr.ib(default=None)
	//	messdienstleistung: Optional[List[Dienstleistung]] = attr.ib(default=None)
	//	messlokationszaehler: Optional[List[Zaehler]] = attr.ib(default=None)
	//
	//	# only one of the following two optional codenr attributes can be set
	//	grundzustaendiger_msb_codenr: Optional[str] = attr.ib(default=None)
	//	grundzustaendiger_msbim_codenr: Optional[str] = attr.ib(default=None)
	//
	//	# only one of the following three optional address attributes can be set
	Messadresse com.Adresse `json:"messadresse"` // address of the melo
	//	geoadresse: Optional[Geokoordinaten] = attr.ib(default=None)
	//	katasterinformation: Optional[Katasteradresse] = attr.ib(default=None)
}
