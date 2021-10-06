// Code generated by "stringer --type NNRechnungstyp"; DO NOT EDIT.

package nnrechnungstyp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ABSCHLUSSRECHNUNG-1]
	_ = x[ABSCHLAGSRECHNUNG-2]
	_ = x[TURNUSRECHNUNG-3]
	_ = x[MONATSRECHNUNG-4]
	_ = x[WIMRECHNUNG-5]
	_ = x[ZWISCHENRECHNUNG-6]
	_ = x[INTEGRIERTE_13TE_RECHNUNG-7]
	_ = x[ZUSAETZLICHE_13TE_RECHNUNG-8]
	_ = x[MEHRMINDERMENGENRECHNUNG-9]
	_ = x[MSBRECHNUNG-10]
}

const _NNRechnungstyp_name = "ABSCHLUSSRECHNUNGABSCHLAGSRECHNUNGTURNUSRECHNUNGMONATSRECHNUNGWIMRECHNUNGZWISCHENRECHNUNGINTEGRIERTE_13TE_RECHNUNGZUSAETZLICHE_13TE_RECHNUNGMEHRMINDERMENGENRECHNUNGMSBRECHNUNG"

var _NNRechnungstyp_index = [...]uint8{0, 17, 34, 48, 62, 73, 89, 114, 140, 164, 175}

func (i NNRechnungstyp) String() string {
	i -= 1
	if i < 0 || i >= NNRechnungstyp(len(_NNRechnungstyp_index)-1) {
		return "NNRechnungstyp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _NNRechnungstyp_name[_NNRechnungstyp_index[i]:_NNRechnungstyp_index[i+1]]
}
