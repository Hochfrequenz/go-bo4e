package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
)

// Energiemenge contains information about consumption at a location
type Energiemenge struct {
	Geschaeftsobjekt
	LokationsId  string                    `json:"lokationsId" example:"DE0123456789012345678901234567890" validate:"alphanum,required,min=11,max=33"` // ID of the location (either a LokationsTyp MALO ID (11 digits) or a LokationsTyp MELO ID (33 alphanum))
	LokationsTyp lokationstyp.Lokationstyp `json:"lokationsTyp" example:"MELO" validate:"required"`                                                    // type of the location in LokationsId
	Verbrauch    []com.Verbrauch           `json:"energieverbrauch" validate:"required,min=1"`                                                         // consumption data (see Verbrauch)
}
