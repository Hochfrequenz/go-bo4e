package bo

import (
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	"github.com/hochfrequenz/go-bo4e/enum/bilanzierungsmethode"
	"github.com/hochfrequenz/go-bo4e/enum/energierichtung"
	"github.com/hochfrequenz/go-bo4e/enum/gasqualitaet"
	"github.com/hochfrequenz/go-bo4e/enum/gebiettyp"
	"github.com/hochfrequenz/go-bo4e/enum/netzebene"
	"github.com/hochfrequenz/go-bo4e/enum/sparte"
	"github.com/hochfrequenz/go-bo4e/enum/verbrauchsart"
	"regexp"
)

// Marktlokation contains information about a market location aka "MaLo"
type Marktlokation struct {
	Geschaeftsobjekt
	MarktlokationsId     string                                    `json:"marktlokationsId" example:"12345678913" validate:"required,numeric,len=11,maloid"` // ID of the market location
	Sparte               sparte.Sparte                             `json:"sparte" validate:"required"`                                                       // Division
	Energierichtung      energierichtung.Energierichtung           `json:"energierichtung" validate:"required"`                                              // Identification of whether energy is supplied out of or fed into the grid.
	Bilanzierungsmethode bilanzierungsmethode.Bilanzierungsmethode `json:"bilanzierungsmethode" validate:"required"`                                         // Accounting method
	Verbrauchsart        verbrauchsart.Verbrauchsart               `json:"verbrauchsart,omitempty"`                                                          // Consumption type
	Unterbrechbar        bool                                      `json:"unterbrechbar,omitempty"`                                                          // Identification whether the supply is interruptible
	Netzebene            netzebene.Netzebene                       `json:"netzebene" validate:"required"`                                                    // Netzebene, in der der Bezug der Energie erfolgt. Bei Strom Spannungsebene der Lieferung, bei Gas Druckstufe.
	Netzbetreibercodenr  string                                    `json:"netzbetreibercodenr,omitempty" validate:"omitempty,numeric"`                       // Code number of the "Netzbetreiber"
	Gebiettyp            gebiettyp.Gebiettyp                       `json:"gebiettyp,omitempty"`                                                              // Type of the "Netzgebiet"
	Netzgebietnr         string                                    `json:"netzgebietnr,omitempty"`                                                           // Number of the "Netzgebiet" in the enet database
	Bilanzierungsgebiet  string                                    `json:"bilanzierungsgebiet,omitempty"`                                                    // Bilanzierungsgebiet, dem das Netzgebiet zugeordnet ist - im Falle eines Strom Netzes.
	Grundversorgercodenr string                                    `json:"grundversorgercodenr,omitempty"`                                                   // Code number of the "Grundversorger" responsible for this market location
	Gasqualitaet         gasqualitaet.Gasqualitaet                 `json:"gasqualitaet,omitempty"`                                                           // Gas Quality
	Endkunde             Geschaeftspartner                         `json:"endkunde,omitempty" validate:"-"`                                                  // Link to the Geschaeftspartner who owns this market location
	// only one of the following three optional address attributes can be set
	Lokationsadresse          *com.Adresse                 `json:"lokationsadresse" validate:"required_without_all=Geoadresse KatasterInformation"` // Address at which the energy supply or feed-in takes place
	Geoadresse                *com.Geokoordinaten          `json:"geoadresse" validate:"required_without_all=Lokationsadresse KatasterInformation"` // Gps coordinates
	Katasterinformation       *com.Katasteradresse         `json:"katasterinformation" validate:"required_without_all=Lokationsadresse Geoadresse"` // Cadastre address
	ZugehoerigeMesslokationen []com.Messlokationszuordnung `json:"zugehoerigemesslokationen,omitempty"`                                             // List of MeLos belonging to this market location
}

// MaloIdFieldLevelValidation validates the Marktlokationsid as specified by the BDEW
func MaloIdFieldLevelValidation(fl validator.FieldLevel) bool {
	maloId := fl.Field().String()
	matched, err := regexp.MatchString("[1-9][\\d]{10}$", maloId)
	if err == nil && matched {
		evenSum := 0
		oddSum := 0
		for index, digitRune := range maloId[0:10] {
			digit := int(digitRune - '0')
			if index%2 == 0 {
				oddSum = oddSum + digit
			} else {
				evenSum = evenSum + digit
			}
		}
		stepC := oddSum + (evenSum * 2)
		checkDigit := (10 - (stepC % 10)) % 10
		checkDigitIsCorrect := int(maloId[10]-'0') == checkDigit
		return checkDigitIsCorrect
	}
	return false
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
