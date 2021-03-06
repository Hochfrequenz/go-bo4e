// Code generated by "stringer --type FehlerCode"; DO NOT EDIT.

package fehlercode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ID_UNBEKANNT-1]
	_ = x[ABSENDER_NICHT_ZUGEORDNET-2]
	_ = x[EMPFAENGER_NICHT_ZUGEORDNET-3]
	_ = x[GERAET_UNBEKANNT-4]
	_ = x[OBIS_UNBEKANNT-5]
	_ = x[REFERENZIERUNG_FEHLERHAFT-6]
	_ = x[TUPEL_UNBEKANNT-7]
	_ = x[ABSENDER_TUPEL_NICHT_ZUGEORDNET-8]
	_ = x[EMPFAENGER_TUPEL_NICHT_ZUGEORDNET-9]
	_ = x[VORKOMMA_ZU_VIELE_STELLEN-10]
	_ = x[ZEITREIHE_UNVOLLSTAENDIG-11]
	_ = x[REFERENZIERTES_TUPEL_UNBEKANNT-12]
	_ = x[MARKTLOKATION_UNBEKANNT-13]
	_ = x[MESSLOKATION_UNBEKANNT-14]
	_ = x[MELDEPUNKT_NICHT_MEHR_IM_NETZ-15]
	_ = x[ERFORDERLICHE_ANGABE_FEHLT-16]
	_ = x[GESCHAEFTSVORFALL_ZURUECKGEWIESEN-17]
	_ = x[ZEITINTERVALL_NEGATIV-18]
	_ = x[FORMAT_NICHT_EINGEHALTEN-19]
	_ = x[GESCHAEFTSVORFALL_ABSENDER-20]
}

const _FehlerCode_name = "ID_UNBEKANNTABSENDER_NICHT_ZUGEORDNETEMPFAENGER_NICHT_ZUGEORDNETGERAET_UNBEKANNTOBIS_UNBEKANNTREFERENZIERUNG_FEHLERHAFTTUPEL_UNBEKANNTABSENDER_TUPEL_NICHT_ZUGEORDNETEMPFAENGER_TUPEL_NICHT_ZUGEORDNETVORKOMMA_ZU_VIELE_STELLENZEITREIHE_UNVOLLSTAENDIGREFERENZIERTES_TUPEL_UNBEKANNTMARKTLOKATION_UNBEKANNTMESSLOKATION_UNBEKANNTMELDEPUNKT_NICHT_MEHR_IM_NETZERFORDERLICHE_ANGABE_FEHLTGESCHAEFTSVORFALL_ZURUECKGEWIESENZEITINTERVALL_NEGATIVFORMAT_NICHT_EINGEHALTENGESCHAEFTSVORFALL_ABSENDER"

var _FehlerCode_index = [...]uint16{0, 12, 37, 64, 80, 94, 119, 134, 165, 198, 223, 247, 277, 300, 322, 351, 377, 410, 431, 455, 481}

func (i FehlerCode) String() string {
	i -= 1
	if i < 0 || i >= FehlerCode(len(_FehlerCode_index)-1) {
		return "FehlerCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _FehlerCode_name[_FehlerCode_index[i]:_FehlerCode_index[i+1]]
}
