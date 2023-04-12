package bo

import (
	"encoding/json"
	"github.com/hochfrequenz/go-bo4e/enum/messtechnischeeinordnung"
	"github.com/hochfrequenz/go-bo4e/enum/sperrstatus"
	"github.com/hochfrequenz/go-bo4e/enum/zeitreihentyp"
	"github.com/hochfrequenz/go-bo4e/internal/unmappeddatamarshaller"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/bilanzierungsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/gasqualitaet"
	"github.com/hochfrequenz/go-bo4e/enum/gebiettyp"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/verbrauchsart"
	"github.com/hochfrequenz/go-bo4e/internal/jsonfieldnames"
)

// Marktlokation contains information about a market location aka "MaLo"
type Marktlokation struct {
	unmappeddatamarshaller.ExtensionData
	Geschaeftsobjekt
	MarktlokationsId     string                                    `json:"marktlokationsId,omitempty" example:"12345678913" validate:"required,numeric,len=11,maloid"` // MarktlokationsId is the ID of the market location
	Sparte               sparte.Sparte                             `json:"sparte,omitempty" validate:"required"`                                                       // Sparte describes the Division
	Energierichtung      energierichtung.Energierichtung           `json:"energierichtung,omitempty" validate:"required"`                                              // Energierichtung describes whether energy is supplied out of or fed into the grid.
	Bilanzierungsmethode bilanzierungsmethode.Bilanzierungsmethode `json:"bilanzierungsmethode,omitempty" validate:"required"`                                         // Bilanzierungsmethode is the accounting method
	Verbrauchsart        verbrauchsart.Verbrauchsart               `json:"verbrauchsart,omitempty"`                                                                    // Verbrauchsart is the consumption type
	Unterbrechbar        *bool                                     `json:"unterbrechbar,omitempty"`                                                                    // Unterbrechbar describes whether the supply is interruptible
	Netzebene            netzebene.Netzebene                       `json:"netzebene,omitempty" validate:"required"`                                                    // Netzebene, in der der Bezug der Energie erfolgt. Bei Strom Spannungsebene der Lieferung, bei Gas Druckstufe.
	Netzbetreibercodenr  string                                    `json:"netzbetreibercodenr,omitempty" validate:"omitempty,numeric,len=13"`                          // Netzbetreibercodenr is the code number of the "Netzbetreiber"
	Gebiettyp            gebiettyp.Gebiettyp                       `json:"gebiettyp,omitempty"`                                                                        // Gebiettyp is the type of the "Netzgebiet"
	Netzgebietnr         string                                    `json:"netzgebietnr,omitempty"`                                                                     // Netzgebietnr is the number of the "Netzgebiet" in the enet database
	Bilanzierungsgebiet  string                                    `json:"bilanzierungsgebiet,omitempty"`                                                              // Bilanzierungsgebiet, dem das Netzgebiet zugeordnet ist - im Falle eines Strom Netzes.
	Grundversorgercodenr string                                    `json:"grundversorgercodenr,omitempty" validate:"omitempty,numeric,len=13"`                         // Grundversorgercodenr is the code number of the "Grundversorger" responsible for this market location
	Gasqualitaet         gasqualitaet.Gasqualitaet                 `json:"gasqualitaet,omitempty"`                                                                     // Gasqualitaet is the gas quality
	Endkunde             *Geschaeftspartner                        `json:"endkunde,omitempty" validate:"-"`                                                            // Endkunde is the Geschaeftspartner who owns this market location
	// only one of the following three optional address attributes can be set
	Lokationsadresse          *com.Adresse                 `json:"lokationsadresse,omitempty" validate:"required_without_all=Geoadresse Katasterinformation"` // Lokationsadresse is the address at which the energy supply or feed-in takes place
	Geoadresse                *com.Geokoordinaten          `json:"geoadresse,omitempty" validate:"required_without_all=Lokationsadresse Katasterinformation"` // Geoadresse are the gps coordinates
	Katasterinformation       *com.Katasteradresse         `json:"katasterinformation,omitempty" validate:"required_without_all=Lokationsadresse Geoadresse"` // Katasterinformation is the Cadastre address
	ZugehoerigeMesslokationen []com.Messlokationszuordnung `json:"zugehoerigemesslokationen,omitempty"`                                                       // ZugehoerigeMesslokationen is a list of MeLos belonging to this market location

	Regelzone   *string `json:"regelzone,omitempty"`
	Marktgebiet *string `json:"marktgebiet,omitempty"`

	Zeitreihentyp                  *zeitreihentyp.Zeitreihentyp                       `json:"zeitreihentyp,omitempty"`
	Zaehlwerke                     []com.Zaehlwerk                                    `json:"zaehlwerke,omitempty"`
	ZaehlwerkeBeteiligteMarktrolle []com.Zaehlwerk                                    `json:"zaehlwerkeBeteiligteMarktrolle,omitempty"`
	Messlokationen                 []Messlokation                                     `json:"messlokationen,omitempty"`
	MesstechnischeEinordnung       *messtechnischeeinordnung.MesstechnischeEinordnung `json:"messtechnischeEinordnung,omitempty"`

	Netznutzungsabrechnungsdaten []com.Netznutzungsabrechnungsdaten `json:"netznutzungsabrechnungsdaten,omitempty"` // Netznutzungsabrechnungsdaten sind Daten für die Prüfung der Netznutzungsabrechnung
	Sperrstatus                  *sperrstatus.Sperrstatus           `json:"sperrstatus,omitempty"`
}

// the marktlokationForUnmarshal type is derived from Marktlokation but uses a different unmarshaler/deserializer so that we don't run into an endless recursion when overriding the UnmarshalJSON func to include our hacky workaround
type marktlokationForUnmarshal Marktlokation

func (malo *Marktlokation) UnmarshalJSON(bytes []byte) (err error) {
	return unmappeddatamarshaller.UnmarshallWithUnmappedData(malo, &malo.ExtensionData, bytes)
}

func (malo Marktlokation) MarshalJSON() (bytes []byte, err error) {
	s := marktlokationForUnmarshal(malo)

	byteArr, err := json.Marshal(s)
	if err != nil {
		return
	}
	return unmappeddatamarshaller.HandleUnmappedDataPropertyMarshalling(byteArr)
}

func (malo Marktlokation) GetDefaultJsonTags() []string {
	// We know we pass a struct here so ignore the error.
	fields, _ := jsonfieldnames.Extract(malo)
	return fields
}

var elevenDigitsRegex = regexp.MustCompile(`^[1-9]\d{10}$`)

// MaloIdFieldLevelValidation validates the Marktlokationsid as specified by the BDEW
// https://bdew-codes.de/Content/Files/MaLo/2017-04-28-BDEW-Anwendungshilfe-MaLo-ID_Version1.0_FINAL.PDF
func MaloIdFieldLevelValidation(fl validator.FieldLevel) bool {
	maloId := fl.Field().String()
	matched := elevenDigitsRegex.MatchString(maloId)
	if matched {
		checkDigit := GetMaLoIdCheckSum(maloId)
		checkDigitIsCorrect := int(maloId[10]-'0') == checkDigit
		return checkDigitIsCorrect
	}
	return false
}

// GetMaLoIdCheckSum returns the checksum (11th character of the malo ID) that matches the first ten characters long provided in maloIdWithoutCheckSum. This is going to crash if the length of the maloIdWithoutCheckSum is <10
func GetMaLoIdCheckSum(maloIdWithoutCheckSum string) int {
	evenSum := 0
	oddSum := 0
	for index, digitRune := range maloIdWithoutCheckSum[0:10] {
		digit := int(digitRune - '0')
		if index%2 == 0 {
			oddSum = oddSum + digit
		} else {
			evenSum = evenSum + digit
		}
	}
	stepC := oddSum + (evenSum * 2)
	return (10 - (stepC % 10)) % 10
}

// XorStructLevelValidation ensures that only one of the possible address types is given
func XorStructLevelValidation(sl validator.StructLevel) {
	malo := sl.Current().Interface().(Marktlokation)
	numberOfAdresses := 0
	addressesExists := []bool{malo.Lokationsadresse != nil, malo.Geoadresse != nil, malo.Katasterinformation != nil}
	for _, adressExists := range addressesExists {
		if adressExists {
			numberOfAdresses++
		}
	}
	if numberOfAdresses > 1 {
		sl.ReportError(Marktlokation{}.Lokationsadresse, "", "", "onlyOneAddressType", "")
	}
}
