// Code generated by jsonenums --type BDEWArtikelnummer; DO NOT EDIT.

package bdewartikelnummer

import (
	"encoding/json"
	"fmt"
)

var (
	_BDEWArtikelnummerNameToValue = map[string]BDEWArtikelnummer{
		"LEISTUNG":                                   LEISTUNG,
		"LEISTUNG_PAUSCHAL":                          LEISTUNG_PAUSCHAL,
		"GRUNDPREIS":                                 GRUNDPREIS,
		"REGELENERGIE_ARBEIT":                        REGELENERGIE_ARBEIT,
		"REGELENERGIE_LEISTUNG":                      REGELENERGIE_LEISTUNG,
		"NOTSTROMLIEFERUNG_ARBEIT":                   NOTSTROMLIEFERUNG_ARBEIT,
		"NOTSTROMLIEFERUNG_LEISTUNG":                 NOTSTROMLIEFERUNG_LEISTUNG,
		"RESERVENETZKAPAZITAET":                      RESERVENETZKAPAZITAET,
		"RESERVELEISTUNG":                            RESERVELEISTUNG,
		"ZUSAETZLICHE_ABLESUNG":                      ZUSAETZLICHE_ABLESUNG,
		"PRUEFGEBUEHREN_AUSSERPLANMAESSIG":           PRUEFGEBUEHREN_AUSSERPLANMAESSIG,
		"WIRKARBEIT":                                 WIRKARBEIT,
		"SINGULAER_GENUTZTE_BETRIEBSMITTEL":          SINGULAER_GENUTZTE_BETRIEBSMITTEL,
		"ABGABE_KWKG":                                ABGABE_KWKG,
		"ABSCHLAG":                                   ABSCHLAG,
		"KONZESSIONSABGABE":                          KONZESSIONSABGABE,
		"ENTGELT_FERNAUSLESUNG":                      ENTGELT_FERNAUSLESUNG,
		"UNTERMESSUNG":                               UNTERMESSUNG,
		"BLINDMEHRARBEIT":                            BLINDMEHRARBEIT,
		"ENTGELT_ABRECHNUNG":                         ENTGELT_ABRECHNUNG,
		"SPERRKOSTEN":                                SPERRKOSTEN,
		"ENTSPERRKOSTEN":                             ENTSPERRKOSTEN,
		"MAHNKOSTEN":                                 MAHNKOSTEN,
		"MEHR_MINDERMENGEN":                          MEHR_MINDERMENGEN,
		"INKASSOKOSTEN":                              INKASSOKOSTEN,
		"BLINDMEHRLEISTUNG":                          BLINDMEHRLEISTUNG,
		"ENTGELT_MESSUNG_ABLESUNG":                   ENTGELT_MESSUNG_ABLESUNG,
		"ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK": ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK,
		"AUSGLEICHSENERGIE":                          AUSGLEICHSENERGIE,
		"ZAEHLEINRICHTUNG":                           ZAEHLEINRICHTUNG,
		"WANDLER_MENGENUMWERTER":                     WANDLER_MENGENUMWERTER,
		"KOMMUNIKATIONSEINRICHTUNG":                  KOMMUNIKATIONSEINRICHTUNG,
		"TECHNISCHE_STEUEREINRICHTUNG":               TECHNISCHE_STEUEREINRICHTUNG,
		"PARAGRAF_19_STROM_NEV_UMLAGE":               PARAGRAF_19_STROM_NEV_UMLAGE,
		"BEFESTIGUNGSEINRICHTUNG":                    BEFESTIGUNGSEINRICHTUNG,
		"OFFSHORE_HAFTUNGSUMLAGE":                    OFFSHORE_HAFTUNGSUMLAGE,
		"FIXE_ARBEITSENTGELTKOMPONENTE":              FIXE_ARBEITSENTGELTKOMPONENTE,
		"FIXE_LEISTUNGSENTGELTKOMPONENTE":            FIXE_LEISTUNGSENTGELTKOMPONENTE,
		"UMLAGE_ABSCHALTBARE_LASTEN":                 UMLAGE_ABSCHALTBARE_LASTEN,
		"MEHRMENGE":                                  MEHRMENGE,
		"MINDERMENGE":                                MINDERMENGE,
		"ENERGIESTEUER":                              ENERGIESTEUER,
		"SMARTMETER_GATEWAY":                         SMARTMETER_GATEWAY,
		"STEUERBOX":                                  STEUERBOX,
		"MSB_INKL_MESSUNG":                           MSB_INKL_MESSUNG,
		"AUSGLEICHSENERGIE_UNTERDECKUNG":             AUSGLEICHSENERGIE_UNTERDECKUNG,
	}

	_BDEWArtikelnummerValueToName = map[BDEWArtikelnummer]string{
		LEISTUNG:                          "LEISTUNG",
		LEISTUNG_PAUSCHAL:                 "LEISTUNG_PAUSCHAL",
		GRUNDPREIS:                        "GRUNDPREIS",
		REGELENERGIE_ARBEIT:               "REGELENERGIE_ARBEIT",
		REGELENERGIE_LEISTUNG:             "REGELENERGIE_LEISTUNG",
		NOTSTROMLIEFERUNG_ARBEIT:          "NOTSTROMLIEFERUNG_ARBEIT",
		NOTSTROMLIEFERUNG_LEISTUNG:        "NOTSTROMLIEFERUNG_LEISTUNG",
		RESERVENETZKAPAZITAET:             "RESERVENETZKAPAZITAET",
		RESERVELEISTUNG:                   "RESERVELEISTUNG",
		ZUSAETZLICHE_ABLESUNG:             "ZUSAETZLICHE_ABLESUNG",
		PRUEFGEBUEHREN_AUSSERPLANMAESSIG:  "PRUEFGEBUEHREN_AUSSERPLANMAESSIG",
		WIRKARBEIT:                        "WIRKARBEIT",
		SINGULAER_GENUTZTE_BETRIEBSMITTEL: "SINGULAER_GENUTZTE_BETRIEBSMITTEL",
		ABGABE_KWKG:                       "ABGABE_KWKG",
		ABSCHLAG:                          "ABSCHLAG",
		KONZESSIONSABGABE:                 "KONZESSIONSABGABE",
		ENTGELT_FERNAUSLESUNG:             "ENTGELT_FERNAUSLESUNG",
		UNTERMESSUNG:                      "UNTERMESSUNG",
		BLINDMEHRARBEIT:                   "BLINDMEHRARBEIT",
		ENTGELT_ABRECHNUNG:                "ENTGELT_ABRECHNUNG",
		SPERRKOSTEN:                       "SPERRKOSTEN",
		ENTSPERRKOSTEN:                    "ENTSPERRKOSTEN",
		MAHNKOSTEN:                        "MAHNKOSTEN",
		MEHR_MINDERMENGEN:                 "MEHR_MINDERMENGEN",
		INKASSOKOSTEN:                     "INKASSOKOSTEN",
		BLINDMEHRLEISTUNG:                 "BLINDMEHRLEISTUNG",
		ENTGELT_MESSUNG_ABLESUNG:          "ENTGELT_MESSUNG_ABLESUNG",
		ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK: "ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK",
		AUSGLEICHSENERGIE:                          "AUSGLEICHSENERGIE",
		ZAEHLEINRICHTUNG:                           "ZAEHLEINRICHTUNG",
		WANDLER_MENGENUMWERTER:                     "WANDLER_MENGENUMWERTER",
		KOMMUNIKATIONSEINRICHTUNG:                  "KOMMUNIKATIONSEINRICHTUNG",
		TECHNISCHE_STEUEREINRICHTUNG:               "TECHNISCHE_STEUEREINRICHTUNG",
		PARAGRAF_19_STROM_NEV_UMLAGE:               "PARAGRAF_19_STROM_NEV_UMLAGE",
		BEFESTIGUNGSEINRICHTUNG:                    "BEFESTIGUNGSEINRICHTUNG",
		OFFSHORE_HAFTUNGSUMLAGE:                    "OFFSHORE_HAFTUNGSUMLAGE",
		FIXE_ARBEITSENTGELTKOMPONENTE:              "FIXE_ARBEITSENTGELTKOMPONENTE",
		FIXE_LEISTUNGSENTGELTKOMPONENTE:            "FIXE_LEISTUNGSENTGELTKOMPONENTE",
		UMLAGE_ABSCHALTBARE_LASTEN:                 "UMLAGE_ABSCHALTBARE_LASTEN",
		MEHRMENGE:                                  "MEHRMENGE",
		MINDERMENGE:                                "MINDERMENGE",
		ENERGIESTEUER:                              "ENERGIESTEUER",
		SMARTMETER_GATEWAY:                         "SMARTMETER_GATEWAY",
		STEUERBOX:                                  "STEUERBOX",
		MSB_INKL_MESSUNG:                           "MSB_INKL_MESSUNG",
		AUSGLEICHSENERGIE_UNTERDECKUNG:             "AUSGLEICHSENERGIE_UNTERDECKUNG",
	}
)

func init() {
	var v BDEWArtikelnummer
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BDEWArtikelnummerNameToValue = map[string]BDEWArtikelnummer{
			interface{}(LEISTUNG).(fmt.Stringer).String():                                   LEISTUNG,
			interface{}(LEISTUNG_PAUSCHAL).(fmt.Stringer).String():                          LEISTUNG_PAUSCHAL,
			interface{}(GRUNDPREIS).(fmt.Stringer).String():                                 GRUNDPREIS,
			interface{}(REGELENERGIE_ARBEIT).(fmt.Stringer).String():                        REGELENERGIE_ARBEIT,
			interface{}(REGELENERGIE_LEISTUNG).(fmt.Stringer).String():                      REGELENERGIE_LEISTUNG,
			interface{}(NOTSTROMLIEFERUNG_ARBEIT).(fmt.Stringer).String():                   NOTSTROMLIEFERUNG_ARBEIT,
			interface{}(NOTSTROMLIEFERUNG_LEISTUNG).(fmt.Stringer).String():                 NOTSTROMLIEFERUNG_LEISTUNG,
			interface{}(RESERVENETZKAPAZITAET).(fmt.Stringer).String():                      RESERVENETZKAPAZITAET,
			interface{}(RESERVELEISTUNG).(fmt.Stringer).String():                            RESERVELEISTUNG,
			interface{}(ZUSAETZLICHE_ABLESUNG).(fmt.Stringer).String():                      ZUSAETZLICHE_ABLESUNG,
			interface{}(PRUEFGEBUEHREN_AUSSERPLANMAESSIG).(fmt.Stringer).String():           PRUEFGEBUEHREN_AUSSERPLANMAESSIG,
			interface{}(WIRKARBEIT).(fmt.Stringer).String():                                 WIRKARBEIT,
			interface{}(SINGULAER_GENUTZTE_BETRIEBSMITTEL).(fmt.Stringer).String():          SINGULAER_GENUTZTE_BETRIEBSMITTEL,
			interface{}(ABGABE_KWKG).(fmt.Stringer).String():                                ABGABE_KWKG,
			interface{}(ABSCHLAG).(fmt.Stringer).String():                                   ABSCHLAG,
			interface{}(KONZESSIONSABGABE).(fmt.Stringer).String():                          KONZESSIONSABGABE,
			interface{}(ENTGELT_FERNAUSLESUNG).(fmt.Stringer).String():                      ENTGELT_FERNAUSLESUNG,
			interface{}(UNTERMESSUNG).(fmt.Stringer).String():                               UNTERMESSUNG,
			interface{}(BLINDMEHRARBEIT).(fmt.Stringer).String():                            BLINDMEHRARBEIT,
			interface{}(ENTGELT_ABRECHNUNG).(fmt.Stringer).String():                         ENTGELT_ABRECHNUNG,
			interface{}(SPERRKOSTEN).(fmt.Stringer).String():                                SPERRKOSTEN,
			interface{}(ENTSPERRKOSTEN).(fmt.Stringer).String():                             ENTSPERRKOSTEN,
			interface{}(MAHNKOSTEN).(fmt.Stringer).String():                                 MAHNKOSTEN,
			interface{}(MEHR_MINDERMENGEN).(fmt.Stringer).String():                          MEHR_MINDERMENGEN,
			interface{}(INKASSOKOSTEN).(fmt.Stringer).String():                              INKASSOKOSTEN,
			interface{}(BLINDMEHRLEISTUNG).(fmt.Stringer).String():                          BLINDMEHRLEISTUNG,
			interface{}(ENTGELT_MESSUNG_ABLESUNG).(fmt.Stringer).String():                   ENTGELT_MESSUNG_ABLESUNG,
			interface{}(ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK).(fmt.Stringer).String(): ENTGELT_EINBAU_BETRIEB_WARTUNG_MESSTECHNIK,
			interface{}(AUSGLEICHSENERGIE).(fmt.Stringer).String():                          AUSGLEICHSENERGIE,
			interface{}(ZAEHLEINRICHTUNG).(fmt.Stringer).String():                           ZAEHLEINRICHTUNG,
			interface{}(WANDLER_MENGENUMWERTER).(fmt.Stringer).String():                     WANDLER_MENGENUMWERTER,
			interface{}(KOMMUNIKATIONSEINRICHTUNG).(fmt.Stringer).String():                  KOMMUNIKATIONSEINRICHTUNG,
			interface{}(TECHNISCHE_STEUEREINRICHTUNG).(fmt.Stringer).String():               TECHNISCHE_STEUEREINRICHTUNG,
			interface{}(PARAGRAF_19_STROM_NEV_UMLAGE).(fmt.Stringer).String():               PARAGRAF_19_STROM_NEV_UMLAGE,
			interface{}(BEFESTIGUNGSEINRICHTUNG).(fmt.Stringer).String():                    BEFESTIGUNGSEINRICHTUNG,
			interface{}(OFFSHORE_HAFTUNGSUMLAGE).(fmt.Stringer).String():                    OFFSHORE_HAFTUNGSUMLAGE,
			interface{}(FIXE_ARBEITSENTGELTKOMPONENTE).(fmt.Stringer).String():              FIXE_ARBEITSENTGELTKOMPONENTE,
			interface{}(FIXE_LEISTUNGSENTGELTKOMPONENTE).(fmt.Stringer).String():            FIXE_LEISTUNGSENTGELTKOMPONENTE,
			interface{}(UMLAGE_ABSCHALTBARE_LASTEN).(fmt.Stringer).String():                 UMLAGE_ABSCHALTBARE_LASTEN,
			interface{}(MEHRMENGE).(fmt.Stringer).String():                                  MEHRMENGE,
			interface{}(MINDERMENGE).(fmt.Stringer).String():                                MINDERMENGE,
			interface{}(ENERGIESTEUER).(fmt.Stringer).String():                              ENERGIESTEUER,
			interface{}(SMARTMETER_GATEWAY).(fmt.Stringer).String():                         SMARTMETER_GATEWAY,
			interface{}(STEUERBOX).(fmt.Stringer).String():                                  STEUERBOX,
			interface{}(MSB_INKL_MESSUNG).(fmt.Stringer).String():                           MSB_INKL_MESSUNG,
			interface{}(AUSGLEICHSENERGIE_UNTERDECKUNG).(fmt.Stringer).String():             AUSGLEICHSENERGIE_UNTERDECKUNG,
		}
	}
}

// MarshalJSON is generated so BDEWArtikelnummer satisfies json.Marshaler.
func (r BDEWArtikelnummer) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BDEWArtikelnummerValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BDEWArtikelnummer: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BDEWArtikelnummer satisfies json.Unmarshaler.
func (r *BDEWArtikelnummer) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BDEWArtikelnummer should be a string, got %s", data)
	}
	v, ok := _BDEWArtikelnummerNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BDEWArtikelnummer %q", s)
	}
	*r = v
	return nil
}
