// Code generated by "stringer --type Rechnungstyp"; DO NOT EDIT.

package rechnungstyp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ABSCHLAGSRECHNUNG-1]
	_ = x[TURNUSRECHNUNG-2]
	_ = x[MONATSRECHNUNG-3]
	_ = x[WIMRECHNUNG-4]
	_ = x[ZWISCHENRECHNUNG-5]
	_ = x[INTEGRIERTE_13TE_RECHNUNG-6]
	_ = x[ZUSAETZLICHE_13TE_RECHNUNG-7]
	_ = x[MEHRMINDERMENGENRECHNUNG-8]
	_ = x[ABSCHLUSSRECHNUNG-9]
	_ = x[MSBRECHNUNG-10]
	_ = x[KAPAZITAETSRECHNUNG-11]
}

const _Rechnungstyp_name = "ABSCHLAGSRECHNUNGTURNUSRECHNUNGMONATSRECHNUNGWIMRECHNUNGZWISCHENRECHNUNGINTEGRIERTE_13TE_RECHNUNGZUSAETZLICHE_13TE_RECHNUNGMEHRMINDERMENGENRECHNUNGABSCHLUSSRECHNUNGMSBRECHNUNGKAPAZITAETSRECHNUNG"

var _Rechnungstyp_index = [...]uint8{0, 17, 31, 45, 56, 72, 97, 123, 147, 164, 175, 194}

func (i Rechnungstyp) String() string {
	i -= 1
	if i < 0 || i >= Rechnungstyp(len(_Rechnungstyp_index)-1) {
		return "Rechnungstyp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Rechnungstyp_name[_Rechnungstyp_index[i]:_Rechnungstyp_index[i+1]]
}
