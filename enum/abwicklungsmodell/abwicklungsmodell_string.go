// Code generated by "stringer --type Abwicklungsmodell"; DO NOT EDIT.

package abwicklungsmodell

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MODELL_1-1]
	_ = x[MODELL_2-2]
}

const _Abwicklungsmodell_name = "MODELL_1MODELL_2"

var _Abwicklungsmodell_index = [...]uint8{0, 8, 16}

func (i Abwicklungsmodell) String() string {
	i -= 1
	if i < 0 || i >= Abwicklungsmodell(len(_Abwicklungsmodell_index)-1) {
		return "Abwicklungsmodell(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Abwicklungsmodell_name[_Abwicklungsmodell_index[i]:_Abwicklungsmodell_index[i+1]]
}
