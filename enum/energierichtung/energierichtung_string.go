// Code generated by "stringer --type Energierichtung"; DO NOT EDIT.

package energierichtung

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AUSSP-1]
	_ = x[EINSP-2]
}

const _Energierichtung_name = "AUSSPEINSP"

var _Energierichtung_index = [...]uint8{0, 5, 10}

func (i Energierichtung) String() string {
	i -= 1
	if i < 0 || i >= Energierichtung(len(_Energierichtung_index)-1) {
		return "Energierichtung(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Energierichtung_name[_Energierichtung_index[i]:_Energierichtung_index[i+1]]
}
