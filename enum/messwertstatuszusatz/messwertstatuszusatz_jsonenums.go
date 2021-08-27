// Code generated by jsonenums --type Messwertstatuszusatz; DO NOT EDIT.

package messwertstatuszusatz

import (
	"encoding/json"
	"fmt"
)

var (
	_MesswertstatuszusatzNameToValue = map[string]Messwertstatuszusatz{
		"Z84_LEERSTAND":                                      Z84_LEERSTAND,
		"Z85_REALERZAEHLERUEBERLAUFGEPRUEFT":                 Z85_REALERZAEHLERUEBERLAUFGEPRUEFT,
		"Z86_PLAUSIBELWGKONTROLLABLESUNG":                    Z86_PLAUSIBELWGKONTROLLABLESUNG,
		"Z87_PLAUSIBELWGKUNDENHINWEIS":                       Z87_PLAUSIBELWGKUNDENHINWEIS,
		"ZC3_AUSTAUSCHDESERSATZWERTES":                       ZC3_AUSTAUSCHDESERSATZWERTES,
		"Z88_VERGLEICHSMESSUNG_GEEICHT":                      Z88_VERGLEICHSMESSUNG_GEEICHT,
		"Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT":                Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT,
		"Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN":          Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN,
		"Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN":     Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN,
		"Z92_INTERPOLATION":                                  Z92_INTERPOLATION,
		"Z93_HALTEWERT":                                      Z93_HALTEWERT,
		"Z94_BILANZIERUNGNETZABSCHNITT":                      Z94_BILANZIERUNGNETZABSCHNITT,
		"Z95_HISTORISCHEMESSWERTE":                           Z95_HISTORISCHEMESSWERTE,
		"ZJ2_STATISTISCHEMETHODE":                            ZJ2_STATISTISCHEMETHODE,
		"Z74_KEINZUGANG":                                     Z74_KEINZUGANG,
		"Z75_KOMMUNIKATIONSSTOERUNG":                         Z75_KOMMUNIKATIONSSTOERUNG,
		"Z76_NETZAUSFALL":                                    Z76_NETZAUSFALL,
		"Z77_SPANNUNGSAUSFALL":                               Z77_SPANNUNGSAUSFALL,
		"Z78_GERAETEWECHSEL":                                 Z78_GERAETEWECHSEL,
		"Z79_KALIBRIERUNG":                                   Z79_KALIBRIERUNG,
		"Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN": Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN,
		"Z81_MESSEINRICHTUNGGESTOERT_DEFEKT":                 Z81_MESSEINRICHTUNGGESTOERT_DEFEKT,
		"Z82_UNSICHERHEITMESSUNG":                            Z82_UNSICHERHEITMESSUNG,
		"Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK":          Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK,
		"Z99_MENGENUMWERTUNGUNVOLLSTAENDIG":                  Z99_MENGENUMWERTUNGUNVOLLSTAENDIG,
		"ZA0_UHRZEITGESTELLT_SYNCHRONISATION":                ZA0_UHRZEITGESTELLT_SYNCHRONISATION,
		"ZA1_MESSWERTUNPLAUSIBEL":                            ZA1_MESSWERTUNPLAUSIBEL,
		"ZC2_TARIFSCHALTGERAETDEFEKT":                        ZC2_TARIFSCHALTGERAETDEFEKT,
		"ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND":               ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND,
		"ZA3_FALSCHERWANDLERFAKTOR":                          ZA3_FALSCHERWANDLERFAKTOR,
		"ZA4_FEHLERHAFTEABLESUNG":                            ZA4_FEHLERHAFTEABLESUNG,
		"ZA5_AENDERUNGDERBERECHNUNG":                         ZA5_AENDERUNGDERBERECHNUNG,
		"ZA6_UMBAUDERMESSLOKATION":                           ZA6_UMBAUDERMESSLOKATION,
		"ZA7_DATENBEARBEITUNGSFEHLER":                        ZA7_DATENBEARBEITUNGSFEHLER,
		"ZA8_BRENNWERTKORREKTUR":                             ZA8_BRENNWERTKORREKTUR,
		"ZA9_ZZAHL_KORREKTUR":                                ZA9_ZZAHL_KORREKTUR,
		"ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG":                 ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG,
		"ZB9_AENDERUNGTARIFSCHALTZEITEN":                     ZB9_AENDERUNGTARIFSCHALTZEITEN,
		"ZG3_UMSTELLUNGGASQUALITAET":                         ZG3_UMSTELLUNGGASQUALITAET,
	}

	_MesswertstatuszusatzValueToName = map[Messwertstatuszusatz]string{
		Z84_LEERSTAND:                                      "Z84_LEERSTAND",
		Z85_REALERZAEHLERUEBERLAUFGEPRUEFT:                 "Z85_REALERZAEHLERUEBERLAUFGEPRUEFT",
		Z86_PLAUSIBELWGKONTROLLABLESUNG:                    "Z86_PLAUSIBELWGKONTROLLABLESUNG",
		Z87_PLAUSIBELWGKUNDENHINWEIS:                       "Z87_PLAUSIBELWGKUNDENHINWEIS",
		ZC3_AUSTAUSCHDESERSATZWERTES:                       "ZC3_AUSTAUSCHDESERSATZWERTES",
		Z88_VERGLEICHSMESSUNG_GEEICHT:                      "Z88_VERGLEICHSMESSUNG_GEEICHT",
		Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT:                "Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT",
		Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN:          "Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN",
		Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN:     "Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN",
		Z92_INTERPOLATION:                                  "Z92_INTERPOLATION",
		Z93_HALTEWERT:                                      "Z93_HALTEWERT",
		Z94_BILANZIERUNGNETZABSCHNITT:                      "Z94_BILANZIERUNGNETZABSCHNITT",
		Z95_HISTORISCHEMESSWERTE:                           "Z95_HISTORISCHEMESSWERTE",
		ZJ2_STATISTISCHEMETHODE:                            "ZJ2_STATISTISCHEMETHODE",
		Z74_KEINZUGANG:                                     "Z74_KEINZUGANG",
		Z75_KOMMUNIKATIONSSTOERUNG:                         "Z75_KOMMUNIKATIONSSTOERUNG",
		Z76_NETZAUSFALL:                                    "Z76_NETZAUSFALL",
		Z77_SPANNUNGSAUSFALL:                               "Z77_SPANNUNGSAUSFALL",
		Z78_GERAETEWECHSEL:                                 "Z78_GERAETEWECHSEL",
		Z79_KALIBRIERUNG:                                   "Z79_KALIBRIERUNG",
		Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN: "Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN",
		Z81_MESSEINRICHTUNGGESTOERT_DEFEKT:                 "Z81_MESSEINRICHTUNGGESTOERT_DEFEKT",
		Z82_UNSICHERHEITMESSUNG:                            "Z82_UNSICHERHEITMESSUNG",
		Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK:          "Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK",
		Z99_MENGENUMWERTUNGUNVOLLSTAENDIG:                  "Z99_MENGENUMWERTUNGUNVOLLSTAENDIG",
		ZA0_UHRZEITGESTELLT_SYNCHRONISATION:                "ZA0_UHRZEITGESTELLT_SYNCHRONISATION",
		ZA1_MESSWERTUNPLAUSIBEL:                            "ZA1_MESSWERTUNPLAUSIBEL",
		ZC2_TARIFSCHALTGERAETDEFEKT:                        "ZC2_TARIFSCHALTGERAETDEFEKT",
		ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND:               "ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND",
		ZA3_FALSCHERWANDLERFAKTOR:                          "ZA3_FALSCHERWANDLERFAKTOR",
		ZA4_FEHLERHAFTEABLESUNG:                            "ZA4_FEHLERHAFTEABLESUNG",
		ZA5_AENDERUNGDERBERECHNUNG:                         "ZA5_AENDERUNGDERBERECHNUNG",
		ZA6_UMBAUDERMESSLOKATION:                           "ZA6_UMBAUDERMESSLOKATION",
		ZA7_DATENBEARBEITUNGSFEHLER:                        "ZA7_DATENBEARBEITUNGSFEHLER",
		ZA8_BRENNWERTKORREKTUR:                             "ZA8_BRENNWERTKORREKTUR",
		ZA9_ZZAHL_KORREKTUR:                                "ZA9_ZZAHL_KORREKTUR",
		ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG:                 "ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG",
		ZB9_AENDERUNGTARIFSCHALTZEITEN:                     "ZB9_AENDERUNGTARIFSCHALTZEITEN",
		ZG3_UMSTELLUNGGASQUALITAET:                         "ZG3_UMSTELLUNGGASQUALITAET",
	}
)

func init() {
	var v Messwertstatuszusatz
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_MesswertstatuszusatzNameToValue = map[string]Messwertstatuszusatz{
			interface{}(Z84_LEERSTAND).(fmt.Stringer).String():                                      Z84_LEERSTAND,
			interface{}(Z85_REALERZAEHLERUEBERLAUFGEPRUEFT).(fmt.Stringer).String():                 Z85_REALERZAEHLERUEBERLAUFGEPRUEFT,
			interface{}(Z86_PLAUSIBELWGKONTROLLABLESUNG).(fmt.Stringer).String():                    Z86_PLAUSIBELWGKONTROLLABLESUNG,
			interface{}(Z87_PLAUSIBELWGKUNDENHINWEIS).(fmt.Stringer).String():                       Z87_PLAUSIBELWGKUNDENHINWEIS,
			interface{}(ZC3_AUSTAUSCHDESERSATZWERTES).(fmt.Stringer).String():                       ZC3_AUSTAUSCHDESERSATZWERTES,
			interface{}(Z88_VERGLEICHSMESSUNG_GEEICHT).(fmt.Stringer).String():                      Z88_VERGLEICHSMESSUNG_GEEICHT,
			interface{}(Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT).(fmt.Stringer).String():                Z89_VERGLEICHSMESSUNG_NICHT_GEEICHT,
			interface{}(Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN).(fmt.Stringer).String():          Z90_MESSWERTNACHBILDUNGAUSGEEICHTENWERTEN,
			interface{}(Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN).(fmt.Stringer).String():     Z91_MESSWERTNACHBILDUNGAUSNICHTGEEICHTENWERTEN,
			interface{}(Z92_INTERPOLATION).(fmt.Stringer).String():                                  Z92_INTERPOLATION,
			interface{}(Z93_HALTEWERT).(fmt.Stringer).String():                                      Z93_HALTEWERT,
			interface{}(Z94_BILANZIERUNGNETZABSCHNITT).(fmt.Stringer).String():                      Z94_BILANZIERUNGNETZABSCHNITT,
			interface{}(Z95_HISTORISCHEMESSWERTE).(fmt.Stringer).String():                           Z95_HISTORISCHEMESSWERTE,
			interface{}(ZJ2_STATISTISCHEMETHODE).(fmt.Stringer).String():                            ZJ2_STATISTISCHEMETHODE,
			interface{}(Z74_KEINZUGANG).(fmt.Stringer).String():                                     Z74_KEINZUGANG,
			interface{}(Z75_KOMMUNIKATIONSSTOERUNG).(fmt.Stringer).String():                         Z75_KOMMUNIKATIONSSTOERUNG,
			interface{}(Z76_NETZAUSFALL).(fmt.Stringer).String():                                    Z76_NETZAUSFALL,
			interface{}(Z77_SPANNUNGSAUSFALL).(fmt.Stringer).String():                               Z77_SPANNUNGSAUSFALL,
			interface{}(Z78_GERAETEWECHSEL).(fmt.Stringer).String():                                 Z78_GERAETEWECHSEL,
			interface{}(Z79_KALIBRIERUNG).(fmt.Stringer).String():                                   Z79_KALIBRIERUNG,
			interface{}(Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN).(fmt.Stringer).String(): Z80_GERAETARBEITETAUSSERHALBDERBETRIEBSBEDINGUNGEN,
			interface{}(Z81_MESSEINRICHTUNGGESTOERT_DEFEKT).(fmt.Stringer).String():                 Z81_MESSEINRICHTUNGGESTOERT_DEFEKT,
			interface{}(Z82_UNSICHERHEITMESSUNG).(fmt.Stringer).String():                            Z82_UNSICHERHEITMESSUNG,
			interface{}(Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK).(fmt.Stringer).String():          Z98_BERUECKSICHTIGUNGSTOERMENGENZAEHLWERK,
			interface{}(Z99_MENGENUMWERTUNGUNVOLLSTAENDIG).(fmt.Stringer).String():                  Z99_MENGENUMWERTUNGUNVOLLSTAENDIG,
			interface{}(ZA0_UHRZEITGESTELLT_SYNCHRONISATION).(fmt.Stringer).String():                ZA0_UHRZEITGESTELLT_SYNCHRONISATION,
			interface{}(ZA1_MESSWERTUNPLAUSIBEL).(fmt.Stringer).String():                            ZA1_MESSWERTUNPLAUSIBEL,
			interface{}(ZC2_TARIFSCHALTGERAETDEFEKT).(fmt.Stringer).String():                        ZC2_TARIFSCHALTGERAETDEFEKT,
			interface{}(ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND).(fmt.Stringer).String():               ZC4_IMPULSWERTIGKEITNICHTAUSREICHEND,
			interface{}(ZA3_FALSCHERWANDLERFAKTOR).(fmt.Stringer).String():                          ZA3_FALSCHERWANDLERFAKTOR,
			interface{}(ZA4_FEHLERHAFTEABLESUNG).(fmt.Stringer).String():                            ZA4_FEHLERHAFTEABLESUNG,
			interface{}(ZA5_AENDERUNGDERBERECHNUNG).(fmt.Stringer).String():                         ZA5_AENDERUNGDERBERECHNUNG,
			interface{}(ZA6_UMBAUDERMESSLOKATION).(fmt.Stringer).String():                           ZA6_UMBAUDERMESSLOKATION,
			interface{}(ZA7_DATENBEARBEITUNGSFEHLER).(fmt.Stringer).String():                        ZA7_DATENBEARBEITUNGSFEHLER,
			interface{}(ZA8_BRENNWERTKORREKTUR).(fmt.Stringer).String():                             ZA8_BRENNWERTKORREKTUR,
			interface{}(ZA9_ZZAHL_KORREKTUR).(fmt.Stringer).String():                                ZA9_ZZAHL_KORREKTUR,
			interface{}(ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG).(fmt.Stringer).String():                 ZB0_STOERUNG_DEFEKTMESSEINRICHTUNG,
			interface{}(ZB9_AENDERUNGTARIFSCHALTZEITEN).(fmt.Stringer).String():                     ZB9_AENDERUNGTARIFSCHALTZEITEN,
			interface{}(ZG3_UMSTELLUNGGASQUALITAET).(fmt.Stringer).String():                         ZG3_UMSTELLUNGGASQUALITAET,
		}
	}
}

// MarshalJSON is generated so Messwertstatuszusatz satisfies json.Marshaler.
func (r Messwertstatuszusatz) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _MesswertstatuszusatzValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Messwertstatuszusatz: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Messwertstatuszusatz satisfies json.Unmarshaler.
func (r *Messwertstatuszusatz) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Messwertstatuszusatz should be a string, got %s", data)
	}
	v, ok := _MesswertstatuszusatzNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Messwertstatuszusatz %q", s)
	}
	*r = v
	return nil
}
