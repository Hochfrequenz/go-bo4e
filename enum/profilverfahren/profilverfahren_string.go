// Code generated by "stringer --type Profilverfahren"; DO NOT EDIT.

package profilverfahren

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SYNTHETISCH-1]
	_ = x[ANALYTISCH-2]
}

const _Profilverfahren_name = "SYNTHETISCHANALYTISCH"

var _Profilverfahren_index = [...]uint8{0, 11, 21}

func (i Profilverfahren) String() string {
	i -= 1
	if i < 0 || i >= Profilverfahren(len(_Profilverfahren_index)-1) {
		return "Profilverfahren(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Profilverfahren_name[_Profilverfahren_index[i]:_Profilverfahren_index[i+1]]
}
