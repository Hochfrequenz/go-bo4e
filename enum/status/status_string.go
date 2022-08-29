// Code generated by "stringer --type Status"; DO NOT EDIT.

package status

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TARIF_1-1]
	_ = x[TARIF_2-2]
	_ = x[TARIF_3-3]
	_ = x[TARIF_4-4]
	_ = x[TARIF_5-5]
	_ = x[TARIF_6-6]
	_ = x[TARIF_7-7]
	_ = x[TARIF_8-8]
	_ = x[TARIF_9-9]
	_ = x[ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT-10]
	_ = x[ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERT-11]
	_ = x[ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG-12]
	_ = x[ZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNG-13]
	_ = x[KEIN_ZUGANG-14]
	_ = x[KOMMUNIKATIONSSTOERUNG-15]
	_ = x[NETZAUSFALL-16]
	_ = x[SPANNUNGSAUSFALL-17]
	_ = x[STATUS_GERAETEWECHSEL-18]
	_ = x[KALIBRIERUNG-19]
	_ = x[GERAET_ARBEITET_AUSSERHALB_DER_BETRIEBSBEDINGUNGEN-20]
	_ = x[MESSEINRICHTUNG_GESTOERT_DEFEKT-21]
	_ = x[UNSICHERHEIT_MESSUNG-22]
	_ = x[KUNDENSELBSTABLESUNG-23]
	_ = x[LEERSTAND-24]
	_ = x[REALER_ZAEHLERUEBERLAUF_GEPRUEFT-25]
	_ = x[PLAUSIBEL_WG_KONTROLLABLESUNG-26]
	_ = x[PLAUSIBEL_WG_KUNDENHINWEIS-27]
	_ = x[VERGLEICHSMESSUNG_GEEICHT-28]
	_ = x[VERGLEICHSMESSUNG_NICHT_GEEICHT-29]
	_ = x[MESSWERTNACHBILDUNG_AUS_GEEICHTEN_WERTEN-30]
	_ = x[MESSWERTNACHBILDUNG_AUS_NICHT_GEEICHTEN_WERTEN-31]
	_ = x[INTERPOLATION-32]
	_ = x[HALTEWERT-33]
	_ = x[BILANZIERUNG_NETZABSCHNITT-34]
	_ = x[HISTORISCHE_MESSWERTE-35]
	_ = x[BERUECKSICHTIGUNG_STOERMENGENZAEHLWERK-36]
	_ = x[MENGENUMWERTUNG_VOLLSTAENDIG-37]
	_ = x[UHRZEIT_GESTELLT_SYNCHRONISATION-38]
	_ = x[MESSWERT_UNPLAUSIBEL-39]
	_ = x[FALSCHER_WANDLERFAKTOR-40]
	_ = x[FEHLERHAFTE_ABLESUNG-41]
	_ = x[AENDERUNG_DER_BERECHNUNG-42]
	_ = x[UMBAU_DER_MESSLOKATION-43]
	_ = x[DATENBEARBEITUNGSFEHLER-44]
	_ = x[BRENNWERTKORREKTUR-45]
	_ = x[Z_ZAHL_KORREKTUR-46]
	_ = x[STOERUNG_DEFEKT_MESSEINRICHTUNG-47]
	_ = x[AENDERUNG_TARIFSCHALTZEITEN-48]
	_ = x[TARIFSCHALTGERAET_DEFEKT-49]
	_ = x[AUSTAUSCH_DES_ERSATZWERTES-50]
	_ = x[IMPULSWERTIGKEIT_NICHT_AUSREICHEND-51]
	_ = x[UMSTELLUNG_GASQUALITAET-52]
	_ = x[STATISTISCHE_METHODE-53]
	_ = x[ENERGIEMENGE_IN_UNGEMESSENEM_ZEITINTERVALL-54]
	_ = x[ENERGIEMENGE_AUS_DEM_UNGEPAIRTEN_ZEITINTERVALL-55]
	_ = x[AUFTEILUNG-56]
	_ = x[VERWENDUNG_VON_WERTEN_DES_STOERMENGENZAEHLWERKS-57]
	_ = x[UMGANGS_UND_KORREKTURMENGEN-58]
	_ = x[WARTUNGSARBEITEN_AN_GEEICHTEM_MESSGERAET-59]
	_ = x[GESTOERTE_WERTE-60]
	_ = x[WARTUNGSARBEITEN_AN_EICHRECHTSKONFORMEN_MESSGERAETEN-61]
	_ = x[KONSISTENZ_UND_SYNCHRONPRUEFUNG-62]
	_ = x[RECHENWERT-63]
	_ = x[ANGABEN_MESSLOKATION-64]
	_ = x[BASIS_MME-65]
	_ = x[GRUND_ANGABEN_MESSLOKATION-66]
}

const _Status_name = "TARIF_1TARIF_2TARIF_3TARIF_4TARIF_5TARIF_6TARIF_7TARIF_8TARIF_9ZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERTZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_VORHANDEN_UND_KOMMUNIZIERTZAEHLERSTAND_ZUM_BEGINN_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNGZAEHLERSTAND_ZUM_ENDE_DER_ANGEGEBENEN_ENERGIEMENGE_NICHT_VORHANDEN_DA_MENGENABGRENZUNGKEIN_ZUGANGKOMMUNIKATIONSSTOERUNGNETZAUSFALLSPANNUNGSAUSFALLSTATUS_GERAETEWECHSELKALIBRIERUNGGERAET_ARBEITET_AUSSERHALB_DER_BETRIEBSBEDINGUNGENMESSEINRICHTUNG_GESTOERT_DEFEKTUNSICHERHEIT_MESSUNGKUNDENSELBSTABLESUNGLEERSTANDREALER_ZAEHLERUEBERLAUF_GEPRUEFTPLAUSIBEL_WG_KONTROLLABLESUNGPLAUSIBEL_WG_KUNDENHINWEISVERGLEICHSMESSUNG_GEEICHTVERGLEICHSMESSUNG_NICHT_GEEICHTMESSWERTNACHBILDUNG_AUS_GEEICHTEN_WERTENMESSWERTNACHBILDUNG_AUS_NICHT_GEEICHTEN_WERTENINTERPOLATIONHALTEWERTBILANZIERUNG_NETZABSCHNITTHISTORISCHE_MESSWERTEBERUECKSICHTIGUNG_STOERMENGENZAEHLWERKMENGENUMWERTUNG_VOLLSTAENDIGUHRZEIT_GESTELLT_SYNCHRONISATIONMESSWERT_UNPLAUSIBELFALSCHER_WANDLERFAKTORFEHLERHAFTE_ABLESUNGAENDERUNG_DER_BERECHNUNGUMBAU_DER_MESSLOKATIONDATENBEARBEITUNGSFEHLERBRENNWERTKORREKTURZ_ZAHL_KORREKTURSTOERUNG_DEFEKT_MESSEINRICHTUNGAENDERUNG_TARIFSCHALTZEITENTARIFSCHALTGERAET_DEFEKTAUSTAUSCH_DES_ERSATZWERTESIMPULSWERTIGKEIT_NICHT_AUSREICHENDUMSTELLUNG_GASQUALITAETSTATISTISCHE_METHODEENERGIEMENGE_IN_UNGEMESSENEM_ZEITINTERVALLENERGIEMENGE_AUS_DEM_UNGEPAIRTEN_ZEITINTERVALLAUFTEILUNGVERWENDUNG_VON_WERTEN_DES_STOERMENGENZAEHLWERKSUMGANGS_UND_KORREKTURMENGENWARTUNGSARBEITEN_AN_GEEICHTEM_MESSGERAETGESTOERTE_WERTEWARTUNGSARBEITEN_AN_EICHRECHTSKONFORMEN_MESSGERAETENKONSISTENZ_UND_SYNCHRONPRUEFUNGRECHENWERTANGABEN_MESSLOKATIONBASIS_MMEGRUND_ANGABEN_MESSLOKATION"

var _Status_index = [...]uint16{0, 7, 14, 21, 28, 35, 42, 49, 56, 63, 142, 219, 307, 393, 404, 426, 437, 453, 474, 486, 536, 567, 587, 607, 616, 648, 677, 703, 728, 759, 799, 845, 858, 867, 893, 914, 952, 980, 1012, 1032, 1054, 1074, 1098, 1120, 1143, 1161, 1177, 1208, 1235, 1259, 1285, 1319, 1342, 1362, 1404, 1450, 1460, 1507, 1534, 1574, 1589, 1641, 1672, 1682, 1702, 1711, 1737}

func (i Status) String() string {
	i -= 1
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
