// Code generated by "stringer --type Fernschaltung"; DO NOT EDIT.

package fernschaltung

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VORHANDEN-1]
	_ = x[NICHT_VORHANDEN-2]
}

const _Fernschaltung_name = "VORHANDENNICHT_VORHANDEN"

var _Fernschaltung_index = [...]uint8{0, 9, 24}

func (i Fernschaltung) String() string {
	i -= 1
	if i < 0 || i >= Fernschaltung(len(_Fernschaltung_index)-1) {
		return "Fernschaltung(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Fernschaltung_name[_Fernschaltung_index[i]:_Fernschaltung_index[i+1]]
}
