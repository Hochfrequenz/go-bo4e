// Code generated by "stringer --type RechnungspositionsZuschlag"; DO NOT EDIT.

package rechnungspositionszuschlag

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UMSPANNUNGSZUSCHLAG-1]
	_ = x[ALLEIN_GENUTZTE_BETRIEBSMITTEL_STROM_NEV-2]
	_ = x[ANPASSUNG_STROM_NEV_19_2-3]
	_ = x[PAUSCHALE_NETZENTGELTREDUZIERUNG_ENWG_14A-4]
}

const _RechnungspositionsZuschlag_name = "UMSPANNUNGSZUSCHLAGALLEIN_GENUTZTE_BETRIEBSMITTEL_STROM_NEVANPASSUNG_STROM_NEV_19_2PAUSCHALE_NETZENTGELTREDUZIERUNG_ENWG_14A"

var _RechnungspositionsZuschlag_index = [...]uint8{0, 19, 59, 83, 124}

func (i RechnungspositionsZuschlag) String() string {
	i -= 1
	if i < 0 || i >= RechnungspositionsZuschlag(len(_RechnungspositionsZuschlag_index)-1) {
		return "RechnungspositionsZuschlag(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RechnungspositionsZuschlag_name[_RechnungspositionsZuschlag_index[i]:_RechnungspositionsZuschlag_index[i+1]]
}
