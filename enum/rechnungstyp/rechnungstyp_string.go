// Code generated by "stringer --type Rechnungstyp"; DO NOT EDIT.

package rechnungstyp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ENDKUNDENRECHNUNG-1]
	_ = x[NETZNUTZUNGSRECHNUNG-2]
	_ = x[MEHRMINDERMENGENRECHNUNG-3]
	_ = x[MESSSTELLENBETRIEBSRECHNUNG-4]
	_ = x[BESCHAFFUNGSRECHNUNG-5]
	_ = x[AUSGLEICHSENERGIERECHNUNG-6]
	_ = x[ABSCHLAGSRECHNUNG-7]
	_ = x[WIMRECHNUNG-8]
	_ = x[SELBSTAUSGESTELLTERECHNUNGMEMI-9]
}

const _Rechnungstyp_name = "ENDKUNDENRECHNUNGNETZNUTZUNGSRECHNUNGMEHRMINDERMENGENRECHNUNGMESSSTELLENBETRIEBSRECHNUNGBESCHAFFUNGSRECHNUNGAUSGLEICHSENERGIERECHNUNGABSCHLAGSRECHNUNGWIMRECHNUNGSELBSTAUSGESTELLTERECHNUNGMEMI"

var _Rechnungstyp_index = [...]uint8{0, 17, 37, 61, 88, 108, 133, 150, 161, 191}

func (i Rechnungstyp) String() string {
	i -= 1
	if i < 0 || i >= Rechnungstyp(len(_Rechnungstyp_index)-1) {
		return "Rechnungstyp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Rechnungstyp_name[_Rechnungstyp_index[i]:_Rechnungstyp_index[i+1]]
}
