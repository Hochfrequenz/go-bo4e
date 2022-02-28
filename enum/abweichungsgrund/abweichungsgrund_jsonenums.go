// Code generated by jsonenums -type=AbweichungsGrund; DO NOT EDIT.

package abweichungsgrund

import (
	"encoding/json"
	"fmt"
)

var (
	_AbweichungsGrundNameToValue = map[string]AbweichungsGrund{
		"PREIS_RECHENREGEL_FALSCH":                          PREIS_RECHENREGEL_FALSCH,
		"FALSCHER_ABRECHNUNGSZEITRAUM":                      FALSCHER_ABRECHNUNGSZEITRAUM,
		"UNBEKANNTE_MARKTLOKATION_MESSLOKATION":             UNBEKANNTE_MARKTLOKATION_MESSLOKATION,
		"SONSTIGER_ABWEICHUNGSGRUND":                        SONSTIGER_ABWEICHUNGSGRUND,
		"DOPPELTE_RECHNUNG":                                 DOPPELTE_RECHNUNG,
		"ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN":         ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
		"ABRECHNUNGSENDE_UNGLEICH_VERTRAGSENDE":             ABRECHNUNGSENDE_UNGLEICH_VERTRAGSENDE,
		"BETRAG_DER_ABSCHLAGSRECHNUNG_FALSCH":               BETRAG_DER_ABSCHLAGSRECHNUNG_FALSCH,
		"VORAUSBEZAHLTER_BETRAG_FALSCH":                     VORAUSBEZAHLTER_BETRAG_FALSCH,
		"ARTIKEL_NICHT_VEREINBART":                          ARTIKEL_NICHT_VEREINBART,
		"NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FEHLEN":        NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FEHLEN,
		"RECHNUNGSNUMMER_BEREITS_ERHALTEN":                  RECHNUNGSNUMMER_BEREITS_ERHALTEN,
		"NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FALSCH":        NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FALSCH,
		"ZEITLICHE_MENGENANGABE_FEHLERHAFT":                 ZEITLICHE_MENGENANGABE_FEHLERHAFT,
		"FALSCHER_BILANZIERUNGSBEGINN":                      FALSCHER_BILANZIERUNGSBEGINN,
		"FALSCHES_NETZNUTZUNGSENDE":                         FALSCHES_NETZNUTZUNGSENDE,
		"BILANZIERTE_MENGE_FEHLT":                           BILANZIERTE_MENGE_FEHLT,
		"BILANZIERTE_MENGE_FALSCH":                          BILANZIERTE_MENGE_FALSCH,
		"NETZNUTZUNGSABRECHNUNG_FEHLT":                      NETZNUTZUNGSABRECHNUNG_FEHLT,
		"REVERSE_CHARGE_ANWENDUNG_FEHLT_ODER_FEHLERHAFT":    REVERSE_CHARGE_ANWENDUNG_FEHLT_ODER_FEHLERHAFT,
		"ALLOKATIONSLISTE_FEHLT":                            ALLOKATIONSLISTE_FEHLT,
		"MEHR_MINDERMENGE_FALSCH":                           MEHR_MINDERMENGE_FALSCH,
		"UNGUELTIGES_RECHNUNGSDATUM":                        UNGUELTIGES_RECHNUNGSDATUM,
		"ZEITINTERVALL_DER_BILANZIERTEN_MENGE_INKONSISTENT": ZEITINTERVALL_DER_BILANZIERTEN_MENGE_INKONSISTENT,
		"RECHNUNGSEMPFAENGER_WIDERSPRICHT_DER_STEUERRECHTLICHEN_EINSCHAETZUNG_DES_RECHNUNGSSTELLERS": RECHNUNGSEMPFAENGER_WIDERSPRICHT_DER_STEUERRECHTLICHEN_EINSCHAETZUNG_DES_RECHNUNGSSTELLERS,
		"ANGEGEBENE_QUOTES_AN_MARKTLOKATION_NICHT_VORHANDEN":                                         ANGEGEBENE_QUOTES_AN_MARKTLOKATION_NICHT_VORHANDEN,
		"RECHNUNGSABWICKLUNG_NICHT_VEREINBART":                                                       RECHNUNGSABWICKLUNG_NICHT_VEREINBART,
		"COMDIS_WIRD_ABGELEHNT":                                                                      COMDIS_WIRD_ABGELEHNT,
	}

	_AbweichungsGrundValueToName = map[AbweichungsGrund]string{
		PREIS_RECHENREGEL_FALSCH:                          "PREIS_RECHENREGEL_FALSCH",
		FALSCHER_ABRECHNUNGSZEITRAUM:                      "FALSCHER_ABRECHNUNGSZEITRAUM",
		UNBEKANNTE_MARKTLOKATION_MESSLOKATION:             "UNBEKANNTE_MARKTLOKATION_MESSLOKATION",
		SONSTIGER_ABWEICHUNGSGRUND:                        "SONSTIGER_ABWEICHUNGSGRUND",
		DOPPELTE_RECHNUNG:                                 "DOPPELTE_RECHNUNG",
		ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN:         "ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN",
		ABRECHNUNGSENDE_UNGLEICH_VERTRAGSENDE:             "ABRECHNUNGSENDE_UNGLEICH_VERTRAGSENDE",
		BETRAG_DER_ABSCHLAGSRECHNUNG_FALSCH:               "BETRAG_DER_ABSCHLAGSRECHNUNG_FALSCH",
		VORAUSBEZAHLTER_BETRAG_FALSCH:                     "VORAUSBEZAHLTER_BETRAG_FALSCH",
		ARTIKEL_NICHT_VEREINBART:                          "ARTIKEL_NICHT_VEREINBART",
		NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FEHLEN:        "NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FEHLEN",
		RECHNUNGSNUMMER_BEREITS_ERHALTEN:                  "RECHNUNGSNUMMER_BEREITS_ERHALTEN",
		NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FALSCH:        "NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FALSCH",
		ZEITLICHE_MENGENANGABE_FEHLERHAFT:                 "ZEITLICHE_MENGENANGABE_FEHLERHAFT",
		FALSCHER_BILANZIERUNGSBEGINN:                      "FALSCHER_BILANZIERUNGSBEGINN",
		FALSCHES_NETZNUTZUNGSENDE:                         "FALSCHES_NETZNUTZUNGSENDE",
		BILANZIERTE_MENGE_FEHLT:                           "BILANZIERTE_MENGE_FEHLT",
		BILANZIERTE_MENGE_FALSCH:                          "BILANZIERTE_MENGE_FALSCH",
		NETZNUTZUNGSABRECHNUNG_FEHLT:                      "NETZNUTZUNGSABRECHNUNG_FEHLT",
		REVERSE_CHARGE_ANWENDUNG_FEHLT_ODER_FEHLERHAFT:    "REVERSE_CHARGE_ANWENDUNG_FEHLT_ODER_FEHLERHAFT",
		ALLOKATIONSLISTE_FEHLT:                            "ALLOKATIONSLISTE_FEHLT",
		MEHR_MINDERMENGE_FALSCH:                           "MEHR_MINDERMENGE_FALSCH",
		UNGUELTIGES_RECHNUNGSDATUM:                        "UNGUELTIGES_RECHNUNGSDATUM",
		ZEITINTERVALL_DER_BILANZIERTEN_MENGE_INKONSISTENT: "ZEITINTERVALL_DER_BILANZIERTEN_MENGE_INKONSISTENT",
		RECHNUNGSEMPFAENGER_WIDERSPRICHT_DER_STEUERRECHTLICHEN_EINSCHAETZUNG_DES_RECHNUNGSSTELLERS: "RECHNUNGSEMPFAENGER_WIDERSPRICHT_DER_STEUERRECHTLICHEN_EINSCHAETZUNG_DES_RECHNUNGSSTELLERS",
		ANGEGEBENE_QUOTES_AN_MARKTLOKATION_NICHT_VORHANDEN:                                         "ANGEGEBENE_QUOTES_AN_MARKTLOKATION_NICHT_VORHANDEN",
		RECHNUNGSABWICKLUNG_NICHT_VEREINBART:                                                       "RECHNUNGSABWICKLUNG_NICHT_VEREINBART",
		COMDIS_WIRD_ABGELEHNT:                                                                      "COMDIS_WIRD_ABGELEHNT",
	}
)

func init() {
	var v AbweichungsGrund
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_AbweichungsGrundNameToValue = map[string]AbweichungsGrund{
			interface{}(PREIS_RECHENREGEL_FALSCH).(fmt.Stringer).String():                                                                   PREIS_RECHENREGEL_FALSCH,
			interface{}(FALSCHER_ABRECHNUNGSZEITRAUM).(fmt.Stringer).String():                                                               FALSCHER_ABRECHNUNGSZEITRAUM,
			interface{}(UNBEKANNTE_MARKTLOKATION_MESSLOKATION).(fmt.Stringer).String():                                                      UNBEKANNTE_MARKTLOKATION_MESSLOKATION,
			interface{}(SONSTIGER_ABWEICHUNGSGRUND).(fmt.Stringer).String():                                                                 SONSTIGER_ABWEICHUNGSGRUND,
			interface{}(DOPPELTE_RECHNUNG).(fmt.Stringer).String():                                                                          DOPPELTE_RECHNUNG,
			interface{}(ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN).(fmt.Stringer).String():                                                  ABRECHNUNGSBEGINN_UNGLEICH_VERTRAGSBEGINN,
			interface{}(ABRECHNUNGSENDE_UNGLEICH_VERTRAGSENDE).(fmt.Stringer).String():                                                      ABRECHNUNGSENDE_UNGLEICH_VERTRAGSENDE,
			interface{}(BETRAG_DER_ABSCHLAGSRECHNUNG_FALSCH).(fmt.Stringer).String():                                                        BETRAG_DER_ABSCHLAGSRECHNUNG_FALSCH,
			interface{}(VORAUSBEZAHLTER_BETRAG_FALSCH).(fmt.Stringer).String():                                                              VORAUSBEZAHLTER_BETRAG_FALSCH,
			interface{}(ARTIKEL_NICHT_VEREINBART).(fmt.Stringer).String():                                                                   ARTIKEL_NICHT_VEREINBART,
			interface{}(NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FEHLEN).(fmt.Stringer).String():                                                 NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FEHLEN,
			interface{}(RECHNUNGSNUMMER_BEREITS_ERHALTEN).(fmt.Stringer).String():                                                           RECHNUNGSNUMMER_BEREITS_ERHALTEN,
			interface{}(NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FALSCH).(fmt.Stringer).String():                                                 NETZNUTZUNGSMESSWERTE_ENERGIEMENGEN_FALSCH,
			interface{}(ZEITLICHE_MENGENANGABE_FEHLERHAFT).(fmt.Stringer).String():                                                          ZEITLICHE_MENGENANGABE_FEHLERHAFT,
			interface{}(FALSCHER_BILANZIERUNGSBEGINN).(fmt.Stringer).String():                                                               FALSCHER_BILANZIERUNGSBEGINN,
			interface{}(FALSCHES_NETZNUTZUNGSENDE).(fmt.Stringer).String():                                                                  FALSCHES_NETZNUTZUNGSENDE,
			interface{}(BILANZIERTE_MENGE_FEHLT).(fmt.Stringer).String():                                                                    BILANZIERTE_MENGE_FEHLT,
			interface{}(BILANZIERTE_MENGE_FALSCH).(fmt.Stringer).String():                                                                   BILANZIERTE_MENGE_FALSCH,
			interface{}(NETZNUTZUNGSABRECHNUNG_FEHLT).(fmt.Stringer).String():                                                               NETZNUTZUNGSABRECHNUNG_FEHLT,
			interface{}(REVERSE_CHARGE_ANWENDUNG_FEHLT_ODER_FEHLERHAFT).(fmt.Stringer).String():                                             REVERSE_CHARGE_ANWENDUNG_FEHLT_ODER_FEHLERHAFT,
			interface{}(ALLOKATIONSLISTE_FEHLT).(fmt.Stringer).String():                                                                     ALLOKATIONSLISTE_FEHLT,
			interface{}(MEHR_MINDERMENGE_FALSCH).(fmt.Stringer).String():                                                                    MEHR_MINDERMENGE_FALSCH,
			interface{}(UNGUELTIGES_RECHNUNGSDATUM).(fmt.Stringer).String():                                                                 UNGUELTIGES_RECHNUNGSDATUM,
			interface{}(ZEITINTERVALL_DER_BILANZIERTEN_MENGE_INKONSISTENT).(fmt.Stringer).String():                                          ZEITINTERVALL_DER_BILANZIERTEN_MENGE_INKONSISTENT,
			interface{}(RECHNUNGSEMPFAENGER_WIDERSPRICHT_DER_STEUERRECHTLICHEN_EINSCHAETZUNG_DES_RECHNUNGSSTELLERS).(fmt.Stringer).String(): RECHNUNGSEMPFAENGER_WIDERSPRICHT_DER_STEUERRECHTLICHEN_EINSCHAETZUNG_DES_RECHNUNGSSTELLERS,
			interface{}(ANGEGEBENE_QUOTES_AN_MARKTLOKATION_NICHT_VORHANDEN).(fmt.Stringer).String():                                         ANGEGEBENE_QUOTES_AN_MARKTLOKATION_NICHT_VORHANDEN,
			interface{}(RECHNUNGSABWICKLUNG_NICHT_VEREINBART).(fmt.Stringer).String():                                                       RECHNUNGSABWICKLUNG_NICHT_VEREINBART,
			interface{}(COMDIS_WIRD_ABGELEHNT).(fmt.Stringer).String():                                                                      COMDIS_WIRD_ABGELEHNT,
		}
	}
}

// MarshalJSON is generated so AbweichungsGrund satisfies json.Marshaler.
func (r AbweichungsGrund) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _AbweichungsGrundValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid AbweichungsGrund: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AbweichungsGrund satisfies json.Unmarshaler.
func (r *AbweichungsGrund) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AbweichungsGrund should be a string, got %s", data)
	}
	v, ok := _AbweichungsGrundNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid AbweichungsGrund %q", s)
	}
	*r = v
	return nil
}
