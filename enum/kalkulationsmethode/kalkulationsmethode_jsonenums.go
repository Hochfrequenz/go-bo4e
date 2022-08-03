// Code generated by jsonenums --type Kalkulationsmethode; DO NOT EDIT.

package kalkulationsmethode

import (
	"encoding/json"
	"fmt"
)

var (
	_KalkulationsmethodeNameToValue = map[string]Kalkulationsmethode{
		"KEINE":                         KEINE,
		"STAFFELN":                      STAFFELN,
		"ZONEN":                         ZONEN,
		"VORZONEN_GP":                   VORZONEN_GP,
		"SIGMOID":                       SIGMOID,
		"BLINDARBEIT_GT_50_PROZENT":     BLINDARBEIT_GT_50_PROZENT,
		"BLINDARBEIT_GT_40_PROZENT":     BLINDARBEIT_GT_40_PROZENT,
		"BLINDARBEIT_MIT_FREIMENGE":     BLINDARBEIT_MIT_FREIMENGE,
		"AP_GP_ZONEN":                   AP_GP_ZONEN,
		"LP_INSTALL_LEISTUNG":           LP_INSTALL_LEISTUNG,
		"AP_TRANSPORT_ODER_VERTEILNETZ": AP_TRANSPORT_ODER_VERTEILNETZ,
		"AP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID": AP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID,
		"LP_JAHRESVERBRAUCH":                                    LP_JAHRESVERBRAUCH,
		"LP_TRANSPORT_ODER_VERTEILNETZ":                         LP_TRANSPORT_ODER_VERTEILNETZ,
		"LP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID": LP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID,
		"FUNKTIONEN": FUNKTIONEN,
		"VERBRAUCH_UEBER_SLP_GRENZE_FUNKTIONSBEZOGEN_WEITERE_BERECHNUNG_ALS_LGK": VERBRAUCH_UEBER_SLP_GRENZE_FUNKTIONSBEZOGEN_WEITERE_BERECHNUNG_ALS_LGK,
	}

	_KalkulationsmethodeValueToName = map[Kalkulationsmethode]string{
		KEINE:                         "KEINE",
		STAFFELN:                      "STAFFELN",
		ZONEN:                         "ZONEN",
		VORZONEN_GP:                   "VORZONEN_GP",
		SIGMOID:                       "SIGMOID",
		BLINDARBEIT_GT_50_PROZENT:     "BLINDARBEIT_GT_50_PROZENT",
		BLINDARBEIT_GT_40_PROZENT:     "BLINDARBEIT_GT_40_PROZENT",
		BLINDARBEIT_MIT_FREIMENGE:     "BLINDARBEIT_MIT_FREIMENGE",
		AP_GP_ZONEN:                   "AP_GP_ZONEN",
		LP_INSTALL_LEISTUNG:           "LP_INSTALL_LEISTUNG",
		AP_TRANSPORT_ODER_VERTEILNETZ: "AP_TRANSPORT_ODER_VERTEILNETZ",
		AP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID: "AP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID",
		LP_JAHRESVERBRAUCH:                                    "LP_JAHRESVERBRAUCH",
		LP_TRANSPORT_ODER_VERTEILNETZ:                         "LP_TRANSPORT_ODER_VERTEILNETZ",
		LP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID: "LP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID",
		FUNKTIONEN: "FUNKTIONEN",
		VERBRAUCH_UEBER_SLP_GRENZE_FUNKTIONSBEZOGEN_WEITERE_BERECHNUNG_ALS_LGK: "VERBRAUCH_UEBER_SLP_GRENZE_FUNKTIONSBEZOGEN_WEITERE_BERECHNUNG_ALS_LGK",
	}
)

func init() {
	var v Kalkulationsmethode
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_KalkulationsmethodeNameToValue = map[string]Kalkulationsmethode{
			interface{}(KEINE).(fmt.Stringer).String():                                                                  KEINE,
			interface{}(STAFFELN).(fmt.Stringer).String():                                                               STAFFELN,
			interface{}(ZONEN).(fmt.Stringer).String():                                                                  ZONEN,
			interface{}(VORZONEN_GP).(fmt.Stringer).String():                                                            VORZONEN_GP,
			interface{}(SIGMOID).(fmt.Stringer).String():                                                                SIGMOID,
			interface{}(BLINDARBEIT_GT_50_PROZENT).(fmt.Stringer).String():                                              BLINDARBEIT_GT_50_PROZENT,
			interface{}(BLINDARBEIT_GT_40_PROZENT).(fmt.Stringer).String():                                              BLINDARBEIT_GT_40_PROZENT,
			interface{}(BLINDARBEIT_MIT_FREIMENGE).(fmt.Stringer).String():                                              BLINDARBEIT_MIT_FREIMENGE,
			interface{}(AP_GP_ZONEN).(fmt.Stringer).String():                                                            AP_GP_ZONEN,
			interface{}(LP_INSTALL_LEISTUNG).(fmt.Stringer).String():                                                    LP_INSTALL_LEISTUNG,
			interface{}(AP_TRANSPORT_ODER_VERTEILNETZ).(fmt.Stringer).String():                                          AP_TRANSPORT_ODER_VERTEILNETZ,
			interface{}(AP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID).(fmt.Stringer).String():                  AP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID,
			interface{}(LP_JAHRESVERBRAUCH).(fmt.Stringer).String():                                                     LP_JAHRESVERBRAUCH,
			interface{}(LP_TRANSPORT_ODER_VERTEILNETZ).(fmt.Stringer).String():                                          LP_TRANSPORT_ODER_VERTEILNETZ,
			interface{}(LP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID).(fmt.Stringer).String():                  LP_TRANSPORT_ODER_VERTEILNETZ_ORTSVERTEILNETZ_SIGMOID,
			interface{}(FUNKTIONEN).(fmt.Stringer).String():                                                             FUNKTIONEN,
			interface{}(VERBRAUCH_UEBER_SLP_GRENZE_FUNKTIONSBEZOGEN_WEITERE_BERECHNUNG_ALS_LGK).(fmt.Stringer).String(): VERBRAUCH_UEBER_SLP_GRENZE_FUNKTIONSBEZOGEN_WEITERE_BERECHNUNG_ALS_LGK,
		}
	}
}

// MarshalJSON is generated so Kalkulationsmethode satisfies json.Marshaler.
func (r Kalkulationsmethode) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _KalkulationsmethodeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Kalkulationsmethode: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Kalkulationsmethode satisfies json.Unmarshaler.
func (r *Kalkulationsmethode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Kalkulationsmethode should be a string, got %s", data)
	}
	v, ok := _KalkulationsmethodeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Kalkulationsmethode %q", s)
	}
	*r = v
	return nil
}
