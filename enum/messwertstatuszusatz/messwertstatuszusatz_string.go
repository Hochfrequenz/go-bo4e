// Code generated by "stringer --type Messwertstatuszusatz"; DO NOT EDIT.

package messwertstatuszusatz

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Z84_LEERSTAND-0]
	_ = x[Z85_REALERZAEHLERUEBERLAUFGEPRUEFT-1]
	_ = x[Z86_PLAUSIBELWGKONTROLLABLESUNG-2]
	_ = x[Z87_PLAUSIBELWGKUNDENHINWEIS-3]
	_ = x[ZC3_AUSTAUSCHDESERSATZWERTES-4]
	_ = x[Z88_VERGLEICHSMESSUNG_GEEICHT-5]
	_ = x[Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT-6]
	_ = x[Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN-7]
	_ = x[Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN-8]
	_ = x[Z92_INTERPOLATION-9]
	_ = x[Z93_HALTEWERT-10]
	_ = x[Z94_BILANZIERUNGNETZABSCHNITT-11]
	_ = x[Z95_HISTORISCHEMESSWERTE-12]
	_ = x[ZJ2_STATISTISCHEMETHODE-13]
	_ = x[Z74_KEINZUGANG-14]
	_ = x[Z75_KOMMUNIKATIONSSTOERUNG-15]
	_ = x[Z76_NETZAUSFALL-16]
	_ = x[Z77_SPANNUNGSAUSFALL-17]
	_ = x[Z78_GERAETEWECHSEL-18]
	_ = x[Z79_KALIBRIERUNG-19]
	_ = x[Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN-20]
	_ = x[Z81_MESSEINRICHTUNGGESTOERT_DEFEKT-21]
	_ = x[Z82_UNSICHERHEITMESSUNG-22]
	_ = x[Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK-23]
	_ = x[Z99_MENGENUMWERTUNGUNVOLLSTAENDIG-24]
	_ = x[ZA0_UHRZEITGESTELLT_SYNCHRONISATION-25]
	_ = x[ZA1_MESSWERTUNPLAUSIBEL-26]
	_ = x[ZC2_TARIFSCHALTGERAETDEFEKT-27]
	_ = x[ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND-28]
	_ = x[ZA3_FALSCHERWANDLERFAKTOR-29]
	_ = x[ZA4_FEHLERHAFTEABLESUNG-30]
	_ = x[ZA5_AENDERUNGDERBERECHNUNG-31]
	_ = x[ZA6_UMBAUDERMESSLOKATION-32]
	_ = x[ZA7_DATENBEARBEITUNGSFEHLER-33]
	_ = x[ZA8_BRENNWERTKORREKTUR-34]
	_ = x[ZA9_ZZAHL_KORREKTUR-35]
	_ = x[ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG-36]
	_ = x[ZB9_AENDERUNGTARIFSCHALTZEITEN-37]
	_ = x[ZG3_UMSTELLUNGGASQUALITAET-38]
}

const _Messwertstatuszusatz_name = "Z84_LEERSTANDZ85_REALERZAEHLERUEBERLAUFGEPRUEFTZ86_PLAUSIBELWGKONTROLLABLESUNGZ87_PLAUSIBELWGKUNDENHINWEISZC3_AUSTAUSCHDESERSATZWERTESZ88_VERGLEICHSMESSUNG_GEEICHTZ89_VERGLEICHSMESSUNG_NICHT_GEEICHTZ90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTENZ91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTENZ92_INTERPOLATIONZ93_HALTEWERTZ94_BILANZIERUNGNETZABSCHNITTZ95_HISTORISCHEMESSWERTEZJ2_STATISTISCHEMETHODEZ74_KEINZUGANGZ75_KOMMUNIKATIONSSTOERUNGZ76_NETZAUSFALLZ77_SPANNUNGSAUSFALLZ78_GERAETEWECHSELZ79_KALIBRIERUNGZ80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGENZ81_MESSEINRICHTUNGGESTOERT_DEFEKTZ82_UNSICHERHEITMESSUNGZ98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERKZ99_MENGENUMWERTUNGUNVOLLSTAENDIGZA0_UHRZEITGESTELLT_SYNCHRONISATIONZA1_MESSWERTUNPLAUSIBELZC2_TARIFSCHALTGERAETDEFEKTZC4_IMPULSWERTIGKEITNICHTAUSREICHENDZA3_FALSCHERWANDLERFAKTORZA4_FEHLERHAFTEABLESUNGZA5_AENDERUNGDERBERECHNUNGZA6_UMBAUDERMESSLOKATIONZA7_DATENBEARBEITUNGSFEHLERZA8_BRENNWERTKORREKTURZA9_ZZAHL_KORREKTURZB0_STOERUNG_DEFEKTMESSEINRICHTUNGZB9_AENDERUNGTARIFSCHALTZEITENZG3_UMSTELLUNGGASQUALITAET"

var _Messwertstatuszusatz_index = [...]uint16{0, 13, 47, 78, 106, 134, 163, 198, 239, 285, 302, 315, 344, 368, 391, 405, 431, 446, 466, 484, 500, 550, 584, 607, 648, 681, 716, 739, 766, 802, 827, 850, 876, 900, 927, 949, 968, 1002, 1032, 1058}

func (i Messwertstatuszusatz) String() string {
	if i < 0 || i >= Messwertstatuszusatz(len(_Messwertstatuszusatz_index)-1) {
		return "Messwertstatuszusatz(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Messwertstatuszusatz_name[_Messwertstatuszusatz_index[i]:_Messwertstatuszusatz_index[i+1]]
}
