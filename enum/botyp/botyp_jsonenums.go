// Code generated by jsonenums --type BOTyp; DO NOT EDIT.

package botyp

import (
	"encoding/json"
	"fmt"
)

var (
	_BOTypNameToValue = map[string]BOTyp{
		"ANGEBOT":                     ANGEBOT,
		"ANSPRECHPARTNER":             ANSPRECHPARTNER,
		"AUSSCHREIBUNG":               AUSSCHREIBUNG,
		"BILANZIERUNG":                BILANZIERUNG,
		"ENERGIEMENGE":                ENERGIEMENGE,
		"GESCHAEFTSPARTNER":           GESCHAEFTSPARTNER,
		"KOSTEN":                      KOSTEN,
		"LASTGANG":                    LASTGANG,
		"MARKTLOKATION":               MARKTLOKATION,
		"MESSLOKATION":                MESSLOKATION,
		"MARKTTEILNEHMER":             MARKTTEILNEHMER,
		"NETZNUTZUNGSRECHNUNG":        NETZNUTZUNGSRECHNUNG,
		"PREISBLATT":                  PREISBLATT,
		"PREISBLATTDIENSTLEISTUNG":    PREISBLATTDIENSTLEISTUNG,
		"PREISBLATTKONZESSIONSABGABE": PREISBLATTKONZESSIONSABGABE,
		"PREISBLATTMESSUNG":           PREISBLATTMESSUNG,
		"PREISBLATTUMLAGEN":           PREISBLATTUMLAGEN,
		"RECHNUNG":                    RECHNUNG,
		"REKLAMATION":                 REKLAMATION,
		"TARIFINFO":                   TARIFINFO,
		"TARIFPREISBLATT":             TARIFPREISBLATT,
		"VERTRAG":                     VERTRAG,
		"ZAEHLER":                     ZAEHLER,
		"ZEITREIHE":                   ZEITREIHE,
		"HANDELSUNSTIMMIGKEIT":        HANDELSUNSTIMMIGKEIT,
		"AVIS":                        AVIS,
		"STATUSBERICHT":               STATUSBERICHT,
	}

	_BOTypValueToName = map[BOTyp]string{
		ANGEBOT:                     "ANGEBOT",
		ANSPRECHPARTNER:             "ANSPRECHPARTNER",
		AUSSCHREIBUNG:               "AUSSCHREIBUNG",
		BILANZIERUNG:                "BILANZIERUNG",
		ENERGIEMENGE:                "ENERGIEMENGE",
		GESCHAEFTSPARTNER:           "GESCHAEFTSPARTNER",
		KOSTEN:                      "KOSTEN",
		LASTGANG:                    "LASTGANG",
		MARKTLOKATION:               "MARKTLOKATION",
		MESSLOKATION:                "MESSLOKATION",
		MARKTTEILNEHMER:             "MARKTTEILNEHMER",
		NETZNUTZUNGSRECHNUNG:        "NETZNUTZUNGSRECHNUNG",
		PREISBLATT:                  "PREISBLATT",
		PREISBLATTDIENSTLEISTUNG:    "PREISBLATTDIENSTLEISTUNG",
		PREISBLATTKONZESSIONSABGABE: "PREISBLATTKONZESSIONSABGABE",
		PREISBLATTMESSUNG:           "PREISBLATTMESSUNG",
		PREISBLATTUMLAGEN:           "PREISBLATTUMLAGEN",
		RECHNUNG:                    "RECHNUNG",
		REKLAMATION:                 "REKLAMATION",
		TARIFINFO:                   "TARIFINFO",
		TARIFPREISBLATT:             "TARIFPREISBLATT",
		VERTRAG:                     "VERTRAG",
		ZAEHLER:                     "ZAEHLER",
		ZEITREIHE:                   "ZEITREIHE",
		HANDELSUNSTIMMIGKEIT:        "HANDELSUNSTIMMIGKEIT",
		AVIS:                        "AVIS",
		STATUSBERICHT:               "STATUSBERICHT",
	}
)

func init() {
	var v BOTyp
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BOTypNameToValue = map[string]BOTyp{
			interface{}(ANGEBOT).(fmt.Stringer).String():                     ANGEBOT,
			interface{}(ANSPRECHPARTNER).(fmt.Stringer).String():             ANSPRECHPARTNER,
			interface{}(AUSSCHREIBUNG).(fmt.Stringer).String():               AUSSCHREIBUNG,
			interface{}(BILANZIERUNG).(fmt.Stringer).String():                BILANZIERUNG,
			interface{}(ENERGIEMENGE).(fmt.Stringer).String():                ENERGIEMENGE,
			interface{}(GESCHAEFTSPARTNER).(fmt.Stringer).String():           GESCHAEFTSPARTNER,
			interface{}(KOSTEN).(fmt.Stringer).String():                      KOSTEN,
			interface{}(LASTGANG).(fmt.Stringer).String():                    LASTGANG,
			interface{}(MARKTLOKATION).(fmt.Stringer).String():               MARKTLOKATION,
			interface{}(MESSLOKATION).(fmt.Stringer).String():                MESSLOKATION,
			interface{}(MARKTTEILNEHMER).(fmt.Stringer).String():             MARKTTEILNEHMER,
			interface{}(NETZNUTZUNGSRECHNUNG).(fmt.Stringer).String():        NETZNUTZUNGSRECHNUNG,
			interface{}(PREISBLATT).(fmt.Stringer).String():                  PREISBLATT,
			interface{}(PREISBLATTDIENSTLEISTUNG).(fmt.Stringer).String():    PREISBLATTDIENSTLEISTUNG,
			interface{}(PREISBLATTKONZESSIONSABGABE).(fmt.Stringer).String(): PREISBLATTKONZESSIONSABGABE,
			interface{}(PREISBLATTMESSUNG).(fmt.Stringer).String():           PREISBLATTMESSUNG,
			interface{}(PREISBLATTUMLAGEN).(fmt.Stringer).String():           PREISBLATTUMLAGEN,
			interface{}(RECHNUNG).(fmt.Stringer).String():                    RECHNUNG,
			interface{}(REKLAMATION).(fmt.Stringer).String():                 REKLAMATION,
			interface{}(TARIFINFO).(fmt.Stringer).String():                   TARIFINFO,
			interface{}(TARIFPREISBLATT).(fmt.Stringer).String():             TARIFPREISBLATT,
			interface{}(VERTRAG).(fmt.Stringer).String():                     VERTRAG,
			interface{}(ZAEHLER).(fmt.Stringer).String():                     ZAEHLER,
			interface{}(ZEITREIHE).(fmt.Stringer).String():                   ZEITREIHE,
			interface{}(HANDELSUNSTIMMIGKEIT).(fmt.Stringer).String():        HANDELSUNSTIMMIGKEIT,
			interface{}(AVIS).(fmt.Stringer).String():                        AVIS,
			interface{}(STATUSBERICHT).(fmt.Stringer).String():               STATUSBERICHT,
		}
	}
}

// MarshalJSON is generated so BOTyp satisfies json.Marshaler.
func (r BOTyp) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BOTypValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BOTyp: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BOTyp satisfies json.Unmarshaler.
func (r *BOTyp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BOTyp should be a string, got %s", data)
	}
	v, ok := _BOTypNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BOTyp %q", s)
	}
	*r = v
	return nil
}
