package market_communication

import (
	"regexp"

	"github.com/hochfrequenz/go-bo4e/bo"
	"github.com/hochfrequenz/go-bo4e/enum/botyp"
)

// receiverKey is the key for BOneyComb.Transaktionsdaten under which the 13 digit ID of the receiving Marktteilnehmer can be found
const receiverKey = "empfaenger"

// senderKey is the key for BOneyComb.Transaktionsdaten under which the 13 digit ID of the sending Marktteilnehmer can be found
const senderKey = "absender"

// GetAbsenderCode returns the 13 digit ID of the sending Marktteilnehmer if present in the Transaktionsdaten; nil otherwise
func (boneyComb *BOneyComb) GetAbsenderCode() *string {
	return boneyComb.getMpCode(senderKey)
}

// SetAbsenderCode sets the 13 digit ID of the sending Marktteilnehmer in the Transaktionsdaten
func (boneyComb *BOneyComb) SetAbsenderCode(mpId string, useBo4eUri bool) {
	boneyComb.initializeTransaktionsdaten()
	boneyComb.Transaktionsdaten[senderKey] = addBo4eMarktteilnehmerPrefixIfNecessary(mpId, useBo4eUri)
}

// GetEmpfaengerCode returns the 13 digit ID of the receiving Marktteilnehmer if present in both Transaktionsdaten; nil otherwise
func (boneyComb *BOneyComb) GetEmpfaengerCode() *string {
	return boneyComb.getMpCode(receiverKey)
}

// SetEmpfaengerCode sets the 13 digit ID of the receiving Marktteilnehmer in the Transaktionsdaten
func (boneyComb *BOneyComb) SetEmpfaengerCode(mpId string, useBo4eUri bool) {
	boneyComb.initializeTransaktionsdaten()
	boneyComb.Transaktionsdaten[receiverKey] = addBo4eMarktteilnehmerPrefixIfNecessary(mpId, useBo4eUri)
}

func addBo4eMarktteilnehmerPrefixIfNecessary(mpId string, useBo4eUri bool) string {
	if useBo4eUri {
		return "bo4e://Marktteilnehmer/" + mpId
	} else {
		return mpId
	}
}

// GetAbsender returns the sending bo.Marktteilnehmer if present in the Transaktionsdaten _and_ Stammdaten; nil otherwise
func (boneyComb *BOneyComb) GetAbsender() *bo.Marktteilnehmer {
	return boneyComb.getMarktteilnehmer(boneyComb.GetAbsenderCode())
}

// SetAbsender sets sending bo.Marktteilnehmer in both the Transaktionsdaten _and_ Stammdaten
func (boneyComb *BOneyComb) SetAbsender(mt bo.Marktteilnehmer, useBo4eUri bool) {
	if mt.Rollencodenummer != nil {
		boneyComb.SetAbsenderCode(*mt.Rollencodenummer, useBo4eUri)
	}
	boneyComb.setMarktteilnehmer(mt)
}

// GetEmpfaenger returns the receiving bo.Marktteilnehmer if present in the Transaktionsdaten _and_ Stammdaten; nil otherwise
func (boneyComb *BOneyComb) GetEmpfaenger() *bo.Marktteilnehmer {
	return boneyComb.getMarktteilnehmer(boneyComb.GetEmpfaengerCode())
}

// SetEmpfaenger sets receiving bo.Marktteilnehmer in both the Transaktionsdaten _and_ Stammdaten
func (boneyComb *BOneyComb) SetEmpfaenger(mt bo.Marktteilnehmer, useBo4eUri bool) {
	if mt.Rollencodenummer != nil {

		boneyComb.SetEmpfaengerCode(*mt.Rollencodenummer, useBo4eUri)
	}
	boneyComb.setMarktteilnehmer(mt)
}

// we expect the 'marktpartner id' to usually be the 13 digits (power/gas sector) or a system internal
// key (as for 'monopol sparten', e.g. water sector)
var bo4eUriRegex = regexp.MustCompile(`(bo4e:\/\/Marktteilnehmer\/)?(?P<mpid>\d{13}|\w+)`)

// getMpCode returns the 13 digit ID of a Marktteilnehmer if present in the Transaktionsdaten
func (boneyComb *BOneyComb) getMpCode(transactionDataKey string) *string {
	mpString := boneyComb.GetTransactionData(transactionDataKey)
	if mpString == nil {
		return nil
	}
	groupNames := bo4eUriRegex.SubexpNames()
	for _, match := range bo4eUriRegex.FindAllStringSubmatch(*mpString, -1) {
		for groupIdx, group := range match {
			name := groupNames[groupIdx]
			if name == "mpid" {
				return &group
			}
		}
	}
	return nil
}

// getMarktteilnehmer returns the Marktteilnehmer from Stammdaten that has the 13 digit ID returned by getMpCode
func (boneyComb *BOneyComb) getMarktteilnehmer(marktteilnehmerId *string) *bo.Marktteilnehmer {
	if marktteilnehmerId == nil || *marktteilnehmerId == "" {
		return nil
	}
	if boneyComb.Stammdaten == nil {
		return nil
	}
	for _, businessObject := range boneyComb.Stammdaten {
		if businessObject.GetBoTyp() == botyp.MARKTTEILNEHMER {
			if *(businessObject.(*bo.Marktteilnehmer).Rollencodenummer) == *marktteilnehmerId {
				mtn := businessObject.(*bo.Marktteilnehmer)
				return mtn
			}
		}
	}
	return nil
}

// setMarktteilnehmer adds the Marktteilnehmer to the Stammdaten slice
func (boneyComb *BOneyComb) setMarktteilnehmer(mt bo.Marktteilnehmer) {
	if boneyComb.Stammdaten == nil {
		boneyComb.Stammdaten = bo.BusinessObjectSlice{}
	}
	boneyComb.Stammdaten = append(boneyComb.Stammdaten, &mt)
}
