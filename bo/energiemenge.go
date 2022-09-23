package bo

import (
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
	"github.com/hochfrequenz/go-bo4e/enum/lokationstyp"
)

// Energiemenge contains information about consumption at a location
type Energiemenge struct {
	Geschaeftsobjekt
	LokationsId  string                    `json:"lokationsId,omitempty" example:"DE0123456789012345678901234567890" validate:"alphanum,required,min=11,max=33"` // LokationsId is the ID of the location (either a LokationsTyp MALO ID (11 digits) or a LokationsTyp MELO ID (33 alphanum))
	LokationsTyp lokationstyp.Lokationstyp `json:"lokationsTyp,omitempty" example:"MELO" validate:"required"`                                                    // LokationsTyp is the type of the location in LokationsId
	Verbrauch    []com.Verbrauch           `json:"energieverbrauch,omitempty" validate:"required,min=1"`                                                         // Verbrauch are consumption data
}

func (_ Energiemenge) GetDefaultJsonTags() []string {
	panic("todo: implement me") // this is needed for (un)marshaling of non-default/unknown json fields
}
func (_ Energiemenge) GetBoTyp() botyp.BOTyp {
	// this is useful for generics to work
	return botyp.ENERGIEMENGE
}
