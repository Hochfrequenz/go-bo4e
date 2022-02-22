package bo

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hochfrequenz/go-bo4e/com"
	aggregationsverantwortung "github.com/hochfrequenz/go-bo4e/enum/aggregationsverwantwortung"
	"github.com/hochfrequenz/go-bo4e/enum/fallgruppenzuordnung"
	"github.com/hochfrequenz/go-bo4e/enum/profiltyp"
	"github.com/hochfrequenz/go-bo4e/enum/prognosegrundlage"
	"github.com/hochfrequenz/go-bo4e/enum/zeitreihentyp"
	"github.com/shopspring/decimal"
	"regexp"
	"time"
)

// Bilanzierung is a business object used for balancing.
// Note that this is not official BO4E standard (yet)!
type Bilanzierung struct {
	Geschaeftsobjekt
	Lastprofile              []com.Lastprofil `json:"lastprofile,omitempty"`                                                         // Lastprofile ist eine Liste der verwendeten Lastprofile (SLP, SLP/TLP, ALP etc.)
	Bilanzierungsbeginn      time.Time        `json:"bilanzierungsbeginn,omitempty" validate:"omitempty,ltefield=Bilanzierungsende"` // Bilanzierungsbeginn ist der inklusive Beginn der Bilanzierung
	Bilanzierungsende        time.Time        `json:"bilanzierungsende,omitempty" validate:"omitempty,gtefield=Bilanzierungsbeginn"` // Bilanzierungsende ist das exklusive Ende der Bilanzierung
	Bilanzkreis              string           `json:"bilanzkreis,omitempty" validate:"omitempty,len=16,eic"`                         // Bilanzkreis ist der EIC-Code des Bilanzkreises
	Jahresverbrauchsprognose *com.Menge       `json:"jahresverbrauchsprognose,omitempty"`                                            // Jahresverbrauchsprognose ist die Jahresverbrauchsprognose
	Kundenwert               *com.Menge       `json:"kundenwert,omitempty"`                                                          // Kundenwert ist der Kundenwert
	// Verbrauchsaufteilung beschreibt, welcher Anteil im SLP- bzw. TLP-Profil steckt
	// 1. [Gemessene Energiemenge der OBIS "nicht Schwachlast"] * [Verbrauchsaufteilung in % / 100%] = [zu verlagernde Energiemenge]
	// 2. [Gemessene Energiemenge der OBIS "Schwachlast"] - [zu verlagernde Energiemenge] = [Ermittelte Energiemenge für Schwachlast]
	// 3. [Gemessene Energiemenge der OBIS "nicht Schwachlast"] + [zu verlagernde Energiemenge] = [Ermittelte Energiemenge für nicht Schwachlast]
	Verbrauchsaufteilung       decimal.NullDecimal                                 `json:"verbrauchsaufteilung,omitempty"`
	Zeitreihentyp              zeitreihentyp.Zeitreihentyp                         `json:"zeitreihentyp,omitempty"`                                // Zeitreihentyp beschreibt den verwendeten Zeitreihentyp (SLS, TLS...)
	Aggregationsverantwortung  aggregationsverantwortung.Aggregationsverantwortung `json:"aggregationsverantwortung,omitempty"`                    // Aggregationsverantwortung benennt, bei wem die Aggregationsverantwortung liegt
	Prognosegrundlage          prognosegrundlage.Prognosegrundlage                 `json:"prognosegrundlage,omitempty"`                            // Die Prognosegrundlage beschreibt die Prognosegrundlage
	DetailsPrognosegrundlage   []profiltyp.Profiltyp                               `json:"detailsPrognosegrundlage,omitempty"`                     // ?
	WahlrechtPrognosegrundlage *bool                                               `json:"wahlrechtPrognosegrundlage,omitempty"`                   // WahlrechtPrognosegrundlage ist true, falls der Lieferant das Wahlrecht hat
	Fallgruppenzuordnung       fallgruppenzuordnung.Fallgruppenzuordnung           `json:"fallgruppenzuordnung,omitempty"`                         // Fallgruppenzuordnung (für Gas RLM)
	Prioritaet                 *int                                                `json:"prioritaet,omitempty"`                                   // Prioritaet ist die Priorität des Bilanzkreises für Gas
	MarktlokationsId           string                                              `json:"marktlokationsId,omitempty" validate:"omitempty,maloid"` // MarktlokationsId referenziert eine Marktlokation
}

func (_ Bilanzierung) GetDefaultJsonTags() []string {
	return bilanzierungJsonKeys
}

// the bilanzierungForMarshal type is derived from Bilanzierung but uses a different unmarshaler/deserializer so that we don't run into an endless recursion when overriding the UnmarshalJSON func to include our hacky workaround
type bilanzierungForUnmarshal Bilanzierung

func (bila *Bilanzierung) UnmarshalJSON(bytes []byte) (err error) {
	// the approach we use here is often referred to as "unmarshal twice"
	// it's a workaround for the not-so-mature/feature rich encoding/json framework in go
	_malo := bilanzierungForUnmarshal{}
	// first we deserialize into the helper/intermediate type. this is to _not_ run into this Unmarshal func in an endless recursion
	if err = json.Unmarshal(bytes, &_malo); err == nil {
		*bila = Bilanzierung(_malo)
		// the bilanzierung contains only those fields that are defined in the Bilanzierung struct by now
	} else {
		return err
	}

	// now we unmarshal the same bytes into the extension data
	if err = json.Unmarshal(bytes, &bila.ExtensionData); err != nil {
		return nil
	}
	// But now the extension data also contain those fields that in fact have a representation in the Bilanzierung struct
	bila.RemoveStronglyTypedFieldsFromExtensionData(bila.GetDefaultJsonTags()) // remove those fields from the extension data that have a representation in the Bilanzierung struct
	return nil
}

// bilanzierungForMarshal is a struct similar to the original Bilanzierung but uses a different Marshaller so that we don't run into an endless recursion
type bilanzierungForMarshal Bilanzierung

//nolint:dupl // This can only be simplified if we use generics. anything else seems overly complicated but maybe it's just me
func (bila Bilanzierung) MarshalJSON() ([]byte, error) {
	if bila.ExtensionData == nil || len(bila.ExtensionData) == 0 {
		// no special handling needed
		return json.Marshal(bilanzierungForMarshal(bila)) // just marshal but use a helper type to not run into an endless recursion
	}
	// there is probably a better way than this, like f.e. creating an adhoc-struct that has an embedded bila-like type and f.e. a map or
	// we first convert the Bilanzierung into a map[string]interface{}
	// we serialize the bilanzierung via our helper type
	bilaDictBytes, bilaMarshalErr := json.Marshal(bilanzierungForMarshal(bila)) // we must use a different type here to not run into an endless recursion
	if bilaMarshalErr != nil {
		return []byte{}, bilaMarshalErr
	}
	// now we deserialize the malo again but we use a generic dictionary as target type
	bilaAsMap := map[string]interface{}{}
	extensionUnmarshalErr := json.Unmarshal(bilaDictBytes, &bilaAsMap)
	if extensionUnmarshalErr != nil {
		return []byte{}, extensionUnmarshalErr
	}
	// now we join/merge the original malo and its extension data (which is already a map[string]interface{} into a single result
	result := map[string]interface{}{}
	for key, value := range bilaAsMap {
		result[key] = value
	}
	for key, value := range bila.ExtensionData {
		result[key] = value
	}
	return json.Marshal(result)
}

// bilanzierungJsonKeys is a list of all keys in the standard bo4e json Bilanzierung.
// It is used to distinguish fields that can be mapped to the Marktlokation struct and those that are moved to Geschaeftsobjekt.ExtensionData
var bilanzierungJsonKeys = []string{
	// https://c.tenor.com/71HGq_GX1pMAAAAC/kill-me-simpsons.gif
	// there has to be a better way than this.
	"boTyp",
	"versionStruktur",
	"lastprofile",
	"bilanzierungsbeginn",
	"bilanzierungsende",
	"bilanzkreis",
	"jahresverbrauchsprognose",
	"kundenwert",
	"verbrauchsaufteilung",
	"aggregationsverantwortung",
	"prognosegrundlage",
	"detailsPrognosegrundlage",
	"wahlrechtPrognosegrundlage",
	"fallgruppenzuordnung",
	"prioritaet",
	"marktlokationsId",
}

var eicRegex = regexp.MustCompile(`(?P<vergabestelle>\d{2})(?P<typ>A|T|V|W|X|Y|Z)([-A-Z\d]{12})(?P<pruefziffer>[A-Z0-9])`)

// EICFieldLevelValidation ensures that the string stored in the annotated/tagged field is a valid EIC code
// Therefore it has to obey the "Bildungsvorschrift" according to https://bdew-codes.de/Content/Files/EIC/Awh_20171218_EIC-Vergabe_V1-0.pdf section 2.2.1
func EICFieldLevelValidation(fl validator.FieldLevel) bool {
	eicCandidate := fl.Field().String()
	matched := eicRegex.MatchString(eicCandidate)
	if matched {
		actualCheckCharacter := eicCandidate[len(eicCandidate)-1:]
		expectedCheckCharacter := getEICCheckCharacter(eicCandidate[0 : len(eicCandidate)-1])
		return actualCheckCharacter == expectedCheckCharacter
	}
	return false
}

// getEICCheckCharacter calculates the check character according to the EIC rules (ENTSO-E)
// See https://eepublicdownloads.entsoe.eu/clean-documents/EDI/Library/cim_based/02%20EIC%20Code%20implementation%20guide_final.pdf line 461 (section 7.1)
func getEICCheckCharacter(eicWithoutChecksum string) string {
	// a priori assumption: eicWithoutChecksum is exactly 15 characters long
	var numericValues []int // step 2
	for _, c := range eicWithoutChecksum {
		numericValue, err := eicCharacterToNumber(string(c))
		if err != nil {
			panic(err.Error())
		}
		numericValues = append(numericValues, numericValue)
	}
	sum := 0
	for i, numericValue := range numericValues {
		// step 3&4
		positionWeight := 16 - i
		numericValues[i] = numericValue * positionWeight
		sum = sum + numericValues[i] // step 5
	}
	var checkCharacter string
	checkNumber := 36 - (sum-1)%37
	if checkNumber < 10 {
		checkCharacter = fmt.Sprint(checkNumber)
	} else {
		var err error
		checkCharacter, err = numberToEicCharacter(rune(checkNumber))
		if err != nil {
			panic(err)
		}
		if checkCharacter == "-" {
			errorMessage := fmt.Sprintf("The check character must not be '-'. Please choose other 15 characters than '%s' to start with", eicWithoutChecksum)
			// https://eepublicdownloads.entsoe.eu/clean-documents/EDI/Library/cim_based/02%20EIC%20Code%20implementation%20guide_final.pdf line 509
			panic(errorMessage)
		}
	}
	return checkCharacter
}

var charToIntEicMap = map[string]int{
	"-": 36,
	// ok, this seems dumb but at least it's explicit and readable
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

// eicCharacterToNumber converts the character s to a number according to the ENTSO-E/EIC algorithm
func eicCharacterToNumber(s string) (int, error) {
	if len(s) != 1 {
		return 0, fmt.Errorf("you must only provide a len=1 string but the length was %d", len(s))
	}
	result, found := charToIntEicMap[s]
	if !found {
		result = int(s[0]) - 55 // in ASCII 65=A but in ENTSO 10=A
	}
	return result, nil
}

// numberToEicCharacter converts the given number/rune to a character according to the ENTSO-E/EIC algorithm
// returns an error if the rune is out of range
func numberToEicCharacter(r rune) (string, error) {
	if r < 10 || r > 36 {
		return "", fmt.Errorf("only the range [10,36] is supported")
	}
	if r == rune(36) {
		return "-", nil
	}
	return string(r + 55), nil
}
