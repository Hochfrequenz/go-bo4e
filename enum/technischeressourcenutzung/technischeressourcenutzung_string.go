// Code generated by "stringer --type Technischeressourcenutzung"; DO NOT EDIT.

package technischeressourcenutzung

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[STROMVERBRAUCHSART-1]
	_ = x[STROMERZEUGUNGSART-2]
	_ = x[SPEICHER-3]
}

const _Technischeressourcenutzung_name = "STROMVERBRAUCHSARTSTROMERZEUGUNGSARTSPEICHER"

var _Technischeressourcenutzung_index = [...]uint8{0, 18, 36, 44}

func (i Technischeressourcenutzung) String() string {
	i -= 1
	if i < 0 || i >= Technischeressourcenutzung(len(_Technischeressourcenutzung_index)-1) {
		return "Technischeressourcenutzung(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Technischeressourcenutzung_name[_Technischeressourcenutzung_index[i]:_Technischeressourcenutzung_index[i+1]]
}
