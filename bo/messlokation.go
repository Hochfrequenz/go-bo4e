package bo

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/enum/gasqualitaet"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
)

// Messlokation contains information about a metering location aka "MeLo"
type Messlokation struct {
	Geschaeftsobjekt
	MesslokationsId              string              `json:"messlokationsId,omitempty" example:"DE0123456789012345678901234567890" validate:"alphanum,required,len=33"` // MesslokationsId is the ID of the metering location
	Sparte                       sparte.Sparte       `json:"sparte,omitempty" validate:"required"`                                                                      // Sparte is the division
	NetzebeneMessung             netzebene.Netzebene `json:"netzebeneMessung,omitempty"`                                                                                // NetzebeneMessung is the grid level of measurement
	MessgebietNr                 string              `json:"messgebietNr,omitempty"`                                                                                    // MessgebietNr is the number of the measurement area in ene't database
	Geraete                      []com.Hardware      `json:"geraete,omitempty"`                                                                                         // Geraete is a list of devices
	Messdienstleistung           *com.Dienstleistung `json:"messdienstleistung,omitempty"`                                                                              // Messdienstleistung is a metering services
	GrundzustaendigerMsbCodeNr   string              `json:"grundzustaendigerMSBCodeNr,omitempty" validate:"omitempty,numeric,len=13"`                                  // GrundzustaendigerMsbCodeNr is the code number of the "grundzuständige Messstellenbetreiber", responsitble for this MeLo
	GrundzustaendigerMsbImCodeNr string              `json:"GrundzustaendigerMsbImCodeNr,omitempty" validate:"omitempty,numeric,len=13"`                                // GrundzustaendigerMsbImCodeNr si the code number of the "grundzuständige Messsstellenbetreiber", responsible for intelligent meters at this MeLo
	Messlokationszaehler         []Zaehler           `json:"messlokationszaehler,omitempty"`                                                                            // Messlokationszaehler meters associated to this Messlokation
	Gasqualitaet                    *gasqualitaet.Gasqualitaet `json:"gasqualitaet"`                                                                                              // gasqualitaet für EDIFACT mapping
	Abrechnungmessstellenbetriebnna *bool                      `json:"abrechnungmessstellenbetriebnna"`                                                                           //Dieser Wert ist true, falls die Abrechnungs des Messstellenbetriebs die Netznutzungsabrechnung enthält. false andernfalls
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

// the messlokationForUnmarshal type is derived from Messlokation but uses a different unmarshaler/deserializer so that we don't run into an endless recursion when overriding the UnmarshalJSON func to include our hacky workaround
type messlokationForUnmarshal Messlokation

func (melo *Messlokation) UnmarshalJSON(bytes []byte) (err error) {
	// the approach we use here is often referred to as "unmarshal twice"
	// it's a workaround for the not-so-mature/feature rich encoding/json framework in go
	_melo := messlokationForUnmarshal{}
	// first we deserialize into the helper/intermediate type. this is to _not_ run into this Unmarshal func in an endless recursion
	if err = json.Unmarshal(bytes, &_melo); err == nil {
		*melo = Messlokation(_melo)
		// the malo contains only those fields that are defined in the Messlokation struct by now
	} else {
		return err
	}

	// now we unmarshal the same bytes into the extension data
	if err = json.Unmarshal(bytes, &melo.ExtensionData); err != nil {
		return nil
	}
	// But now the extension data also contain those fields that in fact have a representation in the Messlokation struct
	melo.RemoveStronglyTypedFieldsFromExtensionData(melo.GetDefaultJsonTags()) // remove those fields from the extension data that have a representation in the Marktlokation struct
	return nil
}

// marktlokationForMarshal is a struct similar to the original Messlokation but uses a different Marshaller so that we don't run into an endless recursion
type messlokationForMarshal Messlokation

//nolint:dupl // This can only be simplified if we use generics. anything else seems overly complicated but maybe it's just me
func (melo Messlokation) MarshalJSON() ([]byte, error) {
	if melo.ExtensionData == nil || len(melo.ExtensionData) == 0 {
		// no special handling needed
		return json.Marshal(messlokationForMarshal(melo)) // just marshal but use a helper type to not run into an endless recursino
	}
	// there is probably a better way than this, like f.e. creating an adhoc-struct that has an embedded malo-like type and f.e. a map or
	// we first convert the Marktlokation into a map[string]interface{}
	// we serialize the malo via our helper type
	maloDictBytes, maloMarhsalErr := json.Marshal(messlokationForMarshal(melo)) // we must use a different type here to not run into an endless recursion
	if maloMarhsalErr != nil {
		return []byte{}, maloMarhsalErr
	}
	// now we deserialize the malo again but we use a generic dictionary as target type
	meloAsMap := map[string]interface{}{}
	extensionUnmarshalErr := json.Unmarshal(maloDictBytes, &meloAsMap)
	if extensionUnmarshalErr != nil {
		return []byte{}, extensionUnmarshalErr
	}
	// now we join/merge the original malo and its extension data (which is already a map[string]interface{} into a single result
	result := map[string]interface{}{}
	for key, value := range meloAsMap {
		result[key] = value
	}
	for key, value := range melo.ExtensionData {
		result[key] = value
	}
	return json.Marshal(result)
}
